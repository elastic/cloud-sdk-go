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
	"strings"

	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_templates"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// UpdateTemplateParams is the parameter of template update sub-command
type UpdateTemplateParams struct {
	*api.API
	ID string
	*models.DeploymentTemplateInfo
	Region string
}

// Validate is the implementation for the ecctl.Validator interface
func (params UpdateTemplateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template update params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if strings.TrimSpace(params.ID) == "" {
		merr = merr.Append(errInvalidTemplateID)
	}

	if params.DeploymentTemplateInfo == nil {
		merr = merr.Append(errDeploymentTemplateMissing)
	}

	if params.DeploymentTemplateInfo != nil {
		merr = merr.Append(params.DeploymentTemplateInfo.Validate(strfmt.Default))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// UpdateTemplate updates a platform deployment template
func UpdateTemplate(params UpdateTemplateParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, _, err := params.V1API.PlatformConfigurationTemplates.SetDeploymentTemplate(
		platform_configuration_templates.NewSetDeploymentTemplateParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithBody(params.DeploymentTemplateInfo).
			WithTemplateID(params.ID),
		params.AuthWriter,
	)

	return apierror.Wrap(err)
}
