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

func TestAddRoleAssignments(t *testing.T) {
	tests := []struct {
		name   string
		params AddRoleAssignmentsParams
		want   models.EmptyResponse
		err    string
	}{
		{
			name: "fails due to parameter validation",
			err:  "invalid user params: 2 errors occurred:\n\t* UserID is not specified and is required for this operation\n\t* api reference is required for the operation\n\n",
		},
		{
			name: "handles successful response",
			params: AddRoleAssignmentsParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "POST",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/users/testuser/role_assignments",
							Body:   mock.NewStringBody(`{"deployment":null,"organization":[{"organization_id":"testorg","role_id":"billing-admin"}],"platform":null}` + "\n"),
						},
						mock.NewStringBody("{}"),
					),
				),
				UserID: "testuser",
				RoleAssignments: models.RoleAssignments{
					Organization: []*models.OrganizationRoleAssignment{
						{
							OrganizationID: ec.String("testorg"),
							RoleID:         ec.String("billing-admin"),
						},
					},
				},
			},
			want: map[string]interface{}{},
		},
		{
			name: "handles failure response",
			params: AddRoleAssignmentsParams{
				API: api.NewMock(
					mock.NewErrorResponse(404, mock.APIError{
						Code:    "user.not_found",
						Message: "user not found",
					}),
				),
				UserID: "testuser",
				RoleAssignments: models.RoleAssignments{
					Organization: []*models.OrganizationRoleAssignment{
						{
							OrganizationID: ec.String("testorg"),
							RoleID:         ec.String("billing-admin"),
						},
					},
				},
			},
			err: "api error: 1 error occurred:\n\t* user.not_found: user not found\n\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := AddRoleAssignments(test.params)
			if err != nil && !assert.EqualError(t, err, test.err) {
				t.Error(err)
			}
			if got != nil && !assert.Equal(t, test.want, *got) {
				t.Error(err)
			}
		})
	}
}
