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

	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
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
	err := StreamFunc(channel, func(res TrackResponse) {
		if _, ok := lastStreamed[res.ID]; !ok {
			lastStreamed[res.ID] = ""
		}
		if res.Err != nil {
			res.Err = apierror.NewJSONError(res.Err)
			lastStreamed[res.ID] = res.Step
			_ = encoder.Encode(res)
		}

		if res.Step != lastStreamed[res.ID] {
			lastStreamed[res.ID] = res.Step
			_ = encoder.Encode(res)
		}
	})

	return multierror.WithFormat(err, "json")
}

// StreamFunc is the underlying function used by Stream and StreamJSON. If used
// directly it allows the user to perform an custom action on each received
// response. Unless the sender closes the channel when it has finished, calling
// this function will block execution forever. If the plan failed, it returns
// the error that made the plan fail.
func StreamFunc(channel <-chan TrackResponse, function func(TrackResponse)) error {
	var merr = multierror.NewPrefixed("found deployment plan errors")
	for res := range channel {
		function(res)

		if res.Err != nil && res.Finished && res.Err != ErrPlanFinished {
			res.Err = apierror.NewJSONError(res.Err)
			merr = merr.Append(res)
		}
	}

	return merr.ErrorOrNil()
}
