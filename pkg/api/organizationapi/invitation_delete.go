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
	"errors"
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/organizations"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"strings"
)

type DeleteInvitationParams struct {
	*api.API

	OrganizationID string

	InvitationTokens []string
}

func (params DeleteInvitationParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid user params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.OrganizationID == "" {
		merr = merr.Append(errors.New("OrganizationID is not specified and is required for this operation"))
	}
	if len(params.InvitationTokens) == 0 {
		merr = merr.Append(errors.New("InvitationTokens are not specified and is required for this operation"))
	}
	return merr.ErrorOrNil()
}

func DeleteInvitation(params DeleteInvitationParams) (models.EmptyResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	invitationTokens := strings.Join(params.InvitationTokens, ",")

	response, err := params.V1API.Organizations.DeleteOrganizationInvitations(
		organizations.NewDeleteOrganizationInvitationsParams().
			WithOrganizationID(params.OrganizationID).
			WithInvitationTokens(invitationTokens),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return response.Payload, nil
}