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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestListKeys(t *testing.T) {
	var listKeysAllUsers = models.APIKeysResponse{
		Keys: []*models.APIKeyResponse{
			{ID: ec.String("10"), UserID: "2"},
			{ID: ec.String("11"), UserID: "3"},
			{ID: ec.String("12"), UserID: "4"},
		},
	}
	var listKeysSomeIDUser = models.APIKeysResponse{
		Keys: []*models.APIKeyResponse{
			{ID: ec.String("10"), UserID: "someid"},
			{ID: ec.String("11"), UserID: "someid"},
			{ID: ec.String("12"), UserID: "someid"},
		},
	}
	type args struct {
		params ListKeysParams
	}
	tests := []struct {
		name string
		args args
		want *models.APIKeysResponse
		err  string
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid user auth admin params",
				apierror.ErrMissingAPI,
				errors.New("one of user id or the all bool set to true must be specified for this operation"),
			).Error(),
		},
		{
			name: "fails due to parameter validation, both all and userid specified",
			args: args{params: ListKeysParams{
				All: true, UserID: "someid",
			}},
			err: multierror.NewPrefixed("invalid user auth admin params",
				apierror.ErrMissingAPI,
				errors.New("user id must not be specified if the all bool is set to true"),
			).Error(),
		},
		{
			name: "fails due to API error on all call",
			args: args{params: ListKeysParams{
				API: api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				All: true,
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "fails due to API error on user call",
			args: args{params: ListKeysParams{
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				UserID: "someid",
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "succeeds listing keys for multiple users",
			args: args{params: ListKeysParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users/auth/keys/_all",
						},
						mock.NewStructBody(listKeysAllUsers),
					),
				),
				All: true,
			}},
			want: &listKeysAllUsers,
		},
		{
			name: "succeeds listing keys for one user",
			args: args{params: ListKeysParams{
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(listKeysSomeIDUser)),
				),
				UserID: "someid",
			}},
			want: &listKeysSomeIDUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListKeys(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
