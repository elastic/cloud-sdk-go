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
	"errors"
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/platformapi/configurationtemplateapi"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

const (
	// DefaultTemplateID is used when there's no template ID specified in
	// the request.
	DefaultTemplateID = "default"

	// DefaultDataSize defines the default node size for data nodes when not
	// specified.
	DefaultDataSize = 4096

	// DefaultDataZoneCount defines the default number of zones a deployment
	// spans.
	DefaultDataZoneCount = 1

	// DefaultElasticsearchRefID is used when the RefID is not specified.
	DefaultElasticsearchRefID = "main-elasticsearch"
)

// NewElasticsearchParams is consumed by NewElasticsearch.
type NewElasticsearchParams struct {
	*api.API

	// Optional region name. Defaults to
	Region string

	// Optional name. If not specified it defaults to the autogeneratd ID.
	Name string

	// Optional RefID for the deployment resource.
	RefID string

	// Required: Version is the Elasticsearch Version.
	Version string
	Plugins []string

	// Required Deployment Template ID.
	TemplateID string

	// Topology settings
	Topology []ElasticsearchTopologyElement
}

func (params *NewElasticsearchParams) fillDefaults() {
	if params.TemplateID == "" {
		params.TemplateID = DefaultTemplateID
	}

	if params.RefID == "" {
		params.RefID = DefaultElasticsearchRefID
	}

	if params.Topology == nil {
		params.Topology = DefaultTopology
	}
}

// Validate ensures the parameters are usable by the consuming function.
func (params *NewElasticsearchParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment elasticsearch")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Region == "" {
		merr = merr.Append(errors.New("region cannot be empty"))
	}

	if params.Version == "" {
		merr = merr.Append(errors.New("version cannot be empty"))
	}

	for i := range params.Topology {
		if err := params.Topology[i].Validate(); err != nil {
			merr = merr.Append(
				multierror.NewPrefixed(fmt.Sprintf("element[%d]", i), err),
			)
		}
	}

	return merr.ErrorOrNil()
}

// NewElasticsearch creates a *models.ElasticsearchPayload from the parameters.
// It relies on a simplified definition of the full API ElasticsearchPayload.
// See BuildElasticsearchTopology for more information on how the construction
// of ElasticsearchPayload works.
func NewElasticsearch(params NewElasticsearchParams) (*models.ElasticsearchPayload, error) {
	params.fillDefaults()
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := configurationtemplateapi.GetTemplate(configurationtemplateapi.GetTemplateParams{
		API:                params.API,
		ID:                 params.TemplateID,
		Region:             params.Region,
		Format:             "cluster",
		ShowInstanceConfig: true,
	})
	if err != nil {
		return nil, err
	}

	var payload = newElasticsearchPayload(params)
	payload.Plan.ClusterTopology, err = BuildElasticsearchTopology(
		BuildElasticsearchTopologyParams{
			ClusterTopology: res.ClusterTemplate.Plan.ClusterTopology,
			TemplateID:      params.TemplateID,
			Topology:        params.Topology,
		},
	)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func newElasticsearchPayload(params NewElasticsearchParams) models.ElasticsearchPayload {
	return models.ElasticsearchPayload{
		DisplayName: params.Name,
		Region:      ec.String(params.Region),
		RefID:       ec.String(params.RefID),
		Plan: &models.ElasticsearchClusterPlan{
			Elasticsearch: &models.ElasticsearchConfiguration{
				Version: params.Version,
			},
			DeploymentTemplate: &models.DeploymentTemplateReference{
				ID: ec.String(params.TemplateID),
			},
		},
	}
}
