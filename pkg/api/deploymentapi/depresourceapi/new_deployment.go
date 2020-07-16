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
	"io"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deptemplateapi"
	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// InstanceParams holds the common instance fields for a resource
type InstanceParams struct {
	Size      int32
	ZoneCount int32
	RefID     string
}

// NewParams is consumed by New()
type NewParams struct {
	*api.API

	Name                     string
	Version                  string
	DeploymentTemplateID     string
	Region                   string
	ApmEnable                bool
	AppsearchEnable          bool
	EnterpriseSearchEnable   bool
	Writer                   io.Writer
	Plugins                  []string
	TopologyElements         []string
	ElasticsearchInstance    InstanceParams
	KibanaInstance           InstanceParams
	ApmInstance              InstanceParams
	AppsearchInstance        InstanceParams
	EnterpriseSearchInstance InstanceParams
}

// New creates the payload for a deployment
func New(params NewParams) (*models.DeploymentCreateRequest, error) {
	res, err := deptemplateapi.Get(deptemplateapi.GetParams{
		API:        params.API,
		TemplateID: params.DeploymentTemplateID,
		Region:     params.Region,
	})
	if err != nil {
		return nil, err
	}

	esPayload, err := ParseElasticsearchInput(ParseElasticsearchInputParams{
		NewElasticsearchParams: NewElasticsearchParams{
			RefID:                    params.ElasticsearchInstance.RefID,
			Version:                  params.Version,
			Plugins:                  params.Plugins,
			Region:                   params.Region,
			TemplateID:               params.DeploymentTemplateID,
			DeploymentTemplateInfoV2: res,
		},
		API:              params.API,
		Size:             params.ElasticsearchInstance.Size,
		ZoneCount:        params.ElasticsearchInstance.ZoneCount,
		Writer:           params.Writer,
		TopologyElements: params.TopologyElements,
	})
	if err != nil {
		return nil, err
	}

	kibanaPayload, err := NewKibana(NewStateless{
		ElasticsearchRefID:       params.ElasticsearchInstance.RefID,
		API:                      params.API,
		RefID:                    params.KibanaInstance.RefID,
		Version:                  params.Version,
		Region:                   params.Region,
		TemplateID:               params.DeploymentTemplateID,
		Size:                     params.KibanaInstance.Size,
		ZoneCount:                params.KibanaInstance.ZoneCount,
		DeploymentTemplateInfoV2: res,
	})
	if err != nil {
		return nil, err
	}

	resources := models.DeploymentCreateResources{
		Elasticsearch: []*models.ElasticsearchPayload{esPayload},
		Kibana:        []*models.KibanaPayload{kibanaPayload},
	}

	if params.ApmEnable {
		apmPayload, err := NewApm(NewStateless{
			ElasticsearchRefID:       params.ElasticsearchInstance.RefID,
			API:                      params.API,
			RefID:                    params.ApmInstance.RefID,
			Version:                  params.Version,
			Region:                   params.Region,
			TemplateID:               params.DeploymentTemplateID,
			Size:                     params.ApmInstance.Size,
			ZoneCount:                params.ApmInstance.ZoneCount,
			DeploymentTemplateInfoV2: res,
		})
		if err != nil {
			return nil, err
		}

		resources.Apm = []*models.ApmPayload{apmPayload}
	}

	if params.AppsearchEnable {
		appsearchPayload, err := NewAppSearch(NewStateless{
			ElasticsearchRefID:       params.ElasticsearchInstance.RefID,
			API:                      params.API,
			RefID:                    params.AppsearchInstance.RefID,
			Version:                  params.Version,
			Region:                   params.Region,
			TemplateID:               params.DeploymentTemplateID,
			Size:                     params.AppsearchInstance.Size,
			ZoneCount:                params.AppsearchInstance.ZoneCount,
			DeploymentTemplateInfoV2: res,
		})
		if err != nil {
			return nil, err
		}

		resources.Appsearch = []*models.AppSearchPayload{appsearchPayload}
	}

	if params.EnterpriseSearchEnable {
		enterpriseSearchPayload, err := NewEnterpriseSearch(NewStateless{
			ElasticsearchRefID:       params.ElasticsearchInstance.RefID,
			API:                      params.API,
			RefID:                    params.EnterpriseSearchInstance.RefID,
			Version:                  params.Version,
			Region:                   params.Region,
			TemplateID:               params.DeploymentTemplateID,
			Size:                     params.EnterpriseSearchInstance.Size,
			ZoneCount:                params.EnterpriseSearchInstance.ZoneCount,
			DeploymentTemplateInfoV2: res,
		})
		if err != nil {
			return nil, err
		}

		resources.EnterpriseSearch = []*models.EnterpriseSearchPayload{enterpriseSearchPayload}
	}

	payload := models.DeploymentCreateRequest{
		Name:      params.Name,
		Resources: &resources,
	}

	return &payload, nil
}
