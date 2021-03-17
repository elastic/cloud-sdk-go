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
	"github.com/go-openapi/strfmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestList(t *testing.T) {
	listResponse := `{
  "values": [
    {
      "comment": {
        "id": "93420f71ca474b79aa4bc2aaa7f37e21",
        "message": "some message",
        "user_id": "root"
      },
      "metadata": {
        "created_time": "2021-03-17T13:05:06.958Z",
        "modified_time": "2021-03-17T13:05:06.958Z",
        "version": "8|14"
      }
    }
  ]
}
`

	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.CommentsWithMetas
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid comment list params",
				errors.New("region not specified and is required for this operation"),
				errors.New("api reference is required for the operation"),
				errors.New("resource type is required"),
				errors.New("resource id is required"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: ListParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id",
					},
					mock.NewStringBody(listResponse),
				)),
			}},
			want: &models.CommentsWithMetas{
				Values: []*models.CommentWithMeta{
					{
						Comment: &models.Comment{
							ID:      ec.String("93420f71ca474b79aa4bc2aaa7f37e21"),
							Message: ec.String("some message"),
							UserID:  ec.String("root"),
						},
						Metadata: &models.Metadata{
							CreatedTime:  parseDateTime("2021-03-17T13:05:06.958Z"),
							ModifiedTime: parseDateTime("2021-03-17T13:05:06.958Z"),
							Version:      ec.String("8|14"),
						},
					},
				},
			},
		},
		{
			name: "fails when api returns an error",
			args: args{params: ListParams{
				Region:       "us-east-1",
				ResourceID:   "some-resource-id",
				ResourceType: "some-resource-type",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/comments/some-resource-type/some-resource-id",
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func parseDateTime(str string) *strfmt.DateTime {
	t, _ := strfmt.ParseDateTime(str)
	return &t
}
