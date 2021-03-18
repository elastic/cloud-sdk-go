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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestCreate(t *testing.T) {
	type args struct {
		params CreateParams
	}
	tests := []struct {
		name string
		args args
		want string
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid comment create params",
				errors.New("region not specified and is required for this operation"),
				errors.New("api reference is required for the operation"),
				errors.New("resource type is required for this operation"),
				errors.New("resource id is required for this operation"),
				errors.New("message is required for this operation"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: CreateParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				Message:      "some-comment-message",
				API: api.NewMock(mock.New201ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id",
						Body:   mock.NewStringBody(`{"message":"some-comment-message"}` + "\n"),
					},
					mock.NewStringBody(`{"id": "random-generated-id"}`),
				)),
			}},
			want: "random-generated-id",
		},
		{
			name: "fails when api returns an error",
			args: args{params: CreateParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				Message:      "some-comment-message",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id",
						Body:   mock.NewStringBody(`{"message":"some-comment-message"}` + "\n"),
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
