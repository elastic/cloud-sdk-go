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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	var constructoGet = `
	{
		"constructor_id": "192.168.44.10",
		"status": {
		  "connected": true,
		  "maintenance_mode": false
		}
	  }`[1:]

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ConstructorInfo
		err  error
	}{
		{
			name: "Get constructor succeeds",
			args: args{params: GetParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(constructoGet),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "GET",
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10",
					},
				}),
			}},
			want: &models.ConstructorInfo{
				ConstructorID: ec.String("192.168.44.10"),
				Status: &models.ConstructorHealthStatus{
					Connected:       ec.Bool(true),
					MaintenanceMode: ec.Bool(false),
				},
			},
		},
		{
			name: "Get constructor fails due to API error",
			args: args{params: GetParams{
				ID:     "192.168.44.10",
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/constructors/192.168.44.10",
				Err: errors.New("error"),
			},
		},
		{
			name: "Get constructor fails due to empty parameter validation",
			err: multierror.NewPrefixed("invalid constructor get params",
				apierror.ErrMissingAPI,
				errors.New("id field cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
