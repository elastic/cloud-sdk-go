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

	"github.com/go-openapi/strfmt"
	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

var (
	allowedKinds = []string{"apm", "elasticsearch", "kibana"}
)

// TrackParams (DEPRECATED) is used as a configuration struct to track a cluster's pending
// plan.
type TrackParams struct {
	*api.API
	ID            string
	Kind          string
	PollFrequency time.Duration
	// If set to > 0, allows up to that number of errors coming from the API.
	MaxRetries uint8
}

// Validate (DEPRECATED) verifies that the parameters are usable by its consumer.
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

// Track (DEPRECATED) iterates over a cluster pending plan and returns timely updates about
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
	response.Duration = strfmt.Duration(time.Since(start))
	if err != nil {
		response.Err = err
		return response
	}

	response.Step, response.Err = GetStepName(log)
	return response
}
