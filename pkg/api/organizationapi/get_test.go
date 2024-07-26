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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOrganization(t *testing.T) {
	tests := []struct {
		name   string
		params GetParams
		want   *models.Organization
		err    string
	}{
		{
			name: "fails due to parameter validation",
			err:  "invalid user params: 2 errors occurred:\n\t* OrganizationID is not specified and is required for this operation\n\t* api reference is required for the operation\n\n",
		},
		{
			name: "handles successful response",
			params: GetParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/organizations/testorg",
						},
						mock.NewStringBody(`
{
   "billing_contacts" : [
      "contact-a", "contact-b"
   ],
   "default_disk_usage_alerts_enabled" : true,
   "id" : "testorg",
   "name" : "testorganization",
   "notifications_allowed_email_domains" : [
      "mail@test"
   ],
   "operational_contacts" : [
      "op@test"
   ]
}`)),
				),
				OrganizationID: "testorg",
			},
			want: &models.Organization{
				ID:                               ec.String("testorg"),
				Name:                             ec.String("testorganization"),
				BillingContacts:                  []string{"contact-a", "contact-b"},
				DefaultDiskUsageAlertsEnabled:    ec.Bool(true),
				NotificationsAllowedEmailDomains: []string{"mail@test"},
				OperationalContacts:              []string{"op@test"},
			},
		},
		{
			name: "handles failure response",
			params: GetParams{
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
			got, err := Get(test.params)
			if err != nil && !assert.EqualError(t, err, test.err) {
				t.Error(err)
			}
			if !assert.Equal(t, test.want, got) {
				t.Error(err)
			}
		})
	}
}
