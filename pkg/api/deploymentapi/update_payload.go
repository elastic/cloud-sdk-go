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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// NewUpdateRequest generates a DeploymentUpdateRequest from a GetResponse.
func NewUpdateRequest(res *models.DeploymentGetResponse) *models.DeploymentUpdateRequest {
	if res == nil {
		return nil
	}

	var esRefID string
	var req = models.DeploymentUpdateRequest{
		Name:         *res.Name,
		Resources:    &models.DeploymentUpdateResources{},
		PruneOrphans: ec.Bool(false),
	}

	for _, r := range res.Resources.Elasticsearch {
		if resource, rID := parseElasticsearchGetResponse(r); resource != nil {
			req.Resources.Elasticsearch = append(req.Resources.Elasticsearch,
				resource,
			)
			esRefID = rID
		}
	}

	for _, r := range res.Resources.Kibana {
		if resource := parseKibanaGetResponse(r, esRefID); resource != nil {
			req.Resources.Kibana = append(req.Resources.Kibana,
				resource,
			)
		}
	}

	for _, r := range res.Resources.Apm {
		if resource := parseApmGetResponse(r, esRefID); resource != nil {
			req.Resources.Apm = append(req.Resources.Apm,
				resource,
			)
		}
	}

	for _, r := range res.Resources.Appsearch {
		if resource := parseAppSearchGetResponse(r, esRefID); resource != nil {
			req.Resources.Appsearch = append(req.Resources.Appsearch,
				resource,
			)
		}
	}

	for _, r := range res.Resources.EnterpriseSearch {
		if resource := parseEnterpriseSearchGetResponse(r, esRefID); resource != nil {
			req.Resources.EnterpriseSearch = append(req.Resources.EnterpriseSearch,
				resource,
			)
		}
	}

	if res.Settings != nil && res.Settings.Observability != nil {
		req.Settings = &models.DeploymentUpdateSettings{
			Observability: res.Settings.Observability,
		}
	}

	return &req
}

func parseElasticsearchGetResponse(r *models.ElasticsearchResourceInfo) (payload *models.ElasticsearchPayload, refID string) {
	plan := r.Info.PlanInfo.Current
	if plan == nil || plan.Plan == nil {
		return nil, ""
	}

	if r.Info.Settings != nil {
		r.Info.Settings.Metadata = nil
	}

	var ct = make([]*models.ElasticsearchClusterTopologyElement, 0, len(plan.Plan.ClusterTopology))
	for _, t := range plan.Plan.ClusterTopology {
		if t.MemoryPerNode > 0 || !nilOZeroToplogySize(t.Size) {
			ct = append(ct, t)
		}
	}
	plan.Plan.ClusterTopology = ct

	return &models.ElasticsearchPayload{
		DisplayName: *r.Info.ClusterName,
		RefID:       r.RefID,
		Region:      r.Region,
		Plan:        plan.Plan,
		Settings:    r.Info.Settings,
	}, *r.RefID
}

func parseKibanaGetResponse(r *models.KibanaResourceInfo, esRefID string) *models.KibanaPayload {
	plan := r.Info.PlanInfo.Current
	if plan == nil || plan.Plan == nil {
		return nil
	}

	if r.Info.Settings != nil {
		r.Info.Settings.Metadata = nil
	}

	var ct = make([]*models.KibanaClusterTopologyElement, 0, len(plan.Plan.ClusterTopology))
	for _, t := range plan.Plan.ClusterTopology {
		if t.MemoryPerNode > 0 || !nilOZeroToplogySize(t.Size) {
			ct = append(ct, t)
		}
	}

	plan.Plan.ClusterTopology = ct
	return &models.KibanaPayload{
		ElasticsearchClusterRefID: &esRefID,
		DisplayName:               *r.Info.ClusterName,
		RefID:                     r.RefID,
		Region:                    r.Region,
		Plan:                      plan.Plan,
		Settings:                  r.Info.Settings,
	}
}

func parseApmGetResponse(r *models.ApmResourceInfo, esRefID string) *models.ApmPayload {
	plan := r.Info.PlanInfo.Current
	if plan == nil || plan.Plan == nil {
		return nil
	}

	if r.Info.Settings != nil {
		r.Info.Settings.Metadata = nil
	}

	var ct = make([]*models.ApmTopologyElement, 0, len(plan.Plan.ClusterTopology))
	for _, t := range plan.Plan.ClusterTopology {
		if t.Size != nil && t.Size.Value != nil && *t.Size.Value > 0 {
			ct = append(ct, t)
		}
	}

	plan.Plan.ClusterTopology = ct
	return &models.ApmPayload{
		ElasticsearchClusterRefID: &esRefID,
		DisplayName:               *r.Info.Name,
		RefID:                     r.RefID,
		Region:                    r.Region,
		Plan:                      plan.Plan,
		Settings:                  r.Info.Settings,
	}
}

func parseAppSearchGetResponse(r *models.AppSearchResourceInfo, esRefID string) *models.AppSearchPayload {
	plan := r.Info.PlanInfo.Current
	if plan == nil || plan.Plan == nil {
		return nil
	}

	if r.Info.Settings != nil {
		r.Info.Settings.Metadata = nil
	}

	var ct = make([]*models.AppSearchTopologyElement, 0, len(plan.Plan.ClusterTopology))
	for _, t := range plan.Plan.ClusterTopology {
		if t.Size != nil && t.Size.Value != nil && *t.Size.Value > 0 {
			ct = append(ct, t)
		}
	}

	plan.Plan.ClusterTopology = ct
	return &models.AppSearchPayload{
		ElasticsearchClusterRefID: &esRefID,
		DisplayName:               *r.Info.Name,
		RefID:                     r.RefID,
		Region:                    r.Region,
		Plan:                      plan.Plan,
		Settings:                  r.Info.Settings,
	}
}

func parseEnterpriseSearchGetResponse(r *models.EnterpriseSearchResourceInfo, esRefID string) *models.EnterpriseSearchPayload {
	plan := r.Info.PlanInfo.Current
	if plan == nil || plan.Plan == nil {
		return nil
	}

	if r.Info.Settings != nil {
		r.Info.Settings.Metadata = nil
	}

	var ct = make([]*models.EnterpriseSearchTopologyElement, 0, len(plan.Plan.ClusterTopology))
	for _, t := range plan.Plan.ClusterTopology {
		if t.Size != nil && t.Size.Value != nil && *t.Size.Value > 0 {
			ct = append(ct, t)
		}
	}

	plan.Plan.ClusterTopology = ct
	return &models.EnterpriseSearchPayload{
		ElasticsearchClusterRefID: &esRefID,
		DisplayName:               *r.Info.Name,
		RefID:                     r.RefID,
		Region:                    r.Region,
		Plan:                      plan.Plan,
		Settings:                  r.Info.Settings,
	}
}

func nilOZeroToplogySize(ts *models.TopologySize) bool {
	return ts == nil || ts.Value == nil || *ts.Value == 0
}
