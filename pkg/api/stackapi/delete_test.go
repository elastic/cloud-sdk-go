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
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
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
			name: "Delete Succeeds",
			args: args{params: DeleteParams{
				Version: "5.6.0",
				Region:  "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       mock.NewStringBody("{}"),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "DELETE",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/stack/versions/5.6.0",
					},
				}),
			}},
		},
		{
			name: "Delete fails due to API error",
			args: args{params: DeleteParams{
				Version: "5.6.0",
				Region:  "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New(`{"error": "some error"}`),
				}),
			}},
			err: &url.Error{
				Op:  "Delete",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/stack/versions/5.6.0",
				Err: errors.New(`{"error": "some error"}`),
			},
		},
		{
			name: "Delete fails due to empty API",
			args: args{params: DeleteParams{
				Version: "5.0.0",
				Region:  "us-east-1",
			}},
			err: multierror.NewPrefixed("invalid stack delete params",
				errors.New("api reference is required for the operation"),
			),
		},
		{
			name: "Delete fails due to empty version",
			args: args{params: DeleteParams{
				API:    new(api.API),
				Region: "us-east-1",
			}},
			err: multierror.NewPrefixed("invalid stack delete params",
				errors.New("version string empty"),
			),
		},
		{
			name: "Delete fails due to empty region",
			args: args{params: DeleteParams{
				API:     new(api.API),
				Version: "5.0.0",
			}},
			err: multierror.NewPrefixed("invalid stack delete params",
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Delete fails due to empty parameters",
			args: args{params: DeleteParams{}},
			err: multierror.NewPrefixed("invalid stack delete params",
				errors.New("api reference is required for the operation"),
				errors.New("version string empty"),
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
