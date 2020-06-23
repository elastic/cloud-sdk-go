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

package filteredgroupapi

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	var proxiesHealth = `
	{
	  "filtered_groups": [
		{
		  "group": {
			"id": "main-proxies",
			"filters": [
			  {
				"key": "proxyType",
				"value": "main-nextgen"
			  }
			],
			"expected_proxies_count": 1
		  },
		  "observed_proxies_count": 7,
		  "status": "Green"
		},
		{
		  "group": {
			"id": "apm-proxies",
			"filters": [
			  {
				"key": "proxyType",
				"value": "apm"
			  }
			],
			"expected_proxies_count": 1
		  },
		  "observed_proxies_count": 0,
		  "status": "Red"
		}
	  ],
	  "expected_proxies_count": 2,
	  "status": "Yellow",
	  "allocations": [
		{
		  "allocations_type": "apm",
		  "max_allocations": 82,
		  "proxies_at_max_allocations": 10
		},
		{
		  "allocations_type": "elasticsearch",
		  "max_allocations": 293,
		  "proxies_at_max_allocations": 3
		},
		{
		  "allocations_type": "sitesearch",
		  "max_allocations": 0,
		  "proxies_at_max_allocations": 10
		},
		{
		  "allocations_type": "appsearch",
		  "max_allocations": 0,
		  "proxies_at_max_allocations": 10
		},
		{
		  "allocations_type": "enterprisesearch",
		  "max_allocations": 0,
		  "proxies_at_max_allocations": 10
		},
		{
		  "allocations_type": "kibana",
		  "max_allocations": 247,
		  "proxies_at_max_allocations": 3
		}
	  ],
	  "observed_proxies_count": 10
	}`
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want []*models.ProxiesFilteredGroupHealth
		err  error
	}{
		{
			name: "Proxies filtered group list succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxiesHealth)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/health",
					},
				}),
			}},
			want: []*models.ProxiesFilteredGroupHealth{
				{
					Group: &models.ProxiesFilteredGroup{
						ExpectedProxiesCount: ec.Int32(1),
						Filters: []*models.ProxiesFilter{
							{
								Key:   ec.String("proxyType"),
								Value: ec.String("main-nextgen"),
							},
						},
						ID: *ec.String("main-proxies"),
					},
					ObservedProxiesCount: ec.Int32(7),
					Status:               ec.String("Green"),
				},
				{
					Group: &models.ProxiesFilteredGroup{
						ExpectedProxiesCount: ec.Int32(1),
						Filters: []*models.ProxiesFilter{
							{
								Key:   ec.String("proxyType"),
								Value: ec.String("apm"),
							},
						},
						ID: *ec.String("apm-proxies"),
					},
					ObservedProxiesCount: ec.Int32(0),
					Status:               ec.String("Red"),
				},
			},
		},
		{
			name: "Proxies filtered group list fails with 403 Forbidden",
			args: args{params: ListParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "Proxies filtered group list fails due validation",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid filtered group params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
