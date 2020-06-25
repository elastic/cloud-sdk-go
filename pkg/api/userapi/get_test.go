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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	const getUserResponse = `{
    "user_name": "admin",
    "builtin": true
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		err     error
	}{
		{
			name:    "Get fails due to parameter validation failure",
			args:    args{},
			wantErr: true,
			err: multierror.NewPrefixed("invalid user params",
				apierror.ErrMissingAPI,
				errors.New("username is not specified and is required for this operation"),
			),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					UserName: "hermenelgilda",
					API:      api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				},
			},
			wantErr: true,
			err:     errors.New(`{"error": "some error"}`),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					UserName: "admin",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getUserResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/users/admin",
						},
					}),
				},
			},
			want: &models.User{
				Builtin:  ec.Bool(true),
				UserName: ec.String("admin"),
			},
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

func TestGetCurrent(t *testing.T) {
	const getCurrentResponse = `{
    "user_name": "admin",
    "builtin": true
}`

	type args struct {
		params GetCurrentParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		err     error
	}{
		{
			name:    "Get fails due to parameter validation failure",
			args:    args{},
			wantErr: true,
			err:     errors.New("api reference is required for the operation"),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetCurrentParams{
					API: api.NewMock(mock.SampleInternalError()),
				},
			},
			wantErr: true,
			err:     mock.MultierrorInternalError,
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetCurrentParams{
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getCurrentResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/user",
						},
					}),
				},
			},
			want: &models.User{
				Builtin:  ec.Bool(true),
				UserName: ec.String("admin"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCurrent(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
