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
	"strings"
)

const (
	// streamFormat with green color
	streamFinishFormat = "\x1b[92;mCluster [%s][%s]: finished running all the plan steps\x1b[0m (Total plan duration: %s)\n"
	// streamFormatOnError with red color
	streamFinishErrFormat = "\x1b[91;1mCluster [%s][%s]: caught error: \"%s\"\x1b[0m (Total plan duration: %s)\n"

	// These formats are used when the plan has not yet finished.
	streamFormat    = "Cluster [%s][%s]: running step \"%s\" (Plan duration %s)...\n"
	streamErrFormat = "Cluster [%s][%s]: running step \"%s\" caught error: \"%s\" (Plan duration %s)...\n"
)

// Stream prints a text formatted line on each TrackResponse received by the
// channel, unless the sender closes the channel when it has finished, calling
// this function will block execution forever.
func Stream(channel <-chan TrackResponse, device io.Writer) error {
	var lastStep string
	return StreamFunc(channel, func(res TrackResponse) {
		kind := strings.Title(res.Kind)
		if res.Finished {
			if res.Err != nil && res.Err != ErrPlanFinished {
				fmt.Fprintf(device, streamFinishErrFormat,
					res.ID, kind, res.Err, res.Duration,
				)
				return
			}
			fmt.Fprintf(device, streamFinishFormat, res.ID, kind, res.Duration)
			return
		}

		if res.Err != nil {
			fmt.Fprintf(device, streamErrFormat,
				res.ID, kind, res.Step, res.Err, res.Duration,
			)
			return
		}
		if res.Step != lastStep {
			fmt.Fprintf(device, streamFormat, res.ID, kind, res.Step, res.Duration.String())
			lastStep = res.Step
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
	var lastStep string
	return StreamFunc(channel, func(res TrackResponse) {
		if res.Err != nil {
			res.Err = &MarshableError{res.Err.Error()}
			lastStep = res.Step
			// nolint
			encoder.Encode(res)
		}

		if res.Step != lastStep {
			lastStep = res.Step
			// nolint
			encoder.Encode(res)
		}
	})
}

// StreamFunc is the underlying function used by Stream and StreamJSON. If used
// directly it allows the user to perform an custom action on each received
// response. Unless the sender closes the channel when it has finished, calling
// this function will block execution forever. If the plan failed, it returns
// the error that made the plan fail.
func StreamFunc(channel <-chan TrackResponse, function func(TrackResponse)) error {
	var err error
	for res := range channel {
		function(res)
		if res.Err != nil && res.Finished && res.Err != ErrPlanFinished {
			err = fmt.Errorf("cluster [%s][%s] %s", res.ID, res.Kind, res.Err.Error())
		}
	}
	return err
}
