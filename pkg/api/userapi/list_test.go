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
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	const listUsersResponse = `{
  "users": [{
    "user_name": "admin",
    "builtin": true
  }, {
    "user_name": "readonly",
    "builtin": true
  }]
}`

	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.UserList
		err  string
	}{
		{
			name: "List fails due to parameter validation failure (missing API)",
			err:  apierror.ErrMissingAPI.Error(),
		},
		{
			name: "List fails due to API failure",
			args: args{params: ListParams{
				API: api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "List succeeds",
			args: args{
				params: ListParams{
					IncludeDisabled: true,
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(listUsersResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users",
							Query: url.Values{
								"include_disabled": {"true"},
							},
						},
					}),
				},
			},
			want: &models.UserList{
				Users: []*models.User{
					{
						Builtin:  ec.Bool(true),
						UserName: ec.String("admin"),
					},
					{
						Builtin:  ec.Bool(true),
						UserName: ec.String("readonly"),
					},
				},
			},
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
