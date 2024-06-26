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

func TestUpdate(t *testing.T) {
	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "Update Succeeds",
			args: args{params: UpdateParams{
				Region: "us-east-1",
				ID:     "kibana",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{"id": "an autogenerated id"}`),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Body:   mock.NewStringBody(`{"allowed_zones":[],"description":"Instance configuration to be used for Kibana","discrete_sizes":{"default_size":1024,"resource":"memory","sizes":[1024,2048,4096,8192]},"id":"kibana","instance_type":"kibana","name":"kibana","storage_multiplier":4}` + "\n"),
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances/kibana",
					},
				}),
				Config: &models.InstanceConfiguration{
					ID:                "kibana",
					Description:       "Instance configuration to be used for Kibana",
					Name:              ec.String("kibana"),
					InstanceType:      "kibana",
					StorageMultiplier: float64(4),
					NodeTypes:         []string{},
					DiscreteSizes: &models.DiscreteSizes{
						DefaultSize: 1024,
						Resource:    "memory",
						Sizes: []int32{
							1024,
							2048,
							4096,
							8192,
						},
					},
					AllowedZones: []string{},
				},
			}},
		},
		{
			name: "Update fails on API error",
			args: args{params: UpdateParams{
				Region: "us-eastt-1",
				ID:     "kibana",
				Config: &models.InstanceConfiguration{
					ID:                "kibana",
					Description:       "Instance configuration to be used for Kibana",
					Name:              ec.String("kibana"),
					InstanceType:      "kibana",
					StorageMultiplier: float64(4),
					NodeTypes:         []string{},
					DiscreteSizes: &models.DiscreteSizes{
						DefaultSize: 1024,
						Resource:    "memory",
						Sizes: []int32{
							1024,
							2048,
							4096,
							8192,
						},
					},
					AllowedZones: []string{},
				},
				API: api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Update fails on parameter validation failure",
			err: multierror.NewPrefixed("invalid instance config update params",
				apierror.ErrMissingAPI,
				errors.New("config not specified and is required for the operation"),
				errors.New("id not specified and is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Update(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
		})
	}
}
