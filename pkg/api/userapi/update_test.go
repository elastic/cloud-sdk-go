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

func TestUpdate(t *testing.T) {
	const successResponse = `{
  "builtin": false,
  "security": {
    "enabled": true,
    "roles": [
      "ece_deployment_viewer"
    ]
  },
  "user_name": "fulgencio"
}`

	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		err     error
	}{
		{
			name: "Update fails due to parameter validation failure",
			args: args{
				params: UpdateParams{},
			},
			wantErr: true,
			err: multierror.NewPrefixed("invalid user params",
				errors.New("update requires a username"),
				apierror.ErrMissingAPI,
			),
		},
		{
			name: "Update fails due to API failure",
			args: args{
				params: UpdateParams{
					UserName: "fulgencio",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_platform_admin"},
					API:      api.NewMock(mock.SampleInternalError()),
				},
			},
			wantErr: true,
			err:     mock.MultierrorInternalError,
		},
		{
			name: "Update succeeds",
			args: args{
				params: UpdateParams{
					UserName: "fulgencio",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_deployment_viewer"},
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(successResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "PATCH",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users/fulgencio",
							Body:   mock.NewStringBody(`{"security":{"password":"supersecretpass","roles":["ece_deployment_viewer"]},"user_name":"fulgencio"}` + "\n"),
						},
					}),
				},
			},
			want: &models.User{
				Builtin:  ec.Bool(false),
				UserName: ec.String("fulgencio"),
				Security: &models.UserSecurity{
					Enabled: ec.Bool(true),
					Roles:   []string{"ece_deployment_viewer"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestUpdateCurrent(t *testing.T) {
	const successResponse = `{
  "builtin": false,
  "security": {
    "enabled": true,
    "roles": [
      "ece_deployment_viewer"
    ]
  }
}`

	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		err     error
	}{
		{
			name: "Update fails due to parameter validation failure (missing API)",
			args: args{
				params: UpdateParams{UserName: "xochitl"},
			},
			wantErr: true,
			err: multierror.NewPrefixed("invalid user params",
				apierror.ErrMissingAPI,
			),
		},
		{
			name: "Update fails due to API failure",
			args: args{
				params: UpdateParams{
					UserName: "xochitl",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_platform_admin"},
					API:      api.NewMock(mock.SampleInternalError()),
				},
			},
			wantErr: true,
			err:     mock.MultierrorInternalError,
		},
		{
			name: "Update succeeds",
			args: args{
				params: UpdateParams{
					UserName: "xochitl",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_deployment_viewer"},
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(successResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "PATCH",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/user",
							Body:   mock.NewStringBody(`{"security":{"password":"supersecretpass","roles":["ece_deployment_viewer"]},"user_name":"xochitl"}` + "\n"),
						},
					}),
				},
			},
			want: &models.User{
				Builtin: ec.Bool(false),
				Security: &models.UserSecurity{
					Enabled: ec.Bool(true),
					Roles:   []string{"ece_deployment_viewer"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateCurrent(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
