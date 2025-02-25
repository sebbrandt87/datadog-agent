// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datadog/remoteconfig/remoteconfig.proto

package pbgo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigMetas struct {
	Roots                []*TopMeta       `protobuf:"bytes,1,rep,name=roots,proto3" json:"roots,omitempty"`
	Timestamp            *TopMeta         `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Snapshot             *TopMeta         `protobuf:"bytes,3,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	TopTargets           *TopMeta         `protobuf:"bytes,4,opt,name=topTargets,proto3" json:"topTargets,omitempty"`
	DelegatedTargets     []*DelegatedMeta `protobuf:"bytes,5,rep,name=delegatedTargets,proto3" json:"delegatedTargets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConfigMetas) Reset()         { *m = ConfigMetas{} }
func (m *ConfigMetas) String() string { return proto.CompactTextString(m) }
func (*ConfigMetas) ProtoMessage()    {}
func (*ConfigMetas) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{0}
}

func (m *ConfigMetas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigMetas.Unmarshal(m, b)
}
func (m *ConfigMetas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigMetas.Marshal(b, m, deterministic)
}
func (m *ConfigMetas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigMetas.Merge(m, src)
}
func (m *ConfigMetas) XXX_Size() int {
	return xxx_messageInfo_ConfigMetas.Size(m)
}
func (m *ConfigMetas) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigMetas.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigMetas proto.InternalMessageInfo

func (m *ConfigMetas) GetRoots() []*TopMeta {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *ConfigMetas) GetTimestamp() *TopMeta {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ConfigMetas) GetSnapshot() *TopMeta {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *ConfigMetas) GetTopTargets() *TopMeta {
	if m != nil {
		return m.TopTargets
	}
	return nil
}

func (m *ConfigMetas) GetDelegatedTargets() []*DelegatedMeta {
	if m != nil {
		return m.DelegatedTargets
	}
	return nil
}

type DirectorMetas struct {
	Roots                []*TopMeta `protobuf:"bytes,1,rep,name=roots,proto3" json:"roots,omitempty"`
	Timestamp            *TopMeta   `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Snapshot             *TopMeta   `protobuf:"bytes,3,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Targets              *TopMeta   `protobuf:"bytes,4,opt,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *DirectorMetas) Reset()         { *m = DirectorMetas{} }
func (m *DirectorMetas) String() string { return proto.CompactTextString(m) }
func (*DirectorMetas) ProtoMessage()    {}
func (*DirectorMetas) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{1}
}

func (m *DirectorMetas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DirectorMetas.Unmarshal(m, b)
}
func (m *DirectorMetas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DirectorMetas.Marshal(b, m, deterministic)
}
func (m *DirectorMetas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DirectorMetas.Merge(m, src)
}
func (m *DirectorMetas) XXX_Size() int {
	return xxx_messageInfo_DirectorMetas.Size(m)
}
func (m *DirectorMetas) XXX_DiscardUnknown() {
	xxx_messageInfo_DirectorMetas.DiscardUnknown(m)
}

var xxx_messageInfo_DirectorMetas proto.InternalMessageInfo

func (m *DirectorMetas) GetRoots() []*TopMeta {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *DirectorMetas) GetTimestamp() *TopMeta {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DirectorMetas) GetSnapshot() *TopMeta {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

func (m *DirectorMetas) GetTargets() *TopMeta {
	if m != nil {
		return m.Targets
	}
	return nil
}

type DelegatedMeta struct {
	Version              uint64   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Role                 string   `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Raw                  []byte   `protobuf:"bytes,3,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DelegatedMeta) Reset()         { *m = DelegatedMeta{} }
func (m *DelegatedMeta) String() string { return proto.CompactTextString(m) }
func (*DelegatedMeta) ProtoMessage()    {}
func (*DelegatedMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{2}
}

func (m *DelegatedMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DelegatedMeta.Unmarshal(m, b)
}
func (m *DelegatedMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DelegatedMeta.Marshal(b, m, deterministic)
}
func (m *DelegatedMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatedMeta.Merge(m, src)
}
func (m *DelegatedMeta) XXX_Size() int {
	return xxx_messageInfo_DelegatedMeta.Size(m)
}
func (m *DelegatedMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatedMeta.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatedMeta proto.InternalMessageInfo

func (m *DelegatedMeta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *DelegatedMeta) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *DelegatedMeta) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

type TopMeta struct {
	Version              uint64   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Raw                  []byte   `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopMeta) Reset()         { *m = TopMeta{} }
func (m *TopMeta) String() string { return proto.CompactTextString(m) }
func (*TopMeta) ProtoMessage()    {}
func (*TopMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{3}
}

func (m *TopMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopMeta.Unmarshal(m, b)
}
func (m *TopMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopMeta.Marshal(b, m, deterministic)
}
func (m *TopMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopMeta.Merge(m, src)
}
func (m *TopMeta) XXX_Size() int {
	return xxx_messageInfo_TopMeta.Size(m)
}
func (m *TopMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TopMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TopMeta proto.InternalMessageInfo

func (m *TopMeta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TopMeta) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

type File struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Raw                  []byte   `protobuf:"bytes,2,opt,name=raw,proto3" json:"raw,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{4}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *File) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

type LatestConfigsRequest struct {
	Hostname     string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	AgentVersion string `protobuf:"bytes,2,opt,name=agentVersion,proto3" json:"agentVersion,omitempty"`
	// timestamp and snapshot versions move in tandem so they are the same.
	CurrentConfigSnapshotVersion uint64    `protobuf:"varint,3,opt,name=current_config_snapshot_version,json=currentConfigSnapshotVersion,proto3" json:"current_config_snapshot_version,omitempty"`
	CurrentConfigRootVersion     uint64    `protobuf:"varint,9,opt,name=current_config_root_version,json=currentConfigRootVersion,proto3" json:"current_config_root_version,omitempty"`
	CurrentDirectorRootVersion   uint64    `protobuf:"varint,8,opt,name=current_director_root_version,json=currentDirectorRootVersion,proto3" json:"current_director_root_version,omitempty"`
	Products                     []string  `protobuf:"bytes,4,rep,name=products,proto3" json:"products,omitempty"`
	NewProducts                  []string  `protobuf:"bytes,5,rep,name=new_products,json=newProducts,proto3" json:"new_products,omitempty"`
	ActiveClients                []*Client `protobuf:"bytes,6,rep,name=active_clients,json=activeClients,proto3" json:"active_clients,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}  `json:"-"`
	XXX_unrecognized             []byte    `json:"-"`
	XXX_sizecache                int32     `json:"-"`
}

func (m *LatestConfigsRequest) Reset()         { *m = LatestConfigsRequest{} }
func (m *LatestConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*LatestConfigsRequest) ProtoMessage()    {}
func (*LatestConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{5}
}

func (m *LatestConfigsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestConfigsRequest.Unmarshal(m, b)
}
func (m *LatestConfigsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestConfigsRequest.Marshal(b, m, deterministic)
}
func (m *LatestConfigsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestConfigsRequest.Merge(m, src)
}
func (m *LatestConfigsRequest) XXX_Size() int {
	return xxx_messageInfo_LatestConfigsRequest.Size(m)
}
func (m *LatestConfigsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestConfigsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LatestConfigsRequest proto.InternalMessageInfo

func (m *LatestConfigsRequest) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *LatestConfigsRequest) GetAgentVersion() string {
	if m != nil {
		return m.AgentVersion
	}
	return ""
}

func (m *LatestConfigsRequest) GetCurrentConfigSnapshotVersion() uint64 {
	if m != nil {
		return m.CurrentConfigSnapshotVersion
	}
	return 0
}

func (m *LatestConfigsRequest) GetCurrentConfigRootVersion() uint64 {
	if m != nil {
		return m.CurrentConfigRootVersion
	}
	return 0
}

func (m *LatestConfigsRequest) GetCurrentDirectorRootVersion() uint64 {
	if m != nil {
		return m.CurrentDirectorRootVersion
	}
	return 0
}

func (m *LatestConfigsRequest) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *LatestConfigsRequest) GetNewProducts() []string {
	if m != nil {
		return m.NewProducts
	}
	return nil
}

func (m *LatestConfigsRequest) GetActiveClients() []*Client {
	if m != nil {
		return m.ActiveClients
	}
	return nil
}

type LatestConfigsResponse struct {
	ConfigMetas          *ConfigMetas   `protobuf:"bytes,1,opt,name=config_metas,json=configMetas,proto3" json:"config_metas,omitempty"`
	DirectorMetas        *DirectorMetas `protobuf:"bytes,2,opt,name=director_metas,json=directorMetas,proto3" json:"director_metas,omitempty"`
	TargetFiles          []*File        `protobuf:"bytes,3,rep,name=target_files,json=targetFiles,proto3" json:"target_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LatestConfigsResponse) Reset()         { *m = LatestConfigsResponse{} }
func (m *LatestConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*LatestConfigsResponse) ProtoMessage()    {}
func (*LatestConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{6}
}

func (m *LatestConfigsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LatestConfigsResponse.Unmarshal(m, b)
}
func (m *LatestConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LatestConfigsResponse.Marshal(b, m, deterministic)
}
func (m *LatestConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LatestConfigsResponse.Merge(m, src)
}
func (m *LatestConfigsResponse) XXX_Size() int {
	return xxx_messageInfo_LatestConfigsResponse.Size(m)
}
func (m *LatestConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LatestConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LatestConfigsResponse proto.InternalMessageInfo

func (m *LatestConfigsResponse) GetConfigMetas() *ConfigMetas {
	if m != nil {
		return m.ConfigMetas
	}
	return nil
}

func (m *LatestConfigsResponse) GetDirectorMetas() *DirectorMetas {
	if m != nil {
		return m.DirectorMetas
	}
	return nil
}

func (m *LatestConfigsResponse) GetTargetFiles() []*File {
	if m != nil {
		return m.TargetFiles
	}
	return nil
}

type Client struct {
	State                *ClientState  `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Id                   string        `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Products             []string      `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	Name                 string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Version              string        `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	IsTracer             bool          `protobuf:"varint,6,opt,name=is_tracer,json=isTracer,proto3" json:"is_tracer,omitempty"`
	TracerClient         *TracerClient `protobuf:"bytes,7,opt,name=tracer_client,json=tracerClient,proto3" json:"tracer_client,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Client) Reset()         { *m = Client{} }
func (m *Client) String() string { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()    {}
func (*Client) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{7}
}

func (m *Client) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Client.Unmarshal(m, b)
}
func (m *Client) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Client.Marshal(b, m, deterministic)
}
func (m *Client) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Client.Merge(m, src)
}
func (m *Client) XXX_Size() int {
	return xxx_messageInfo_Client.Size(m)
}
func (m *Client) XXX_DiscardUnknown() {
	xxx_messageInfo_Client.DiscardUnknown(m)
}

var xxx_messageInfo_Client proto.InternalMessageInfo

func (m *Client) GetState() *ClientState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Client) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Client) GetProducts() []string {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *Client) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Client) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Client) GetIsTracer() bool {
	if m != nil {
		return m.IsTracer
	}
	return false
}

func (m *Client) GetTracerClient() *TracerClient {
	if m != nil {
		return m.TracerClient
	}
	return nil
}

type TracerClient struct {
	RuntimeId            string   `protobuf:"bytes,1,opt,name=runtime_id,json=runtimeId,proto3" json:"runtime_id,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	TracerVersion        string   `protobuf:"bytes,3,opt,name=tracer_version,json=tracerVersion,proto3" json:"tracer_version,omitempty"`
	Service              string   `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	Env                  string   `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	AppVersion           string   `protobuf:"bytes,6,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TracerClient) Reset()         { *m = TracerClient{} }
func (m *TracerClient) String() string { return proto.CompactTextString(m) }
func (*TracerClient) ProtoMessage()    {}
func (*TracerClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{8}
}

func (m *TracerClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TracerClient.Unmarshal(m, b)
}
func (m *TracerClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TracerClient.Marshal(b, m, deterministic)
}
func (m *TracerClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TracerClient.Merge(m, src)
}
func (m *TracerClient) XXX_Size() int {
	return xxx_messageInfo_TracerClient.Size(m)
}
func (m *TracerClient) XXX_DiscardUnknown() {
	xxx_messageInfo_TracerClient.DiscardUnknown(m)
}

var xxx_messageInfo_TracerClient proto.InternalMessageInfo

func (m *TracerClient) GetRuntimeId() string {
	if m != nil {
		return m.RuntimeId
	}
	return ""
}

func (m *TracerClient) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *TracerClient) GetTracerVersion() string {
	if m != nil {
		return m.TracerVersion
	}
	return ""
}

func (m *TracerClient) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *TracerClient) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *TracerClient) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

type Config struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version              uint64   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{9}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Config) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type ClientState struct {
	RootVersion          uint64    `protobuf:"varint,1,opt,name=root_version,json=rootVersion,proto3" json:"root_version,omitempty"`
	TargetsVersion       uint64    `protobuf:"varint,2,opt,name=targets_version,json=targetsVersion,proto3" json:"targets_version,omitempty"`
	Configs              []*Config `protobuf:"bytes,3,rep,name=configs,proto3" json:"configs,omitempty"`
	HasError             bool      `protobuf:"varint,4,opt,name=has_error,json=hasError,proto3" json:"has_error,omitempty"`
	Error                string    `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{10}
}

func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientState.Unmarshal(m, b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return xxx_messageInfo_ClientState.Size(m)
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

func (m *ClientState) GetRootVersion() uint64 {
	if m != nil {
		return m.RootVersion
	}
	return 0
}

func (m *ClientState) GetTargetsVersion() uint64 {
	if m != nil {
		return m.TargetsVersion
	}
	return 0
}

func (m *ClientState) GetConfigs() []*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *ClientState) GetHasError() bool {
	if m != nil {
		return m.HasError
	}
	return false
}

func (m *ClientState) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type TargetFileHash struct {
	Algorithm            string   `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TargetFileHash) Reset()         { *m = TargetFileHash{} }
func (m *TargetFileHash) String() string { return proto.CompactTextString(m) }
func (*TargetFileHash) ProtoMessage()    {}
func (*TargetFileHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{11}
}

func (m *TargetFileHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetFileHash.Unmarshal(m, b)
}
func (m *TargetFileHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetFileHash.Marshal(b, m, deterministic)
}
func (m *TargetFileHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetFileHash.Merge(m, src)
}
func (m *TargetFileHash) XXX_Size() int {
	return xxx_messageInfo_TargetFileHash.Size(m)
}
func (m *TargetFileHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetFileHash.DiscardUnknown(m)
}

var xxx_messageInfo_TargetFileHash proto.InternalMessageInfo

func (m *TargetFileHash) GetAlgorithm() string {
	if m != nil {
		return m.Algorithm
	}
	return ""
}

func (m *TargetFileHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type TargetFileMeta struct {
	Path                 string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Length               int64             `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
	Hashes               []*TargetFileHash `protobuf:"bytes,3,rep,name=hashes,proto3" json:"hashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TargetFileMeta) Reset()         { *m = TargetFileMeta{} }
func (m *TargetFileMeta) String() string { return proto.CompactTextString(m) }
func (*TargetFileMeta) ProtoMessage()    {}
func (*TargetFileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{12}
}

func (m *TargetFileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetFileMeta.Unmarshal(m, b)
}
func (m *TargetFileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetFileMeta.Marshal(b, m, deterministic)
}
func (m *TargetFileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetFileMeta.Merge(m, src)
}
func (m *TargetFileMeta) XXX_Size() int {
	return xxx_messageInfo_TargetFileMeta.Size(m)
}
func (m *TargetFileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetFileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_TargetFileMeta proto.InternalMessageInfo

func (m *TargetFileMeta) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *TargetFileMeta) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *TargetFileMeta) GetHashes() []*TargetFileHash {
	if m != nil {
		return m.Hashes
	}
	return nil
}

type ClientGetConfigsRequest struct {
	Client               *Client           `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	CachedTargetFiles    []*TargetFileMeta `protobuf:"bytes,2,rep,name=cached_target_files,json=cachedTargetFiles,proto3" json:"cached_target_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClientGetConfigsRequest) Reset()         { *m = ClientGetConfigsRequest{} }
func (m *ClientGetConfigsRequest) String() string { return proto.CompactTextString(m) }
func (*ClientGetConfigsRequest) ProtoMessage()    {}
func (*ClientGetConfigsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{13}
}

func (m *ClientGetConfigsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientGetConfigsRequest.Unmarshal(m, b)
}
func (m *ClientGetConfigsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientGetConfigsRequest.Marshal(b, m, deterministic)
}
func (m *ClientGetConfigsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientGetConfigsRequest.Merge(m, src)
}
func (m *ClientGetConfigsRequest) XXX_Size() int {
	return xxx_messageInfo_ClientGetConfigsRequest.Size(m)
}
func (m *ClientGetConfigsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientGetConfigsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClientGetConfigsRequest proto.InternalMessageInfo

func (m *ClientGetConfigsRequest) GetClient() *Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *ClientGetConfigsRequest) GetCachedTargetFiles() []*TargetFileMeta {
	if m != nil {
		return m.CachedTargetFiles
	}
	return nil
}

type ClientGetConfigsResponse struct {
	Roots                []*TopMeta `protobuf:"bytes,1,rep,name=roots,proto3" json:"roots,omitempty"`
	Targets              *TopMeta   `protobuf:"bytes,2,opt,name=targets,proto3" json:"targets,omitempty"`
	TargetFiles          []*File    `protobuf:"bytes,3,rep,name=target_files,json=targetFiles,proto3" json:"target_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ClientGetConfigsResponse) Reset()         { *m = ClientGetConfigsResponse{} }
func (m *ClientGetConfigsResponse) String() string { return proto.CompactTextString(m) }
func (*ClientGetConfigsResponse) ProtoMessage()    {}
func (*ClientGetConfigsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b7d6c4df67ae5e, []int{14}
}

func (m *ClientGetConfigsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientGetConfigsResponse.Unmarshal(m, b)
}
func (m *ClientGetConfigsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientGetConfigsResponse.Marshal(b, m, deterministic)
}
func (m *ClientGetConfigsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientGetConfigsResponse.Merge(m, src)
}
func (m *ClientGetConfigsResponse) XXX_Size() int {
	return xxx_messageInfo_ClientGetConfigsResponse.Size(m)
}
func (m *ClientGetConfigsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientGetConfigsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClientGetConfigsResponse proto.InternalMessageInfo

func (m *ClientGetConfigsResponse) GetRoots() []*TopMeta {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *ClientGetConfigsResponse) GetTargets() *TopMeta {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *ClientGetConfigsResponse) GetTargetFiles() []*File {
	if m != nil {
		return m.TargetFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigMetas)(nil), "datadog.config.ConfigMetas")
	proto.RegisterType((*DirectorMetas)(nil), "datadog.config.DirectorMetas")
	proto.RegisterType((*DelegatedMeta)(nil), "datadog.config.DelegatedMeta")
	proto.RegisterType((*TopMeta)(nil), "datadog.config.TopMeta")
	proto.RegisterType((*File)(nil), "datadog.config.File")
	proto.RegisterType((*LatestConfigsRequest)(nil), "datadog.config.LatestConfigsRequest")
	proto.RegisterType((*LatestConfigsResponse)(nil), "datadog.config.LatestConfigsResponse")
	proto.RegisterType((*Client)(nil), "datadog.config.Client")
	proto.RegisterType((*TracerClient)(nil), "datadog.config.TracerClient")
	proto.RegisterType((*Config)(nil), "datadog.config.Config")
	proto.RegisterType((*ClientState)(nil), "datadog.config.ClientState")
	proto.RegisterType((*TargetFileHash)(nil), "datadog.config.TargetFileHash")
	proto.RegisterType((*TargetFileMeta)(nil), "datadog.config.TargetFileMeta")
	proto.RegisterType((*ClientGetConfigsRequest)(nil), "datadog.config.ClientGetConfigsRequest")
	proto.RegisterType((*ClientGetConfigsResponse)(nil), "datadog.config.ClientGetConfigsResponse")
}

func init() {
	proto.RegisterFile("datadog/remoteconfig/remoteconfig.proto", fileDescriptor_68b7d6c4df67ae5e)
}

var fileDescriptor_68b7d6c4df67ae5e = []byte{
	// 958 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x92, 0xdb, 0x44,
	0x14, 0x2d, 0x59, 0xb6, 0x6c, 0x5f, 0xd9, 0x66, 0x68, 0x86, 0x89, 0x6a, 0x1e, 0x64, 0x50, 0x15,
	0x95, 0x59, 0x80, 0x87, 0x4c, 0x2a, 0x64, 0x15, 0xaa, 0x92, 0x4c, 0x80, 0x54, 0xf1, 0xaa, 0x8e,
	0x8b, 0x05, 0x1b, 0x55, 0x47, 0xea, 0x48, 0x2a, 0x6c, 0xb5, 0xe8, 0x6e, 0xcf, 0x7c, 0x0b, 0x0b,
	0xfe, 0x80, 0x2d, 0x4b, 0xf8, 0x0c, 0x76, 0x7c, 0x06, 0x7b, 0xaa, 0x5f, 0x1a, 0xc9, 0x9e, 0x0c,
	0x03, 0x2b, 0x76, 0x7d, 0xbb, 0xcf, 0x3d, 0xee, 0xbe, 0xf7, 0x9c, 0x2b, 0xc3, 0xbd, 0x8c, 0x48,
	0x92, 0xb1, 0xfc, 0x94, 0xd3, 0x15, 0x93, 0x34, 0x65, 0xd5, 0xeb, 0xb2, 0x1b, 0xcc, 0x6b, 0xce,
	0x24, 0x43, 0x33, 0x0b, 0x9c, 0x9b, 0xdd, 0xf8, 0x97, 0x1e, 0x84, 0xcf, 0xf4, 0xf2, 0x2b, 0x2a,
	0x89, 0x40, 0x1f, 0xc1, 0x80, 0x33, 0x26, 0x45, 0xe4, 0x1d, 0xfb, 0x27, 0xe1, 0xd9, 0x9d, 0x79,
	0x17, 0x3f, 0x5f, 0xb0, 0x5a, 0x01, 0xb1, 0x41, 0xa1, 0x87, 0x30, 0x96, 0xe5, 0x8a, 0x0a, 0x49,
	0x56, 0x75, 0xd4, 0x3b, 0xf6, 0x6e, 0x4a, 0xb9, 0x42, 0xa2, 0x07, 0x30, 0x12, 0x15, 0xa9, 0x45,
	0xc1, 0x64, 0xe4, 0xdf, 0x9c, 0xd5, 0x00, 0xd1, 0x23, 0x00, 0xc9, 0xea, 0x05, 0xe1, 0x39, 0x95,
	0x22, 0xea, 0xdf, 0x9c, 0xd6, 0x82, 0xa2, 0x17, 0xb0, 0x93, 0xd1, 0x25, 0xcd, 0x89, 0xa4, 0x99,
	0x4b, 0x1f, 0xe8, 0xe7, 0x1d, 0x6d, 0xa6, 0x9f, 0x3b, 0x9c, 0x26, 0xd9, 0x4a, 0x8b, 0xff, 0xf4,
	0x60, 0x7a, 0x5e, 0x72, 0x9a, 0x4a, 0xc6, 0xff, 0xf7, 0x05, 0xbb, 0x0f, 0x43, 0x79, 0xbb, 0x6a,
	0x39, 0x5c, 0xfc, 0x0d, 0x4c, 0x3b, 0x25, 0x40, 0x11, 0x0c, 0x2f, 0x28, 0x17, 0x25, 0xab, 0x22,
	0xef, 0xd8, 0x3b, 0xe9, 0x63, 0x17, 0x22, 0x04, 0x7d, 0xce, 0x96, 0x54, 0x3f, 0x62, 0x8c, 0xf5,
	0x1a, 0xed, 0x80, 0xcf, 0xc9, 0xa5, 0xbe, 0xe1, 0x04, 0xab, 0x65, 0xfc, 0x10, 0x86, 0xf6, 0x47,
	0x6e, 0xa0, 0xb2, 0x69, 0xbd, 0xab, 0xb4, 0x0f, 0xa1, 0xff, 0x59, 0xb9, 0xa4, 0xea, 0x47, 0x6a,
	0x22, 0x0b, 0x9d, 0x30, 0xc6, 0x7a, 0x7d, 0x0d, 0xfa, 0x67, 0x1f, 0x76, 0xbf, 0x24, 0x92, 0x0a,
	0x69, 0xa4, 0x2c, 0x30, 0xfd, 0x71, 0x4d, 0x85, 0x44, 0xfb, 0x30, 0x2a, 0x98, 0x90, 0x15, 0x59,
	0x51, 0x4b, 0xd1, 0xc4, 0x28, 0x86, 0x09, 0xc9, 0x69, 0x25, 0xbf, 0xb3, 0x77, 0x32, 0xef, 0xe8,
	0xec, 0xa1, 0xe7, 0x70, 0x37, 0x5d, 0x73, 0x4e, 0x2b, 0x99, 0x98, 0x8a, 0x25, 0xae, 0xb8, 0x89,
	0x7b, 0x8a, 0xaf, 0x9f, 0x72, 0x68, 0x61, 0xe6, 0xf7, 0x5f, 0x5a, 0x90, 0xa3, 0x79, 0x0c, 0x07,
	0x1b, 0x34, 0x4a, 0x0c, 0x0d, 0xc5, 0x58, 0x53, 0x44, 0x1d, 0x0a, 0xcc, 0xae, 0xd2, 0x9f, 0xc0,
	0x91, 0x4b, 0xcf, 0xac, 0xf6, 0xba, 0x04, 0x23, 0x4d, 0xb0, 0x6f, 0x41, 0x4e, 0x9f, 0x6d, 0x8a,
	0x7d, 0x18, 0xd5, 0x9c, 0x65, 0xeb, 0x54, 0x6b, 0xc1, 0x57, 0x85, 0x70, 0x31, 0x7a, 0x1f, 0x26,
	0x15, 0xbd, 0x4c, 0x9a, 0xf3, 0x81, 0x3e, 0x0f, 0x2b, 0x7a, 0xf9, 0xad, 0x83, 0x3c, 0x86, 0x19,
	0x49, 0x65, 0x79, 0x41, 0x93, 0x74, 0x59, 0xd2, 0x4a, 0x8a, 0x28, 0xd0, 0x6a, 0xdf, 0xdb, 0x14,
	0xd4, 0x33, 0x7d, 0x8c, 0xa7, 0x06, 0x6d, 0x22, 0x11, 0xff, 0xe1, 0xc1, 0xbb, 0x1b, 0xfd, 0x11,
	0x35, 0xab, 0x04, 0x45, 0x9f, 0xc2, 0xc4, 0x56, 0x64, 0xa5, 0xdc, 0xa4, 0x9b, 0x14, 0x9e, 0x1d,
	0x6c, 0xd1, 0x5e, 0x4d, 0x28, 0x1c, 0xa6, 0xad, 0x71, 0x75, 0x0e, 0xb3, 0xa6, 0x24, 0x86, 0xc1,
	0x78, 0x6a, 0xdb, 0xd8, 0x6d, 0xd3, 0xe2, 0x69, 0xd6, 0xf1, 0xf0, 0x23, 0x98, 0x18, 0x03, 0x24,
	0xaf, 0xcb, 0x25, 0x15, 0x91, 0xaf, 0x1f, 0xb7, 0xbb, 0xc9, 0xa1, 0x14, 0x89, 0x43, 0x83, 0x54,
	0x6b, 0x11, 0xff, 0xe5, 0x41, 0x60, 0x1e, 0x89, 0xee, 0xc3, 0x40, 0x48, 0x22, 0xe9, 0x1b, 0x9f,
	0xa0, 0x61, 0x2f, 0x15, 0x04, 0x1b, 0x24, 0x9a, 0x41, 0xaf, 0xcc, 0xac, 0xee, 0x7a, 0x65, 0xd6,
	0x69, 0x92, 0xbf, 0xd1, 0x24, 0x04, 0x7d, 0xad, 0xe2, 0xbe, 0x31, 0x82, 0x56, 0x70, 0xcb, 0x50,
	0x03, 0xbd, 0xdd, 0x18, 0xea, 0x00, 0xc6, 0xa5, 0x48, 0x24, 0x27, 0x29, 0xe5, 0x51, 0x70, 0xec,
	0x9d, 0x8c, 0xf0, 0xa8, 0x14, 0x0b, 0x1d, 0xa3, 0x27, 0x30, 0x35, 0x27, 0xb6, 0x99, 0xd1, 0x50,
	0xdf, 0xf8, 0x70, 0x6b, 0x38, 0x68, 0x90, 0xed, 0xe8, 0x44, 0xb6, 0xa2, 0xf8, 0x37, 0x0f, 0x26,
	0xed, 0x63, 0x74, 0x04, 0xc0, 0xd7, 0x95, 0x9a, 0x57, 0x49, 0x99, 0x59, 0xab, 0x8d, 0xed, 0xce,
	0x0b, 0xfd, 0xb2, 0x25, 0xa9, 0xf2, 0x35, 0xc9, 0xdd, 0xbc, 0x68, 0x62, 0xf4, 0x01, 0xcc, 0xec,
	0x75, 0xda, 0x96, 0x1a, 0x63, 0x7b, 0x49, 0xa7, 0xe0, 0x08, 0x86, 0x82, 0xf2, 0x8b, 0x32, 0x75,
	0x35, 0x70, 0xa1, 0x9a, 0x07, 0xb4, 0xba, 0xb0, 0x25, 0x50, 0x4b, 0x74, 0x17, 0x42, 0x52, 0xd7,
	0x0d, 0x5f, 0xa0, 0x4f, 0x80, 0xd4, 0xb5, 0x25, 0x8b, 0xcf, 0x20, 0x30, 0x92, 0xb2, 0x3d, 0xf0,
	0x9a, 0x1e, 0xb4, 0x6a, 0xda, 0xeb, 0x0c, 0xa9, 0xf8, 0x77, 0x0f, 0xc2, 0x56, 0x13, 0x95, 0x6d,
	0x3a, 0x26, 0x34, 0x33, 0x2d, 0xe4, 0x2d, 0xd7, 0xdd, 0x83, 0xb7, 0xec, 0x60, 0x4d, 0xba, 0xa4,
	0x33, 0xbb, 0xed, 0x80, 0x1f, 0xc3, 0xd0, 0x14, 0xdd, 0x69, 0x6f, 0xef, 0x7a, 0x07, 0x60, 0x07,
	0x53, 0x1d, 0x2e, 0x88, 0x48, 0x28, 0xe7, 0x8c, 0xeb, 0x82, 0x8c, 0xf0, 0xa8, 0x20, 0xe2, 0xb9,
	0x8a, 0xd1, 0x2e, 0x0c, 0xcc, 0x81, 0xa9, 0x89, 0x09, 0xe2, 0xa7, 0x30, 0x5b, 0x34, 0xda, 0xfd,
	0x82, 0x88, 0x02, 0x1d, 0xc2, 0x98, 0x2c, 0x73, 0xc6, 0x4b, 0x59, 0xac, 0x5c, 0xd3, 0x9a, 0x0d,
	0x25, 0xb9, 0x82, 0x88, 0xc2, 0x0e, 0x5a, 0xbd, 0x8e, 0x65, 0x9b, 0x43, 0x4f, 0xf5, 0xeb, 0x26,
	0xf4, 0x1e, 0x04, 0x4b, 0x5a, 0xe5, 0xd2, 0xe4, 0xfa, 0xd8, 0x46, 0xe8, 0x13, 0x08, 0x14, 0x4b,
	0xe3, 0xb0, 0xf7, 0xb6, 0x24, 0xd7, 0xb9, 0x1f, 0xb6, 0xe8, 0xf8, 0x27, 0x0f, 0xee, 0x98, 0xd2,
	0x7f, 0x4e, 0x37, 0x47, 0xfc, 0x1c, 0x02, 0x2b, 0x63, 0x63, 0xbc, 0x37, 0x8d, 0x24, 0x8b, 0x42,
	0x5f, 0xc3, 0x3b, 0x29, 0x49, 0x0b, 0x9a, 0x25, 0x1d, 0xcb, 0xf7, 0xfe, 0xe9, 0x42, 0xfa, 0x3b,
	0xf9, 0xb6, 0x49, 0x5d, 0xb4, 0x46, 0xc0, 0xaf, 0x1e, 0x44, 0xdb, 0x77, 0xb3, 0xe3, 0xed, 0x5f,
	0xfe, 0x39, 0x68, 0x7d, 0xb0, 0x7b, 0xb7, 0xfb, 0x60, 0xff, 0xe7, 0xd1, 0xf5, 0x74, 0xe7, 0xfb,
	0x59, 0xfd, 0x43, 0x7e, 0xaa, 0xff, 0x15, 0x9e, 0xd6, 0xaf, 0x72, 0xf6, 0x2a, 0xd0, 0xeb, 0x07,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x38, 0x16, 0x76, 0x8b, 0x4c, 0x0a, 0x00, 0x00,
}
