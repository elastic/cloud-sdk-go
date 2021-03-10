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

package stackapi

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	urlError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/stack/versions?show_deleted=false",
		Err: errors.New(`{"error": "some error"}`),
	}

	deleteURLError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/stack/versions?show_deleted=true",
		Err: errors.New(`{"error": "some error"}`),
	}
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.StackVersionConfigs
		err  string
	}{
		{
			name: "List succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body: mock.NewStructBody(models.StackVersionConfigs{
							Stacks: []*models.StackVersionConfig{
								{
									Deleted: ec.Bool(false),
									Version: "6.0.0",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
								{
									Deleted: ec.Bool(false),
									Version: "6.1.0",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
								{
									Deleted: ec.Bool(false),
									Version: "6.2.0",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
								{
									Deleted: ec.Bool(false),
									Version: "5.6.0",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
								{
									Deleted: ec.Bool(false),
									Version: "7.9.3",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
								{
									Deleted: ec.Bool(false),
									Version: "7.10.0",
									Kibana: &models.StackVersionKibanaConfig{
										CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
											Max: ec.Int32(8192),
											Min: ec.Int32(1024),
										},
									},
								},
							},
						}),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/stack/versions",
						Query: url.Values{
							"show_deleted": []string{"false"},
						},
					},
				}),
			}},
			want: &models.StackVersionConfigs{
				Stacks: []*models.StackVersionConfig{
					{
						Deleted: ec.Bool(false),
						Version: "7.10.0",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
					{
						Deleted: ec.Bool(false),
						Version: "7.9.3",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
					{
						Deleted: ec.Bool(false),
						Version: "6.2.0",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
					{
						Deleted: ec.Bool(false),
						Version: "6.1.0",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
					{
						Deleted: ec.Bool(false),
						Version: "6.0.0",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
					{
						Deleted: ec.Bool(false),
						Version: "5.6.0",
						Kibana: &models.StackVersionKibanaConfig{
							CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
								Max: ec.Int32(8192),
								Min: ec.Int32(1024),
							},
						},
					},
				},
			},
		},
		{
			name: "List fails due to API error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New(`{"error": "some error"}`),
				}),
			}},
			err: urlError.Error(),
		},
		{
			name: "List deleted fails due to API error",
			args: args{params: ListParams{
				Deleted: true,
				Region:  "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New(`{"error": "some error"}`),
				}),
			}},
			err: deleteURLError.Error(),
		},
		{
			name: "List fails due to validation",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid stack list params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("List() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	type fields struct {
		Version1 string
		Version2 string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Compare versions different major versions",
			fields: fields{
				Version1: "6.0.0",
				Version2: "5.9.9",
			},
			want: true,
		},
		{
			name: "Compare versions different minor versions",
			fields: fields{
				Version1: "5.5.0",
				Version2: "5.4.9",
			},
			want: true,
		},
		{
			name: "Compare versions different patch versions",
			fields: fields{
				Version1: "5.5.5",
				Version2: "5.5.4",
			},
			want: true,
		},
		{
			name: "Compare versions different patch versions greater than 9",
			fields: fields{
				Version1: "5.5.10",
				Version2: "5.5.9",
			},
			want: true,
		},
		{
			name: "Compare versions different patch versions greater than 9",
			fields: fields{
				Version1: "7.10.0",
				Version2: "7.9.3",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := compareVersions(tt.fields.Version1, tt.fields.Version2)
			if res != tt.want {
				t.Errorf("CompareVersions() want = %v, actual %v", tt.want, res)
			}
		})
	}
}
