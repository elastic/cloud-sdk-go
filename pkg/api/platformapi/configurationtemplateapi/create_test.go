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
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCreateTemplate(t *testing.T) {
	tests := []struct {
		name string
		args CreateTemplateParams
		want string
		err  error
	}{
		{
			name: "Platform deployment template create succeeds",
			args: CreateTemplateParams{
				DeploymentTemplateInfo: deploymentTemplateModel(),
				Region:                 "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{"id": "9362b09d838640b2beea21b3343b4686"}"`),
						StatusCode: 201,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Body:   mock.NewStringBody(`{"cluster_template":{"plan":{"cluster_topology":[{"instance_configuration_id":"default-elasticsearch","node_type":{"data":true,"master":true},"size":{"resource":"memory","value":1024}}],"elasticsearch":{"version":"6.2.3"}}},"kibana_deeplink":null,"metadata":[{"key":"trial","value":"true"}],"name":"(Trial) Default Elasticsearch","source":{"action":"deployments.create-template","admin_id":"admin","date":"2018-04-19T18:16:57.297Z","facilitator":"adminconsole","remote_addresses":["52.205.1.231"],"user_id":"1"},"system_owned":false}` + "\n"),
						Path:   "/api/v1/regions/us-east-1/platform/configuration/templates/deployments",
					},
				}),
			},
			want: "9362b09d838640b2beea21b3343b4686",
		},
		{
			name: "Platform deployment template create succeeds specifying template ID",
			args: CreateTemplateParams{
				DeploymentTemplateInfo: deploymentTemplateModelWithID("template-id"),
				Region:                 "us-east-1",
				API: api.NewMock(mock.Response{Response: http.Response{
					Body:       mock.NewStringBody(`{"id": "template-id"}"`),
					StatusCode: 201,
				}}),
			},
			want: "template-id",
		},
		{
			name: "Platform deployment template create fails due to API error",
			args: CreateTemplateParams{
				DeploymentTemplateInfo: deploymentTemplateModel(),
				Region:                 "us-east-1",
				API:                    api.NewMock(mock.Response{Error: errors.New("error")}),
			},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/templates/deployments",
				Err: errors.New("error"),
			},
		},
		{
			name: "Platform deployment template create fails with empty params",
			err: multierror.NewPrefixed("invalid deployment template create params",
				apierror.ErrMissingAPI,
				errors.New("deployment template is missing"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTemplate(tt.args)

			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deploymentTemplateModelWithID(id string) *models.DeploymentTemplateInfo {
	var model = deploymentTemplateModel()
	model.ID = id

	return model
}
func deploymentTemplateModel() *models.DeploymentTemplateInfo {
	sourceDate, _ := strfmt.ParseDateTime("2018-04-19T18:16:57.297Z")

	template := models.DeploymentTemplateInfo{Name: ec.String("(Trial) Default Elasticsearch"),
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
		ClusterTemplate: &models.DeploymentTemplateDefinitionRequest{
			Plan: &models.ElasticsearchClusterPlan{
				Elasticsearch: &models.ElasticsearchConfiguration{
					Version: "6.2.3",
				},
				ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
					InstanceConfigurationID: "default-elasticsearch",
					Size: &models.TopologySize{
						Value:    ec.Int32(1024),
						Resource: ec.String("memory"),
					},
					NodeType: &models.ElasticsearchNodeType{
						Master: ec.Bool(true),
						Data:   ec.Bool(true),
					},
				},
				},
			},
		}}

	return &template
}
