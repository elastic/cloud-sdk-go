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
	"fmt"
	"strings"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployment_templates"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// ListParams is consumed by the List function.
type ListParams struct {
	*api.API

	MetadataFilter string
	Region         string
	StackVersion   string

	ShowHidden                 bool
	HideInstanceConfigurations bool
}

// Validate ensures the parameters are usable by List.
func (params ListParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template list params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.MetadataFilter != "" && len(strings.Split(params.MetadataFilter, ":")) != 2 {
		merr = merr.Append(fmt.Errorf(
			`invalid metadata filter "%s", must be formatted in the form of (key:value)`,
			params.MetadataFilter,
		))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// List returns a list of the available deployment templates.
func List(params ListParams) ([]*models.DeploymentTemplateInfoV2, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.V1API.DeploymentTemplates.GetDeploymentTemplatesV2(
		listParams(params), params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}

// listParams creates the API params from the intermediary params struct.
// This is required for fields which are *<type> instead of directly using
// With<Param>() builders which don't work well for pointer of types on default
// values (i.e. "" => <string>, 0 => <int>, false => <bool>).
// and the function makes it easier to use since all the logic to create and
// set params is contained here.
func listParams(params ListParams) *deployment_templates.GetDeploymentTemplatesV2Params {
	var apiParams = deployment_templates.NewGetDeploymentTemplatesV2Params().
		WithShowInstanceConfigurations(ec.Bool(!params.HideInstanceConfigurations)).
		WithShowHidden(&params.ShowHidden).
		WithRegion(params.Region)

	if params.MetadataFilter != "" {
		apiParams.SetMetadata(&params.MetadataFilter)
	}

	if params.StackVersion != "" {
		apiParams.SetStackVersion(&params.StackVersion)
	}

	return apiParams
}
