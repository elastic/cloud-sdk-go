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

package organizationapi

import (
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/organizations"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

type ListParams struct {
	*api.API
}

func (params ListParams) Validate() error {
	var err = multierror.NewPrefixed("invalid user params")
	if params.API == nil {
		err = err.Append(apierror.ErrMissingAPI)
	}
	return err.ErrorOrNil()
}

func List(params ListParams) ([]*models.Organization, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	response, err := params.V1API.Organizations.ListOrganizations(
		organizations.NewListOrganizationsParams(),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return response.Payload.Organizations, nil
}
