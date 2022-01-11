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
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/elastic/cloud-sdk-go/pkg/util/testutils"
)

var (
	validTemplateID = "84e0bd6d69bb44e294809d89cea88a7e"
)

func TestGetTemplate(t *testing.T) {
	var sourceDate = testutils.ParseDate(t, "2018-04-19T18:16:57.297Z")

	var templateFormatDeployment = `
	{
  "name": "(Trial) Default Elasticsearch",
  "source": {
	"user_id": "1",
	"facilitator": "adminconsole",
	"date": "2018-04-19T18:16:57.297Z",
	"admin_id": "admin",
	"action": "deployments.create-template",
	"remote_addresses": ["52.205.1.231"]
  },
  "description": "Test default Elasticsearch trial template",
  "id": "` + validTemplateID + `",
  "metadata": [{
	"key": "trial",
	"value": "true"
	}],
	"deployment_template": {
        "resources": {}
    },
	"system_owned": false
}`
	urlError := url.Error{
		Op:  "Get",
		URL: `https://mock.elastic.co/api/v1/deployments/templates/84e0bd6d69bb44e294809d89cea88a7e?region=us-east-1&show_instance_configurations=false`,
		Err: errors.New("error"),
	}
	tests := []struct {
		name string
		args GetTemplateParams
		want *models.DeploymentTemplateInfoV2
		err  string
	}{
		{
			name: "Platform deployment template succeeds",
			args: GetTemplateParams{
				ID: validTemplateID,
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(templateFormatDeployment),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Query: url.Values{
							"region":                       {"us-east-1"},
							"show_instance_configurations": {"false"},
						},
						Path: "/api/v1/deployments/templates/84e0bd6d69bb44e294809d89cea88a7e",
					},
				}),
				Region: "us-east-1",
			},
			want: &models.DeploymentTemplateInfoV2{
				Name:        ec.String("(Trial) Default Elasticsearch"),
				ID:          &validTemplateID,
				Description: "Test default Elasticsearch trial template",
				SystemOwned: ec.Bool(false),
				Metadata: []*models.MetadataItem{{

					Value: ec.String("true"),
					Key:   ec.String("trial"),
				}},
				Source: &models.ChangeSourceInfo{
					UserID:          "1",
					Facilitator:     ec.String("adminconsole"),
					Date:            &sourceDate,
					AdminID:         "admin",
					Action:          ec.String("deployments.create-template"),
					RemoteAddresses: []string{"52.205.1.231"},
				},
				DeploymentTemplate: &models.DeploymentCreateRequest{
					Resources: &models.DeploymentCreateResources{},
				},
			},
		},
		{
			name: "Platform deployment template show fails due to API error",
			args: GetTemplateParams{
				ID:     validTemplateID,
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			},
			err: urlError.Error(),
		},
		{
			name: "Platform deployment template show fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment template get params",
				errors.New("api reference is required for the operation"),
				errors.New("template ID not specified and is required for this operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTemplate(tt.args)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
