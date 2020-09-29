// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package api

import (
	"io"
	"net/http"
	"sync/atomic"
)

// NewDebugTransport Deprecated factory for DebugTransport.
func NewDebugTransport(transport http.RoundTripper, o io.Writer, obscure bool) *DebugTransport {
	if t, ok := transport.(*DebugTransport); ok {
		return t
	}

	if transport == nil {
		transport = http.DefaultTransport
	}

	return &DebugTransport{
		transport:  transport,
		output:     o,
		count:      -1,
		redactAuth: obscure,
	}
}

// DebugTransport Deprecated is an http.RoundTripper that keeps track of the
// in-flight request and implements hooks to report HTTP tracing events.
type DebugTransport struct {
	output    io.Writer
	transport http.RoundTripper

	count      int64
	redactAuth bool
}

// RoundTrip Deprecated wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *DebugTransport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	count := atomic.AddInt64(&t.count, 1)

	handleVerboseRequest(t.output, req, count, t.redactAuth)
	res, err = t.transport.RoundTrip(req)
	if res != nil {
		handleVerboseResponse(t.output, res, count)
	}

	return res, err
}
