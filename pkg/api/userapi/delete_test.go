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

package userapi

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

func TestDelete(t *testing.T) {
	type args struct {
		params DeleteParams
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name:    "Delete fails due to parameter validation failure",
			args:    args{},
			wantErr: true,
			err: multierror.NewPrefixed("invalid user params",
				apierror.ErrMissingAPI,
				errors.New("username is not specified and is required for this operation"),
			),
		},
		{
			name: "Delete fails due to API failure",
			args: args{
				params: DeleteParams{
					UserName: "user bob",
					API: api.NewMock(mock.Response{Response: http.Response{
						Body:       mock.NewStringBody(`{"error": "some error"}`),
						StatusCode: 400,
					}}),
				},
			},
			wantErr: true,
			err:     errors.New(`{"error": "some error"}`),
		},
		{
			name: "Delete succeeds",
			args: args{
				params: DeleteParams{
					UserName: "userbob",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(""),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "DELETE",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users/userbob",
						},
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Delete(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
