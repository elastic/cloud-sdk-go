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
	"fmt"
	"strings"

	"github.com/go-openapi/strfmt"
)

// TrackResponse is returned by Track and indicates the progress of a pending
// plan.
type TrackResponse struct {
	ID   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Step string `json:"step,omitempty"`
	Err  error  `json:"err,omitempty"`

	// Introduced as part of the Deployment Plan Tracker
	DeploymentID string          `json:"deployment_id,omitempty"`
	RefID        string          `json:"ref_id,omitempty"`
	Duration     strfmt.Duration `json:"duration,omitempty"`

	// Introduced as part of the plan failure categorization
	FailureDetails *FailureDetails `json:"failure_details,omitempty"`

	Finished    bool `json:"finished,omitempty"`
	runningStep bool `json:"-"`
}

// FailureDetails contains the details for a failure.
type FailureDetails struct {
	// A map with details for the log about what happened during the step execution. Keys and values for are always both strings, representing the name of the detail and its value, respectively.
	Details map[string]string `json:"details"`

	// The failure type, in case the step failed
	FailureType string `json:"failure_type,omitempty"`

	// A json object with sensitive details for the log, visible only to admins. May contain nested json objects.
	Internal interface{} `json:"internal,omitempty"`
}

func (res TrackResponse) Error() string {
	if res.Err == nil {
		return ""
	}

	return fmt.Sprintf(
		"deployment [%s] - [%s][%s]: caught error: \"%s\"",
		res.DeploymentID, res.Kind, res.ID, res.Err.Error(),
	)
}

func (res TrackResponse) String() string {
	kind := strings.Title(strings.Replace(res.Kind, "_", " ", 1))

	if msg := formatFinishedStep(res, kind); msg != "" {
		return msg
	}

	if msg := formatErrStep(res, kind); msg != "" {
		return msg
	}

	return formatRunningStep(res, kind)
}

func formatFinishedStep(res TrackResponse, kind string) string {
	if !res.Finished {
		return ""
	}

	if res.Err != nil && res.Err != ErrPlanFinished {
		return fmt.Sprintf(streamFinishErrFormat,
			res.DeploymentID, kind, res.ID, res.Err, res.Duration,
		)
	}

	return fmt.Sprintf(streamFinishFormat,
		res.DeploymentID, kind, res.ID, res.Duration,
	)
}

func formatErrStep(res TrackResponse, kind string) string {
	if res.Err == nil {
		return ""
	}

	return fmt.Sprintf(streamErrFormat, res.DeploymentID,
		kind, res.ID, res.Step, res.Err, res.Duration,
	)
}

func formatRunningStep(res TrackResponse, kind string) string {
	// If the runningStep is not populated, no string is returned.
	// This is used to internally for the text plan streamer.
	if !res.runningStep {
		return ""
	}

	return fmt.Sprintf(streamFormat, res.DeploymentID,
		kind, res.ID, res.Step, res.Duration,
	)
}
