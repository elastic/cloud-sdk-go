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

package depresourceapi

import (
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

const (
	// DefaultEnterpriseSearchRefID is used when the RefID is not specified.
	DefaultEnterpriseSearchRefID = "main-enterprise_search"
)

// NewEnterpriseSearch creates a *models.EnterpriseSearchPayload from the parameters.
// It relies on a simplified single dimension memory size and zone count to
// construct the EnterpriseSearch's topology.
func NewEnterpriseSearch(params NewStateless) (*models.EnterpriseSearchPayload, error) {
	params.fillDefaults(DefaultEnterpriseSearchRefID)
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// When either not set, we obtain the values from the running deployment.
	// Overriding either or both is done at the end of the if.
	if err := getTemplateAndRefID(&params); err != nil {
		return nil, err
	}

	if params.DeploymentTemplateInfo.DeploymentTemplate.Resources.EnterpriseSearch == nil {
		return nil, fmt.Errorf("deployment: the %s template is not configured for Enterprise Search. Please use another template if you wish to start Enterprise Search instances",
			params.TemplateID)
	}

	var clusterTopology = params.DeploymentTemplateInfo.DeploymentTemplate.Resources.EnterpriseSearch[0].Plan.ClusterTopology
	var topology = models.EnterpriseSearchTopologyElement{Size: new(models.TopologySize)}
	if len(clusterTopology) > 0 {
		topology = *clusterTopology[0]
	}
	var payload = newEnterpriseSearchPayload(params, topology)

	return &payload, nil
}

func newEnterpriseSearchPayload(params NewStateless, topology models.EnterpriseSearchTopologyElement) models.EnterpriseSearchPayload {
	if params.Size > 0 {
		topology.Size.Value = ec.Int32(params.Size)
	}
	if params.ZoneCount > 0 {
		topology.ZoneCount = params.ZoneCount
	}

	return models.EnterpriseSearchPayload{
		ElasticsearchClusterRefID: ec.String(params.ElasticsearchRefID),
		DisplayName:               params.Name,
		Region:                    ec.String(params.Region),
		RefID:                     ec.String(params.RefID),
		Plan: &models.EnterpriseSearchPlan{
			EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: params.Version},
			ClusterTopology:  []*models.EnterpriseSearchTopologyElement{&topology},
		},
	}
}
