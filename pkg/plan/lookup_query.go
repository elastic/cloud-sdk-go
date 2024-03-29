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
	"fmt"
	"github.com/elastic/cloud-sdk-go/pkg/util"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// NewReverseLookupQuery can be used to look up a deployment's ID by specifying
// the resource kind and ID (elasticsearch, 6779ce55fc0646309ef812d007bb2526).
func NewReverseLookupQuery(resourceID, kind string) *models.SearchRequest {
	return &models.SearchRequest{Query: &models.QueryContainer{
		Nested: &models.NestedQuery{
			Path: ec.String(fmt.Sprint("resources.", kind)),
			Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
				fmt.Sprint("resources.", kind, ".id"): {Query: &resourceID},
			}},
		},
	}}
}

// LookupByResourceIdQuery can be used to find a deployment by a resource-id (can be any kind e.g. elasticsearch, kibana, etc.)
// (Builds a query that searches all possible kinds for the resource-id)
func LookupByResourceIdQuery(resourceID string) *models.SearchRequest {
	queries := []*models.QueryContainer{}

	for _, kind := range util.AllKinds {
		queries = append(queries, &models.QueryContainer{
			Nested: &models.NestedQuery{
				Path: ec.String(fmt.Sprint("resources.", kind)),
				Query: &models.QueryContainer{Match: map[string]models.MatchQuery{
					fmt.Sprint("resources.", kind, ".id"): {Query: &resourceID},
				}},
			},
		})
	}

	return &models.SearchRequest{Query: &models.QueryContainer{
		Bool: &models.BoolQuery{
			MinimumShouldMatch: 1,
			Should:             queries,
		},
	}}
}
