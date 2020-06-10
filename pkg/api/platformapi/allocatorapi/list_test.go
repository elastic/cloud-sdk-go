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
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	var listAllocatorsSuccess = `{
  "zones": [
	{
	  "allocators": [
		{
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
		},
		{
		  "allocator_id": "i-09a0e797fb3af6861",
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
		  "host_ip": "172.25.61.201",
		  "instances": [],
		  "metadata": [
			{
			  "key": "version",
			  "value": "2017-09-30"
			},
			{
			  "key": "instanceId",
			  "value": "i-09a0e797fb3af6861"
			},
			{
			  "key": "architecture",
			  "value": "x86_64"
			},
			{
			  "key": "instanceType",
			  "value": "i3.large"
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
			  "value": "172.25.61.200"
			},
			{
			  "key": "region",
			  "value": "us-east-1"
			}
		  ],
		  "public_hostname": "172.25.61.200",
		  "settings": {},
		  "status": {
			"connected": true,
			"healthy": true,
			"maintenance_mode": false
		  },
		  "zone_id": "us-east-1a"
		},
		{
		  "allocator_id": "i-09a0e797fb3a12345",
		  "capacity": {
			"memory": {
			  "total": 16384
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
		  "host_ip": "172.25.61.202",
		  "instances": [],
		  "metadata": [],
		  "public_hostname": "172.25.61.200",
		  "settings": {},
		  "status": {
			"connected": true,
			"healthy": true,
			"maintenance_mode": false
		  },
		  "zone_id": "us-east-1a"
		}
	  ]
	}
  ]
}`
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.AllocatorOverview
		err  error
	}{
		{
			name: "List fails due to parameter validation failure (missing API)",
			args: args{},
			err: multierror.NewPrefixed("invalid allocator list params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "List fails due to API error",
			args: args{params: ListParams{
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				Region: "us-east-1",
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "List Succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					Body:       mock.NewStringBody(listAllocatorsSuccess),
					StatusCode: 200,
				}}),
			}},
			want: &models.AllocatorOverview{
				Zones: []*models.AllocatorZoneInfo{
					{
						Allocators: []*models.AllocatorInfo{
							{
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
							{
								AllocatorID: ec.String("i-09a0e797fb3af6861"),
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
								HostIP:    ec.String("172.25.61.201"),
								Instances: []*models.AllocatedInstanceStatus{},
								Metadata: []*models.MetadataItem{
									{
										Key:   ec.String("version"),
										Value: ec.String("2017-09-30"),
									},
									{
										Key:   ec.String("instanceId"),
										Value: ec.String("i-09a0e797fb3af6861"),
									},
									{
										Key:   ec.String("architecture"),
										Value: ec.String("x86_64"),
									},
									{
										Key:   ec.String("instanceType"),
										Value: ec.String("i3.large"),
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
										Value: ec.String("172.25.61.200"),
									},
									{
										Key:   ec.String("region"),
										Value: ec.String("us-east-1"),
									},
								},
								PublicHostname: ec.String("172.25.61.200"),
								Settings:       &models.AllocatorSettings{},
								Status: &models.AllocatorHealthStatus{
									Connected:       ec.Bool(true),
									Healthy:         ec.Bool(true),
									MaintenanceMode: ec.Bool(false),
								},
								ZoneID: ec.String("us-east-1a"),
							},
							{
								AllocatorID: ec.String("i-09a0e797fb3a12345"),
								Capacity: &models.AllocatorCapacity{Memory: &models.AllocatorCapacityMemory{
									Total: ec.Int32(16384),
								}},
								Features: []string{
									"apm",
									"elasticsearch",
									"elasticsearch_data",
									"ssd",
									"templates",
									"tinyauth",
								},
								HostIP:         ec.String("172.25.61.202"),
								Instances:      []*models.AllocatedInstanceStatus{},
								Metadata:       []*models.MetadataItem{},
								PublicHostname: ec.String("172.25.61.200"),
								Settings:       &models.AllocatorSettings{},
								Status: &models.AllocatorHealthStatus{
									Connected:       ec.Bool(true),
									Healthy:         ec.Bool(true),
									MaintenanceMode: ec.Bool(false),
								},
								ZoneID: ec.String("us-east-1a"),
							},
						},
					},
				},
			},
		},
		{
			name: "List Succeeds with one filter tag",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					Body:       mock.NewStringBody(listAllocatorsSuccess),
					StatusCode: 200,
				}}),
				FilterTags: "[instanceType:i3.8xlarge]",
			}},
			want: &models.AllocatorOverview{
				Zones: []*models.AllocatorZoneInfo{
					{
						Allocators: []*models.AllocatorInfo{
							{
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
					},
				},
			},
		},
		{
			name: "List Succeeds with two filter tags",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(listAllocatorsSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/allocators",
					},
				}),
				FilterTags: "[instanceType:i3.8xlarge,instanceType:i3.large]",
			}},
			want: &models.AllocatorOverview{
				Zones: []*models.AllocatorZoneInfo{
					{
						Allocators: []*models.AllocatorInfo{
							{
								AllocatorID: ec.String("i-09a0e797fb3af6861"),
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
								HostIP:    ec.String("172.25.61.201"),
								Instances: []*models.AllocatedInstanceStatus{},
								Metadata: []*models.MetadataItem{
									{
										Key:   ec.String("version"),
										Value: ec.String("2017-09-30"),
									},
									{
										Key:   ec.String("instanceId"),
										Value: ec.String("i-09a0e797fb3af6861"),
									},
									{
										Key:   ec.String("architecture"),
										Value: ec.String("x86_64"),
									},
									{
										Key:   ec.String("instanceType"),
										Value: ec.String("i3.large"),
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
										Value: ec.String("172.25.61.200"),
									},
									{
										Key:   ec.String("region"),
										Value: ec.String("us-east-1"),
									},
								},
								PublicHostname: ec.String("172.25.61.200"),
								Settings:       &models.AllocatorSettings{},
								Status: &models.AllocatorHealthStatus{
									Connected:       ec.Bool(true),
									Healthy:         ec.Bool(true),
									MaintenanceMode: ec.Bool(false),
								},
								ZoneID: ec.String("us-east-1a"),
							},
						},
					},
				},
			},
		},
		{
			name: "List Succeeds with unknown filter tags",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					Body:       mock.NewStringBody(listAllocatorsSuccess),
					StatusCode: 200,
				}}),
				FilterTags: "[unkowntag:withvalue,instanceType:i3.large]",
			}},
			want: &models.AllocatorOverview{Zones: []*models.AllocatorZoneInfo{
				{Allocators: nil},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("List() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}
