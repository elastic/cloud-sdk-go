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
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var systemAlerts = ec.Int64(5)

// GetParams is consumed by get resource functions
type GetParams struct {
	// Required API instance.
	*api.API

	// Required Deployment identifier.
	DeploymentID string

	// Optional parameters
	deputil.QueryParams

	// Optionally convert the legacy plans to the current deployment format.
	ConvertLegacyPlans bool

	// RefID, when specified, skips auto-discovering the deployment resource
	// RefID and instead uses the one that's passed.
	RefID string
}

// Validate ensures that the parameters are usable by the consuming function.
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("deployment get")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if len(params.DeploymentID) != 32 {
		merr = merr.Append(deputil.NewInvalidDeploymentIDError(params.DeploymentID))
	}

	return merr.ErrorOrNil()
}

// Get returns info about a deployment.
func Get(params GetParams) (*models.DeploymentGetResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeployment(
		deployments.NewGetDeploymentParams().
			WithDeploymentID(params.DeploymentID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowPlanHistory(ec.Bool(params.ShowPlanHistory)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)).
			WithConvertLegacyPlans(ec.Bool(params.ConvertLegacyPlans)).
			WithShowSystemAlerts(systemAlerts),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}

// GetApm returns info about an apm resource belonging to a given deployment.
func GetApm(params GetParams) (*models.ApmResourceInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeploymentApmResourceInfo(
		deployments.NewGetDeploymentApmResourceInfoParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}
	return res.Payload, nil
}

// GetAppSearch returns info about an appsearch resource belonging to a given deployment.
func GetAppSearch(params GetParams) (*models.AppSearchResourceInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeploymentAppsearchResourceInfo(
		deployments.NewGetDeploymentAppsearchResourceInfoParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}
	return res.Payload, nil
}

// GetElasticsearch returns info about an elasticsearch resource belonging to a given deployment.
func GetElasticsearch(params GetParams) (*models.ElasticsearchResourceInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeploymentEsResourceInfo(
		deployments.NewGetDeploymentEsResourceInfoParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)).
			WithShowSystemAlerts(systemAlerts),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}
	return res.Payload, nil
}

// GetEnterpriseSearch returns info about an Enterprise Search resource belonging to a given deployment.
func GetEnterpriseSearch(params GetParams) (*models.EnterpriseSearchResourceInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeploymentEnterpriseSearchResourceInfo(
		deployments.NewGetDeploymentEnterpriseSearchResourceInfoParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}
	return res.Payload, nil
}

// GetKibana returns info about an kibana resource belonging to a given deployment.
func GetKibana(params GetParams) (*models.KibanaResourceInfo, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Deployments.GetDeploymentKibResourceInfo(
		deployments.NewGetDeploymentKibResourceInfoParams().
			WithDeploymentID(params.DeploymentID).
			WithRefID(params.RefID).
			WithShowPlans(ec.Bool(params.ShowPlans)).
			WithShowPlanDefaults(ec.Bool(params.ShowPlanDefaults)).
			WithShowPlanLogs(ec.Bool(params.ShowPlanLogs)).
			WithShowMetadata(ec.Bool(params.ShowMetadata)).
			WithShowSettings(ec.Bool(params.ShowSettings)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}
	return res.Payload, nil
}

// GetElasticsearchID returns the deployment's elasticsearch resource ID
func GetElasticsearchID(params GetParams) (string, error) {
	res, err := Get(params)
	if err != nil {
		return "", err
	}

	esID := *res.Resources.Elasticsearch[0].ID
	return esID, nil
}
