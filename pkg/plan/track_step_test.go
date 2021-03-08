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
	"errors"
	"reflect"
	"testing"

	planmock "github.com/elastic/cloud-sdk-go/pkg/plan/mock"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

func TestGetStepName(t *testing.T) {
	type args struct {
		log []*models.ClusterPlanStepInfo
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "Get logs that have an pending item",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					planmock.NewPlanStep("step1", "success"),
					planmock.NewPlanStep("step2", "pending"),
				},
			},
			want: "step2",
			err:  nil,
		},
		{
			name: "Get logs for a plan that has finished",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					planmock.NewPlanStep("step1", "success"),
					planmock.NewPlanStep("step2", "success"),
					planmock.NewPlanStep(planCompleted, "success"),
				},
			},
			want: planCompleted,
			err:  ErrPlanFinished,
		},
		{
			name: `Get logs for a plan that has errored but the last step isn't "plan-completed"`,
			args: args{
				log: []*models.ClusterPlanStepInfo{
					planmock.NewPlanStep("step1", "success"),
					planmock.NewPlanStep("step2", "success"),
					planmock.NewPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{{
						Message: ec.String(planStepLogErrorMessage),
					}}),
				},
			},
			want: "step3",
			err:  nil,
		},
		{
			name: "Get the last step when it is an error, ignores the previous error step",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					planmock.NewPlanStep("step1", "success"),
					planmock.NewPlanStep("step2", "success"),
					planmock.NewPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{{
						Message: ec.String(planStepLogErrorMessage),
					}}),
					planmock.NewPlanStepWithDetailsAndError(planCompleted, []*models.ClusterPlanStepLogMessageInfo{{
						Message: ec.String(planFinishedErrorMessage),
					}}),
				},
			},
			want: "plan-completed",
			err:  errors.New(planFinishedErrorMessage),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStepName(tt.args.log)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("GetStepName() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("GetStepName() = %v, want %v", got, tt.want)
			}
		})
	}
}
