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
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCreate(t *testing.T) {
	successResponse := `{
  "builtin": false,
  "security": {
    "enabled": true,
    "roles": [
      "ece_deployment_viewer"
    ]
  },
  "user_name": "bob"
}`

	type args struct {
		params CreateParams
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
		err     error
	}{
		{
			name: "Create fails due to parameter validation failure (missing API)",
			args: args{
				params: CreateParams{
					Email: "hi",
				},
			},
			wantErr: true,
			err: multierror.NewPrefixed("invalid user params",
				errors.New("api reference is required for the operation"),
				errors.New("username is not specified and is required for this operation"),
				errors.New("a password with a minimum of 8 characters is required for this operation"),
				errors.New("a minimum of 1 role is required for this operation"),
				errors.New("hi is not a valid email address format"),
			),
		},
		{
			name: "Create fails due to API failure",
			args: args{
				params: CreateParams{
					UserName: "bob",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_platform_admin"},
					API:      api.NewMock(mock.New404Response(mock.NewStringBody(`{"error": "some error"}`))),
				},
			},
			wantErr: true,
			err:     errors.New(`{"error": "some error"}`),
		},
		{
			name: "Create succeeds",
			args: args{
				params: CreateParams{
					UserName: "bob",
					Password: []byte("supersecretpass"),
					Roles:    []string{"ece_deployment_viewer"},
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(successResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "POST",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/users",
							Body:   mock.NewStringBody(`{"security":{"enabled":true,"password":"supersecretpass","roles":["ece_deployment_viewer"]},"user_name":"bob"}` + "\n"),
						},
					}),
				},
			},
			want: &models.User{
				Builtin:  ec.Bool(false),
				UserName: ec.String("bob"),
				Security: &models.UserSecurity{
					Enabled: ec.Bool(true),
					Roles:   []string{"ece_deployment_viewer"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if err != nil && !assert.Equal(t, tt.err.Error(), err.Error()) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
