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

package userauthadminapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestDeleteKey(t *testing.T) {
	type args struct {
		params DeleteKeyParams
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid user auth admin params",
				apierror.ErrMissingAPI,
				errors.New("key id is not specified and is required for this operation"),
				errors.New("user id is not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "fails due to API error",
			args: args{params: DeleteKeyParams{
				API: api.NewMock(mock.NewErrorResponse(404, mock.APIError{
					Code: "key.not_found", Message: "key not found",
				})),
				ID:     "somekey",
				UserID: "someid",
			}},
			err: multierror.NewPrefixed("api error",
				errors.New("key.not_found: key not found"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: DeleteKeyParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "DELETE",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users/someid/auth/keys/somekey",
						},
						mock.NewStringBody("")),
				),
				ID:     "somekey",
				UserID: "someid",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteKey(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
		})
	}
}
