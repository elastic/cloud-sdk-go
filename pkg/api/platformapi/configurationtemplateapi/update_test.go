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
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestUpdateTemplate(t *testing.T) {
	urlError := url.Error{
		Op:  "Put",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/templates/deployments/84e0bd6d69bb44e294809d89cea88a7e?create_only=false",
		Err: errors.New("error"),
	}
	tests := []struct {
		name string
		args UpdateTemplateParams
		err  string
	}{
		{
			name: "Platform deployment template update succeeds",
			args: UpdateTemplateParams{
				DeploymentTemplateInfo: deploymentTemplateModel(),
				ID:                     validTemplateID,
				Region:                 "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{}`),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Body:   mock.NewStringBody(`{"cluster_template":{"plan":{"cluster_topology":[{"instance_configuration_id":"default-elasticsearch","node_roles":null,"node_type":{"data":true,"master":true},"size":{"resource":"memory","value":1024}}],"elasticsearch":{"version":"6.2.3"}}},"kibana_deeplink":null,"metadata":[{"key":"trial","value":"true"}],"name":"(Trial) Default Elasticsearch","source":{"action":"deployments.create-template","admin_id":"admin","date":"2018-04-19T18:16:57.297Z","facilitator":"adminconsole","remote_addresses":["52.205.1.231"],"user_id":"1"},"system_owned":false}` + "\n"),
						Query: url.Values{
							"create_only": {"false"},
						},
						Path: "/api/v1/regions/us-east-1/platform/configuration/templates/deployments/84e0bd6d69bb44e294809d89cea88a7e",
					},
				}),
			},
		},
		{
			name: "Platform deployment template update fails due to API error",
			args: UpdateTemplateParams{
				DeploymentTemplateInfo: deploymentTemplateModel(),
				ID:                     validTemplateID,
				API:                    api.NewMock(mock.Response{Error: errors.New("error")}),
				Region:                 "us-east-1",
			},
			err: urlError.Error(),
		},
		{
			name: "Platform deployment template update fails with empty params",
			err: multierror.NewPrefixed("invalid deployment template update params",
				apierror.ErrMissingAPI,
				errInvalidTemplateID,
				errors.New("deployment template is missing"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateTemplate(tt.args)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
		})
	}
}
