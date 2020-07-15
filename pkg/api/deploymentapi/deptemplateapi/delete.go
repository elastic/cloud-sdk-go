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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// DeleteParams is consumed by the Delete function.
type DeleteParams struct {
	*api.API

	TemplateID string
	Region     string
}

// Validate ensures the parameters are usable by Delete.
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template delete params")

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

// Delete removes an existing deployment temeplate.
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.V1API.DeploymentTemplates.DeleteDeploymentTemplateV2(
			deployment_templates.NewDeleteDeploymentTemplateV2Params().
				WithRegion(params.Region).
				WithTemplateID(params.TemplateID),
			params.AuthWriter,
		),
	)
}
