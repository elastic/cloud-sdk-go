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

// NewPayloadParams is consumed by NewPayload()
type NewPayloadParams struct {
	*api.API

	// Optional deployment name
	Name string

	// Optional Elastic Stack version, Latest version will be used if empty.
	Version string

	// Required deployment template ID to use to create the deployment.
	DeploymentTemplateID string

	// Required region where the deployment will be created.
	Region string

	// Optionally enable an APM instance.
	ApmEnable bool

	// Optionally enable an AppSearch instance.
	AppsearchEnable bool

	// Optionally enable an Enterprise Search instance.
	EnterpriseSearchEnable bool

	// Do not use. The field will be removed once an API bug has been resolved.
	DeploymentTemplateAsList bool

	// Optional io.Writer where notices will be written.
	Writer io.Writer

	// Optional Elasticsearch plug-ins to be enabled.
	Plugins []string

	// Optional Elasticsearch Topology Elements to compose a complex multi-
	// topology Elasticsearch cluster. Required when ElasticsearchInstance is
	// empty.
	TopologyElements []string

	// Optional ElasticsearchInstance template for single topology element
	// Elasticsearch clusters.
	ElasticsearchInstance InstanceParams

	// Required KibanaInstance template for the Kibana instance.
	KibanaInstance InstanceParams

	// Optional ApmInstance unless AppsearchEnable is set.
	ApmInstance InstanceParams

	// Optional AppsearchInstance unless AppsearchEnable is set.
	AppsearchInstance InstanceParams

	// Optional EnterpriseSearchInstance unless EnterpriseSearchEnable is set.
	EnterpriseSearchInstance InstanceParams
}

// NewPayload creates the payload for a deployment
// // * Auto-discovers the latest Stack version if Version is not specified.
func NewPayload(params NewPayloadParams) (*models.DeploymentCreateRequest, error) {
	res, err := deptemplateapi.Get(deptemplateapi.GetParams{
		API:        params.API,
		TemplateID: params.DeploymentTemplateID,
		Region:     params.Region,
		AsList:     params.DeploymentTemplateAsList,
	})
	if err != nil {
		return nil, err
	}

	// Version Discovery
	version, err := LatestStackVersion(LatestStackVersionParams{
		Writer:  params.Writer,
		API:     params.API,
		Version: params.Version,
		Region:  params.Region,
	})
	if err != nil {
		return nil, err
	}

	esPayload, err := ParseElasticsearchInput(ParseElasticsearchInputParams{
		NewElasticsearchParams: NewElasticsearchParams{
			RefID:                    params.ElasticsearchInstance.RefID,
			Version:                  version,
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
		Version:                  version,
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
			Version:                  version,
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
			Version:                  version,
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
		enterpriseSearchPayload, err := NewEnterpriseSearch(
			NewStateless{
				ElasticsearchRefID:       params.ElasticsearchInstance.RefID,
				API:                      params.API,
				RefID:                    params.EnterpriseSearchInstance.RefID,
				Version:                  version,
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
