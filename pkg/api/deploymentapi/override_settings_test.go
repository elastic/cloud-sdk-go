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

package deploymentapi

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestOverrideCreateOrUpdateRequest(t *testing.T) {
	var eceRegion = "ece-region"
	var overriddenRegion = "overridden-region"
	type args struct {
		req       interface{}
		overrides *PayloadOverrides
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "set name override",
			args: args{
				req: &models.DeploymentCreateRequest{
					Name:      "Some",
					Resources: &models.DeploymentCreateResources{},
				},
				overrides: &PayloadOverrides{Name: "some other"},
			},
			want: &models.DeploymentCreateRequest{
				Name:      "some other",
				Resources: &models.DeploymentCreateResources{},
			},
		},
		{
			name: "set name, version, region and ref_id override",
			args: args{
				overrides: &PayloadOverrides{
					Name:               "some other",
					Version:            "7.4.1",
					Region:             eceRegion,
					ElasticsearchRefID: "main-elasticsearch",
					OverrideRefIDs:     true,
				},
				req: &models.DeploymentCreateRequest{
					Name: "Some",
					Resources: &models.DeploymentCreateResources{
						Apm: []*models.ApmPayload{
							{
								Plan: &models.ApmPlan{
									Apm: &models.ApmConfiguration{Version: "1.2.3"},
								},
							},
						},
						Appsearch: []*models.AppSearchPayload{
							{
								Plan: &models.AppSearchPlan{
									Appsearch: &models.AppSearchConfiguration{Version: "1.2.3"},
								},
							},
						},
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									Elasticsearch: &models.ElasticsearchConfiguration{Version: "1.2.3"},
									ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
										{
											ID: "hot_content",
											NodeType: &models.ElasticsearchNodeType{
												Master: ec.Bool(true),
												Data:   ec.Bool(true),
												Ingest: ec.Bool(true),
												Ml:     ec.Bool(false),
											},
											NodeRoles: []string{
												"master",
												"ingest",
												"remote_cluster_client",
												"data_hot",
												"transform",
												"data_content",
											},
										},
									},
								},
							},
						},
						EnterpriseSearch: []*models.EnterpriseSearchPayload{
							{
								Plan: &models.EnterpriseSearchPlan{
									EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "1.2.3"},
								},
							},
						},
						Kibana: []*models.KibanaPayload{
							{
								Plan: &models.KibanaClusterPlan{
									Kibana: &models.KibanaConfiguration{Version: "1.2.3"},
								},
							},
						},
					},
				},
			},
			want: &models.DeploymentCreateRequest{
				Name: "some other",
				Resources: &models.DeploymentCreateResources{
					Apm: []*models.ApmPayload{
						{
							Region: &eceRegion,
							Plan: &models.ApmPlan{
								Apm: &models.ApmConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-apm"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Appsearch: []*models.AppSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.AppSearchPlan{
								Appsearch: &models.AppSearchConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-appsearch"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Elasticsearch: []*models.ElasticsearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.4.1"},
								ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
									{
										ID: "hot_content",
										NodeType: &models.ElasticsearchNodeType{
											Master: ec.Bool(true),
											Data:   ec.Bool(true),
											Ingest: ec.Bool(true),
											Ml:     ec.Bool(false),
										},
									},
								},
							},
							RefID: ec.String("main-elasticsearch"),
						},
					},
					EnterpriseSearch: []*models.EnterpriseSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.EnterpriseSearchPlan{
								EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-enterprise_search"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Kibana: []*models.KibanaPayload{
						{
							Region: &eceRegion,
							Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-kibana"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
				},
			},
		},
		{
			name: "set region and ref_id override (no version leaves the NodeRoles field)",
			args: args{
				overrides: &PayloadOverrides{
					Region:             eceRegion,
					ElasticsearchRefID: "main-elasticsearch",
					OverrideRefIDs:     true,
				},
				req: &models.DeploymentCreateRequest{
					Name: "Some",
					Resources: &models.DeploymentCreateResources{
						Apm: []*models.ApmPayload{
							{
								Plan: &models.ApmPlan{
									Apm: &models.ApmConfiguration{},
								},
							},
						},
						Appsearch: []*models.AppSearchPayload{
							{
								Plan: &models.AppSearchPlan{
									Appsearch: &models.AppSearchConfiguration{},
								},
							},
						},
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									Elasticsearch: &models.ElasticsearchConfiguration{},
									ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
										{
											ID: "hot_content",
											NodeType: &models.ElasticsearchNodeType{
												Master: ec.Bool(true),
												Data:   ec.Bool(true),
												Ingest: ec.Bool(true),
												Ml:     ec.Bool(false),
											},
											NodeRoles: []string{
												"master",
												"ingest",
												"remote_cluster_client",
												"data_hot",
												"transform",
												"data_content",
											},
										},
									},
								},
							},
						},
						EnterpriseSearch: []*models.EnterpriseSearchPayload{
							{
								Plan: &models.EnterpriseSearchPlan{
									EnterpriseSearch: &models.EnterpriseSearchConfiguration{},
								},
							},
						},
						Kibana: []*models.KibanaPayload{
							{
								Plan: &models.KibanaClusterPlan{
									Kibana: &models.KibanaConfiguration{},
								},
							},
						},
					},
				},
			},
			want: &models.DeploymentCreateRequest{
				Name: "Some",
				Resources: &models.DeploymentCreateResources{
					Apm: []*models.ApmPayload{
						{
							Region: &eceRegion,
							Plan: &models.ApmPlan{
								Apm: &models.ApmConfiguration{},
							},
							RefID:                     ec.String("main-apm"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Appsearch: []*models.AppSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.AppSearchPlan{
								Appsearch: &models.AppSearchConfiguration{},
							},
							RefID:                     ec.String("main-appsearch"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Elasticsearch: []*models.ElasticsearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{},
								ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
									{
										ID: "hot_content",
										NodeRoles: []string{
											"master",
											"ingest",
											"remote_cluster_client",
											"data_hot",
											"transform",
											"data_content",
										},
									},
								},
							},
							RefID: ec.String("main-elasticsearch"),
						},
					},
					EnterpriseSearch: []*models.EnterpriseSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.EnterpriseSearchPlan{
								EnterpriseSearch: &models.EnterpriseSearchConfiguration{},
							},
							RefID:                     ec.String("main-enterprise_search"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Kibana: []*models.KibanaPayload{
						{
							Region: &eceRegion,
							Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{},
							},
							RefID:                     ec.String("main-kibana"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
				},
			},
		},
		{
			name: "set name, version, region and ref_id override",
			args: args{
				overrides: &PayloadOverrides{
					Name:               "some other",
					Version:            "7.11.2",
					Region:             eceRegion,
					ElasticsearchRefID: "main-elasticsearch",
					OverrideRefIDs:     true,
				},
				req: &models.DeploymentCreateRequest{
					Name: "Some",
					Resources: &models.DeploymentCreateResources{
						Apm: []*models.ApmPayload{
							{
								Plan: &models.ApmPlan{
									Apm: &models.ApmConfiguration{Version: "1.2.3"},
								},
							},
						},
						Appsearch: []*models.AppSearchPayload{
							{
								Plan: &models.AppSearchPlan{
									Appsearch: &models.AppSearchConfiguration{Version: "1.2.3"},
								},
							},
						},
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									Elasticsearch: &models.ElasticsearchConfiguration{Version: "1.2.3"},
									ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
										{
											ID: "hot_content",
											NodeType: &models.ElasticsearchNodeType{
												Master: ec.Bool(true),
												Data:   ec.Bool(true),
												Ingest: ec.Bool(true),
												Ml:     ec.Bool(false),
											},
											NodeRoles: []string{
												"master",
												"ingest",
												"remote_cluster_client",
												"data_hot",
												"transform",
												"data_content",
											},
										},
									},
								},
							},
						},
						EnterpriseSearch: []*models.EnterpriseSearchPayload{
							{
								Plan: &models.EnterpriseSearchPlan{
									EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "1.2.3"},
								},
							},
						},
						Kibana: []*models.KibanaPayload{
							{
								Plan: &models.KibanaClusterPlan{
									Kibana: &models.KibanaConfiguration{Version: "1.2.3"},
								},
							},
						},
					},
				},
			},
			want: &models.DeploymentCreateRequest{
				Name: "some other",
				Resources: &models.DeploymentCreateResources{
					Apm: []*models.ApmPayload{
						{
							Region: &eceRegion,
							Plan: &models.ApmPlan{
								Apm: &models.ApmConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("main-apm"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Appsearch: []*models.AppSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.AppSearchPlan{
								Appsearch: &models.AppSearchConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("main-appsearch"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Elasticsearch: []*models.ElasticsearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.11.2"},
								ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
									{
										ID: "hot_content",
										NodeRoles: []string{
											"master",
											"ingest",
											"remote_cluster_client",
											"data_hot",
											"transform",
											"data_content",
										},
									},
								},
							},
							RefID: ec.String("main-elasticsearch"),
						},
					},
					EnterpriseSearch: []*models.EnterpriseSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.EnterpriseSearchPlan{
								EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("main-enterprise_search"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Kibana: []*models.KibanaPayload{
						{
							Region: &eceRegion,
							Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("main-kibana"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
				},
			},
		},
		{
			name: "set name, version, region with no ref_id override",
			args: args{
				overrides: &PayloadOverrides{
					Name:               "some other",
					Version:            "7.11.2",
					Region:             eceRegion,
					ElasticsearchRefID: "main-elasticsearch",
				},
				req: &models.DeploymentCreateRequest{
					Name: "Some",
					Resources: &models.DeploymentCreateResources{
						Apm: []*models.ApmPayload{
							{
								Plan: &models.ApmPlan{
									Apm: &models.ApmConfiguration{Version: "1.2.3"},
								},
								RefID: ec.String("apm"),
							},
						},
						Appsearch: []*models.AppSearchPayload{
							{
								Plan: &models.AppSearchPlan{
									Appsearch: &models.AppSearchConfiguration{Version: "1.2.3"},
								},
								RefID: ec.String("appsearch"),
							},
						},
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									Elasticsearch: &models.ElasticsearchConfiguration{Version: "1.2.3"},
									ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
										{
											ID: "hot_content",
											NodeType: &models.ElasticsearchNodeType{
												Master: ec.Bool(true),
												Data:   ec.Bool(true),
												Ingest: ec.Bool(true),
												Ml:     ec.Bool(false),
											},
											NodeRoles: []string{
												"master",
												"ingest",
												"remote_cluster_client",
												"data_hot",
												"transform",
												"data_content",
											},
										},
									},
								},
							},
						},
						EnterpriseSearch: []*models.EnterpriseSearchPayload{
							{
								Plan: &models.EnterpriseSearchPlan{
									EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "1.2.3"},
								},
								RefID: ec.String("enterprise_search"),
							},
						},
						Kibana: []*models.KibanaPayload{
							{
								Plan: &models.KibanaClusterPlan{
									Kibana: &models.KibanaConfiguration{Version: "1.2.3"},
								},
								RefID: ec.String("kibana"),
							},
						},
					},
				},
			},
			want: &models.DeploymentCreateRequest{
				Name: "some other",
				Resources: &models.DeploymentCreateResources{
					Apm: []*models.ApmPayload{
						{
							Region: &eceRegion,
							Plan: &models.ApmPlan{
								Apm: &models.ApmConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("apm"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Appsearch: []*models.AppSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.AppSearchPlan{
								Appsearch: &models.AppSearchConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("appsearch"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Elasticsearch: []*models.ElasticsearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.11.2"},
								ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
									{
										ID: "hot_content",
										NodeRoles: []string{
											"master",
											"ingest",
											"remote_cluster_client",
											"data_hot",
											"transform",
											"data_content",
										},
									},
								},
							},
							RefID: ec.String("main-elasticsearch"),
						},
					},
					EnterpriseSearch: []*models.EnterpriseSearchPayload{
						{
							Region: &eceRegion,
							Plan: &models.EnterpriseSearchPlan{
								EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("enterprise_search"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Kibana: []*models.KibanaPayload{
						{
							Region: &eceRegion,
							Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{Version: "7.11.2"},
							},
							RefID:                     ec.String("kibana"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
				},
			},
		},
		{
			name: "set region override on a DeploymentUpdateRequest",
			args: args{
				overrides: &PayloadOverrides{
					Region:             "overridden-region",
					ElasticsearchRefID: "main-elasticsearch",
					OverrideRefIDs:     true,
				},
				req: &models.DeploymentUpdateRequest{
					Resources: &models.DeploymentUpdateResources{
						Apm: []*models.ApmPayload{
							{
								Plan: &models.ApmPlan{
									Apm: &models.ApmConfiguration{Version: "7.4.1"},
								},
							},
						},
						Appsearch: []*models.AppSearchPayload{
							{
								Plan: &models.AppSearchPlan{
									Appsearch: &models.AppSearchConfiguration{Version: "7.4.1"},
								},
							},
						},
						EnterpriseSearch: []*models.EnterpriseSearchPayload{
							{
								Plan: &models.EnterpriseSearchPlan{
									EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "7.4.1"},
								},
							},
						},
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.4.1"},
								},
							},
						},
						Kibana: []*models.KibanaPayload{
							{
								Plan: &models.KibanaClusterPlan{
									Kibana: &models.KibanaConfiguration{Version: "7.4.1"},
								},
							},
						},
					},
				},
			},
			want: &models.DeploymentUpdateRequest{
				Resources: &models.DeploymentUpdateResources{
					Apm: []*models.ApmPayload{
						{
							Region: &overriddenRegion,
							Plan: &models.ApmPlan{
								Apm: &models.ApmConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-apm"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Appsearch: []*models.AppSearchPayload{
						{
							Region: &overriddenRegion,
							Plan: &models.AppSearchPlan{
								Appsearch: &models.AppSearchConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-appsearch"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Elasticsearch: []*models.ElasticsearchPayload{
						{
							Region: &overriddenRegion,
							Plan: &models.ElasticsearchClusterPlan{
								Elasticsearch: &models.ElasticsearchConfiguration{Version: "7.4.1"},
							},
							RefID: ec.String("main-elasticsearch"),
						},
					},
					EnterpriseSearch: []*models.EnterpriseSearchPayload{
						{
							Region: &overriddenRegion,
							Plan: &models.EnterpriseSearchPlan{
								EnterpriseSearch: &models.EnterpriseSearchConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-enterprise_search"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
					Kibana: []*models.KibanaPayload{
						{
							Region: &overriddenRegion,
							Plan: &models.KibanaClusterPlan{
								Kibana: &models.KibanaConfiguration{Version: "7.4.1"},
							},
							RefID:                     ec.String("main-kibana"),
							ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req = tt.args.req
			OverrideCreateOrUpdateRequest(req, tt.args.overrides)

			if !assert.Equal(t, tt.want, req) {
				t.Errorf("setOverrides() = %v, want %v", req, tt.want)
			}
		})
	}
}
