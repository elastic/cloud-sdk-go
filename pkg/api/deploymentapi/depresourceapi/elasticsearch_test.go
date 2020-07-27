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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var defaultESTopologies = []*models.ElasticsearchClusterTopologyElement{
	{
		InstanceConfigurationID: "default.data",
		Size: &models.TopologySize{
			Resource: ec.String("memory"),
			Value:    ec.Int32(1024),
		},
		NodeType: &models.ElasticsearchNodeType{
			Data: ec.Bool(true),
		},
	},
	{
		InstanceConfigurationID: "default.master",
		Size: &models.TopologySize{
			Resource: ec.String("memory"),
			Value:    ec.Int32(1024),
		},
		NodeType: &models.ElasticsearchNodeType{
			Master: ec.Bool(true),
		},
	},
	{
		InstanceConfigurationID: "default.ml",
		Size: &models.TopologySize{
			Resource: ec.String("memory"),
			Value:    ec.Int32(1024),
		},
		NodeType: &models.ElasticsearchNodeType{
			Ml: ec.Bool(true),
		},
	},
}

var elasticsearchTemplateResponse = models.DeploymentTemplateInfoV2{
	ID: ec.String("default"),
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

func TestNewElasticsearch(t *testing.T) {
	type args struct {
		params NewElasticsearchParams
	}
	tests := []struct {
		name string
		args args
		want *models.ElasticsearchPayload
		err  error
	}{
		{
			name: "fails due to parameter validation",
			args: args{params: NewElasticsearchParams{
				Topology: []ElasticsearchTopologyElement{
					{},
				},
			}},
			err: multierror.NewPrefixed("invalid deployment resource params",
				errors.New("deployment template info is not specified and is required for the operation"),
				errors.New("region cannot be empty"),
				errors.New("version cannot be empty"),
				errors.New("element[0]: elasticsearch topology: node_type cannot be empty"),
				errors.New("element[0]: elasticsearch topology: size cannot be empty"),
			),
		},
		{
			name: "fails due to unknown desired topology",
			args: args{params: NewElasticsearchParams{
				Region:                   "ece-region",
				Version:                  "7.4.2",
				TemplateID:               "default",
				DeploymentTemplateInfoV2: &elasticsearchTemplateResponse,
				Topology: []ElasticsearchTopologyElement{
					{NodeType: "some", Size: 1024},
				},
			}},
			err: errors.New(`deployment topology: failed to obtain desired topology names ([{NodeType:some ZoneCount:0 Size:1024}]) in deployment template id "default"`),
		},
		{
			name: "fails due to unknown invalid template",
			args: args{params: NewElasticsearchParams{
				Region:     "ece-region",
				Version:    "7.4.2",
				TemplateID: "default",
				DeploymentTemplateInfoV2: &models.DeploymentTemplateInfoV2{
					ID: ec.String("default"),
					DeploymentTemplate: &models.DeploymentCreateRequest{
						Resources: &models.DeploymentCreateResources{},
					},
				},
				Topology: []ElasticsearchTopologyElement{
					{NodeType: "some", Size: 1024},
				},
			}},
			err: errors.New("deployment: the default template is not configured for Elasticsearch. Please use another template if you wish to start Elasticsearch instances"),
		},
		{
			name: "Returns the default topology",
			args: args{params: NewElasticsearchParams{
				Region:                   "ece-region",
				Version:                  "7.4.2",
				TemplateID:               "default",
				DeploymentTemplateInfoV2: &elasticsearchTemplateResponse,
			}},
			want: &models.ElasticsearchPayload{
				DisplayName: "",
				Region:      ec.String("ece-region"),
				RefID:       ec.String(DefaultElasticsearchRefID),
				Plan: &models.ElasticsearchClusterPlan{
					Elasticsearch: &models.ElasticsearchConfiguration{
						Version: "7.4.2",
					},
					DeploymentTemplate: &models.DeploymentTemplateReference{
						ID: ec.String(DefaultTemplateID),
					},
					ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
						{
							ZoneCount:               1,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(4096),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						},
					},
				},
			},
		},
		{
			name: "Returns a custom topology",
			args: args{params: NewElasticsearchParams{
				Region:                   "ece-region",
				Version:                  "7.4.2",
				TemplateID:               "default",
				DeploymentTemplateInfoV2: &elasticsearchTemplateResponse,
				Topology: []ElasticsearchTopologyElement{
					{NodeType: DataNode, Size: 8192, ZoneCount: 2},
					{NodeType: MasterNode, Size: 1024, ZoneCount: 1},
					{NodeType: MLNode, Size: 2048, ZoneCount: 1},
				},
			}},
			want: &models.ElasticsearchPayload{
				DisplayName: "",
				Region:      ec.String("ece-region"),
				RefID:       ec.String(DefaultElasticsearchRefID),
				Plan: &models.ElasticsearchClusterPlan{
					Elasticsearch: &models.ElasticsearchConfiguration{
						Version: "7.4.2",
					},
					DeploymentTemplate: &models.DeploymentTemplateReference{
						ID: ec.String(DefaultTemplateID),
					},
					ClusterTopology: []*models.ElasticsearchClusterTopologyElement{
						{
							ZoneCount:               2,
							InstanceConfigurationID: "default.data",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(8192),
							},
							NodeType: &models.ElasticsearchNodeType{
								Data: ec.Bool(true),
							},
						},
						{
							ZoneCount:               1,
							InstanceConfigurationID: "default.master",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(1024),
							},
							NodeType: &models.ElasticsearchNodeType{
								Master: ec.Bool(true),
							},
						},
						{
							ZoneCount:               1,
							InstanceConfigurationID: "default.ml",
							Size: &models.TopologySize{
								Resource: ec.String("memory"),
								Value:    ec.Int32(2048),
							},
							NodeType: &models.ElasticsearchNodeType{
								Ml: ec.Bool(true),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewElasticsearch(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
