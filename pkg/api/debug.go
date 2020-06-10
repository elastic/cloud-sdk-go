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
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"regexp"
	"sync/atomic"
)

// NewDebugTransport factory for DebugTransport
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

// DebugTransport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type DebugTransport struct {
	output    io.Writer
	transport http.RoundTripper

	count      int64
	redactAuth bool
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *DebugTransport) RoundTrip(rreq *http.Request) (*http.Response, error) {
	var (
		err error
		res *http.Response
	)

	count := atomic.AddInt64(&t.count, 1)
	t.handleVerboseRequest(rreq, count)
	res, err = t.transport.RoundTrip(rreq)
	if res != nil {
		t.handleVerboseResponse(res, count)
	}

	return res, err
}

func (t *DebugTransport) handleVerboseRequest(req *http.Request, count int64) {
	b, _ := httputil.DumpRequestOut(req, req.Body != nil)

	fmt.Fprintf(t.output, "==================== Start of Request #%d ====================\n", count)
	fmt.Fprintln(t.output, redactAuth(string(b), t.redactAuth))
	fmt.Fprintf(t.output, "====================  End of Request #%d  ====================\n", count)
}

func (t *DebugTransport) handleVerboseResponse(res *http.Response, count int64) {
	b, _ := httputil.DumpResponse(res, res.Body != nil)
	fmt.Fprintf(t.output, "==================== Start of Response #%d ====================\n", count)
	fmt.Fprintln(t.output, string(b))
	fmt.Fprintf(t.output, "====================  End of Response #%d  ====================\n", count)
}

func redactAuth(request string, obscure bool) string {
	if !obscure {
		return request
	}

	return regexp.MustCompile(`Authorization: .*\n`).ReplaceAllString(
		request, "Authorization: [REDACTED]\r\n",
	)
}
