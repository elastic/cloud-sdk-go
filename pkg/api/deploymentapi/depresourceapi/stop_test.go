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

package depresourceapi

import (
	"errors"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestStopInstancesParams_Validate(t *testing.T) {
	tests := []struct {
		name    string
		params  StopInstancesParams
		wantErr bool
		err     error
	}{
		{
			name:   "validate should return all possible errors",
			params: StopInstancesParams{},
			err: multierror.NewPrefixed("deployment stop",
				errors.New("at least 1 instance ID must be provided"),
				multierror.NewPrefixed("deployment resource", multierror.NewPrefixed("deployment resource",
					apierror.ErrMissingAPI,
					errors.New("id \"\" is invalid"),
					errors.New("resource kind cannot be empty"),
					errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
					errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
				)),
			),
			wantErr: true,
		},
		{
			name: "validate should return error on missing instance IDs",
			params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          &api.API{},
						DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
			},
			err: multierror.NewPrefixed("deployment stop",
				errors.New("at least 1 instance ID must be provided"),
			),
			wantErr: true,
		},
		{
			name: "validate should pass if all params are properly set",
			params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          &api.API{},
						DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.err == nil {
				t.Errorf("Validate() expected errors = '%v' but no errors returned", tt.err)
			}

			if tt.wantErr && err.Error() != tt.err.Error() {
				t.Errorf("Validate() expected errors = '%v' but got %v", tt.err, err)
			}
		})
	}
}

func TestStop(t *testing.T) {
	type args struct {
		params StopParams
	}

	tests := []struct {
		name string
		args args
		want models.DeploymentResourceCommandResponse
		err  error
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("deployment resource", multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				errors.New("id \"\" is invalid"),
				errors.New("resource kind cannot be empty"),
				errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
				errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
			)),
		},
		{
			name: "fails due to API error",
			args: args{params: StopParams{
				Params: Params{
					API:          api.NewMock(mock.SampleNotFoundError()),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
					RefID:        "main-elasticsearch",
				},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "fails due to RefID discovery",
			args: args{params: StopParams{
				Params: Params{
					API: api.NewMock(mock.New500Response(mock.NewStructBody(&models.BasicFailedReply{
						Errors: []*models.BasicFailedReplyElement{
							{Code: ec.String("deployment.missing")},
						},
					}))),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
				},
			}},
			err: multierror.NewPrefixed("deployment resource",
				errors.New("failed auto-discovering the resource ref id: api error: deployment.missing: unknown"),
			),
		},
		{
			name: "succeeds",
			args: args{params: StopParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
					RefID:        "main-elasticsearch",
				},
			}},
		},
		{
			name: "succeeds when RefID is not set",
			args: args{params: StopParams{
				Params: Params{
					API: api.NewMock(
						mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
							Healthy: ec.Bool(true),
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Elasticsearch: []*models.ElasticsearchResourceInfo{{
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID: ec.String("elasticsearch"),
								}},
							},
						})),
						mock.New200Response(mock.NewStringBody("")),
					),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Stop(tt.args.params)
			if tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopInstances(t *testing.T) {
	type args struct {
		params StopInstancesParams
	}

	tests := []struct {
		name string
		args args
		want models.DeploymentResourceCommandResponse
		err  error
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("deployment stop",
				errors.New("at least 1 instance ID must be provided"),
				multierror.NewPrefixed("deployment resource", multierror.NewPrefixed("deployment resource",
					apierror.ErrMissingAPI,
					errors.New("id \"\" is invalid"),
					errors.New("resource kind cannot be empty"),
					errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
					errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
				)),
			),
		},
		{
			name: "fails due to API error",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.SampleNotFoundError()),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "fails due to RefID discovery",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API: api.NewMock(mock.New500Response(mock.NewStructBody(&models.BasicFailedReply{
							Errors: []*models.BasicFailedReplyElement{
								{Code: ec.String("deployment.missing")},
							},
						}))),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
			err: multierror.NewPrefixed("deployment stop",
				errors.New("deployment resource: failed auto-discovering the resource ref id: api error: deployment.missing: unknown"),
			),
		},
		{
			name: "succeeds",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
		},
		{
			name: "succeeds when RefID is not set",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API: api.NewMock(
							mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
								Healthy: ec.Bool(true),
								ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
								Resources: &models.DeploymentResources{
									Elasticsearch: []*models.ElasticsearchResourceInfo{{
										ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
										RefID: ec.String("elasticsearch"),
									}},
								},
							})),
							mock.New200Response(mock.NewStringBody("")),
						),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopInstances(tt.args.params)
			if tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStopAllOrSpecified(t *testing.T) {
	type args struct {
		params StopInstancesParams
	}

	tests := []struct {
		name string
		args args
		want models.DeploymentResourceCommandResponse
		err  error
	}{
		{
			name: "fails due to API error when all is not set",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.SampleNotFoundError()),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "fails due to API error when all is set to true",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.SampleNotFoundError()),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
					All: true,
				},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "succeeds when all is not set",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
				},
				InstanceIDs: []string{"instance-0000000001", "instance-0000000002"},
			}},
		},
		{
			name: "succeeds when all is set to true",
			args: args{params: StopInstancesParams{
				StopParams: StopParams{
					Params: Params{
						API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
						DeploymentID: mock.ValidClusterID,
						Kind:         "elasticsearch",
						RefID:        "main-elasticsearch",
					},
					All: true,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StopAllOrSpecified(tt.args.params)
			if tt.err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Stop() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stop() = %v, want %v", got, tt.want)
			}
		})
	}
}
