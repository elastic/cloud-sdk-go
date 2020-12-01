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

// UpdateParams is consumed by the Update function.
type UpdateParams struct {
	*api.API

	Region     string
	TemplateID string
	Request    *models.DeploymentTemplateRequestBody
}

// Validate ensures the parameters are usable by Update.
func (params UpdateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template update params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Request == nil {
		merr = merr.Append(errors.New("required template request definition not provided"))
	}

	if params.TemplateID == "" {
		merr = merr.Append(errors.New("required template ID not provided"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Update updates an existing deployment template from a definition.
func Update(params UpdateParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, _, err := params.V1API.DeploymentTemplates.SetDeploymentTemplateV2(
		deployment_templates.NewSetDeploymentTemplateV2Params().
			WithTemplateID(params.TemplateID).
			WithCreateOnly(ec.Bool(false)).
			WithRegion(params.Region).
			WithBody(params.Request),
		params.AuthWriter,
	)

	return apierror.Wrap(err)
}
