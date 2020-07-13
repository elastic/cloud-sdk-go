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
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var enterpriseSearchTemplateResponse = models.DeploymentTemplateInfo{
	ID: "default.enterprise-search",
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

var defaultESTemplateResponse = models.DeploymentTemplateInfo{
	ID: "default",
	DeploymentTemplate: &models.DeploymentCreateRequest{
		Resources: &models.DeploymentCreateResources{
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

func TestNewEnterpriseSearch(t *testing.T) {
	var getResponse = models.DeploymentGetResponse{
		Resources: &models.DeploymentResources{
			Elasticsearch: []*models.ElasticsearchResourceInfo{{
				RefID: ec.String("main-elasticsearch"),
				Info: &models.ElasticsearchClusterInfo{
					PlanInfo: &models.ElasticsearchClusterPlansInfo{
						Current: &models.ElasticsearchClusterPlanInfo{
							Plan: &models.ElasticsearchClusterPlan{
								DeploymentTemplate: &models.DeploymentTemplateReference{
									ID: ec.String("an ID"),
								},
							},
						},
					},
				},
			}},
		},
	}

	type args struct {
		params NewStateless
	}
	tests := []struct {
		name string
		args args
		want *models.EnterpriseSearchPayload
		err  error
	}{
		{
			name: "fails due to parameter validation",
			args: args{params: NewStateless{DeploymentID: "invalidID"}},
			err: multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
				errors.New("topology: region cannot be empty"),
			),
		},
		{
			name: "fails obtaining the deployment info",
			args: args{params: NewStateless{
				DeploymentID: mock.ValidClusterID,
				API:          api.NewMock(mock.SampleInternalError()),
				Region:       "ece-region",
			}},
			err: mock.MultierrorInternalError,
		},
		{
			name: "obtains the deployment info but fails getting the template ID info",
			args: args{params: NewStateless{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
						Resources: &models.DeploymentResources{
							Elasticsearch: []*models.ElasticsearchResourceInfo{{
								Info: &models.ElasticsearchClusterInfo{
									PlanInfo: &models.ElasticsearchClusterPlansInfo{},
								},
							}},
						},
					})),
				),
				Region: "ece-region",
			}},
			err: errors.New("unable to obtain deployment template ID from existing deployment ID, please specify a one"),
		},
		{
			name: "obtains the deployment info but fails getting the template ID info from the API",
			args: args{params: NewStateless{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(getResponse)),
					mock.SampleInternalError(),
				),
				Region: "ece-region",
			}},
			err: mock.MultierrorInternalError,
		},
		{
			name: "obtains the deployment template but it's an invalid template for enterprise search",
			args: args{params: NewStateless{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(getResponse)),
					mock.New200Response(mock.NewStructBody(defaultESTemplateResponse)),
				),
				Region: "ece-region",
			}},
			err: errors.New("deployment: the an ID template is not configured for Enterprise Search. Please use another template if you wish to start Enterprise Search instances"),
		},
		{
			name: "succeeds with no argument override",
			args: args{params: NewStateless{
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(getResponse)),
					mock.New200Response(mock.NewStructBody(enterpriseSearchTemplateResponse)),
				),
				Region: "ece-region",
			}},
			want: &models.EnterpriseSearchPayload{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				Region:                    ec.String("ece-region"),
				RefID:                     ec.String("main-enterprise_search"),
				Plan: &models.EnterpriseSearchPlan{
					EnterpriseSearch: &models.EnterpriseSearchConfiguration{},
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
		{
			name: "succeeds with argument overrides",
			args: args{params: NewStateless{
				Size:         4096,
				ZoneCount:    3,
				DeploymentID: mock.ValidClusterID,
				API: api.NewMock(
					mock.New200Response(mock.NewStructBody(getResponse)),
					mock.New200Response(mock.NewStructBody(enterpriseSearchTemplateResponse)),
				),
				Region: "ece-region",
			}},
			want: &models.EnterpriseSearchPayload{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				Region:                    ec.String("ece-region"),
				RefID:                     ec.String("main-enterprise_search"),
				Plan: &models.EnterpriseSearchPlan{
					EnterpriseSearch: &models.EnterpriseSearchConfiguration{},
					ClusterTopology: []*models.EnterpriseSearchTopologyElement{
						{
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(4096),
							},
							ZoneCount: 3,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEnterpriseSearch(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("NewEnterpriseSearch() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				g, _ := json.Marshal(got)
				w, _ := json.Marshal(tt.want)
				println(string(g))
				println(string(w))
				t.Errorf("NewEnterpriseSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
