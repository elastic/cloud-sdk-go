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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

const (
	planCompleted = "plan-completed"
	errorStatus   = "error"
)

// GetStepName analyzes the last step in a plan and returns the step id and:
// 1. If the ID is plan-completed, and the status == "error", returns the error.
// 2. If the ID is plan-completed, and the status != "error" returns the ErrPlanFinished.
// 3. If none of the above, returns no error.
func GetStepName(log []*models.ClusterPlanStepInfo) (string, error) {
	stepLog, err := lastLog(log)
	if err != nil {
		return "", err
	}

	stepID := *stepLog.StepID
	stepStatus := *stepLog.Status
	if stepID == planCompleted {
		if stepStatus == errorStatus {
			return stepID, StepErrorOrUnknownError(stepLog)
		}
		return stepID, ErrPlanFinished
	}

	return stepID, nil
}

// StepErrorOrUnknownError returns the last step message as an error except when
// the step InfoLog is empty, in which case it returns errorPlanFailedUnknown.
func StepErrorOrUnknownError(step *models.ClusterPlanStepInfo) error {
	if len(step.InfoLog) == 0 {
		return errors.New("plan failed due to unknown error")
	}

	return errors.New(*step.InfoLog[len(step.InfoLog)-1].Message)
}

func lastLog(log []*models.ClusterPlanStepInfo) (*models.ClusterPlanStepInfo, error) {
	if len(log) == 0 {
		return nil, errors.New("invalid plan step info")
	}

	return log[len(log)-1], nil
}
