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
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestDelete(t *testing.T) {
	type args struct {
		params DeleteParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Delete succeeds",
			args: args{params: DeleteParams{
				Region: "us-east-1",
				ID:     "data.highstorage",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(getInstanceConfigsSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "DELETE",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances/data.highstorage",
					},
				}),
			}},
		},
		{
			name: "Delete succeeds on kibana ID",
			args: args{params: DeleteParams{
				Region: "us-east-1",
				ID:     "kibana",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(getInstanceConfigsSuccessKibana),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "DELETE",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/configuration/instances/kibana",
					},
				}),
			}},
		},
		{
			name: "Delete fails on API error",
			args: args{params: DeleteParams{
				Region: "us-east-1",
				ID:     "kibana",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "Delete fails on parameter validation failure",
			err: multierror.NewPrefixed("invalid instance config delete params",
				apierror.ErrMissingAPI,
				errors.New("id must not be empty"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
