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
	"fmt"
	"time"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

const (
	planCompleted = "plan-completed"
	successStatus = "success"
	errorStatus   = "error"
)

var (
	// ErrPlanFinished is returned when a cluster has no plan PlanStepInfo
	ErrPlanFinished = errors.New("finished all the plan steps")
	allowedKinds    = []string{"apm", "elasticsearch", "kibana"}
)

// TrackParams is used as a configuration struct to track a cluster's pending
// plan.
type TrackParams struct {
	*api.API
	ID            string
	Kind          string
	PollFrequency time.Duration
	// If set to > 0, allows up to that number of errors coming from the API.
	MaxRetries uint8
}

// Validate verifies that the parameters are usable by its consumer.
func (params TrackParams) Validate() error {
	var merr = new(multierror.Error)
	if len(params.ID) != 32 {
		merr = multierror.Append(merr, errors.New("plan Track: invalid ID"))
	}

	if !slice.HasString(allowedKinds, params.Kind) {
		merr = multierror.Append(merr, fmt.Errorf("plan Track: invalid kind %s", params.Kind))
	}

	if params.API == nil {
		merr = multierror.Append(merr, errors.New("plan Track: API cannot be nil"))
	}

	if params.PollFrequency.Nanoseconds() < 1 {
		merr = multierror.Append(merr, errors.New("plan Track: poll frequency must be at least 1 nanosecond"))
	}

	return merr.ErrorOrNil()
}

// TrackResponse is returned by Track and indicates the progress of a pending
// plan.
type TrackResponse struct {
	ID       string        `json:"id,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	Step     string        `json:"step,omitempty"`
	Finished bool          `json:"finished,omitempty"`
	Duration time.Duration `json:"duration,omitempty"`
	Err      error         `json:"err,omitempty"`
}

// Track iterates over a cluster pending plan and returns timely updates about
// the status of the plan. When all of the updates have been sent, the channel
// is automatically closed by the poller function.
func Track(params TrackParams) (<-chan TrackResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	var out = make(chan TrackResponse)
	go track(params, out, time.NewTicker(params.PollFrequency))

	return out, nil
}

// track requests the plan steps to the params.API until the client returns
// an error, meaning that the plan has finished. It updates the channel with
// each response it receives.
func track(params TrackParams, c chan<- TrackResponse, ticker *time.Ticker) {
	var start = time.Now()

	// We try to obtain the current plan logs, to see if any error occurred
	// between the time that we last polled and the plan finished.
	defer func() {
		c <- buildTrackerResponse(params, false, start)
		close(c)
	}()

	for range ticker.C {
		// Get the pending plan logs, if an error is returned, the plan has
		// finished.
		res := buildTrackerResponse(params, true, start)

		// If the step name and error are empty, it means that the call failed
		// nothing needs to be sent to the channel and we're returning.
		// Also if the plan has finished, it needs to break the loop and stop
		// the ticker.
		clientError := res.Step == "" && res.Err != nil
		planFinished := res.Step == planCompleted
		if clientError || planFinished {
			ticker.Stop()
			return
		}
		c <- res
	}
}

// buildTrackerResponse bulids a TrackResponse from a set of parameters using
// the API.Get to obtain the plan log. If the client returns an error, the
// function still returns a response, only with an empty name.
func buildTrackerResponse(params TrackParams, pending bool, start time.Time) TrackResponse {
	var response = TrackResponse{
		ID:       params.ID,
		Kind:     params.Kind,
		Finished: !pending,
	}

	// Get the pending plan logs, if an error is returned, the plan has
	// finished. Returning an empty TrackResponse
	log, err := Get(GetParams{
		API:        params.API,
		ID:         params.ID,
		Kind:       params.Kind,
		Pending:    pending,
		MaxRetries: params.MaxRetries,
		Cooldown:   params.PollFrequency,
	})
	response.Duration = time.Since(start)
	if err != nil {
		response.Err = err
		return response
	}

	response.Step, response.Err = GetStepName(log)
	return response
}

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
		// If the step is not "error" or "success"
		if *step.Status != successStatus {
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
