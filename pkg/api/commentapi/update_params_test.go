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

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/stretchr/testify/assert"
)

func TestUpdateParams_Validate(t *testing.T) {
	tests := []struct {
		name   string
		params UpdateParams
		err    error
	}{
		{
			name:   "should return all validation errors",
			params: UpdateParams{},
			err: multierror.NewPrefixed("invalid comment update params",
				errors.New("region not specified and is required for this operation"),
				errors.New("api reference is required for the operation"),
				errors.New("resource type is required for this operation"),
				errors.New("resource id is required for this operation"),
				errors.New("message is required for this operation"),
				errors.New("comment id is required for this operation"),
			),
		},
		{
			name: "should return no validation errors",
			params: UpdateParams{
				API:          &api.API{},
				ResourceType: "some-type",
				ResourceID:   "some-id",
				Message:      "some-message",
				Region:       "some-region",
				CommentID:    "some-comment-id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if err != nil && tt.err == nil {
				t.Fatalf("expected no errors but got %+v", err)
			}
			if err == nil && tt.err != nil {
				t.Fatalf("expected errors %+v but got no errors", tt.err)
			}
			if err != nil && tt.err != nil {
				assert.EqualError(t, err, tt.err.Error())
			}
		})
	}
}
