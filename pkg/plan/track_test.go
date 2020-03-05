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
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
)

const (
	planStepLogErrorMessage  = "Unexpected error during step: [perform-snapshot]: [no.found.constructor.models.TimeoutException: Timeout]"
	planFinishedErrorMessage = "[ClusterFailure:NoAvailableInstanceFound]: Could not find an available instance while attempting: [suspend-snapshotting]... Please validate the cluster is in a healthy status and retry."
)

func newStringPointer(s string) *string { return &s }

func newPlanStep(name, status string) *models.ClusterPlanStepInfo {
	started := strfmt.DateTime(time.Now())
	return &models.ClusterPlanStepInfo{
		StepID:  &name,
		Started: &started,
		Status:  &status,
	}
}

func newPlanStepWithDetailsAndError(name string, details []*models.ClusterPlanStepLogMessageInfo) *models.ClusterPlanStepInfo {
	step := newPlanStep(name, "error")
	step.InfoLog = details
	return step
}

func newPlanFinishedResponse() mock.Response {
	return mock.Response{Response: http.Response{
		Body:       mock.NewStringBody(""),
		StatusCode: 404,
	}}
}

func TestGetStepName(t *testing.T) {
	type args struct {
		log []*models.ClusterPlanStepInfo
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "Get logs that have an pending item",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					newPlanStep("step1", "success"),
					newPlanStep("step2", "pending"),
				},
			},
			want: "step2",
			err:  nil,
		},
		{
			name: "Get logs for a plan that has finished",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					newPlanStep("step1", "success"),
					newPlanStep("step2", "success"),
					newPlanStep(planCompleted, "success"),
				},
			},
			want: planCompleted,
			err:  ErrPlanFinished,
		},
		{
			name: "Get logs for a plan that has errored",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					newPlanStep("step1", "success"),
					newPlanStep("step2", "success"),
					newPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{{
						Message: newStringPointer(planStepLogErrorMessage),
					}}),
				},
			},
			want: "step3",
			err:  errors.New(planStepLogErrorMessage),
		},
		{
			name: "Get the last step when it is an error, ignores the previous error step",
			args: args{
				log: []*models.ClusterPlanStepInfo{
					newPlanStep("step1", "success"),
					newPlanStep("step2", "success"),
					newPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{{
						Message: newStringPointer(planStepLogErrorMessage),
					}}),
					newPlanStepWithDetailsAndError(planCompleted, []*models.ClusterPlanStepLogMessageInfo{{
						Message: newStringPointer(planFinishedErrorMessage),
					}}),
				},
			},
			want: "plan-completed",
			err:  errors.New(planFinishedErrorMessage),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStepName(tt.args.log)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("GetStepName() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("GetStepName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newPollerBody(t *testing.T, pending, current *models.ElasticsearchClusterPlanInfo) io.ReadCloser {
	payload := &models.ElasticsearchClusterPlansInfo{Pending: pending, Current: current}
	var response, err = json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	response = append(response, []byte("\n\n")...)
	return ioutil.NopCloser(bytes.NewReader(response))
}

func newKibanaPollerBody(t *testing.T, pending, current *models.KibanaClusterPlanInfo) io.ReadCloser {
	payload := &models.KibanaClusterPlansInfo{Pending: pending, Current: current}
	var response, err = json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	response = append(response, []byte("\n\n")...)
	return ioutil.NopCloser(bytes.NewReader(response))
}

func newApmPollerBody(t *testing.T, pending, current *models.ApmPlanInfo) io.ReadCloser {
	payload := &models.ApmPlansInfo{Pending: pending, Current: current}
	var response, err = json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	response = append(response, []byte("\n\n")...)
	return ioutil.NopCloser(bytes.NewReader(response))
}

func TestTrack(t *testing.T) {
	type args struct {
		params TrackParams
	}
	tests := []struct {
		name string
		args args
		want []TrackResponse
		err  error
	}{
		{
			name: "Track returns an error on invalid params",
			args: args{
				params: TrackParams{},
			},
			want: nil,
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan Track: invalid ID"),
					errors.New("plan Track: invalid kind "),
					errors.New("plan Track: API cannot be nil"),
					errors.New("plan Track: poll frequency must be at least 1 nanosecond"),
				},
			},
		},
		{
			name: "Track returns an error on invalid params",
			args: args{
				params: TrackParams{
					Kind: "kibana",
				},
			},
			want: nil,
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan Track: invalid ID"),
					errors.New("plan Track: API cannot be nil"),
					errors.New("plan Track: poll frequency must be at least 1 nanosecond"),
				},
			},
		},
		{
			name: "Track returns an error on invalid params",
			args: args{
				params: TrackParams{
					Kind: "kibana",
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
				},
			},
			want: nil,
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan Track: API cannot be nil"),
					errors.New("plan Track: poll frequency must be at least 1 nanosecond"),
				},
			},
		},
		{
			name: "Track returns an error on invalid params",
			args: args{
				params: TrackParams{
					Kind:          "elasticsearch",
					ID:            "9e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
				},
			},
			want: nil,
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan Track: API cannot be nil"),
				},
			},
		},
		{
			name: "Track Elasticsearch returns a channel whick streams the changes and is closed on change finish",
			args: args{
				params: TrackParams{
					Kind:          "elasticsearch",
					ID:            "9e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								nil,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step2",
				},
				{
					ID:       "9e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "elasticsearch",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
		{
			name: "Track kibana returns a channel whick streams the changes and is closed on change finish",
			args: args{
				params: TrackParams{
					Kind:          "kibana",
					ID:            "2e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newKibanaPollerBody(t,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newKibanaPollerBody(t,
								nil,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "2e9c997ff4d0bfc273da17f549e45e76",
					Kind: "kibana",
					Step: "step2",
				},
				{
					ID:       "2e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "kibana",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
		{
			name: "Track apm returns a channel whick streams the changes and is closed on change finish",
			args: args{
				params: TrackParams{
					Kind:          "apm",
					ID:            "4e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newApmPollerBody(t,
								&models.ApmPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newApmPollerBody(t,
								nil,
								&models.ApmPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "4e9c997ff4d0bfc273da17f549e45e76",
					Kind: "apm",
					Step: "step2",
				},
				{
					ID:       "4e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "apm",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
		{
			name: "Track returns a channel whick streams the changes and is closed on change finish (Poller error)",
			args: args{
				params: TrackParams{
					Kind:          "elasticsearch",
					ID:            "9e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// Pending plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								nil,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step2",
				},
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step3",
				},
				{
					ID:       "9e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "elasticsearch",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
		{
			name: "Track returns a channel whick streams the changes and is closed on change finish (Catch pending plan error)",
			args: args{
				params: TrackParams{
					Kind:          "elasticsearch",
					ID:            "9e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// Pending plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// Got an error in plan
						mock.Response{Response: http.Response{
							StatusCode: 200,
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{
										{Message: newStringPointer("stuff we don't want to see")},
										{Message: newStringPointer(planStepLogErrorMessage)},
									}),
								}},
								nil,
							),
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								nil,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetailsAndError("next-step", []*models.ClusterPlanStepLogMessageInfo{
										{Message: newStringPointer("stuff we don't want to see")},
										{Message: newStringPointer(planStepLogErrorMessage)},
									}),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step2",
				},
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step3",
				},
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step3",
					Err:  errors.New(planStepLogErrorMessage),
				},
				{
					ID:       "9e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "elasticsearch",
					Step:     "next-step",
					Finished: true,
					Err:      errors.New(planStepLogErrorMessage),
				},
			},
		},
		{
			name: "Track returns a channel whick streams the changes and is closed on change finish (Catch current plan error)",
			args: args{
				params: TrackParams{
					Kind:          "elasticsearch",
					ID:            "9e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// Pending plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep("step3", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newPollerBody(t,
								nil,
								&models.ElasticsearchClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStepWithDetailsAndError("step3", []*models.ClusterPlanStepLogMessageInfo{
										{Message: newStringPointer("stuff we don't want to see")},
										{Message: newStringPointer(planStepLogErrorMessage)},
									}),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step2",
				},
				{
					ID:   "9e9c997ff4d0bfc273da17f549e45e76",
					Kind: "elasticsearch",
					Step: "step3",
				},
				{
					ID:       "9e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "elasticsearch",
					Step:     "step3",
					Finished: true,
					Err:      errors.New(planStepLogErrorMessage),
				},
			},
		},
		{
			name: "Track returns a channel whick streams the changes and is closed on change finish (With 1 retry)",
			args: args{
				params: TrackParams{
					Kind:          "kibana",
					ID:            "2e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					MaxRetries:    1,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newKibanaPollerBody(t,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newKibanaPollerBody(t,
								nil,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "2e9c997ff4d0bfc273da17f549e45e76",
					Kind: "kibana",
					Step: "step2",
				},
				{
					ID:       "2e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "kibana",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
		{
			name: "Track returns a channel whick streams the changes and is closed on change finish (With 5 retry)",
			args: args{
				params: TrackParams{
					Kind:          "kibana",
					ID:            "2e9c997ff4d0bfc273da17f549e45e76",
					PollFrequency: time.Microsecond,
					MaxRetries:    5,
					API: api.NewMock(
						mock.Response{Response: http.Response{
							// Pending plan
							Body: newKibanaPollerBody(t,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "pending"),
								}},
								nil,
							),
							StatusCode: 200,
						}},
						// 404 means that the plan has finished
						newPlanFinishedResponse(),
						newPlanFinishedResponse(),
						newPlanFinishedResponse(),
						newPlanFinishedResponse(),
						newPlanFinishedResponse(),
						newPlanFinishedResponse(),
						// Current plan
						mock.Response{Response: http.Response{
							Body: newKibanaPollerBody(t,
								nil,
								&models.KibanaClusterPlanInfo{PlanAttemptLog: []*models.ClusterPlanStepInfo{
									newPlanStep("step1", "success"),
									newPlanStep("step2", "success"),
									newPlanStep(planCompleted, "success"),
								}},
							),
							StatusCode: 200,
						}},
					),
				},
			},
			want: []TrackResponse{
				{
					ID:   "2e9c997ff4d0bfc273da17f549e45e76",
					Kind: "kibana",
					Step: "step2",
				},
				{
					ID:       "2e9c997ff4d0bfc273da17f549e45e76",
					Kind:     "kibana",
					Finished: true,
					Step:     planCompleted,
					Err:      ErrPlanFinished,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := Track(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Track() error = %v, wantErr %v", err, tt.err)
				return
			}

			var got []TrackResponse
			if c != nil {
				for response := range c {
					response.Duration = 0
					got = append(got, response)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Track() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}
