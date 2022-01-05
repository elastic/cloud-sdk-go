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

package configurationtemplateapi

import (
	"context"
	"errors"
	"strings"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployment_templates"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

const defaultTemplateFormat = "cluster"

var (
	errInvalidTemplateID     = errors.New("template ID not specified and is required for this operation")
	errInvalidTemplateFormat = errors.New("template format not specified and is required for this operation")
)

// GetTemplateParams is the parameter of template show sub-command
type GetTemplateParams struct {
	*api.API

	ID     string
	Region string

	// If true, will return details for each instance configuration referenced by the template.
	ShowInstanceConfig bool
}

// Validate is the implementation for the ecctl.Validator interface
func (params *GetTemplateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template get params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if strings.TrimSpace(params.ID) == "" {
		merr = merr.Append(errInvalidTemplateID)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// GetTemplate obtains information about a specific platform deployment template
func GetTemplate(params GetTemplateParams) (*models.DeploymentTemplateInfoV2, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.DeploymentTemplates.GetDeploymentTemplateV2(
		deployment_templates.NewGetDeploymentTemplateV2Params().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithShowInstanceConfigurations(ec.Bool(params.ShowInstanceConfig)).
			WithRegion(params.Region).
			WithTemplateID(params.ID),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
