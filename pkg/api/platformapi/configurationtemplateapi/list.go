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

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_templates"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// ListTemplateParams is the parameter of template list sub-command
type ListTemplateParams struct {
	*api.API
	Region string

	// If true, will return details for each instance configuration referenced by the template.
	ShowInstanceConfig bool

	// If present, it will cause the returned deployment templates to be adapted to return only the elements allowed
	// in that version.
	StackVersion string

	// An optional key/value pair in the form of (key:value) that will act as a filter and exclude any templates
	// that do not have a matching metadata item associated.
	Metadata string

	// If cluster is specified populates cluster_template in the response,
	// if deployment is specified populates deployment_template in the response
	Format string
}

// Validate is the implementation for the ecctl.Validator interface
func (params *ListTemplateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template list params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if strings.TrimSpace(params.Format) == "" {
		merr = merr.Append(errInvalidTemplateFormat)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

func (params *ListTemplateParams) fillDefaults() {
	if strings.TrimSpace(params.Format) == "" {
		params.Format = defaultTemplateFormat
	}
}

// ListTemplates obtains all the configured platform deployment templates
func ListTemplates(params ListTemplateParams) ([]*models.DeploymentTemplateInfo, error) {
	params.fillDefaults()

	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.PlatformConfigurationTemplates.GetDeploymentTemplates(
		platform_configuration_templates.NewGetDeploymentTemplatesParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithStackVersion(ec.String(params.StackVersion)).
			WithMetadata(ec.String(params.Metadata)).
			WithFormat(ec.String(params.Format)).
			WithShowInstanceConfigurations(ec.Bool(params.ShowInstanceConfig)),
		params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
