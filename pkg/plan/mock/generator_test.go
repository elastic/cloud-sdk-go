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

package planmock

import (
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGenerate(t *testing.T) {
	var deploymentID = ec.RandomResourceID()
	var apmID = ec.RandomResourceID()
	var esID = ec.RandomResourceID()
	var kibanaID = ec.RandomResourceID()
	var appsearchID = ec.RandomResourceID()
	var enterpriseSearchID = ec.RandomResourceID()
	var ApmCurrentLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("completed")},
	}
	var EsCurrentLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("completed")},
	}
	var kibanaCurrentLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("completed")},
	}
	var AppsearchCurrentLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("completed")},
	}
	var EnterpriseSearchCurrentLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("completed")},
	}
	var ApmPendingLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("in_progress")},
	}
	var EsPendingLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("in_progress")},
	}
	var kibanaPendingLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("in_progress")},
	}
	var AppsearchPendingLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("in_progress")},
	}
	var EnterpriseSearchPendingLog = []*models.ClusterPlanStepInfo{
		{Stage: ec.String("in_progress")},
	}
	type args struct {
		cfg GenerateConfig
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentGetResponse
	}{
		{
			name: "Generate a DeploymentGetResponse",
			args: args{cfg: GenerateConfig{
				ID: deploymentID,
				Apm: []GeneratedResourceConfig{
					{
						ID:         apmID,
						CurrentLog: ApmCurrentLog,
						PendingLog: ApmPendingLog,
					},
				},
				Elasticsearch: []GeneratedResourceConfig{
					{
						ID:         esID,
						CurrentLog: EsCurrentLog,
						PendingLog: EsPendingLog,
					},
				},
				Kibana: []GeneratedResourceConfig{
					{
						ID:         kibanaID,
						CurrentLog: kibanaCurrentLog,
						PendingLog: kibanaPendingLog,
					},
				},
				Appsearch: []GeneratedResourceConfig{
					{
						ID:         appsearchID,
						CurrentLog: AppsearchCurrentLog,
						PendingLog: AppsearchPendingLog,
					},
				},
				EnterpriseSearch: []GeneratedResourceConfig{
					{
						ID:         enterpriseSearchID,
						CurrentLog: EnterpriseSearchCurrentLog,
						PendingLog: EnterpriseSearchPendingLog,
					},
				},
			}},
			want: &models.DeploymentGetResponse{
				ID: ec.String(deploymentID),
				Resources: &models.DeploymentResources{
					Apm: generateApmResourceInfo([]GeneratedResourceConfig{
						{
							ID:         apmID,
							CurrentLog: ApmCurrentLog,
							PendingLog: ApmPendingLog,
						},
					}),
					Elasticsearch: generateElasticsearchResourceInfo([]GeneratedResourceConfig{
						{
							ID:         esID,
							CurrentLog: EsCurrentLog,
							PendingLog: EsPendingLog,
						},
					}),
					Kibana: generateKibanaResourceInfo([]GeneratedResourceConfig{
						{
							ID:         kibanaID,
							CurrentLog: kibanaCurrentLog,
							PendingLog: kibanaPendingLog,
						},
					}),
					Appsearch: generateAppSearchResourceInfo([]GeneratedResourceConfig{
						{
							ID:         appsearchID,
							CurrentLog: AppsearchCurrentLog,
							PendingLog: AppsearchPendingLog,
						},
					}),
					EnterpriseSearch: generateEnterpriseSearchResourceInfo([]GeneratedResourceConfig{
						{
							ID:         enterpriseSearchID,
							CurrentLog: EnterpriseSearchCurrentLog,
							PendingLog: EnterpriseSearchPendingLog,
						},
					}),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
