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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCreateTemplate(t *testing.T) {
	urlError := url.Error{
		Op:  "Post",
		URL: "https://mock.elastic.co/api/v1/deployments/templates?region=us-east-1",
		Err: errors.New("error"),
	}
	tests := []struct {
		name string
		args CreateTemplateParams
		want string
		err  string
	}{
		{
			name: "Platform deployment template create succeeds",
			args: CreateTemplateParams{
				DeploymentTemplateRequestBody: deploymentTemplateModel("us-east-1"),
				Region:                        "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(`{"id": "9362b09d838640b2beea21b3343b4686"}"`),
						StatusCode: 201,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Query: url.Values{
							"region": {"us-east-1"},
						},
						Body: mock.NewStringBody(`{"deployment_template":{"resources":{"apm":null,"appsearch":null,"elasticsearch":[{"plan":{"cluster_topology":[{"instance_configuration_id":"default-elasticsearch","node_roles":null,"node_type":{"data":true,"master":true},"size":{"resource":"memory","value":1024}}],"elasticsearch":{"version":"6.2.3"}},"ref_id":"main-elasticsearch","region":"us-east-1"}],"enterprise_search":null,"integrations_server":null,"kibana":null}},"kibana_deeplink":null,"metadata":[{"key":"trial","value":"true"}],"name":"(Trial) Default Elasticsearch","system_owned":false}` + "\n"),
						Path: "/api/v1/deployments/templates",
					},
				}),
			},
			want: "9362b09d838640b2beea21b3343b4686",
		},
		{
			name: "Platform deployment template create fails due to API error",
			args: CreateTemplateParams{
				DeploymentTemplateRequestBody: deploymentTemplateModel("us-east-1"),
				Region:                        "us-east-1",
				API:                           api.NewMock(mock.Response{Error: errors.New("error")}),
			},
			err: urlError.Error(),
		},
		{
			name: "Platform deployment template create fails with empty params",
			err: multierror.NewPrefixed("invalid deployment template create params",
				apierror.ErrMissingAPI,
				errors.New("deployment template is missing"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTemplate(tt.args)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func deploymentTemplateModel(region string) *models.DeploymentTemplateRequestBody {
	var refId = "main-elasticsearch"
	template := models.DeploymentTemplateRequestBody{Name: ec.String("(Trial) Default Elasticsearch"),
		SystemOwned: ec.Bool(false),
		Metadata: []*models.MetadataItem{{
			Value: ec.String("true"),
			Key:   ec.String("trial"),
		}},
		DeploymentTemplate: &models.DeploymentCreateRequest{
			Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  &refId,
					Region: &region,
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
				},
				},
			},
		},
	}

	return &template
}
