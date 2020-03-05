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
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-multierror"
)

// Stream prints a text formatted line on each TrackResponse received by the
// channel, unless the sender closes the channel when it has finished, calling
// this function will block execution until the received channel is closed.
func Stream(channel <-chan TrackResponse, device io.Writer) error {
	var lastStreamed = make(map[string]string)
	return StreamFunc(channel, func(res TrackResponse) {
		if _, ok := lastStreamed[res.ID]; !ok {
			lastStreamed[res.ID] = ""
		}

		res.runningStep = res.Step != lastStreamed[res.ID]
		if msg := res.String(); msg != "" {
			lastStreamed[res.ID] = res.Step
			fmt.Fprint(device, msg)
		}
	})
}

// MarshableError wraps any incoming error inside this struct so that it can
// be correctly  marshaled to JSON.
type MarshableError struct {
	Message string `json:"message,omitempty"`
}

// Error complies with the error interface
func (me MarshableError) Error() string { return me.Message }

// StreamJSON prints a json formatted line for on each TrackResponse received
// by the channel, if pretty is set to true, the message will be intended with
// 2 spaces. Unless the sender closes the channel when it has finished, calling
// this function will block execution forever.
func StreamJSON(channel <-chan TrackResponse, device io.Writer, pretty bool) error {
	var encoder = json.NewEncoder(device)
	if pretty {
		encoder.SetIndent("", "  ")
	}

	var lastStreamed = make(map[string]string)
	return StreamFunc(channel, func(res TrackResponse) {
		if _, ok := lastStreamed[res.ID]; !ok {
			lastStreamed[res.ID] = ""
		}
		if res.Err != nil {
			res.Err = &MarshableError{res.Err.Error()}
			lastStreamed[res.ID] = res.Step
			_ = encoder.Encode(res)
		}

		if res.Step != lastStreamed[res.ID] {
			lastStreamed[res.ID] = res.Step
			_ = encoder.Encode(res)
		}
	})
}

// StreamFunc is the underlying function used by Stream and StreamJSON. If used
// directly it allows the user to perform an custom action on each received
// response. Unless the sender closes the channel when it has finished, calling
// this function will block execution forever. If the plan failed, it returns
// the error that made the plan fail.
func StreamFunc(channel <-chan TrackResponse, function func(TrackResponse)) error {
	var err = new(multierror.Error)
	for res := range channel {
		function(res)
		if res.Err != nil && res.Finished && res.Err != ErrPlanFinished {
			err = multierror.Append(err, res.Error())
		}
	}

	return compatibleError(err)
}

// compatibleError is a small utility to ensure that the otuput of StreamFunc
// remains the same as long as the TrackResponse is kept the same.
func compatibleError(e *multierror.Error) error {
	if e != nil {
		if e.Len() == 1 {
			return e.Errors[0]
		}
	}

	return e.ErrorOrNil()
}
