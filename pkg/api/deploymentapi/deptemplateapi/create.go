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

type idResponse interface {
	GetPayload() *models.IDResponse
}

// CreateParams is consumed by the Create function.
type CreateParams struct {
	*api.API

	Region     string
	TemplateID string
	Request    *models.DeploymentTemplateRequestBody
}

// Validate ensures the parameters are usable by Create.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template create params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Request == nil {
		merr = merr.Append(errors.New("required template request definition not provided"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Create creates a deployment template from a definition. If a TemplateID is
// specified in the params then it uses the "PUT" version of the call which
// will create a deployment template with the specified id, if the ID already
// exists, it'll return an error.
func Create(params CreateParams) (string, error) {
	if err := params.Validate(); err != nil {
		return "", err
	}

	var res idResponse
	var err error
	if params.TemplateID == "" {
		_, res, err = params.V1API.DeploymentTemplates.CreateDeploymentTemplateV2(
			deployment_templates.NewCreateDeploymentTemplateV2Params().
				WithRegion(params.Region).
				WithBody(params.Request),
			params.AuthWriter,
		)
	} else {
		_, res, err = params.V1API.DeploymentTemplates.SetDeploymentTemplateV2(
			deployment_templates.NewSetDeploymentTemplateV2Params().
				WithTemplateID(params.TemplateID).
				WithCreateOnly(ec.Bool(true)).
				WithRegion(params.Region).
				WithBody(params.Request),
			params.AuthWriter,
		)
	}

	if err != nil {
		return "", apierror.Wrap(err)
	}

	return *res.GetPayload().ID, nil
}
