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
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestListInvitations(t *testing.T) {
	dateTime := strfmt.DateTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC))
	tests := []struct {
		name   string
		params ListInvitationsParams
		want   *models.OrganizationInvitations
		err    string
	}{
		{
			name: "fails due to parameter validation",
			err:  "invalid user params: 2 errors occurred:\n\t* OrganizationID is not specified and is required for this operation\n\t* api reference is required for the operation\n\n",
		},
		{
			name: "handles successful response",
			params: ListInvitationsParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/organizations/testorg/invitations",
						},
						mock.NewStringBody(`
{
   "invitations" : [
      {
         "accepted_at" : "2019-01-01T00:00:00Z",
         "created_at" : "2019-01-01T00:00:00Z",
         "email" : "test@mail",
         "expired" : false,
         "expires_at" : "2019-01-01T00:00:00Z",
         "organization" : {
            "billing_contacts" : [
               "billing@mail"
            ],
            "default_disk_usage_alerts_enabled" : true,
            "id" : "testorg",
            "name" : "testorganization",
            "notifications_allowed_email_domains" : [
               "allowed@mail"
            ],
            "operational_contacts" : [
               "op@mail"
            ]
         },
         "role_assignments": {
           "organization": [
             {
               "organization_id": "testorg",
               "role_id": "billing-admin"
             }
           ]
         },
         "token" : "token"
      }
   ]
}`)),
				),
				OrganizationID: "testorg",
			},
			want: &models.OrganizationInvitations{
				Invitations: []*models.OrganizationInvitation{
					{
						Email:      ec.String("test@mail"),
						CreatedAt:  &dateTime,
						AcceptedAt: dateTime,
						ExpiresAt:  &dateTime,
						Expired:    ec.Bool(false),
						Organization: &models.Organization{
							ID:                               ec.String("testorg"),
							Name:                             ec.String("testorganization"),
							BillingContacts:                  []string{"billing@mail"},
							DefaultDiskUsageAlertsEnabled:    ec.Bool(true),
							NotificationsAllowedEmailDomains: []string{"allowed@mail"},
							OperationalContacts:              []string{"op@mail"},
						},
						RoleAssignments: &models.RoleAssignments{
							Organization: []*models.OrganizationRoleAssignment{
								{
									OrganizationID: ec.String("testorg"),
									RoleID:         ec.String("billing-admin"),
								},
							},
						},
						Token: ec.String("token"),
					},
				},
			},
		},
		{
			name: "handles failure response",
			params: ListInvitationsParams{
				API: api.NewMock(
					mock.NewErrorResponse(404, mock.APIError{
						Code:    "organization.not_found",
						Message: "organization not found",
					}),
				),
				OrganizationID: "testorg",
			},
			err: "api error: 1 error occurred:\n\t* organization.not_found: organization not found\n\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ListInvitations(test.params)
			if test.err != "" {
				assert.EqualError(t, err, test.err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
