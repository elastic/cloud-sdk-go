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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestStartMaintenance(t *testing.T) {
	type args struct {
		params MaintenanceParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Start maintenance fails due to parameter validation (Missing API)",
			args: args{
				params: MaintenanceParams{
					ID: "an ID",
				},
			},
			err: multierror.NewPrefixed("invalid allocator maintenance params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Start maintenance fails due to parameter validation (Missing ID)",
			args: args{
				params: MaintenanceParams{
					API: new(api.API),
				},
			},
			err: multierror.NewPrefixed("invalid allocator maintenance params",
				errors.New("id cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Start maintenance succeeds",
			args: args{
				params: MaintenanceParams{
					ID:     "an ID",
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(""),
							StatusCode: 202,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "POST",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/infrastructure/allocators/an ID/maintenance-mode/_start",
						},
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StartMaintenance(tt.args.params)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestStopMaintenance(t *testing.T) {
	type args struct {
		params MaintenanceParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Stop maintenance fails due to parameter validation (Missing API)",
			args: args{
				params: MaintenanceParams{
					ID: "an ID",
				},
			},
			err: multierror.NewPrefixed("invalid allocator maintenance params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Stop maintenance fails due to parameter validation (Missing ID)",
			args: args{
				params: MaintenanceParams{
					API: new(api.API),
				},
			},
			err: multierror.NewPrefixed("invalid allocator maintenance params",
				errors.New("id cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Stop maintenance succeeds",
			args: args{
				params: MaintenanceParams{
					ID:     "an ID",
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(""),
							StatusCode: 202,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "POST",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/infrastructure/allocators/an ID/maintenance-mode/_stop",
						},
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := StopMaintenance(tt.args.params)
			assert.Equal(t, tt.err, err)
		})
	}
}
