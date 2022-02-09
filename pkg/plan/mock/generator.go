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

package planmock

import (
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// GenerateConfig is a helper used to create a DeploymentGetResponse to be
// used for tests. It provides a generic implementation so all deployment
// resoource types can be generated a plan log and tracked.
type GenerateConfig struct {
	// If omitted, a random ID will be auto-generated.
	ID string

	Elasticsearch      []GeneratedResourceConfig
	EnterpriseSearch   []GeneratedResourceConfig
	Kibana             []GeneratedResourceConfig
	Apm                []GeneratedResourceConfig
	IntegrationsServer []GeneratedResourceConfig
	Appsearch          []GeneratedResourceConfig
}

// GeneratedResourceConfig is used to construct a deployment resource plan
// with the plan log steps specified.
type GeneratedResourceConfig struct {
	// If omitted, a random ID will be auto-generated.
	ID string

	// If omitted, it'll default to "main-<kind>"
	RefID string

	// Current plan logs.
	CurrentLog []*models.ClusterPlanStepInfo

	// Pending plan logs.
	PendingLog []*models.ClusterPlanStepInfo

	// Histoory plan logs.
	HistoryLog []*models.ClusterPlanStepInfo
}

// Generate creates a DeploymentGetResponse to mock the plan tracker. See the
// configuration options in the GenerateConfig struct.
func Generate(cfg GenerateConfig) *models.DeploymentGetResponse {
	var id = cfg.ID
	if id == "" {
		id = ec.RandomResourceID()
	}
	return &models.DeploymentGetResponse{
		ID: ec.String(id),
		Resources: &models.DeploymentResources{
			Apm:                generateApmResourceInfo(cfg.Apm),
			IntegrationsServer: generateIntegrationsServerResourceInfo(cfg.IntegrationsServer),
			Kibana:             generateKibanaResourceInfo(cfg.Kibana),
			Elasticsearch:      generateElasticsearchResourceInfo(cfg.Elasticsearch),
			Appsearch:          generateAppSearchResourceInfo(cfg.Appsearch),
			EnterpriseSearch:   generateEnterpriseSearchResourceInfo(cfg.EnterpriseSearch),
		},
	}
}

func generateApmResourceInfo(c []GeneratedResourceConfig) []*models.ApmResourceInfo {
	var resInfo = make([]*models.ApmResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.ApmResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-apm")
		}

		info.Info = &models.ApmInfo{PlanInfo: &models.ApmPlansInfo{
			Pending: &models.ApmPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.ApmPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.ApmPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}

func generateIntegrationsServerResourceInfo(c []GeneratedResourceConfig) []*models.IntegrationsServerResourceInfo {
	var resInfo = make([]*models.IntegrationsServerResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.IntegrationsServerResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-integrations_server")
		}

		info.Info = &models.IntegrationsServerInfo{PlanInfo: &models.IntegrationsServerPlansInfo{
			Pending: &models.IntegrationsServerPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.IntegrationsServerPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.IntegrationsServerPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}

func generateAppSearchResourceInfo(c []GeneratedResourceConfig) []*models.AppSearchResourceInfo {
	var resInfo = make([]*models.AppSearchResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.AppSearchResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-appsearch")
		}

		info.Info = &models.AppSearchInfo{PlanInfo: &models.AppSearchPlansInfo{
			Pending: &models.AppSearchPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.AppSearchPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.AppSearchPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}

func generateEnterpriseSearchResourceInfo(c []GeneratedResourceConfig) []*models.EnterpriseSearchResourceInfo {
	var resInfo = make([]*models.EnterpriseSearchResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.EnterpriseSearchResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-enterprise_search")
		}

		info.Info = &models.EnterpriseSearchInfo{PlanInfo: &models.EnterpriseSearchPlansInfo{
			Pending: &models.EnterpriseSearchPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.EnterpriseSearchPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.EnterpriseSearchPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}

func generateKibanaResourceInfo(c []GeneratedResourceConfig) []*models.KibanaResourceInfo {
	var resInfo = make([]*models.KibanaResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.KibanaResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-kibana")
		}

		info.Info = &models.KibanaClusterInfo{PlanInfo: &models.KibanaClusterPlansInfo{
			Pending: &models.KibanaClusterPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.KibanaClusterPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.KibanaClusterPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}

func generateElasticsearchResourceInfo(c []GeneratedResourceConfig) []*models.ElasticsearchResourceInfo {
	var resInfo = make([]*models.ElasticsearchResourceInfo, 0, len(c))
	for _, gen := range c {
		var info models.ElasticsearchResourceInfo
		info.ID = &gen.ID
		if gen.ID == "" {
			info.ID = ec.String(ec.RandomResourceID())
		}
		info.RefID = &gen.RefID
		if gen.RefID == "" {
			info.RefID = ec.String("main-elasticsearch")
		}

		info.Info = &models.ElasticsearchClusterInfo{PlanInfo: &models.ElasticsearchClusterPlansInfo{
			Pending: &models.ElasticsearchClusterPlanInfo{PlanAttemptLog: gen.PendingLog},
			Current: &models.ElasticsearchClusterPlanInfo{PlanAttemptLog: gen.CurrentLog},
			History: []*models.ElasticsearchClusterPlanInfo{{PlanAttemptLog: gen.HistoryLog}},
		}}

		resInfo = append(resInfo, &info)
	}

	return resInfo
}
