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
				if tt.want.runtime.BasePath != got.runtime.BasePath {
					t.Errorf("NewCloudClientRuntime() runtime = %v, want %v",
						got.runtime.BasePath, tt.want.runtime.BasePath,
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
	var mocknewRuntimeFunc = func(r string) *runtimeclient.Runtime {
		return AddTypeConsumers(runtimeclient.NewWithClient(
			"cloud.elastic.co", fmt.Sprintf(RegionPrefix, r),
			[]string{"https"}, nil,
		))
	}
	var regionless = AddTypeConsumers(runtimeclient.NewWithClient(
		"cloud.elastic.co",
		RegionlessPrefix,
		[]string{"https"}, nil,
	))
	type fields struct {
		newRegionRuntime newRuntimeFunc
		runtime          *runtimeclient.Runtime
		region           string
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
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/deployments",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/"},
		},
		{
			name: "/deployment/someid/notes operation uses the region path",
			fields: fields{
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/deployments/someid/notes",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-1"},
		},
		{
			name: "/platform operation uses the regioned path",
			fields: fields{
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/platform",
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-1"},
		},
		{
			name: "/platform operation uses the regioned path",
			fields: fields{
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/platform",
				Context:     context.Background(),
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-1"},
		},
		{
			name: "/platform operation uses the regioned path obtained from the region context",
			fields: fields{
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/platform",
				Context:     WithRegion(context.Background(), "us-west-1"),
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-west-1"},
		},
		{
			name: "/platform operation uses a different regioned path obtained from the region context",
			fields: fields{
				region:           "us-east-1",
				newRegionRuntime: mocknewRuntimeFunc,
				runtime:          regionless,
			},
			args: args{op: &runtime.ClientOperation{
				PathPattern: "/platform",
				Context:     WithRegion(context.Background(), "us-east-2"),
			}},
			want: &runtimeclient.Runtime{BasePath: "/api/v1/regions/us-east-2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &CloudClientRuntime{
				newRegionRuntime: tt.fields.newRegionRuntime,
				runtime:          tt.fields.runtime,
				region:           tt.fields.region,
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
