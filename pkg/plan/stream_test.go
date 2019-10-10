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
	"bytes"
	"errors"
	"reflect"
	"testing"
	"time"
)

func sendTrackResponses(responses []TrackResponse, c chan<- TrackResponse) {
	for _, res := range responses {
		c <- res
	}
	close(c)
}

func TestStream(t *testing.T) {
	type args struct {
		contents []TrackResponse
	}
	tests := []struct {
		name       string
		args       args
		wantDevice string
		err        error
	}{
		{
			name: "Stream succeeds with successful finish",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     planCompleted,
						Finished: true,
						Duration: time.Second * 3,
					},
				},
			},
			// nolint
			wantDevice: `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
[92;mCluster [1234567890][Elasticseach]: finished running all the plan steps[0m (Total plan duration: 3s)
`[1:],
		},
		{
			name: "Stream succeeds with error finish",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 3,
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			//nolint
			wantDevice: `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
[91;1mCluster [1234567890][Elasticseach]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 3s)
`[1:],
		},
		{
			name: "Stream succeeds with error finish and step error",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 3,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 4,
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			//nolint
			wantDevice: `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
Cluster [1234567890][Elasticseach]: running step "step2" caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]" (Plan duration 3s)...
[91;1mCluster [1234567890][Elasticseach]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 4s)
`[1:],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device := new(bytes.Buffer)
			channel := make(chan TrackResponse)

			// Simulate sender
			go sendTrackResponses(tt.args.contents, channel)
			err := Stream(channel, device)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Stream() error = \n%v, want \n%v", err, tt.err)
			}
			if gotDevice := device.String(); gotDevice != tt.wantDevice {
				t.Errorf("Stream() = \n%v, want \n%v", gotDevice, tt.wantDevice)
			}
		})
	}
}

func TestStreamJSON(t *testing.T) {
	type args struct {
		contents []TrackResponse
		pretty   bool
	}
	tests := []struct {
		name       string
		args       args
		wantDevice string
		err        error
	}{
		{
			name: "Stream succeeds with successful finish",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     planCompleted,
						Finished: true,
						Duration: time.Second * 3,
					},
				},
			},
			wantDevice: `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":1000000000}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":2000000000}
{"id":"1234567890","kind":"elasticseach","step":"plan-completed","finished":true,"duration":3000000000}
`[1:],
		},
		{
			name: "Stream succeeds with error finish",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 3,
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":1000000000}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":2000000000}
{"id":"1234567890","kind":"elasticseach","step":"step2","finished":true,"duration":3000000000,"err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"}}
`[1:],
		},
		{
			name: "Stream succeeds with error finish and step error",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 3,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 4,
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":1000000000}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":2000000000}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":3000000000,"err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"}}
{"id":"1234567890","kind":"elasticseach","step":"step2","finished":true,"duration":4000000000,"err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"}}
`[1:],
		},
		{
			name: "Stream succeeds with error finish and step error (Pretty format)",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: time.Second,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: time.Second * 2,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 3,
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: time.Second * 4,
					},
				},
				pretty: true,
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: `
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step1",
  "duration": 1000000000
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "duration": 2000000000
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "duration": 3000000000,
  "err": {
    "message": "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
  }
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "finished": true,
  "duration": 4000000000,
  "err": {
    "message": "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
  }
}
`[1:],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device := new(bytes.Buffer)
			channel := make(chan TrackResponse)

			// Simulate sender
			go sendTrackResponses(tt.args.contents, channel)
			err := StreamJSON(channel, device, tt.args.pretty)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("StreamJSON() error = \n%v, want \n%v", err, tt.err)
			}

			if gotDevice := device.String(); gotDevice != tt.wantDevice {
				t.Errorf("StreamJSON() = %v, want %v", gotDevice, tt.wantDevice)
			}
		})
	}
}
