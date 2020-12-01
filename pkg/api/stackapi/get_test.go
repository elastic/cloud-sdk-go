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

func TestGet(t *testing.T) {
	urlError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/stack/versions/6.0.0",
		Err: errors.New(`{"error": "some error"}`),
	}
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.StackVersionConfig
		err  string
	}{
		{
			name: "Get Succeeds",
			args: args{params: GetParams{
				Version: "6.0.0",
				Region:  "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body: mock.NewStructBody(models.StackVersionConfig{
							Deleted: ec.Bool(false),
							Version: "6.0.0",
							Kibana: &models.StackVersionKibanaConfig{
								CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
									Max: ec.Int32(8192),
									Min: ec.Int32(1024),
								},
							},
						}),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/stack/versions/6.0.0",
					},
				}),
			}},
			want: &models.StackVersionConfig{
				Deleted: ec.Bool(false),
				Version: "6.0.0",
				Kibana: &models.StackVersionKibanaConfig{
					CapacityConstraints: &models.StackVersionInstanceCapacityConstraint{
						Max: ec.Int32(8192),
						Min: ec.Int32(1024),
					},
				},
			},
		},
		{
			name: "Get fails due to API error",
			args: args{params: GetParams{
				Version: "6.0.0",
				Region:  "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New(`{"error": "some error"}`),
				}),
			}},
			err: urlError.Error(),
		},
		{
			name: "Get fails due to empty parameters",
			args: args{params: GetParams{}},
			err: multierror.NewPrefixed("invalid stack get params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
				errors.New("version string empty"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
