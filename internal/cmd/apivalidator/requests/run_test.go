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

package requests

import (
	"errors"
	"net/http"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestRun(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name: "Fails due to validation",
			args: args{config: Config{}},
			want: 1,
			wantErr: multierror.NewPrefixed("api validation configuration",
				errors.New("an api specification file must be specified"),
				errors.New("a host for connecting to the validation proxy must be specified"),
				errors.New("a port for connecting to the validation proxy must be specified"),
				errors.New("an http client must be specified"),
			),
		},
		{
			name: "Fails due to corrupted file",
			args: args{config: Config{
				Source: "test/corrupted.json",
				Host:   "http://0.0.0.0",
				Port:   "1234",
				Client: new(http.Client),
			}},
			want:    4,
			wantErr: errors.New("invalid character 'h' looking for beginning of value"),
		},
		{
			name: "Fails due to connection error",
			args: args{config: Config{
				Source: "test/apidocs-valid.json",
				Host:   "http://0.0.0.0",
				Port:   "1234",
				Client: mock.NewClient(mock.Response{Error: errors.New("connection refused")}),
			}},
			want: 5,
			wantErr: multierror.NewPrefixed("api spec validation",
				//nolint - error strings should not be capitalized or end with punctuation or a newline
				errors.New(`Post "http://0.0.0.0:1234/deployments/_search": connection refused`),
			),
		},
		{
			name: "Fails due failed file download",
			args: args{config: Config{
				Source: "https://example.co/hola",
				Host:   "http://0.0.0.0",
				Port:   "1234",
				Client: mock.NewClient(mock.Response{Error: errors.New("no such host")}),
			}},
			want: 2,
			//nolint - error strings should not be capitalized or end with punctuation or a newline
			wantErr: errors.New(`Get "https://example.co/hola": no such host`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Run(tt.args.config)
			if err != nil && (err.Error() != tt.wantErr.Error()) {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
