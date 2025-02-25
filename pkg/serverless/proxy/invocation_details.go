// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package proxy

import (
	"time"
)

// InvocationStartDetails stores information about the start of an invocation.
// This structure is passed to the onInvokeStart method of the invocationProcessor interface
type InvocationStartDetails struct {
	StartTime          time.Time
	InvokeHeaders      map[string][]string
	InvokeEventPayload string
}

// InvocationEndDetails stores information about the end of an invocation.
// This structure is passed to the onInvokeEnd method of the invocationProcessor interface
type InvocationEndDetails struct {
	EndTime time.Time
	IsError bool
}
