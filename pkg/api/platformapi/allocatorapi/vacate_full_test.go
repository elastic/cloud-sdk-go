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

package allocatorapi

import (
	"errors"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/output"
	"github.com/elastic/cloud-sdk-go/pkg/plan"
	sdkSync "github.com/elastic/cloud-sdk-go/pkg/sync"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestVacate(t *testing.T) {
	type args struct {
		params *VacateParams
		buf    *sdkSync.Buffer
	}
	tests := []struct {
		name string
		args args
		err  string
		want string
	}{
		{
			name: "Succeeds moving a single ES cluster from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region: "us-east-1",
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							elasticsearch: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving a single kibana instance from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region: "us-east-1",
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							kibana: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving a single APM cluster from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region: "us-east-1",
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							apm: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Apm][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Apm][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving a single Appsearch cluster from a single allocator (without tracking)",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region:       "us-east-1",
					skipTracking: true,
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							appsearch: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
		},
		{
			name: "Succeeds moving a single Appsearch cluster from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region: "us-east-1",
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							appsearch: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Appsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Appsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving a single Enterprise Search cluster from a single allocator (without tracking)",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region:       "us-east-1",
					skipTracking: true,
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							EnterpriseSearch: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
		},
		{
			name: "Succeeds moving a single Enterprise Search cluster from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{
					region: "us-east-1",
					topology: []vacateCaseClusters{
						{
							Allocator: "allocatorID",
							EnterpriseSearch: []vacateCaseClusterConfig{
								{
									ID: "3ee11eb40eda22cac0cce259625c6734",
									steps: [][]*models.ClusterPlanStepInfo{
										{
											newPlanStep("step1", "success"),
											newPlanStep("step2", "pending"),
										},
									},
									plan: []*models.ClusterPlanStepInfo{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("plan-completed", "success"),
									},
								},
							},
						},
					},
				}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Enterprise Search][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Enterprise Search][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving a multiple clusters from a single allocator",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "3ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Succeeds moving multiple clusters from multiple allocators",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "3ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
		},
		{
			name: "Moving multiple clusters from multiple allocators that contain track failures",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "3ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStepWithDetails("step2", "success", nil),
										newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
											Message: ec.String(planStepLogErrorMessage),
										}}),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
									newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				`Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: running step "step3" (Plan duration )...`,
				"\x1b[91;1mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: caught error: \"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]\"\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
			err: `vacate error: 1 error occurred:
	* found deployment plan errors: deployment [DISCOVERED_DEPLOYMENT_ID] - [elasticsearch][5ee11eb40eda22cac0cce259625c6734]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"

`,
		},
		{
			name: "Moving an Elasticsearch clusters from an allocator where the plan finishes too fast and an error is thrown",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{{
							ID:    "3ee11eb40eda22cac0cce259625c6734",
							steps: [][]*models.ClusterPlanStepInfo{},
							plan: []*models.ClusterPlanStepInfo{
								newPlanStep("step1", "success"),
								newPlanStep("step2", "success"),
								newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
									Message: ec.String(planStepLogErrorMessage),
								}}),
							},
						}},
					},
				}}),
			},
			want: newOutputResponses(
				"\x1b[91;1mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: caught error: \"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]\"\x1b[0m (Total plan duration )",
			),
			err: `vacate error: 1 error occurred:
	* found deployment plan errors: deployment [DISCOVERED_DEPLOYMENT_ID] - [elasticsearch][3ee11eb40eda22cac0cce259625c6734]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"

`,
		},
		{
			name: "Moving multiple clusters from multiple allocators that fail to move",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID:   "3ee11eb40eda22cac0cce259625c6734",
								fail: true,
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
											Message: ec.String(planStepLogErrorMessage),
										}}),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
									newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][2ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				`Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: running step "step3" (Plan duration )...`,
				"\x1b[91;1mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][5ee11eb40eda22cac0cce259625c6734]: caught error: \"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]\"\x1b[0m (Total plan duration )",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Kibana][4ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
			err: `vacate error: 2 errors occurred:
	* allocator allocatorID-1: resource id [3ee11eb40eda22cac0cce259625c6734][elasticsearch]: failed vacating: a message (a code)
	* found deployment plan errors: deployment [DISCOVERED_DEPLOYMENT_ID] - [elasticsearch][5ee11eb40eda22cac0cce259625c6734]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"

`,
		},
		{
			name: "Moving multiple clusters from multiple allocators that fail to move with JSON format",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{outputFormat: "json", region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID:   "3ee11eb40eda22cac0cce259625c6734",
								fail: true,
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
											Message: ec.String(planStepLogErrorMessage),
										}}),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
									newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step3","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"plan-completed","err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s","finished":true}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
			),
			err: `{
  "errors": [
    {
      "allocator_id": "allocatorID-1",
      "context": "failed vacating",
      "error": {
        "message": "a message (a code)"
      },
      "kind": "elasticsearch",
      "resource_id": "3ee11eb40eda22cac0cce259625c6734"
    },
    {
      "deployment_id": "DISCOVERED_DEPLOYMENT_ID",
      "err": {
        "message": "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
      },
      "finished": true,
      "id": "5ee11eb40eda22cac0cce259625c6734",
      "kind": "elasticsearch",
      "ref_id": "main-elasticsearch",
      "step": "plan-completed"
    }
  ]
}
`,
		},
		{
			name: "Moving multiple clusters from multiple allocators that fail to move with JSON format with internal details",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{outputFormat: "json", region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID:   "3ee11eb40eda22cac0cce259625c6734",
								fail: true,
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
											Message: ec.String(planStepLogErrorMessage),
										}}),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
									newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message:     ec.String("Failed to detect running cluster - instance was not detected as running in time. Check the health of the cluster, and look at the instance and/or allocator logs to determine if there were any issues starting."),
										FailureType: "ClusterFailure:InstanceDidNotStartWhileWaitingForRunning",
										InternalDetails: map[string]interface{}{
											"details": "The state did not become the desired one before [600000 milliseconds] elapsed. Last error was: [Instance is not running [instance-0000000038]. Please check allocator/docker logs.]",
										},
									}}),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step3","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"plan-completed","err":{"message":"Failed to detect running cluster - instance was not detected as running in time. Check the health of the cluster, and look at the instance and/or allocator logs to determine if there were any issues starting."},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s":"The state did not become the desired one before [600000 milliseconds] elapsed. Last error was: [Instance is not running [instance-0000000038]. Please check allocator/docker logs.]"}},"finished":true}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
			),
			err: `{
  "errors": [
    {
      "allocator_id": "allocatorID-1",
      "context": "failed vacating",
      "error": {
        "message": "a message (a code)"
      },
      "kind": "elasticsearch",
      "resource_id": "3ee11eb40eda22cac0cce259625c6734"
    },
    {
      "deployment_id": "DISCOVERED_DEPLOYMENT_ID",
      "err": {
        "message": "Failed to detect running cluster - instance was not detected as running in time. Check the health of the cluster, and look at the instance and/or allocator logs to determine if there were any issues starting."
      },
      "failure_details": {
        "details": null,
        "failure_type": "ClusterFailure:InstanceDidNotStartWhileWaitingForRunning",
        "internal": {
          "details": "The state did not become the desired one before [600000 milliseconds] elapsed. Last error was: [Instance is not running [instance-0000000038]. Please check allocator/docker logs.]"
        }
      },
      "finished": true,
      "id": "5ee11eb40eda22cac0cce259625c6734",
      "kind": "elasticsearch",
      "ref_id": "main-elasticsearch",
      "step": "plan-completed"
    }
  ]
}
`,
		},
		{
			name: "Moving multiple clusters from multiple allocators that fail to move with JSON format with details",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: newVacateTestCase(t, vacateCase{outputFormat: "json", region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID-1",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID:   "3ee11eb40eda22cac0cce259625c6734",
								fail: true,
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
					{
						Allocator: "allocatorID-2",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "5ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
											Message: ec.String(planStepLogErrorMessage),
										}}),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetails("step3", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message: ec.String(planStepLogErrorMessage),
									}}),
									newPlanStepWithDetails("plan-completed", "error", []*models.ClusterPlanStepLogMessageInfo{{
										Message:     ec.String("An unexpected error was encountered during step [rolling-upgrade]. Ensure there are no issue with your deployment and then attempt to re-run the plan."),
										FailureType: "UnknownFailure:UnknownErrorEncountered",
										Details:     map[string]string{"stepId": "rolling-upgrade"},
									}}),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "4ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("plan-completed", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"2ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"step3","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s"}`,
				`{"id":"5ee11eb40eda22cac0cce259625c6734","kind":"elasticsearch","step":"plan-completed","err":{"message":"An unexpected error was encountered during step [rolling-upgrade]. Ensure there are no issue with your deployment and then attempt to re-run the plan."},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-elasticsearch","duration":"s":{"stepId":"rolling-upgrade"},"failure_type":"UnknownFailure:UnknownErrorEncountered"},"finished":true}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"step2","deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s"}`,
				`{"id":"4ee11eb40eda22cac0cce259625c6734","kind":"kibana","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"DISCOVERED_DEPLOYMENT_ID","ref_id":"main-kibana","duration":"s","finished":true}`,
			),
			err: `{
  "errors": [
    {
      "allocator_id": "allocatorID-1",
      "context": "failed vacating",
      "error": {
        "message": "a message (a code)"
      },
      "kind": "elasticsearch",
      "resource_id": "3ee11eb40eda22cac0cce259625c6734"
    },
    {
      "deployment_id": "DISCOVERED_DEPLOYMENT_ID",
      "err": {
        "message": "An unexpected error was encountered during step [rolling-upgrade]. Ensure there are no issue with your deployment and then attempt to re-run the plan."
      },
      "failure_details": {
        "details": {
          "stepId": "rolling-upgrade"
        },
        "failure_type": "UnknownFailure:UnknownErrorEncountered"
      },
      "finished": true,
      "id": "5ee11eb40eda22cac0cce259625c6734",
      "kind": "elasticsearch",
      "ref_id": "main-elasticsearch",
      "step": "plan-completed"
    }
  ]
}
`,
		},
		{
			name: "Vacate has some failures",
			args: args{
				buf: sdkSync.NewBuffer(),
				params: &VacateParams{
					Allocators:     []string{"allocatorID"},
					Concurrency:    1,
					Region:         "us-east-1",
					MaxPollRetries: 1,
					TrackFrequency: time.Nanosecond,
					API: api.NewMock(mock.Response{Response: http.Response{
						Body:       newKibanaMoveFailure(t, "3ee11eb40eda22cac0cce259625c6734", "allocatorID"),
						StatusCode: 202,
					}}),
				},
			},
			err: multierror.NewPrefixed("vacate error",
				VacateError{
					AllocatorID: "allocatorID",
					ResourceID:  "3ee11eb40eda22cac0cce259625c6734",
					Kind:        "kibana",
					Ctx:         "failed vacating",
					Err:         errors.New("failed for reason (some code)"),
				},
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.buf != nil {
				tt.args.params.Output = output.NewDevice(tt.args.buf)
			}

			if err := Vacate(tt.args.params); err != nil {
				// Set duration to 0
				var merr *multierror.Prefixed
				if errors.As(err, &merr) {
					for i, e := range merr.Errors {
						var tr plan.TrackResponse
						if errors.As(e, &tr) {
							tr.Duration = 0
							merr.Errors[i] = tr
						}
					}
				}
				if !assert.EqualError(t, err, tt.err) {
					t.Errorf("Vacate() error = %v, wantErr %v", err, tt.err)
				}
			}

			var got string
			if tt.args.buf != nil {
				if tt.args.params.OutputFormat == "json" {
					got = regexp.MustCompile(`"duration".*s"`).
						ReplaceAllString(tt.args.buf.String(), `"duration":"s"`)
				} else {
					got = regexp.MustCompile(`duration.*\)`).
						ReplaceAllString(tt.args.buf.String(), "duration )")
				}
			}

			if tt.args.buf != nil && !assert.Equal(t, tt.want, got) {
				t.Errorf("VacateCluster() output = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}

func TestVacateInterrupt(t *testing.T) {
	type args struct {
		params *VacateParams
	}
	tests := []struct {
		name string
		args args
		err  string
		want string
	}{
		{
			name: "Interrupts the vacate",
			args: args{
				params: newVacateTestCase(t, vacateCase{region: "us-east-1", topology: []vacateCaseClusters{
					{
						Allocator: "allocatorID",
						elasticsearch: []vacateCaseClusterConfig{
							{
								ID: "3ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("step3", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("step3", "success"),
										newPlanStep("step4", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "success"),
									newPlanStep("step4", "success"),
								},
							},
						},
						kibana: []vacateCaseClusterConfig{
							{
								ID: "2ee11eb40eda22cac0cce259625c6734",
								steps: [][]*models.ClusterPlanStepInfo{
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("step3", "pending"),
									},
									{
										newPlanStep("step1", "success"),
										newPlanStep("step2", "success"),
										newPlanStep("step3", "success"),
										newPlanStep("step4", "pending"),
									},
								},
								plan: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "success"),
									newPlanStep("step4", "success"),
								},
							},
						},
					},
				}}),
			},
			want: newOutputResponses(
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step2\" (Plan duration )...",
				"pool: received interrupt, stopping pool...",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step3\" (Plan duration )...",
				"Deployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: running step \"step4\" (Plan duration )...",
				"\x1b[92;mDeployment [DISCOVERED_DEPLOYMENT_ID] - [Elasticsearch][3ee11eb40eda22cac0cce259625c6734]: finished running all the plan steps\x1b[0m (Total plan duration )",
			),
			err: multierror.NewPrefixed("vacate error",
				errors.New("allocator allocatorID: resource id [2ee11eb40eda22cac0cce259625c6734][kibana]: was either cancelled or not processed, follow up accordingly"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var once sync.Once
			// Sends a signal once the first write is received.
			buf := sdkSync.NewBuffer(func() {
				once.Do(func() {
					p, err := os.FindProcess(os.Getpid())
					if err != nil {
						t.Fatal(err)
					}

					if err := p.Signal(os.Interrupt); err != nil {
						t.Fatal(err)
					}
				})
			})
			tt.args.params.Output = output.NewDevice(buf)

			if err := Vacate(tt.args.params); err != nil && err.Error() != tt.err {
				t.Errorf("Vacate() error = %v, wantErr %v", err, tt.err)
			}

			var got string
			if buf != nil {
				got = regexp.MustCompile(`duration.*\)`).
					ReplaceAllString(buf.String(), "duration )")
			}

			if buf != nil && tt.want != got {
				wantLines := strings.Split(tt.want, "\n")
				var matched bool
				for _, want := range wantLines {
					if strings.Contains(got, want) {
						matched = true
					} else {
						matched = false
					}
				}

				if !matched {
					t.Errorf("VacateCluster() output = \n%v, want \n%v", got, tt.want)
				}
			}
		})
	}
}
