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

func TestDeleteInvitation(t *testing.T) {
	tests := []struct {
		name   string
		params DeleteInvitationParams
		err    string
	}{
		{
			name: "fails due to parameter validation",
			err:  "invalid user params: 3 errors occurred:\n\t* InvitationTokens are not specified and is required for this operation\n\t* OrganizationID is not specified and is required for this operation\n\t* api reference is required for the operation\n\n",
		},
		{
			name: "handles successful response",
			params: DeleteInvitationParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "DELETE",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/organizations/testorg/invitations/invtoken",
						},
						mock.NewStringBody("{}")),
				),
				OrganizationID:   "testorg",
				InvitationTokens: []string{"invtoken"},
			},
		},
		{
			name: "handles failure response",
			params: DeleteInvitationParams{
				API: api.NewMock(
					mock.NewErrorResponse(400, mock.APIError{
						Code:    "root.invalid_data",
						Message: "No valid invitation token was supplied",
					}),
				),
				OrganizationID:   "testorg",
				InvitationTokens: []string{"invtooken"},
			},
			err: "api error: 1 error occurred:\n\t* root.invalid_data: No valid invitation token was supplied\n\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := DeleteInvitation(test.params)
			if test.err != "" {
				assert.EqualError(t, err, test.err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}
