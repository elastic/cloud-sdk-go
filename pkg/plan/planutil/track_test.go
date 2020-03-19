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
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/plan"
	planmock "github.com/elastic/cloud-sdk-go/pkg/plan/mock"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestTrackChange(t *testing.T) {
	var textBufErr = new(bytes.Buffer)
	var textBufSuccess = new(bytes.Buffer)
	var textBufErrJSON = new(bytes.Buffer)
	var textBufSuccessJSON = new(bytes.Buffer)
	var deploymentID = ec.RandomResourceID()
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
						Message: ec.String("some nasty error"),
					}}),
				),
			},
		},
	})
	var wantBufErr = fmt.Sprintf(
		"Deployment [%s] - [Elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: running step \"step-4\"\n\x1b[91;1mDeployment [%s] - [Elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: caught error: \"some nasty error\"\x1b[0m\n",
		deploymentID, deploymentID,
	)
	var currentPlan = planmock.Generate(planmock.GenerateConfig{
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
					planmock.NewPlanStep("plan-completed", "success"),
				),
			},
		},
	})
	var wantBufSuccess = fmt.Sprintf(
		"Deployment [%s] - [Elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: running step \"step-4\"\n\x1b[92;mDeployment [%s] - [Elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: finished running all the plan steps\x1b[0m\n",
		deploymentID, deploymentID,
	)
	var wantBufErrJSON = fmt.Sprintf(`{"id":"cde7b6b605424a54ce9d56316eab13a1","kind":"elasticsearch","step":"step-4","deployment_id":"%s","ref_id":"main-elasticsearch"}
{"id":"cde7b6b605424a54ce9d56316eab13a1","kind":"elasticsearch","step":"plan-completed","err":{"message":"some nasty error"},"deployment_id":"%s","ref_id":"main-elasticsearch","finished":true}`,
		deploymentID, deploymentID,
	) + "\n"
	var wantBufSuccessJSON = fmt.Sprintf(`{"id":"cde7b6b605424a54ce9d56316eab13a1","kind":"elasticsearch","step":"step-4","deployment_id":"%s","ref_id":"main-elasticsearch"}
{"id":"cde7b6b605424a54ce9d56316eab13a1","kind":"elasticsearch","step":"plan-completed","err":{"message":"finished all the plan steps"},"deployment_id":"%s","ref_id":"main-elasticsearch","finished":true}`,
		deploymentID, deploymentID,
	) + "\n"
	type args struct {
		params TrackChangeParams
	}
	tests := []struct {
		name    string
		args    args
		err     error
		wantBuf string
	}{
		{
			name: "returns error on parameter validation failure (downstream)",
			args: args{params: TrackChangeParams{}},
			err: &multierror.Error{Errors: []error{
				errors.New("plan track change: API cannot be nil"),
				errors.New("plan track change: one of DeploymentID or ResourceID must be specified"),
				errors.New("plan track change: Kind cannot be empty"),
			}},
		},
		{
			name: "returns error on parameter validation failure",
			args: args{params: TrackChangeParams{
				Format: "some",
			}},
			err: &multierror.Error{Errors: []error{
				errors.New("planutil track change: writer needs to be specified when format is not empty"),
				errors.New(`planutil track change: invalid format "some"`),
			}},
		},
		{
			name: "tracks a cluster with text format and error result",
			args: args{params: TrackChangeParams{
				TrackChangeParams: plan.TrackChangeParams{
					API: api.NewMock(
						mock.New200StructResponse(pendingPlan),
						mock.New200StructResponse(currentPlanWithError),
						mock.New200StructResponse(currentPlanWithError),
					),
					DeploymentID: deploymentID,
				},
				Format: "text",
				Writer: textBufErr,
			}},
			err: fmt.Errorf(
				`deployment [%s] - [elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: caught error: "some nasty error"`,
				deploymentID,
			),
			wantBuf: wantBufErr,
		},
		{
			name: "tracks a cluster with text format and success",
			args: args{params: TrackChangeParams{
				TrackChangeParams: plan.TrackChangeParams{
					API: api.NewMock(
						mock.New200StructResponse(pendingPlan),
						mock.New200StructResponse(currentPlan),
						mock.New200StructResponse(currentPlan),
					),
					DeploymentID: deploymentID,
				},
				Format: "text",
				Writer: textBufSuccess,
			}},
			wantBuf: wantBufSuccess,
		},
		{
			name: "tracks a cluster with json format and error result",
			args: args{params: TrackChangeParams{
				TrackChangeParams: plan.TrackChangeParams{
					API: api.NewMock(
						mock.New200StructResponse(pendingPlan),
						mock.New200StructResponse(currentPlanWithError),
						mock.New200StructResponse(currentPlanWithError),
					),
					DeploymentID: deploymentID,
				},
				Format: "json",
				Writer: textBufErrJSON,
			}},
			err: fmt.Errorf(
				`deployment [%s] - [elasticsearch][cde7b6b605424a54ce9d56316eab13a1]: caught error: "some nasty error"`,
				deploymentID,
			),
			wantBuf: wantBufErrJSON,
		},
		{
			name: "tracks a cluster with text format and success",
			args: args{params: TrackChangeParams{
				TrackChangeParams: plan.TrackChangeParams{
					API: api.NewMock(
						mock.New200StructResponse(pendingPlan),
						mock.New200StructResponse(currentPlan),
						mock.New200StructResponse(currentPlan),
					),
					DeploymentID: deploymentID,
				},
				Format: "json",
				Writer: textBufSuccessJSON,
			}},
			wantBuf: wantBufSuccessJSON,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := TrackChange(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("TrackChange() error = %v, wantErr %v", err, tt.err)
			}
			// Remove all of the duration timestamps.
			pattern := `(?mi).\(.*plan duration.*`
			if tt.args.params.Format == "json" {
				pattern = `,"duration":".*\.\w+ms"`
			}
			re, err := regexp.Compile(pattern)
			if err != nil {
				t.Fatal(err)
			}
			if buf, ok := tt.args.params.Writer.(*bytes.Buffer); ok {
				out := re.ReplaceAllString(buf.String(), "")
				if out != tt.wantBuf {
					t.Errorf("TrackChange() output = %v, wantOutput %v", out, tt.wantBuf)
				}
			}
		})
	}
}
