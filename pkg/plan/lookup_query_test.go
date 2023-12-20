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

package plan

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestNewReverseLookupQuery(t *testing.T) {
	var elasticsearchRandID = ec.RandomResourceID()
	var kibanaRandID = ec.RandomResourceID()
	var apmRandID = ec.RandomResourceID()
	var appsearchRandID = ec.RandomResourceID()
	var enterprisesearchRandID = ec.RandomResourceID()
	var integrationsserverRandID = ec.RandomResourceID()
	type args struct {
		resourceID string
		kind       string
	}
	tests := []struct {
		name string
		args args
		want *models.SearchRequest
	}{
		{
			name: "creates a lookup query for an Elasticsearch resource",
			args: args{resourceID: elasticsearchRandID, kind: "elasticsearch"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.elasticsearch"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.elasticsearch.id": {Query: &elasticsearchRandID},
					}},
				},
			}},
		},
		{
			name: "creates a lookup query for a Kibana resource",
			args: args{resourceID: kibanaRandID, kind: "kibana"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.kibana"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.kibana.id": {Query: &kibanaRandID},
					}},
				},
			}},
		},
		{
			name: "creates a lookup query for an Apm resource",
			args: args{resourceID: apmRandID, kind: "apm"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.apm"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.apm.id": {Query: &apmRandID},
					}},
				},
			}},
		},
		{
			name: "creates a lookup query for an Appsearch resource",
			args: args{resourceID: appsearchRandID, kind: "appsearch"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.appsearch"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.appsearch.id": {Query: &appsearchRandID},
					}},
				},
			}},
		},
		{
			name: "creates a lookup query for an Enterprise-Search resource",
			args: args{resourceID: enterprisesearchRandID, kind: "enterprise_search"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.enterprise_search"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.enterprise_search.id": {Query: &enterprisesearchRandID},
					}},
				},
			}},
		},
		{
			name: "creates a lookup query for an integrations-server resource",
			args: args{resourceID: integrationsserverRandID, kind: "integrations_server"},
			want: &models.SearchRequest{Query: &models.QueryContainer{
				Nested: &models.NestedQuery{
					Path: ec.String("resources.integrations_server"),
					Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
						"resources.integrations_server.id": {Query: &integrationsserverRandID},
					}},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewReverseLookupQuery(tt.args.resourceID, tt.args.kind); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReverseLookupQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookupByResourceIdQuery(t *testing.T) {
	id := "resource-id"
	result := LookupByResourceIdQuery(id)

	expected := &models.SearchRequest{Query: &models.QueryContainer{
		Bool: &models.BoolQuery{
			MinimumShouldMatch: 1,
			Should: []*models.QueryContainer{
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.elasticsearch"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.elasticsearch.id": {Query: &id},
						}},
					},
				},
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.kibana"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.kibana.id": {Query: &id},
						}},
					},
				},
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.apm"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.apm.id": {Query: &id},
						}},
					},
				},
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.appsearch"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.appsearch.id": {Query: &id},
						}},
					},
				},
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.enterprise_search"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.enterprise_search.id": {Query: &id},
						}},
					},
				},
				{
					Nested: &models.NestedQuery{
						Path: ec.String("resources.integrations_server"),
						Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
							"resources.integrations_server.id": {Query: &id},
						}},
					},
				},
			},
		},
	}}
	assert.EqualValuesf(t, expected, result, "TestLookupByResourceIdQuery() failed")
}
