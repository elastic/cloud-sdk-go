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
	"net/http"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
)

type mockRT struct{}

func (*mockRT) RoundTrip(*http.Request) (*http.Response, error) { return nil, nil }

func TestNewTransport(t *testing.T) {
	type args struct {
		rt  http.RoundTripper
		cfg TransportConfig
		buf *bytes.Buffer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "An empty rt with no config returns the default transport",
			args: args{buf: new(bytes.Buffer)},
		},
		{
			name: "An empty rt with SkipTLSVerify returns the default transport",
			args: args{cfg: TransportConfig{
				SkipTLSVerify: true,
			}, buf: new(bytes.Buffer)},
		},
		{
			name: "An empty rt with SkipTLSVerify returns the default transport",
			args: args{cfg: TransportConfig{
				VerboseSettings: VerboseSettings{Verbose: true},
				SkipTLSVerify:   true,
			}, buf: new(bytes.Buffer)},
		},
		{
			name: "DebugTransport is returned",
			args: args{
				buf: new(bytes.Buffer),
				cfg: TransportConfig{
					VerboseSettings: VerboseSettings{Verbose: true},
					SkipTLSVerify:   true,
				},
				rt: new(DebugTransport),
			},
		},
		{
			name: "mock.RoundTripper is returned",
			args: args{
				buf: new(bytes.Buffer),
				rt:  new(mock.RoundTripper),
			},
		},
		{
			name: "Receives a message when the rt is not known",
			args: args{
				buf: new(bytes.Buffer),
				cfg: TransportConfig{
					VerboseSettings: VerboseSettings{Verbose: true},
					SkipTLSVerify:   true,
				},
				rt: new(mockRT),
			},
		},
		{
			name: "Receives no error when the UserAgentTransport",
			args: args{
				buf: new(bytes.Buffer),
				cfg: TransportConfig{
					VerboseSettings: VerboseSettings{Verbose: true},
					SkipTLSVerify:   true,
				},
				rt: new(UserAgentTransport),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.buf != nil {
				tt.args.cfg.ErrorDevice = tt.args.buf
			}

			// Not possible to assert the http.RoundTripper type.
			NewTransport(tt.args.rt, tt.args.cfg)

			if tt.args.buf != nil {
				if got := tt.args.buf.String(); got != tt.want {
					t.Errorf("NewTransport() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
