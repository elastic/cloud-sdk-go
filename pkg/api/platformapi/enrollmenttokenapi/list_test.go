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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	urlError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens",
		Err: errors.New("error"),
	}
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.ListEnrollmentTokenReply
		err  string
	}{
		{
			name: "List Succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: 200,
						Body: mock.NewStructBody(models.ListEnrollmentTokenReply{
							Tokens: []*models.ListEnrollmentTokenElement{
								{TokenID: ec.String("token-1")},
								{TokenID: ec.String("token-2"), Roles: []string{"role"}},
							},
						}),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "GET",
						Path:   "/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens",
					},
				}),
			}},
			want: &models.ListEnrollmentTokenReply{
				Tokens: []*models.ListEnrollmentTokenElement{
					{TokenID: ec.String("token-1")},
					{TokenID: ec.String("token-2"), Roles: []string{"role"}},
				},
			},
		},
		{
			name: "List fails due to API error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New("error"),
				}),
			}},
			err: urlError.Error(),
		},
		{
			name: "List fails due to missing parameters",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid enrollment-token list params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("List() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
