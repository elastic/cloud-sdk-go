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
	successStatus = "success"
	errorStatus   = "error"
	pendingStatus = "pending"
)

// GetStepName goes over the pending plan logs to locate:
//
// 1. Errors that have occurred in the plan
//
// 2. Steps that are not "success"
//
// 3. If the ID for the last step in the log is plan-completed, it will return the ErrPlanFinished
//
// If none of the above are found, it returns the last step ID with the trackpayload.
func GetStepName(log []*models.ClusterPlanStepInfo) (string, error) {
	// Obtain the last step in the plan log and if its status is "error",
	// return the plan step log ID with the detailed error message.
	if stepLog, _ := lastLog(log); stepLog != nil {
		if *stepLog.Status == errorStatus {
			return *stepLog.StepID, StepErrorOrUnknownError(stepLog)
		}
	}

	for _, step := range log {
		if *step.Status == errorStatus {
			return *step.StepID, StepErrorOrUnknownError(step)
		}
		// If the step is not "error" or "success" and is not pending (Cancelled plans)
		if *step.Status != successStatus && *step.Status != pendingStatus {
			return *step.StepID, nil
		}
	}

	var stepName = lastLogStepID(log)
	var err error
	if stepName == planCompleted {
		err = ErrPlanFinished
	}

	return stepName, err
}

func lastLogStepID(log []*models.ClusterPlanStepInfo) string {
	if len(log) == 0 {
		return ""
	}

	return *log[len(log)-1].StepID
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
