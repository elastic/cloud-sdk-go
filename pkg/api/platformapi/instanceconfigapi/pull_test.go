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

package instanceconfigapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestPullToDirectory(t *testing.T) {
	type args struct {
		params PullToDirectoryParams
	}
	tests := []struct {
		name string
		args args
		err  error
		want map[string]string
	}{
		{
			name: "fails due to param validation",
			err: multierror.NewPrefixed("invalid instance config pull params",
				apierror.ErrMissingAPI,
				errors.New("folder not specified and is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "fails listing the configs due to API error",
			args: args{params: PullToDirectoryParams{
				Region:    "us-east-1",
				Directory: "some",
				API:       api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/instances",
				Err: errors.New("error"),
			},
		},
		{
			name: "pulls instance configs successfully",
			args: args{params: PullToDirectoryParams{
				Region:    "us-east-1",
				Directory: "some-folder",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(listInstanceConfigsSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances",
					},
				}),
			}},
			want: map[string]string{
				"some-folder/data.highstorage.json": `{
  "description": "Instance configuration to be used for a higher disk/memory ratio",
  "discrete_sizes": {
    "default_size": 1024,
    "resource": "memory",
    "sizes": [
      1024,
      2048,
      4096,
      8192,
      16384,
      32768,
      65536,
      131072,
      262144
    ]
  },
  "id": "data.highstorage",
  "instance_type": "elasticsearch",
  "name": "data.highstorage",
  "node_types": [
    "data",
    "ingest",
    "master"
  ],
  "storage_multiplier": 32
}
`,
				"some-folder/kibana.json": `{
  "description": "Instance configuration to be used for Kibana",
  "discrete_sizes": {
    "default_size": 1024,
    "resource": "memory",
    "sizes": [
      1024,
      2048,
      4096,
      8192
    ]
  },
  "id": "kibana",
  "instance_type": "kibana",
  "name": "kibana",
  "storage_multiplier": 4
}
`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PullToDirectory(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}

			if p := tt.args.params.Directory; p != "" {
				matches, err := filepath.Glob(filepath.Join(p, "*.json"))
				if err != nil {
					t.Fatal(err)
				}

				for _, m := range matches {
					want, ok := tt.want[m]
					if !ok {
						t.Error("didn't find template", m, "in want")
					}

					gotV, err := ioutil.ReadFile(m)
					if err != nil {
						t.Error(err)
					}

					if got := string(gotV); got != want {
						t.Error("got", got, "!=", "want", want)
					}
				}

				if err := os.RemoveAll(p); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
