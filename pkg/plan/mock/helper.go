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
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewPlanStepLog combines a list of models.ClusterPlanStepInfo into a slice.
func NewPlanStepLog(l ...*models.ClusterPlanStepInfo) []*models.ClusterPlanStepInfo {
	return l
}

// NewPlanStep creates a new plan step with the specified fields.
func NewPlanStep(name, status string) *models.ClusterPlanStepInfo {
	started := strfmt.DateTime(time.Now())
	return &models.ClusterPlanStepInfo{
		StepID:  &name,
		Started: &started,
		Status:  &status,
	}
}

// NewPlanStepWithDetailsAndError creates a new plan step with the specified fields
// with an error status.
func NewPlanStepWithDetailsAndError(name string, details []*models.ClusterPlanStepLogMessageInfo) *models.ClusterPlanStepInfo {
	step := NewPlanStep(name, "error")
	step.InfoLog = details
	return step
}
