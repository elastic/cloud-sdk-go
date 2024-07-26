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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteMember(t *testing.T) {
	tests := []struct {
		name   string
		params DeleteMemberParams
		err    string
	}{
		{
			name: "fails due to parameter validation",
			err:  "invalid user params: 3 errors occurred:\n\t* OrganizationID is not specified and is required for this operation\n\t* UserIDs is not specified and is required for this operation\n\t* api reference is required for the operation\n\n",
		},
		{
			name: "handles successful response",
			params: DeleteMemberParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "DELETE",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/organizations/testorg/members/orgmember1,orgmember2",
						},
						mock.NewStringBody("{}")),
				),
				OrganizationID: "testorg",
				UserIDs:        []string{"orgmember1", "orgmember2"},
			},
		},
		{
			name: "handles failure response",
			params: DeleteMemberParams{
				API: api.NewMock(
					mock.NewErrorResponse(404, mock.APIError{
						Code:    "organization.membership_not_found",
						Message: "Organization membership not found",
					}),
				),
				OrganizationID: "testorg",
				UserIDs:        []string{"orgmember"},
			},
			err: "api error: 1 error occurred:\n\t* organization.membership_not_found: Organization membership not found\n\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := DeleteMember(test.params)
			if err != nil && !assert.EqualError(t, err, test.err) {
				t.Error(err)
			}
			if err == nil && !assert.NotNil(t, got) {
				t.Error(err)
			}
		})
	}
}
