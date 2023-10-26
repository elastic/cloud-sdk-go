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
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.InstanceConfiguration
		err  string
	}{
		{
			name: "Get succeeds",
			args: args{params: GetParams{
				Region: "us-east-1",
				ID:     "data.highstorage",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(getInstanceConfigsSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances/data.highstorage",
					},
				}),
			}},
			want: &models.InstanceConfiguration{
				ID:                "data.highstorage",
				Description:       "Instance configuration to be used for a higher disk/memory ratio",
				Name:              ec.String("data.highstorage"),
				InstanceType:      "elasticsearch",
				StorageMultiplier: float64(32),
				NodeTypes:         []string{"data", "ingest", "master"},
				DiscreteSizes: &models.DiscreteSizes{
					DefaultSize: ec.Int32(1024),
					Resource:    ec.String("memory"),
					Sizes: []int32{
						1024,
						2048,
						4096,
						8192,
						16384,
						32768,
						65536,
						131072,
						262144,
					},
				},
			},
		},
		{
			name: "Get succeeds on kibana ID",
			args: args{params: GetParams{
				Region: "us-east-1",
				ID:     "kibana",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(getInstanceConfigsSuccessKibana),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances/kibana",
					},
				}),
			}},
			want: &models.InstanceConfiguration{
				ID:                "kibana",
				Description:       "Instance configuration to be used for Kibana",
				Name:              ec.String("kibana"),
				InstanceType:      "kibana",
				StorageMultiplier: float64(4),
				DiscreteSizes: &models.DiscreteSizes{
					DefaultSize: ec.Int32(1024),
					Resource:    ec.String("memory"),
					Sizes: []int32{
						1024,
						2048,
						4096,
						8192,
					},
				},
			},
		},
		{
			name: "Get fails on API error",
			args: args{
				params: GetParams{
					Region: "us-east-1",
					ID:     "kibana",
					API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				},
			},
			err: `{"error": "some error"}`,
		},
		{
			name: "Get fails on parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("invalid instance config get params",
				apierror.ErrMissingAPI,
				errors.New("id not specified and is required for the operation"),
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
			assert.Equal(t, tt.want, got)
		})
	}
}
