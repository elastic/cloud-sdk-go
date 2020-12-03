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

func TestGet(t *testing.T) {
	var proxiesFilteredGroup = `
	{
      "expected_proxies_count": 5,
      "filters": [
        {
          "key": "proxyType",
          "value": "main-nextgen"
        }
      ],
      "id": "test1"
	}`
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxiesFilteredGroup
		err  string
	}{
		{
			name: "Proxies filtered group get succeeds",
			args: args{params: GetParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxiesFilteredGroup)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/filtered-groups/test1",
					},
				}),
				ID: "test1",
			}},
			want: &models.ProxiesFilteredGroup{
				ExpectedProxiesCount: ec.Int32(5),
				Filters: []*models.ProxiesFilter{
					{
						Key:   ec.String("proxyType"),
						Value: ec.String("main-nextgen"),
					},
				},
				ID: *ec.String("test1"),
			},
		},
		{
			name: "Proxies filtered group get fails with 403 Forbidden",
			args: args{params: GetParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				ID:     "test1",
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Proxies filtered group get fails due validation",
			args: args{params: GetParams{}},
			err: multierror.NewPrefixed("invalid filtered group params",
				errors.New("id is not specified and is required for the operation"),
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
