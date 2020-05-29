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
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
)

func TestNewCloudClientRuntime(t *testing.T) {
	type args struct {
		c Config
	}
	tests := []struct {
		name string
		args args
		want *CloudClientRuntime
		err  error
	}{
		{
			name: "returns an error when it can't parse the host",
			args: args{c: Config{
				Host: " https://cloud.elastic.co",
			}},
			err: &url.Error{
				Op:  "parse",
				URL: " https://cloud.elastic.co",
				Err: errors.New("first path segment in URL cannot contain colon"),
			},
		},
		{
			name: "when region is specified the structure has two different runtimes",
			args: args{c: Config{
				Host:   "https://cloud.elastic.co",
				Region: "us-east-1",
			}},
			want: &CloudClientRuntime{
				regionRuntime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					fmt.Sprintf(RegionPrefix, "us-east-1"),
					[]string{"https"}, nil,
				)),
				runtime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					RegionlessPrefix,
					[]string{"https"}, nil,
				)),
			},
		},
		{
			name: "when region is not specified the structure has two equal runtimes",
			args: args{c: Config{
				Host: "https://cloud.elastic.co",
			}},
			want: &CloudClientRuntime{
				regionRuntime: &runtimeclient.Runtime{
					Host:     "cloud.elastic.co",
					BasePath: RegionlessPrefix,
				},
				runtime: &runtimeclient.Runtime{
					Host:     "cloud.elastic.co",
					BasePath: RegionlessPrefix,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCloudClientRuntime(tt.args.c)
			if !reflect.DeepEqual(tt.err, err) {
				t.Errorf("NewCloudClientRuntime() error = %v, wantErr %v", err, tt.err)
				return
			}
			if tt.want != nil && got != nil {
				if tt.want.regionRuntime.BasePath != got.regionRuntime.BasePath {
					t.Errorf("NewCloudClientRuntime() regionRuntime = %v, want %v",
						got.regionRuntime.BasePath, tt.want.regionRuntime.BasePath,
					)
				}
				if tt.want.runtime.BasePath != got.runtime.BasePath {
					t.Errorf("NewCloudClientRuntime() runtime = %v, want %v",
						got.runtime.BasePath, tt.want.runtime.BasePath,
					)
				}
				if tt.want.regionRuntime.Host != got.regionRuntime.Host {
					t.Errorf("NewCloudClientRuntime() regionRuntime = %v, want %v",
						got.regionRuntime.Host, tt.want.regionRuntime.Host,
					)
				}
				if tt.want.runtime.Host != got.runtime.Host {
					t.Errorf("NewCloudClientRuntime() runtime = %v, want %v",
						got.runtime.Host, tt.want.runtime.Host,
					)
				}
			}
		})
	}
}

func TestCloudClientRuntime_getRuntime(t *testing.T) {
	type fields struct {
		regionRuntime *runtimeclient.Runtime
		runtime       *runtimeclient.Runtime
	}
	type args struct {
		op *runtime.ClientOperation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *runtimeclient.Runtime
	}{
		{
			name: "/deployment operation uses the regionless path",
			fields: fields{
				regionRuntime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					fmt.Sprintf(RegionPrefix, "us-east-1"),
					[]string{"https"}, nil,
				)),
				runtime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					RegionlessPrefix,
					[]string{"https"}, nil,
				)),
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/deployments",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1"},
		},
		{
			name: "/deployment/someid/notes operation uses the region path",
			fields: fields{
				regionRuntime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					fmt.Sprintf(RegionPrefix, "us-east-1"),
					[]string{"https"}, nil,
				)),
				runtime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					RegionlessPrefix,
					[]string{"https"}, nil,
				)),
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/deployments/someid/notes",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-1"},
		},
		{
			name: "/platform operation uses the regioned path",
			fields: fields{
				regionRuntime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					fmt.Sprintf(RegionPrefix, "us-east-1"),
					[]string{"https"}, nil,
				)),
				runtime: AddTypeConsumers(runtimeclient.NewWithClient(
					"cloud.elastic.co",
					RegionlessPrefix,
					[]string{"https"}, nil,
				)),
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/platform",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CloudClientRuntime{
				regionRuntime: tt.fields.regionRuntime,
				runtime:       tt.fields.runtime,
			}
			got := r.getRuntime(tt.args.op)
			if tt.want.BasePath != got.BasePath {
				t.Errorf("NewCloudClientRuntime() = %v, want %v",
					got.BasePath, tt.want.BasePath,
				)
			}
		})
	}
}

func Test_overrideJSONProducer(t *testing.T) {
	type args struct {
		r       *runtimeclient.Runtime
		opID    string
		content string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		callback bool
	}{
		{
			name: "doesn't do anything when the operation ID doesn't match the expectation",
			args: args{
				r: &runtimeclient.Runtime{
					Producers: map[string]runtime.Producer{
						runtime.JSONMime: runtime.JSONProducer(),
					},
				},
				opID:    "some-id",
				content: `{"some":"content"}`,
			},
			want: `"{\"some\":\"content\"}"` + "\n",
		},
		{
			name: "doesn't do anything when the operation ID doesn't match the expectation (callback: true)",
			args: args{
				r: &runtimeclient.Runtime{
					Producers: map[string]runtime.Producer{
						runtime.JSONMime: runtime.JSONProducer(),
					},
				},
				opID:    "some-id",
				content: `{"some":"content"}`,
			},
			callback: true,
			want:     `"{\"some\":\"content\"}"` + "\n",
		},
		{
			name: "changes the producer",
			args: args{
				r: &runtimeclient.Runtime{
					Producers: map[string]runtime.Producer{
						runtime.JSONMime: runtime.JSONProducer(),
					},
				},
				opID:    "set-es-cluster-metadata-raw",
				content: `{"some":"content"}`,
			},
			want: `{"some":"content"}`,
		},
		{
			name: "resets the producer even when changed",
			args: args{
				r: &runtimeclient.Runtime{
					Producers: map[string]runtime.Producer{
						runtime.JSONMime: runtime.JSONProducer(),
					},
				},
				opID:    "set-es-cluster-metadata-raw",
				content: `{"some":"content"}`,
			},
			callback: true,
			want:     `"{\"some\":\"content\"}"` + "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if f := overrideJSONProducer(tt.args.r, tt.args.opID); tt.callback {
				f()
			}

			var buf = new(bytes.Buffer)
			tt.args.r.Producers[runtime.JSONMime].Produce(buf, tt.args.content)
			if buf.String() != tt.want {
				t.Errorf("overrideJSONProducer() = %v, want %v", buf.String(), tt.want)
			}
		})
	}
}
