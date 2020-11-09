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

package deptemplateapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployment_templates"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// GetParams is consumed by the Get function.
type GetParams struct {
	*api.API

	TemplateID   string
	Region       string
	StackVersion string

	HideInstanceConfigurations bool
}

// Validate ensures the parameters are usable by Get.
func (params GetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template get params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.TemplateID == "" {
		merr = merr.Append(errors.New("required template ID not provided"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Get returns the specified deployment template.
func Get(params GetParams) (*models.DeploymentTemplateInfoV2, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	return get(params)
}

func get(params GetParams) (*models.DeploymentTemplateInfoV2, error) {
	res, err := params.V1API.DeploymentTemplates.GetDeploymentTemplateV2(
		getParams(params), params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Unwrap(err)
	}

	return res.Payload, nil
}

// getParams creates the API params from the intermediary params struct.
// This is required for fields which are *<type> instead of directly using
// With<Param>() builders which don't work well for pointer of types on default
// values (i.e. "" => <string>, 0 => <int>, false => <bool>).
// and the function makes it easier to use since all the logic to create and
// set params is contained here.
func getParams(params GetParams) *deployment_templates.GetDeploymentTemplateV2Params {
	var apiParams = deployment_templates.NewGetDeploymentTemplateV2Params().
		WithShowInstanceConfigurations(ec.Bool(!params.HideInstanceConfigurations)).
		WithRegion(params.Region).
		WithTemplateID(params.TemplateID)

	if params.StackVersion != "" {
		apiParams.SetStackVersion(&params.StackVersion)
	}

	return apiParams
}
