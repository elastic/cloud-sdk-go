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
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewDebugTransport(t *testing.T) {
	var debugBuffer = new(bytes.Buffer)
	var debugTransport = &DebugTransport{
		transport: http.DefaultTransport,
		output:    new(bytes.Buffer),
		count:     -1,
	}
	type args struct {
		transport http.RoundTripper
		obscure   bool
	}
	tests := []struct {
		name string
		args args
		want *DebugTransport
	}{
		{
			name: "Returns a new DebugTransport when there's an empty transport",
			args: args{},
			want: &DebugTransport{
				transport: http.DefaultTransport,
				output:    debugBuffer,
				count:     -1,
			},
		},
		{
			name: "Returns the same debugTransport when a DebugTransport is sent as a parameter",
			args: args{
				transport: debugTransport,
			},
			want: &DebugTransport{
				transport: http.DefaultTransport,
				output:    debugBuffer,
				count:     -1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &bytes.Buffer{}
			if got := NewDebugTransport(tt.args.transport, o, tt.args.obscure); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestDebugTransport_handleVerboseRequest(t *testing.T) {
	u, err := url.Parse("https://elastic.co")
	if err != nil {
		t.Fatal(err)
	}
	wantNonObscuredReq := `==================== Start of Request #0 ====================
GET / HTTP/1.1
Host: elastic.co
User-Agent: cloud-sdk-go/2.5.0-ms36
Authorization: Bearer SomeBearerValue
Accept-Encoding: gzip


====================  End of Request #0  ====================
`
	wantObscuredReq := `==================== Start of Request #0 ====================
GET / HTTP/1.1
Host: elastic.co
User-Agent: cloud-sdk-go/2.5.0-ms36
Authorization: [REDACTED]
Accept-Encoding: gzip


====================  End of Request #0  ====================
`

	type fields struct {
		output              io.Writer
		transport           http.RoundTripper
		count               int64
		redactAuthorization bool
	}
	type args struct {
		req   *http.Request
		count int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Prints the request as is",
			fields: fields{
				transport:           mock.NewRoundTripper(mock.New200Response(mock.NewStringBody(`{}`))),
				output:              new(bytes.Buffer),
				redactAuthorization: false,
			},
			args: args{req: &http.Request{
				URL: u,
				Header: http.Header{
					"User-Agent":    []string{"cloud-sdk-go/2.5.0-ms36"},
					"Authorization": []string{"Bearer SomeBearerValue"},
				},
			}},
			want: wantNonObscuredReq,
		},
		{
			name: "Prints the request obscuring the headers",
			fields: fields{
				transport:           mock.NewRoundTripper(mock.New200Response(mock.NewStringBody(`{}`))),
				output:              new(bytes.Buffer),
				redactAuthorization: true,
			},
			args: args{req: &http.Request{
				URL: u,
				Header: http.Header{
					"User-Agent":    []string{"cloud-sdk-go/2.5.0-ms36"},
					"Authorization": []string{"Bearer SomeBearerValue"},
				},
			}},
			want: wantObscuredReq,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transp := &DebugTransport{
				output:     tt.fields.output,
				transport:  tt.fields.transport,
				count:      tt.fields.count,
				redactAuth: tt.fields.redactAuthorization,
			}
			transp.handleVerboseRequest(tt.args.req, tt.args.count)
			if buf, ok := transp.output.(*bytes.Buffer); ok {
				if got := strings.ReplaceAll(buf.String(), "\r", ""); got != tt.want {
					assert.EqualValues(t, tt.want, got)
				}
			}
		})
	}
}
