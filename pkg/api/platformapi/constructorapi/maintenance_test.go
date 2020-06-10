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

package constructorapi

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestEnableMaintenace(t *testing.T) {
	type args struct {
		params MaintenanceParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Set constructor to maintenance mode succeeds",
			args: args{params: MaintenanceParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{}`),
						StatusCode: 202,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "POST",
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10/maintenance-mode/_start",
					},
				}),
			}},
		},
		{
			name: "Set constructor to maintenance mode fails",
			args: args{params: MaintenanceParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10/maintenance-mode/_start",
				Err: errors.New("error"),
			},
		},
		{
			name: "Set constructor to maintenance mode fails due to param validation",
			err: multierror.NewPrefixed("invalid constructor maintenance params",
				apierror.ErrMissingAPI,
				errors.New("id field cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := EnableMaintenace(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}

func TestDisableMaintenance(t *testing.T) {
	type args struct {
		params MaintenanceParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Set constructor maintenance mode to false succeeds",
			args: args{params: MaintenanceParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{}`),
						StatusCode: 202,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "POST",
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10/maintenance-mode/_stop",
					},
				}),
			}},
		},
		{
			name: "Set constructor maintenance mode to false fails",
			args: args{params: MaintenanceParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10/maintenance-mode/_stop",
				Err: errors.New("error"),
			},
		},
		{
			name: "Set constructor maintenance mode to false fails due to param validation",
			err: multierror.NewPrefixed("invalid constructor maintenance params",
				apierror.ErrMissingAPI,
				errors.New("id field cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DisableMaintenance(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
