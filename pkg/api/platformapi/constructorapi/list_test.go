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

func TestList(t *testing.T) {
	var constructorListSuccess = `
	{
		"constructors": [
		  {
			"constructor_id": "192.168.44.10",
			"status": {
			  "connected": true,
			  "maintenance_mode": false
			}
		  }
		]
	}`[1:]
	urlError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/constructors",
		Err: errors.New("error"),
	}
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.ConstructorOverview
		err  string
	}{
		{
			name: "Constructor list succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(constructorListSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "GET",
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/constructors",
					},
				}),
			}},
			want: &models.ConstructorOverview{
				Constructors: []*models.ConstructorInfo{
					{
						ConstructorID: ec.String("192.168.44.10"),
						Status: &models.ConstructorHealthStatus{
							Connected:       ec.Bool(true),
							MaintenanceMode: ec.Bool(false),
						},
					},
				},
			},
		},
		{
			name: "Constructor list fails due to API error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: urlError.Error(),
		},
		{
			name: "Constructor list fails due to parameter validation",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid constructor list params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
