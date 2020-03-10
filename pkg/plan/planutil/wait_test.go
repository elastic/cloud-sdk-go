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

package planutil

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/plan"
	planmock "github.com/elastic/cloud-sdk-go/pkg/plan/mock"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestWait(t *testing.T) {
	deploymentID := ec.RandomResourceID()
	var pendingPlan = planmock.Generate(planmock.GenerateConfig{
		ID: deploymentID,
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
	var currentPlanWithError = planmock.Generate(planmock.GenerateConfig{
		ID: deploymentID,
		Elasticsearch: []planmock.GeneratedResourceConfig{
			{
				ID: "cde7b6b605424a54ce9d56316eab13a1",
				CurrentLog: planmock.NewPlanStepLog(
					planmock.NewPlanStep("step-1", "success"),
					planmock.NewPlanStep("step-2", "success"),
					planmock.NewPlanStep("step-3", "success"),
					planmock.NewPlanStep("step-4", "success"),
					planmock.NewPlanStep("step-5", "success"),
					planmock.NewPlanStepWithDetailsAndError("plan-completed", []*models.ClusterPlanStepLogMessageInfo{{
						Message: ec.String("some nasty  error"),
					}}),
				),
			},
		},
	})
	type args struct {
		params plan.TrackChangeParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "waits",
			args: args{params: plan.TrackChangeParams{
				API:          api.NewMock(),
				DeploymentID: ec.RandomResourceID(),
			}},
		},
		{
			name: "waits and returns error",
			args: args{params: plan.TrackChangeParams{
				API: api.NewMock(
					mock.New200StructResponse(pendingPlan),
					mock.New200StructResponse(currentPlanWithError),
					mock.New200StructResponse(currentPlanWithError),
				),
				DeploymentID: deploymentID,
			}},
			err: fmt.Errorf(
				fmt.Sprintf(`deployment [%s] - [elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: caught error: "some nasty  error"`, deploymentID),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if err := Wait(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Wait() error = %v, wantErr %v", err, tt.err)
			}

			expectedRunDuration := tt.args.params.Config.PollFrequency * time.Duration(tt.args.params.Config.MaxRetries+1)
			if time.Since(start) < expectedRunDuration {
				t.Errorf("BROKEN")
			}
		})
	}
}
