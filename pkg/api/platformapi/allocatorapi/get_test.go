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

package allocatorapi

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
	var getAllocatorSuccess = `{
  "allocator_id": "i-09a0e797fb3af6864",
  "capacity": {
	"memory": {
	  "total": 236544
	}
  },
  "features": [
	"apm",
	"elasticsearch",
	"elasticsearch_data",
	"ssd",
	"templates",
	"tinyauth"
  ],
  "host_ip": "172.25.61.100",
  "instances": [],
  "metadata": [
	{
	  "key": "version",
	  "value": "2017-09-30"
	},
	{
	  "key": "instanceId",
	  "value": "i-09a0e797fb3af6864"
	},
	{
	  "key": "architecture",
	  "value": "x86_64"
	},
	{
	  "key": "instanceType",
	  "value": "i3.8xlarge"
	},
	{
	  "key": "availabilityZone",
	  "value": "us-east-1a"
	},
	{
	  "key": "pendingTime",
	  "value": "2018-05-18T13:24:21Z"
	},
	{
	  "key": "imageId",
	  "value": "ami-ba0a51c0"
	},
	{
	  "key": "privateIp",
	  "value": "172.25.61.100"
	},
	{
	  "key": "region",
	  "value": "us-east-1"
	}
  ],
  "public_hostname": "172.25.61.100",
  "settings": {},
  "status": {
	"connected": true,
	"healthy": true,
	"maintenance_mode": false
  },
  "zone_id": "us-east-1a"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.AllocatorInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure (missing API)",
			args: args{
				params: GetParams{
					ID: "i-09a0e797fb3af6864",
				},
			},
			err: multierror.NewPrefixed("invalid allocator get params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "Get fails due to parameter validation failure (missing ID)",
			args: args{
				params: GetParams{
					API: new(api.API),
				},
			},
			err: multierror.NewPrefixed("invalid allocator get params",
				errors.New("id cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					ID:     "an id",
					API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
					Region: "some-region",
				},
			},
			err: `{"error": "some error"}`,
		},
		{
			name: "Get Succeeds",
			args: args{
				params: GetParams{
					ID: "i-09a0e797fb3af6864",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getAllocatorSuccess),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/infrastructure/allocators/i-09a0e797fb3af6864",
						},
					}),
					Region: "us-east-1",
				},
			},
			want: &models.AllocatorInfo{
				AllocatorID: ec.String("i-09a0e797fb3af6864"),
				Capacity: &models.AllocatorCapacity{Memory: &models.AllocatorCapacityMemory{
					Total: ec.Int32(236544),
				}},
				Features: []string{
					"apm",
					"elasticsearch",
					"elasticsearch_data",
					"ssd",
					"templates",
					"tinyauth",
				},
				HostIP:    ec.String("172.25.61.100"),
				Instances: []*models.AllocatedInstanceStatus{},
				Metadata: []*models.MetadataItem{
					{
						Key:   ec.String("version"),
						Value: ec.String("2017-09-30"),
					},
					{
						Key:   ec.String("instanceId"),
						Value: ec.String("i-09a0e797fb3af6864"),
					},
					{
						Key:   ec.String("architecture"),
						Value: ec.String("x86_64"),
					},
					{
						Key:   ec.String("instanceType"),
						Value: ec.String("i3.8xlarge"),
					},
					{
						Key:   ec.String("availabilityZone"),
						Value: ec.String("us-east-1a"),
					},
					{
						Key:   ec.String("pendingTime"),
						Value: ec.String("2018-05-18T13:24:21Z"),
					},
					{
						Key:   ec.String("imageId"),
						Value: ec.String("ami-ba0a51c0"),
					},
					{
						Key:   ec.String("privateIp"),
						Value: ec.String("172.25.61.100"),
					},
					{
						Key:   ec.String("region"),
						Value: ec.String("us-east-1"),
					},
				},
				PublicHostname: ec.String("172.25.61.100"),
				Settings:       &models.AllocatorSettings{},
				Status: &models.AllocatorHealthStatus{
					Connected:       ec.Bool(true),
					Healthy:         ec.Bool(true),
					MaintenanceMode: ec.Bool(false),
				},
				ZoneID: ec.String("us-east-1a"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}
