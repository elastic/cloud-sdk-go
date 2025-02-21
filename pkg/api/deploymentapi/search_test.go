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

package deploymentapi

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestSearch(t *testing.T) {
	type args struct {
		params SearchParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentsSearchResponse
		err  string
	}{
		{
			name: "fails on parameter validation",
			err: multierror.NewPrefixed("deployment search",
				apierror.ErrMissingAPI,
				errors.New("request cannot be empty"),
			).Error(),
		},
		{
			name: "fails on API error",
			args: args{params: SearchParams{
				API:     api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				Request: &models.SearchRequest{},
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Succeeds",
			args: args{params: SearchParams{
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/_search",
						Query:  url.Values{},
						Body:   mock.NewStructBody(models.SearchRequest{}),
					},
					mock.NewStructBody(models.DeploymentsSearchResponse{
						Deployments: []*models.DeploymentSearchResponse{
							{ID: ec.String("123")},
						},
					}))),
				Request: &models.SearchRequest{},
			}},
			want: &models.DeploymentsSearchResponse{Deployments: []*models.DeploymentSearchResponse{
				{ID: ec.String("123")},
			}},
		},
		{
			name: "Adds minimal metadata",
			args: args{params: SearchParams{
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/_search",
						Query: url.Values{
							"minimal_metadata": {"id,name"},
						},
						Body: mock.NewStructBody(models.SearchRequest{}),
					},
					mock.NewStringBody(`
{
  "return_count" : 2,
  "match_count" : 8642,
  "deployments" : [
  ],
  "minimal_metadata" : [
    {
      "name" : "My deployment",
      "id" : "d5a14995cd014268819e004337fbcd83"
    },
    {
      "name" : "Another deployment",
      "id" : "4c9f5af7c2a64210bdbb88a1fdd57680"
    }
  ],
  "cursor" : "dY7RDoIwDEX/pc8LARQx/IohS4USlmyga0lUsn93jcY3n3rTe9t7dmDCOMzQ7V9lcRKK0F3qojRVeyyr+nRuDm1b9wbYvQi6qikN3DeKTz27rqvXGTaWfLbDQiw06uqXoYdjYVWTI589iMTrFgfigjyyuOHTXrgRUjJwQ5n/pzSSC1EU12t7SplOCeyyKkWf9EtwiwvobSBBO6JgdiA3GFgwEPTpDQ=="
}`),
				)),
				Request:         &models.SearchRequest{},
				MinimalMetadata: []string{"id", "name"},
			}},
			want: &models.DeploymentsSearchResponse{
				ReturnCount: ec.Int32(2),
				MatchCount:  8642,
				Deployments: []*models.DeploymentSearchResponse{},
				MinimalMetadata: []interface{}{
					map[string]interface{}{
						"name": "My deployment",
						"id":   "d5a14995cd014268819e004337fbcd83",
					},
					map[string]interface{}{
						"name": "Another deployment",
						"id":   "4c9f5af7c2a64210bdbb88a1fdd57680",
					},
				},
				Cursor: "dY7RDoIwDEX/pc8LARQx/IohS4USlmyga0lUsn93jcY3n3rTe9t7dmDCOMzQ7V9lcRKK0F3qojRVeyyr+nRuDm1b9wbYvQi6qikN3DeKTz27rqvXGTaWfLbDQiw06uqXoYdjYVWTI589iMTrFgfigjyyuOHTXrgRUjJwQ5n/pzSSC1EU12t7SplOCeyyKkWf9EtwiwvobSBBO6JgdiA3GFgwEPTpDQ==",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
