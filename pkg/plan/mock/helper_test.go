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

func TestNewPlanStepLog(t *testing.T) {
	var step1 = NewPlanStep("step-1", "success")
	var step2 = NewPlanStep("step-1", "success")
	var step3 = NewPlanStep("step-1", "success")
	type args struct {
		l []*models.ClusterPlanStepInfo
	}
	tests := []struct {
		name string
		args args
		want []*models.ClusterPlanStepInfo
	}{
		{
			name: "one step",
			args: args{l: []*models.ClusterPlanStepInfo{
				step1,
			}},
			want: []*models.ClusterPlanStepInfo{
				step1,
			},
		},
		{
			name: "two steps",
			args: args{l: []*models.ClusterPlanStepInfo{
				step1, step2,
			}},
			want: []*models.ClusterPlanStepInfo{
				step1, step2,
			},
		},
		{
			name: "three steps",
			args: args{l: []*models.ClusterPlanStepInfo{
				step1, step2, step3,
			}},
			want: []*models.ClusterPlanStepInfo{
				step1, step2, step3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlanStepLog(tt.args.l...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlanStepLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlanStep(t *testing.T) {
	type args struct {
		name   string
		status string
	}
	tests := []struct {
		name string
		args args
		want *models.ClusterPlanStepInfo
	}{
		{
			name: "adds a success step",
			args: args{name: "somestep", status: "success"},
			want: &models.ClusterPlanStepInfo{
				StepID: ec.String("somestep"),
				Status: ec.String("success"),
			},
		},
		{
			name: "adds an error step",
			args: args{name: "somestep", status: "error"},
			want: &models.ClusterPlanStepInfo{
				StepID: ec.String("somestep"),
				Status: ec.String("error"),
			},
		},
		{
			name: "adds an in_progress step",
			args: args{name: "somestep", status: "in_progress"},
			want: &models.ClusterPlanStepInfo{
				StepID: ec.String("somestep"),
				Status: ec.String("in_progress"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlanStep(tt.args.name, tt.args.status)
			if got.Started == nil {
				t.Errorf("NewPlanStep() = %v, want %v", got.Started, "not nil")
			}
			got.Started = nil
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlanStep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPlanStepWithDetailsAndError(t *testing.T) {
	type args struct {
		name    string
		details []*models.ClusterPlanStepLogMessageInfo
	}
	tests := []struct {
		name string
		args args
		want *models.ClusterPlanStepInfo
	}{
		{
			name: "adds a step with an error",
			args: args{name: "somename", details: []*models.ClusterPlanStepLogMessageInfo{
				{Message: ec.String("some error message")},
			}},
			want: &models.ClusterPlanStepInfo{
				StepID: ec.String("somename"),
				Status: ec.String("error"),
				InfoLog: []*models.ClusterPlanStepLogMessageInfo{
					{Message: ec.String("some error message")},
				},
			},
		},
		{
			name: "adds a step with a different error",
			args: args{name: "somename", details: []*models.ClusterPlanStepLogMessageInfo{
				{Message: ec.String("some different error message")},
			}},
			want: &models.ClusterPlanStepInfo{
				StepID: ec.String("somename"),
				Status: ec.String("error"),
				InfoLog: []*models.ClusterPlanStepLogMessageInfo{
					{Message: ec.String("some different error message")},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewPlanStepWithDetailsAndError(tt.args.name, tt.args.details)
			got.Started = nil
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlanStepWithDetailsAndError() = %v, want %v", got, tt.want)
			}
		})
	}
}
