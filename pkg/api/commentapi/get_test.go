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

package commentapi

import (
	"errors"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestGet(t *testing.T) {
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.Comment
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid comment get params",
				errors.New("region not specified and is required for this operation"),
				errors.New("api reference is required for the operation"),
				errors.New("resource type is required"),
				errors.New("resource id is required"),
				errors.New("comment id is required"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: GetParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				CommentID:    "some-comment-id",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id/some-comment-id",
					},
					mock.NewStringBody(`{"id": "random-generated-id", "user_id": "lord-of-comments", "message":"decrypted"}`),
				)),
			}},
			want: &models.Comment{
				ID:      ec.String("random-generated-id"),
				Message: ec.String("decrypted"),
				UserID:  ec.String("lord-of-comments"),
			},
		},
		{
			name: "fails when api returns an error",
			args: args{params: GetParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				CommentID:    "some-comment-id",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id/some-comment-id",
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
