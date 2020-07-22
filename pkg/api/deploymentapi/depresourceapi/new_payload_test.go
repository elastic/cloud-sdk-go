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

package depresourceapi

import (
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var apmKibanaTemplateResponse = models.DeploymentTemplateInfoV2{
	ID: ec.String("default"),
	DeploymentTemplate: &models.DeploymentCreateRequest{
		Resources: &models.DeploymentCreateResources{
			Apm: []*models.ApmPayload{
				{
					Plan: &models.ApmPlan{
						ClusterTopology: []*models.ApmTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Kibana: []*models.KibanaPayload{
				{
					Plan: &models.KibanaClusterPlan{
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Elasticsearch: []*models.ElasticsearchPayload{
				{
					Plan: &models.ElasticsearchClusterPlan{
						ClusterTopology: defaultESTopologies,
					},
				},
			},
		},
	},
}

var appsearchKibanaTemplateResponse = models.DeploymentTemplateInfoV2{
	ID: ec.String("default"),
	DeploymentTemplate: &models.DeploymentCreateRequest{
		Resources: &models.DeploymentCreateResources{
			Appsearch: []*models.AppSearchPayload{
				{
					Plan: &models.AppSearchPlan{
						ClusterTopology: []*models.AppSearchTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Kibana: []*models.KibanaPayload{
				{
					Plan: &models.KibanaClusterPlan{
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Elasticsearch: []*models.ElasticsearchPayload{
				{
					Plan: &models.ElasticsearchClusterPlan{
						ClusterTopology: defaultESTopologies,
					},
				},
			},
		},
	},
}

var enterpriseSearchKibanaTemplateResponse = models.DeploymentTemplateInfoV2{
	ID: ec.String("default"),
	DeploymentTemplate: &models.DeploymentCreateRequest{
		Resources: &models.DeploymentCreateResources{
			EnterpriseSearch: []*models.EnterpriseSearchPayload{
				{
					Plan: &models.EnterpriseSearchPlan{
						ClusterTopology: []*models.EnterpriseSearchTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Kibana: []*models.KibanaPayload{
				{
					Plan: &models.KibanaClusterPlan{
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				},
			},
			Elasticsearch: []*models.ElasticsearchPayload{
				{
					Plan: &models.ElasticsearchClusterPlan{
						ClusterTopology: defaultESTopologies,
					},
				},
			},
		},
	},
}

func TestNewPayload(t *testing.T) {
	type args struct {
		params NewPayloadParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentCreateRequest
		err  error
	}{
		{
			name: "Fails due to API error",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				API: api.NewMock(mock.New500Response(
					mock.NewStringBody("error"),
				)),
			}},
			err: errors.New("error"),
		},
		{
			name: "Fails to create a deployment payload with ES and Kibana instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(defaultTemplateResponse)),
				),
			}},
			err: errors.New("deployment: the default template is not configured for Kibana. Please use another template if you wish to start Kibana instances"),
		},
		{
			name: "Fails to create a deployment payload with ES, Kibana and APM instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				ApmInstance: InstanceParams{
					RefID:     "main-apm",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				ApmEnable:            true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(appsearchKibanaTemplateResponse)),
				),
			}},
			err: errors.New("deployment: the default template is not configured for APM. Please use another template if you wish to start APM instances"),
		},
		{
			name: "Fails to create a deployment payload with ES, Kibana and App Search instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				AppsearchInstance: InstanceParams{
					RefID:     "main-appsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				AppsearchEnable:      true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(apmKibanaTemplateResponse)),
				),
			}},
			err: errors.New("deployment: the default template is not configured for App Search. Please use another template if you wish to start App Search instances"),
		},
		{
			name: "Fails to create a deployment payload with ES, Kibana and App Search instances with version auto-discover",
			args: args{params: NewPayloadParams{
				Region: "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				AppsearchInstance: InstanceParams{
					RefID:     "main-appsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				AppsearchEnable:      true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(apmKibanaTemplateResponse)),
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Host:   api.DefaultMockHost,
							Header: api.DefaultReadMockHeaders,
							Path:   "/api/v1/regions/ece-region/stack/versions",
							Method: "GET",
							Query: url.Values{
								"show_deleted":  {"false"},
								"show_unusable": {"false"},
							},
						},
						mock.NewStructBody(models.StackVersionConfigs{Stacks: []*models.StackVersionConfig{
							{Version: "7.8.0"},
						}}),
					),
				),
			}},
			err: errors.New("deployment: the default template is not configured for App Search. Please use another template if you wish to start App Search instances"),
		},
		{
			name: "Succeeds to create a deployment payload with ES and Kibana instances with version auto-discovery",
			args: args{params: NewPayloadParams{
				Region: "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(kibanaTemplateResponse)),
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Host:   api.DefaultMockHost,
							Header: api.DefaultReadMockHeaders,
							Path:   "/api/v1/regions/ece-region/stack/versions",
							Method: "GET",
							Query: url.Values{
								"show_deleted":  {"false"},
								"show_unusable": {"false"},
							},
						},
						mock.NewStructBody(models.StackVersionConfigs{Stacks: []*models.StackVersionConfig{
							{Version: "7.8.0"},
						}}),
					),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.8.0",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.8.0",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
		{
			name: "Succeeds to create a deployment payload with ES and Kibana instances (AsList)",
			args: args{params: NewPayloadParams{
				Version:                  "7.6.1",
				Region:                   "ece-region",
				DeploymentTemplateAsList: true,
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody([]models.DeploymentTemplateInfoV2{kibanaTemplateResponse})),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.6.1",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
		{
			name: "Succeeds to create a deployment payload with ES, Kibana and APM instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				ApmInstance: InstanceParams{
					RefID:     "main-apm",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				ApmEnable:            true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(apmKibanaTemplateResponse)),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.6.1",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
				Apm: []*models.ApmPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-apm"),
					Plan: &models.ApmPlan{
						Apm: &models.ApmConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.ApmTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
		{
			name: "Succeeds to create a deployment payload with ES, Kibana and APM instances (AsList)",
			args: args{params: NewPayloadParams{
				Version:                  "7.6.1",
				Region:                   "ece-region",
				DeploymentTemplateAsList: true,
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				ApmInstance: InstanceParams{
					RefID:     "main-apm",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				ApmEnable:            true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody([]models.DeploymentTemplateInfoV2{apmKibanaTemplateResponse})),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.6.1",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
				Apm: []*models.ApmPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-apm"),
					Plan: &models.ApmPlan{
						Apm: &models.ApmConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.ApmTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
		{
			name: "Succeeds to create a deployment payload with ES, Kibana and Appsearch instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				AppsearchInstance: InstanceParams{
					RefID:     "main-appsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID: "default",
				AppsearchEnable:      true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(appsearchKibanaTemplateResponse)),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.6.1",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
				Appsearch: []*models.AppSearchPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-appsearch"),
					Plan: &models.AppSearchPlan{
						Appsearch: &models.AppSearchConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.AppSearchTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
		{
			name: "Succeeds to create a deployment payload with ES, Kibana and EnterpriseSearch instances",
			args: args{params: NewPayloadParams{
				Version: "7.6.1",
				Region:  "ece-region",
				ElasticsearchInstance: InstanceParams{
					RefID:     "main-elasticsearch",
					Size:      1024,
					ZoneCount: 1,
				},
				KibanaInstance: InstanceParams{
					RefID:     "main-kibana",
					Size:      1024,
					ZoneCount: 1,
				},
				EnterpriseSearchInstance: InstanceParams{
					RefID:     "main-enterprise_search",
					Size:      1024,
					ZoneCount: 1,
				},
				DeploymentTemplateID:   "default",
				EnterpriseSearchEnable: true,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(enterpriseSearchKibanaTemplateResponse)),
				),
			}},
			want: &models.DeploymentCreateRequest{Resources: &models.DeploymentCreateResources{
				Elasticsearch: []*models.ElasticsearchPayload{{
					RefID:  ec.String("main-elasticsearch"),
					Region: ec.String("ece-region"),
					Plan: &models.ElasticsearchClusterPlan{
						Elasticsearch: &models.ElasticsearchConfiguration{
							Version: "7.6.1",
						},
						DeploymentTemplate: &models.DeploymentTemplateReference{
							ID: ec.String("default"),
						},
						ClusterTopology: []*models.ElasticsearchClusterTopologyElement{{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						}},
					}},
				},
				Kibana: []*models.KibanaPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-kibana"),
					Plan: &models.KibanaClusterPlan{
						Kibana: &models.KibanaConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.KibanaClusterTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
				EnterpriseSearch: []*models.EnterpriseSearchPayload{{
					ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
					Region:                    ec.String("ece-region"),
					RefID:                     ec.String("main-enterprise_search"),
					Plan: &models.EnterpriseSearchPlan{
						EnterpriseSearch: &models.EnterpriseSearchConfiguration{
							Version: "7.6.1",
						},
						ClusterTopology: []*models.EnterpriseSearchTopologyElement{
							{
								Size: &models.TopologySize{
									Resource: ec.String("memory"),
									Value:    ec.Int32(1024),
								},
								ZoneCount: 1,
							},
						},
					},
				}},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPayload(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
