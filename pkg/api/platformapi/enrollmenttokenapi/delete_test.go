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

package enrollmenttokenapi

import (
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestDelete(t *testing.T) {
	urlError := url.Error{
		Op:  "Delete",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens/atoken",
		Err: errors.New("error"),
	}
	type args struct {
		params DeleteParams
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "Create fails due to missing token",
			args: args{params: DeleteParams{
				API: new(api.API),
			}},
			err: multierror.NewPrefixed("invalid enrollment-token delete params",
				errors.New("token cannot be empty"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "Create fails due to missing API",
			args: args{params: DeleteParams{
				Token: "token",
			}},
			err: multierror.NewPrefixed("invalid enrollment-token delete params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "Delete fails due to API error",
			args: args{params: DeleteParams{
				Token:  "atoken",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New("error"),
				}),
			}},
			err: urlError.Error(),
		},
		{
			name: "Delete Succeeds with persistent token",
			args: args{params: DeleteParams{
				Token:  "atoken",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: 200,
						Body:       mock.NewStructBody(struct{}{}),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "DELETE",
						Path:   "/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens/atoken",
					},
				}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Delete(tt.args.params)
			if !assert.EqualError(t, err, tt.err) {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
