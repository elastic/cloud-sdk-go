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
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestNewCustomTransport(t *testing.T) {
	type args struct {
		cfg CustomTransportCfg
	}
	tests := []struct {
		name string
		args args
		want *CustomTransport
		err  error
	}{
		{
			name: "fails creating a custom transport due to invalid config",
			args: args{cfg: CustomTransportCfg{
				Verbose: true,
			}},
			err: multierror.NewPrefixed("invalid custom transport settings",
				errors.New("roundtripper cannot be nil"),
				errors.New("verbose set to true, but no writer has been set"),
			),
		},
		{
			name: "creates a new transport with all defaults",
			args: args{cfg: CustomTransportCfg{
				RoundTripper: http.DefaultTransport,
			}},
			want: &CustomTransport{
				rt:      http.DefaultTransport,
				agent:   DefaultUserAgent,
				backoff: defaultBackoff,
			},
		},
		{
			name: "creates a new transport with custom configuration",
			args: args{cfg: CustomTransportCfg{
				RoundTripper: http.DefaultTransport,
				UserAgent:    "someagent",
				Retries:      5,
				Backoff:      time.Millisecond,
				Verbose:      true,
				Writer:       ioutil.Discard,
			}},
			want: &CustomTransport{
				rt:      http.DefaultTransport,
				agent:   "someagent",
				backoff: time.Millisecond,
				retries: 5,
				verbose: true,
				writer:  ioutil.Discard,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCustomTransport(tt.args.cfg)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type timeoutError struct{}

func (timeoutError) Error() string   { return "timeout error" }
func (timeoutError) Timeout() bool   { return true }
func (timeoutError) Temporary() bool { return true }

// nolint
func TestCustomTransport_RoundTrip(t *testing.T) {
	req, err := http.NewRequest("method", "http://localhost", nil)
	req.Header.Add("Authorization", "mocked")
	if err != nil {
		t.Fatal(err)
	}

	var sucessBuf = &bytes.Buffer{}
	sucessBuf.WriteString("{}")
	type fields struct {
		rt         http.RoundTripper
		agent      string
		retries    int
		backoff    time.Duration
		verbose    bool
		count      int64
		redactAuth bool
		writer     io.Writer
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantOut string
		err     error
	}{
		{
			name: "times out after the maximum retries have been reached",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
				),
				verbose: true,
				writer:  new(bytes.Buffer),
				retries: 1,
				backoff: time.Nanosecond,
			},
			args:    args{req: req},
			err:     context.DeadlineExceeded,
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\nrequest timed out, giving up.\n",
		},
		{
			name: "returns a different error after the maximum retries have been reached",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: errors.New("some other error"),
					},
				),
				retries: 1,
				backoff: time.Nanosecond,
				verbose: true,
				writer:  new(bytes.Buffer),
			},
			args:    args{req: req},
			err:     errors.New("some other error"),
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\n",
		},
		{
			name: "returns a different error after the maximum retries have been reached (multiple errors complying with interface)",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: &timeoutError{},
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: os.ErrDeadlineExceeded,
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						// Complies with the interface but Timeout() == false.
						Error: &net.AddrError{},
					},
				),
				retries: 3,
				backoff: time.Nanosecond,
				verbose: true,
				writer:  new(bytes.Buffer),
			},
			args:    args{req: req},
			err:     &net.AddrError{},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\nrequest timed out, retrying...\n==================== Start of Request #3 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #3  ====================\nrequest timed out, retrying...\n==================== Start of Request #4 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #4  ====================\n",
		},
		{
			name: "succeeds after retrying the request",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.New200Response(ioutil.NopCloser(sucessBuf)),
				),
				retries: 1,
				verbose: true,
				writer:  new(bytes.Buffer),
				backoff: time.Nanosecond,
			},
			args: args{req: req},
			want: &http.Response{
				StatusCode: 200,
			},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: mocked\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\n==================== Start of Response #2 ====================\nHTTP/0.0 200 OK\r\n\r\n{}\n====================  End of Response #2  ====================\n",
		},
		{
			name: "succeeds after retrying the request and redacts the Authorization header",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.New200Response(ioutil.NopCloser(sucessBuf)),
				),
				retries:    1,
				verbose:    true,
				redactAuth: true,
				writer:     new(bytes.Buffer),
				backoff:    time.Nanosecond,
			},
			args: args{req: req},
			want: &http.Response{
				StatusCode: 200,
			},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\n==================== Start of Response #2 ====================\nHTTP/0.0 200 OK\r\nContent-Length: 0\r\n\r\n\n====================  End of Response #2  ====================\n",
		},
		{
			name: "succeeds after retrying the request and redacts the Authorization header",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.Response{
						Response: http.Response{
							StatusCode: 500,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.Response{
						Response: http.Response{
							StatusCode: 404,
							Body:       mock.NewStringBody("{}"),
						},
						Error: context.DeadlineExceeded,
					},
					mock.New201Response(ioutil.NopCloser(sucessBuf)),
				),
				retries:    2,
				verbose:    true,
				redactAuth: true,
				writer:     new(bytes.Buffer),
				backoff:    time.Nanosecond,
			},
			args: args{req: req},
			want: &http.Response{
				StatusCode: 201,
			},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\nrequest timed out, retrying...\n==================== Start of Request #2 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #2  ====================\nrequest timed out, retrying...\n==================== Start of Request #3 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #3  ====================\n==================== Start of Response #3 ====================\nHTTP/0.0 201 Created\r\nContent-Length: 0\r\n\r\n\n====================  End of Response #3  ====================\n",
		},
		{
			name: "succeeds directly (Ensures that no retries are performed when err != e.Timeout())",
			fields: fields{
				rt: mock.NewRoundTripper(
					mock.New202Response(ioutil.NopCloser(sucessBuf)),
				),
				retries:    2,
				verbose:    true,
				redactAuth: true,
				writer:     new(bytes.Buffer),
				backoff:    time.Nanosecond,
			},
			args: args{req: req},
			want: &http.Response{
				StatusCode: 202,
			},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\n==================== Start of Response #1 ====================\nHTTP/0.0 202 Accepted\r\nContent-Length: 0\r\n\r\n\n====================  End of Response #1  ====================\n",
		},
		{
			name: "fails directly (Ensures that no retries are performed when err != e.Timeout())",
			fields: fields{
				rt: mock.NewRoundTripper(mock.Response{
					Response: http.Response{
						StatusCode: 500,
						Body:       mock.NewStringBody("{}"),
					},
					Error: context.Canceled,
				}),
				retries:    2,
				verbose:    true,
				redactAuth: true,
				writer:     new(bytes.Buffer),
				backoff:    time.Nanosecond,
			},
			args:    args{req: req},
			wantOut: "==================== Start of Request #1 ====================\nmethod / HTTP/1.1\r\nHost: localhost\r\nAuthorization: [REDACTED]\r\nAccept-Encoding: gzip\r\n\r\n\n====================  End of Request #1  ====================\n",
			err:     context.Canceled,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ct := &CustomTransport{
				rt:         tt.fields.rt,
				agent:      tt.fields.agent,
				retries:    tt.fields.retries,
				backoff:    tt.fields.backoff,
				verbose:    tt.fields.verbose,
				count:      tt.fields.count,
				redactAuth: tt.fields.redactAuth,
				writer:     tt.fields.writer,
			}
			got, err := ct.RoundTrip(tt.args.req)
			assert.Equal(t, tt.err, err)

			if got != nil {
				got.Body = nil
			}
			assert.Equal(t, tt.want, got)

			if buf, ok := ct.writer.(*bytes.Buffer); ok {
				assert.Equal(t, tt.wantOut, buf.String())
			}
		})
	}
}

func Test_backoff(t *testing.T) {
	type args struct {
		d time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "second backoff",
			args: args{d: time.Second},
			want: time.Second,
		},
		{
			name: "two second backoff",
			args: args{d: 2 * time.Second},
			want: 2 * time.Second,
		},
		{
			name: "ten second backoff",
			args: args{d: 10 * time.Second},
			want: 10 * time.Second,
		},
		{
			name: "millisecond backoff",
			args: args{d: time.Millisecond},
			want: time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := backoff(tt.args.d)
			assert.NotZero(t, got)
			assert.LessOrEqual(t, int64(got), int64(tt.want))
		})
	}
}
