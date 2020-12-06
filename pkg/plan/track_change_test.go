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

package plan

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	planmock "github.com/elastic/cloud-sdk-go/pkg/plan/mock"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestTrackChange(t *testing.T) {
	var foundDeploymentIDResponse = models.DeploymentsSearchResponse{
		Deployments: []*models.DeploymentSearchResponse{
			{ID: ec.String("cbb4bc6c09684c86aa5de54c05ea1d38")},
		},
	}
	var pendingPlan = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
	})
	var pendingPlanWithDownstreamChanges = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Kibana: []planmock.GeneratedResourceConfig{
			{
				ID: "4de9b2b605424a54ce9d56316eab13a8",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
		Apm: []planmock.GeneratedResourceConfig{
			{
				ID: "5de9b2b605424a54ce9d56316eab13a5",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
		Appsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
		EnterpriseSearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "pending"),
				),
			},
		},
	})
	var secondPendingPlanWithDownstreamChanges = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Kibana: []planmock.GeneratedResourceConfig{
			{
				ID: "4de9b2b605424a54ce9d56316eab13a8",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
		Apm: []planmock.GeneratedResourceConfig{
			{
				ID: "5de9b2b605424a54ce9d56316eab13a5",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
		Appsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
		EnterpriseSearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
	})
	var secondPendingPlan = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				PendingLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "pending"),
				),
			},
		},
	})
	var noMorePendingPlan = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID:         "cde7b6b605424a54ce9d56316eab13a1",
				PendingLog: nil,
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
	})
	var currentPlan = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
	})
	var historyPlan = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				HistoryLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
	})
	var currentPlanWithDownstream = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Kibana: []planmock.GeneratedResourceConfig{
			{
				ID: "4de9b2b605424a54ce9d56316eab13a8",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Apm: []planmock.GeneratedResourceConfig{
			{
				ID: "5de9b2b605424a54ce9d56316eab13a5",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Appsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		EnterpriseSearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
	})
	var historyPlanWithDownstream = planmock.Generate(planmock.GenerateConfig{
		ID: "cbb4bc6c09684c86aa5de54c05ea1d38",
		Kibana: []planmock.GeneratedResourceConfig{
			{
				ID: "4de9b2b605424a54ce9d56316eab13a8",
				HistoryLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Apm: []planmock.GeneratedResourceConfig{
			{
				ID: "5de9b2b605424a54ce9d56316eab13a5",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Appsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		EnterpriseSearch: []planmock.GeneratedResourceConfig{
			{
				ID: "6de9b2b605424a54ce9d56316eab13a6",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				HistoryLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				),
			},
		},
	})

	type args struct {
		params TrackChangeParams
	}
	tests := []struct {
		name string
		args args
		want []TrackResponse
		err  string
	}{
		{
			name: "errors out on parameter validation",
			args: args{},
			err: multierror.NewPrefixed("plan track change",
				errors.New("API cannot be nil"),
				errors.New("one of DeploymentID or ResourceID must be specified"),
				errors.New("kind cannot be empty"),
			).Error(),
		},
		{
			name: "errors out when both DeploymentID and ResourceID are specified",
			args: args{params: TrackChangeParams{
				ResourceID:   "cde7b6b605424a54ce9d56316eab13a1",
				DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38",
			}},
			err: multierror.NewPrefixed("plan track change",
				errors.New("API cannot be nil"),
				errors.New("cannot specify both DeploymentID and ResourceID"),
			).Error(),
		},
		{
			name: "fails looking up the deploymentID",
			args: args{params: TrackChangeParams{
				API: api.NewMock(mock.New404Response(
					mock.NewStringBody("some error"),
				)),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			err: "some error",
		},
		{
			name: "fails looking up the deploymentID (No DeploymentID found)",
			args: args{params: TrackChangeParams{
				API: api.NewMock(mock.New200Response(
					mock.NewStructBody(models.DeploymentsSearchResponse{
						Deployments: nil,
					}),
				)),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			err: "plan track change: couldn't find a deployment containing Kind elasticsearch with ID cde7b6b605424a54ce9d56316eab13a1",
		},
		{
			name: "looks up the deploymentID and tracks the change",
			args: args{params: TrackChangeParams{
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlan),
					mock.New200StructResponse(secondPendingPlan),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(currentPlan),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change obtaining the plan history",
			args: args{params: TrackChangeParams{
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlan),
					mock.New200StructResponse(secondPendingPlan),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(historyPlan),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change allowing for an extra error to happen",
			args: args{params: TrackChangeParams{
				Config: TrackFrequencyConfig{
					MaxRetries: 2,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlan),
					mock.New200StructResponse(secondPendingPlan),
					mock.New500Response(mock.NewStringBody("some error")),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(currentPlan),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change, ignoring downstream changes",
			args: args{params: TrackChangeParams{
				IgnoreDownstream: true,
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlanWithDownstreamChanges),
					mock.New200StructResponse(secondPendingPlanWithDownstreamChanges),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(currentPlanWithDownstream),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change, without ignoring downstream changes",
			args: args{params: TrackChangeParams{
				IgnoreDownstream: false,
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlanWithDownstreamChanges),
					mock.New200StructResponse(secondPendingPlanWithDownstreamChanges),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(currentPlanWithDownstream),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana"},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana"},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change, ignoring downstream changes (Gets plan history at the end)",
			args: args{params: TrackChangeParams{
				IgnoreDownstream: true,
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlanWithDownstreamChanges),
					mock.New200StructResponse(secondPendingPlanWithDownstreamChanges),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(historyPlanWithDownstream),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change, without ignoring downstream changes (Gets plan history at the end)",
			args: args{params: TrackChangeParams{
				IgnoreDownstream: false,
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlanWithDownstreamChanges),
					mock.New200StructResponse(secondPendingPlanWithDownstreamChanges),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New200StructResponse(historyPlanWithDownstream),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana"},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana"},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch"},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "4de9b2b605424a54ce9d56316eab13a8", Kind: "kibana", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-kibana", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "5de9b2b605424a54ce9d56316eab13a5", Kind: "apm", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-apm", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "appsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-appsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
				{ID: "6de9b2b605424a54ce9d56316eab13a6", Kind: "enterprise_search", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-enterprise_search", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
		{
			name: "looks up the deploymentID and tracks the change, retrying the plan until MaxRetries is hit",
			args: args{params: TrackChangeParams{
				Config: TrackFrequencyConfig{
					MaxRetries: 1,
				},
				API: api.NewMock(
					mock.New200StructResponse(foundDeploymentIDResponse),
					mock.New200StructResponse(pendingPlan),
					mock.New200StructResponse(secondPendingPlan),
					mock.New200StructResponse(noMorePendingPlan),
					mock.New500Response(mock.NewStringBody("error")),
					mock.New200StructResponse(currentPlan),
				),
				ResourceID: "cde7b6b605424a54ce9d56316eab13a1",
				Kind:       "elasticsearch",
			}},
			want: []TrackResponse{
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-4", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", Step: "step-5", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch"},
				{ID: "cde7b6b605424a54ce9d56316eab13a1", Kind: "elasticsearch", DeploymentID: "cbb4bc6c09684c86aa5de54c05ea1d38", RefID: "main-elasticsearch", Step: "plan-completed", Finished: true, Err: ErrPlanFinished},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TrackChange(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("TrackChange() error = %v, wantErr %v", err, tt.err)
				return
			}

			if got != nil {
				var gotResponses []TrackResponse
				var gotBuf = new(bytes.Buffer)
				for res := range got {
					// nullify duration
					res.Duration = 0
					json.NewEncoder(gotBuf).Encode(res)
					gotResponses = append(gotResponses, res)
				}

				if !reflect.DeepEqual(gotResponses, tt.want) {
					var wantBuf = new(bytes.Buffer)
					for _, w := range tt.want {
						json.NewEncoder(wantBuf).Encode(w)
					}
					t.Errorf("TrackChange() = \n%v, want \n%v", gotBuf.String(), wantBuf.String())
				}
			}
		})
	}
}
