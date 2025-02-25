#include "tracer.h"
#include "bpf_helpers.h"
#include "ip.h"
#include "ipv6.h"
#include "sock.h"
#include "sockfd.h"
#include "tags-types.h"
#include "port_range.h"
#include "http.h"
#include "https.h"

#define HTTPS_PORT 443
#define SO_SUFFIX_SIZE 3

static __always_inline void read_into_buffer_skb(char *buffer, struct __sk_buff* skb, skb_info_t *info) {
    u64 offset = (u64)info->data_off;

#pragma unroll
    for (int i = 0; i < HTTP_BUFFER_SIZE; i++) {
        if (offset < skb->len) {
            asm("r8 = *(u64 *)%[offset]\n\t"
                "r0 = 0\n\t"
                "r0 = *(u8 *)skb[r8]\n\t"
                "*(u8 *)%[buffer] = r0\n\t"
                : [buffer]"=m"(buffer[i])
                : [offset]"m"(offset)
                : "r0", "r1", "r2", "r3", "r4", "r5", "r8");
        }
        offset++;
    }
}

SEC("socket/http_filter")
int socket__http_filter(struct __sk_buff* skb) {
    skb_info_t skb_info;
    http_transaction_t http;
    __builtin_memset(&http, 0, sizeof(http));

    if (!read_conn_tuple_skb(skb, &skb_info, &http.tup)) {
        return 0;
    }

    // If the socket is for https and it is finishing,
    // make sure we pass it on to `http_process` to ensure that any ongoing transaction is flushed.
    // Otherwise, don't bother to inspect packet contents
    // when there is no chance we're dealing with plain HTTP (or a finishing HTTPS socket)
    if (!(http.tup.metadata&CONN_TYPE_TCP)) {
        return 0;
    }
    if ((http.tup.sport == HTTPS_PORT || http.tup.dport == HTTPS_PORT) && !(skb_info.tcp_flags & TCPHDR_FIN)) {
        return 0;
    }

    // src_port represents the source port number *before* normalization
    // for more context please refer to http-types.h comment on `owned_by_src_port` field
    http.owned_by_src_port = http.tup.sport;
    normalize_tuple(&http.tup);

    read_into_buffer_skb((char *)http.request_fragment, skb, &skb_info);
    http_process(&http, &skb_info);
    return 0;
}

// This kprobe is used to send batch completion notification to userspace
// because perf events can't be sent from socket filter programs
SEC("kretprobe/tcp_sendmsg")
int kretprobe__tcp_sendmsg(struct pt_regs* ctx) {
    http_notify_batch(ctx);
    return 0;
}

// this uprobe is essentially creating an index mapping a SSL context to a conn_tuple_t
SEC("uprobe/SSL_set_fd")
int uprobe__SSL_set_fd(struct pt_regs* ctx) {
    void *ssl_ctx = (void *)PT_REGS_PARM1(ctx);
    u32 socket_fd = (u32)PT_REGS_PARM2(ctx);
    init_ssl_sock(ssl_ctx, socket_fd);
    return 0;
}

SEC("uprobe/BIO_new_socket")
int uprobe__BIO_new_socket(struct pt_regs* ctx) {
    u64 pid_tgid = bpf_get_current_pid_tgid();
    u32 socket_fd = (u32)PT_REGS_PARM1(ctx);
    bpf_map_update_elem(&bio_new_socket_args, &pid_tgid, &socket_fd, BPF_ANY);
    return 0;
}

SEC("uretprobe/BIO_new_socket")
int uretprobe__BIO_new_socket(struct pt_regs* ctx) {
    u64 pid_tgid = bpf_get_current_pid_tgid();
    u32 *socket_fd = bpf_map_lookup_elem(&bio_new_socket_args, &pid_tgid);
    if (socket_fd == NULL) {
        return 0;
    }

    void *bio = (void *)PT_REGS_RC(ctx);
    if (bio == NULL) {
        goto cleanup;
    }
    u32 fd = *socket_fd; // copy map value into stack (required by older Kernels)
    bpf_map_update_elem(&fd_by_ssl_bio, &bio, &fd, BPF_ANY);
 cleanup:
    bpf_map_delete_elem(&bio_new_socket_args, &pid_tgid);
    return 0;
}

SEC("uprobe/SSL_set_bio")
int uprobe__SSL_set_bio(struct pt_regs* ctx) {
    void *ssl_ctx = (void *)PT_REGS_PARM1(ctx);
    void *bio = (void *)PT_REGS_PARM2(ctx);
    u32 *socket_fd = bpf_map_lookup_elem(&fd_by_ssl_bio, &bio);
    if (socket_fd == NULL)  {
        return 0;
    }
    init_ssl_sock(ssl_ctx, *socket_fd);
    bpf_map_delete_elem(&fd_by_ssl_bio, &bio);
    return 0;
}

SEC("uprobe/SSL_read")
int uprobe__SSL_read(struct pt_regs* ctx) {
    ssl_read_args_t args = {0};
    args.ctx = (void *)PT_REGS_PARM1(ctx);
    args.buf = (void *)PT_REGS_PARM2(ctx);
    u64 pid_tgid = bpf_get_current_pid_tgid();
    bpf_map_update_elem(&ssl_read_args, &pid_tgid, &args, BPF_ANY);
    return 0;
}

SEC("uretprobe/SSL_read")
int uretprobe__SSL_read(struct pt_regs* ctx) {
    u64 pid_tgid = bpf_get_current_pid_tgid();
    ssl_read_args_t *args = bpf_map_lookup_elem(&ssl_read_args, &pid_tgid);
    if (args == NULL) {
        return 0;
    }

    void *ssl_ctx = args->ctx;
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_ctx, pid_tgid);
    if (t == NULL) {
        goto cleanup;
    }

    u32 len = (u32)PT_REGS_RC(ctx);
    https_process(t, args->buf, len, LIBSSL);
 cleanup:
    bpf_map_delete_elem(&ssl_read_args, &pid_tgid);
    return 0;
}

SEC("uprobe/SSL_write")
int uprobe__SSL_write(struct pt_regs* ctx) {
    void *ssl_ctx = (void *)PT_REGS_PARM1(ctx);
    u64 pid_tgid = bpf_get_current_pid_tgid();
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_ctx, pid_tgid);
    if (t == NULL) {
        return 0;
    }

    void *ssl_buffer = (void *)PT_REGS_PARM2(ctx);
    size_t len = (size_t)PT_REGS_PARM3(ctx);
    https_process(t, ssl_buffer, len, LIBSSL);
    return 0;
}

SEC("uprobe/SSL_shutdown")
int uprobe__SSL_shutdown(struct pt_regs* ctx) {
    void *ssl_ctx = (void *)PT_REGS_PARM1(ctx);
    u64 pid_tgid = bpf_get_current_pid_tgid();
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_ctx, pid_tgid);
    if (t == NULL) {
        return 0;
    }

    https_finish(t);
    return 0;
}

// void gnutls_transport_set_int (gnutls_session_t session, int fd)
// Note: this function is implemented as a macro in gnutls
// that calls gnutls_transport_set_int2, so no uprobe is needed

// void gnutls_transport_set_int2 (gnutls_session_t session, int recv_fd, int send_fd)
SEC("uprobe/gnutls_transport_set_int2")
int uprobe__gnutls_transport_set_int2(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);
    // Use the recv_fd and ignore the send_fd;
    // in most real-world scenarios, they are the same.
    int recv_fd = (int)PT_REGS_PARM2(ctx);

    init_ssl_sock(ssl_session, (u32)recv_fd);
    return 0;
}

// void gnutls_transport_set_ptr (gnutls_session_t session, gnutls_transport_ptr_t ptr)
// "In berkeley style sockets this function will set the connection descriptor."
SEC("uprobe/gnutls_transport_set_ptr")
int uprobe__gnutls_transport_set_ptr(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);
    // This is a void*, but it might contain the socket fd cast as a pointer.
    int fd = (int)PT_REGS_PARM2(ctx);

    init_ssl_sock(ssl_session, (u32)fd);
    return 0;
}

// void gnutls_transport_set_ptr2 (gnutls_session_t session, gnutls_transport_ptr_t recv_ptr, gnutls_transport_ptr_t send_ptr)
// "In berkeley style sockets this function will set the connection descriptor."
SEC("uprobe/gnutls_transport_set_ptr2")
int uprobe__gnutls_transport_set_ptr2(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);
    // Use the recv_ptr and ignore the send_ptr;
    // in most real-world scenarios, they are the same.
    // This is a void*, but it might contain the socket fd cast as a pointer.
    int recv_fd = (int)PT_REGS_PARM2(ctx);

    init_ssl_sock(ssl_session, (u32)recv_fd);
    return 0;
}

// ssize_t gnutls_record_recv (gnutls_session_t session, void * data, size_t data_size)
SEC("uprobe/gnutls_record_recv")
int uprobe__gnutls_record_recv(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);
    void *data = (void *)PT_REGS_PARM2(ctx);

    // Re-use the map for SSL_read
    ssl_read_args_t args = {
        .ctx = ssl_session,
        .buf = data,
    };
    u64 pid_tgid = bpf_get_current_pid_tgid();
    bpf_map_update_elem(&ssl_read_args, &pid_tgid, &args, BPF_ANY);
    return 0;
}

// ssize_t gnutls_record_recv (gnutls_session_t session, void * data, size_t data_size)
SEC("uretprobe/gnutls_record_recv")
int uretprobe__gnutls_record_recv(struct pt_regs* ctx) {
    ssize_t read_len = (ssize_t)PT_REGS_RC(ctx);

    u64 pid_tgid = bpf_get_current_pid_tgid();
    // Re-use the map for SSL_read
    ssl_read_args_t *args = bpf_map_lookup_elem(&ssl_read_args, &pid_tgid);
    if (args == NULL) {
        return 0;
    }

    void *ssl_session = args->ctx;
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_session, pid_tgid);
    if (t == NULL) {
        goto cleanup;
    }

    https_process(t, args->buf, read_len, LIBGNUTLS);
 cleanup:
    bpf_map_delete_elem(&ssl_read_args, &pid_tgid);
    return 0;
}

// ssize_t gnutls_record_send (gnutls_session_t session, const void * data, size_t data_size)
SEC("uprobe/gnutls_record_send")
int uprobe__gnutls_record_send(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);
    void *data = (void *)PT_REGS_PARM2(ctx);
    size_t data_size = (size_t)PT_REGS_PARM3(ctx);

    u64 pid_tgid = bpf_get_current_pid_tgid();
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_session, pid_tgid);
    if (t == NULL) {
        return 0;
    }

    https_process(t, data, data_size, LIBGNUTLS);
    return 0;
}

// int gnutls_bye (gnutls_session_t session, gnutls_close_request_t how)
SEC("uprobe/gnutls_bye")
int uprobe__gnutls_bye(struct pt_regs* ctx) {
    void *ssl_session = (void *)PT_REGS_PARM1(ctx);

    u64 pid_tgid = bpf_get_current_pid_tgid();
    conn_tuple_t *t = tup_from_ssl_ctx(ssl_session, pid_tgid);
    if (t == NULL) {
        return 0;
    }

    https_finish(t);
    bpf_map_delete_elem(&ssl_sock_by_ctx, &ssl_session);
    return 0;
}

SEC("kprobe/do_sys_open")
int kprobe__do_sys_open(struct pt_regs* ctx) {
    char *path_argument = (char *)PT_REGS_PARM2(ctx);
    lib_path_t path = {0};
    bpf_probe_read(path.buf, sizeof(path.buf), path_argument);

    // Find the null character and clean up the garbage following it
#pragma unroll
    for (int i = 0; i < LIB_PATH_MAX_SIZE; i++) {
        if (path.len) {
            path.buf[i] = 0;
        } else if (path.buf[i] == 0) {
            path.len = i;
        }
    }

    // Bail out if the path size is larger than our buffer
    if (!path.len) {
        return 0;
    }

    u64 pid_tgid = bpf_get_current_pid_tgid();
    path.pid = pid_tgid >> 32;
    bpf_map_update_elem(&open_at_args, &pid_tgid, &path, BPF_ANY);
    return 0;
}

SEC("kretprobe/do_sys_open")
int kretprobe__do_sys_open(struct pt_regs* ctx) {
    u64 pid_tgid = bpf_get_current_pid_tgid();

    // If file couldn't be opened, bail out
    if (!(long)PT_REGS_RC(ctx)) {
        goto cleanup;
    }

    lib_path_t *path = bpf_map_lookup_elem(&open_at_args, &pid_tgid);
    if (path == NULL) {
        return 0;
    }

    // Detect whether the file being opened is a shared library
    bool is_shared_library = false;
#pragma unroll
    for (int i = 0; i < LIB_PATH_MAX_SIZE - SO_SUFFIX_SIZE; i++) {
        if (path->buf[i] == '.' && path->buf[i+1] == 's' && path->buf[i+2] == 'o') {
            is_shared_library = true;
            break;
        }
    }

    if (!is_shared_library) {
        goto cleanup;
    }

    // Copy map value into eBPF stack
    lib_path_t lib_path;
    __builtin_memcpy(&lib_path, path, sizeof(lib_path));

    u32 cpu = bpf_get_smp_processor_id();
    bpf_perf_event_output(ctx, &shared_libraries, cpu, &lib_path, sizeof(lib_path));
 cleanup:
    bpf_map_delete_elem(&open_at_args, &pid_tgid);
    return 0;
}

// This number will be interpreted by elf-loader to set the current running kernel version
__u32 _version SEC("version") = 0xFFFFFFFE; // NOLINT(bugprone-reserved-identifier)

char _license[] SEC("license") = "GPL"; // NOLINT(bugprone-reserved-identifier)
