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

	"github.com/go-openapi/strfmt"
)

func sendTrackResponses(responses []TrackResponse, c chan<- TrackResponse) {
	for _, res := range responses {
		c <- res
	}
	close(c)
}

func TestStream(t *testing.T) {
	// nolint
	var sucessESLegacyFmt = `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
[92;mCluster [1234567890][Elasticseach]: finished running all the plan steps[0m (Total plan duration: 3s)
`[1:]
	// nolint
	var failureESLegacyFmt = `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
[91;1mCluster [1234567890][Elasticseach]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 3s)
`[1:]
	// nolint
	var failureInPlanESLegacyFmt = `
Cluster [1234567890][Elasticseach]: running step "step1" (Plan duration 1s)...
Cluster [1234567890][Elasticseach]: running step "step2" (Plan duration 2s)...
Cluster [1234567890][Elasticseach]: running step "step2" caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]" (Plan duration 3s)...
[91;1mCluster [1234567890][Elasticseach]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 4s)
`[1:]
	// nolint
	var sucessESFmt = `
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step1" (Plan duration 1s)...
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step2" (Plan duration 2s)...
[92;mDeployment [0987654321] - [Elasticseach][1234567890]: finished running all the plan steps[0m (Total plan duration: 3s)
`[1:]
	// nolint
	var failureESFmt = `
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step1" (Plan duration 1s)...
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step2" (Plan duration 2s)...
[91;1mDeployment [0987654321] - [Elasticseach][1234567890]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 3s)
`[1:]
	// nolint
	var failureInPlanESFmt = `
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step1" (Plan duration 1s)...
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step2" (Plan duration 2s)...
Deployment [0987654321] - [Elasticseach][1234567890]: running step "step2" caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]" (Plan duration 3s)...
[91;1mDeployment [0987654321] - [Elasticseach][1234567890]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"[0m (Total plan duration: 4s)
`[1:]
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
			name: "Stream succeeds with successful finish (legacy)",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     planCompleted,
						Finished: true,
						Duration: strfmt.Duration(time.Second * 3),
					},
				},
			},
			// nolint
			wantDevice: sucessESLegacyFmt,
		},
		{
			name: "Stream succeeds with error finish (legacy)",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 3),
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			//nolint
			wantDevice: failureESLegacyFmt,
		},
		{
			name: "Stream succeeds with error finish and step error (legacy)",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 3),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 4),
					},
				},
			},
			err: errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			//nolint
			wantDevice: failureInPlanESLegacyFmt,
		},
		{
			name: "Stream succeeds with successful finish",
			args: args{
				contents: []TrackResponse{
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step1",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second * 2),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         planCompleted,
						Finished:     true,
						Duration:     strfmt.Duration(time.Second * 3),
					},
				},
			},
			// nolint
			wantDevice: sucessESFmt,
		},
		{
			name: "Stream succeeds with error finish",
			args: args{
				contents: []TrackResponse{
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step1",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second * 2),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     true,
						Err:          errors.New(planStepLogErrorMessage),
						Duration:     strfmt.Duration(time.Second * 3),
					},
				},
			},
			err: errors.New(`deployment [0987654321] - [elasticseach][1234567890]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"`),
			//nolint
			wantDevice: failureESFmt,
		},
		{
			name: "Stream succeeds with error finish and step error",
			args: args{
				contents: []TrackResponse{
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step1",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second * 2),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     false,
						Err:          errors.New(planStepLogErrorMessage),
						Duration:     strfmt.Duration(time.Second * 3),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     true,
						Err:          errors.New(planStepLogErrorMessage),
						Duration:     strfmt.Duration(time.Second * 4),
					},
				},
			},
			err: errors.New(`deployment [0987654321] - [elasticseach][1234567890]: caught error: "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"`),
			//nolint
			wantDevice: failureInPlanESFmt,
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
	var wantPrettyOut = `
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step1",
  "duration": "1s"
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "duration": "2s"
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "err": {
    "message": "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
  },
  "duration": "3s"
}
{
  "id": "1234567890",
  "kind": "elasticseach",
  "step": "step2",
  "err": {
    "message": "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
  },
  "duration": "4s",
  "finished": true
}
`[1:]
	var wantSucessLegacy = `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":"1s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":"2s"}
{"id":"1234567890","kind":"elasticseach","step":"plan-completed","duration":"3s","finished":true}
`[1:]
	var wantSucess = `
{"id":"1234567890","kind":"elasticseach","step":"step1","deployment_id":"0987654321","ref_id":"main-elasticsearch","duration":"1s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","deployment_id":"0987654321","ref_id":"main-elasticsearch","duration":"2s"}
{"id":"1234567890","kind":"elasticseach","step":"plan-completed","deployment_id":"0987654321","ref_id":"main-elasticsearch","duration":"3s","finished":true}
`[1:]
	var wantSuccessWithErrCatch = `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":"1s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":"2s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"},"duration":"3s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"},"duration":"4s","finished":true}
`[1:]
	var wantSuccessWithErrFinish = `
{"id":"1234567890","kind":"elasticseach","step":"step1","duration":"1s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","duration":"2s"}
{"id":"1234567890","kind":"elasticseach","step":"step2","err":{"message":"Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"},"duration":"3s","finished":true}
`[1:]
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
			name: "Stream succeeds with successful finish (legacy)",
			args: args{
				contents: []TrackResponse{
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step1",
						Finished: false,
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     planCompleted,
						Finished: true,
						Duration: strfmt.Duration(time.Second * 3),
					},
				},
			},
			wantDevice: wantSucessLegacy,
		},
		{
			name: "Stream succeeds with successful finish",
			args: args{
				contents: []TrackResponse{
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step1",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         "step2",
						Finished:     false,
						Duration:     strfmt.Duration(time.Second * 2),
					},
					{
						DeploymentID: "0987654321",
						RefID:        "main-elasticsearch",
						ID:           "1234567890",
						Kind:         "elasticseach",
						Step:         planCompleted,
						Finished:     true,
						Duration:     strfmt.Duration(time.Second * 3),
					},
				},
			},
			wantDevice: wantSucess,
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
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 3),
					},
				},
			},
			err:        errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: wantSuccessWithErrFinish,
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
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 3),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 4),
					},
				},
			},
			err:        errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: wantSuccessWithErrCatch,
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
						Duration: strfmt.Duration(time.Second),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Duration: strfmt.Duration(time.Second * 2),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: false,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 3),
					},
					{
						ID:       "1234567890",
						Kind:     "elasticseach",
						Step:     "step2",
						Finished: true,
						Err:      errors.New(planStepLogErrorMessage),
						Duration: strfmt.Duration(time.Second * 4),
					},
				},
				pretty: true,
			},
			err:        errors.New("cluster [1234567890][elasticseach] Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"),
			wantDevice: wantPrettyOut,
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
