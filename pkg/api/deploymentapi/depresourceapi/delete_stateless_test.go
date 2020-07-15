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
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestDeleteStateless(t *testing.T) {
	type args struct {
		params DeleteStatelessParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentResourceUpgradeResponse
		err  error
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("deployment resource delete", multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				errors.New("id \"\" is invalid"),
				errors.New("resource kind cannot be empty"),
				errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
				errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
			)),
		},
		{
			name: "fails due to parameter validation on invalid kind",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					Kind: "elasticsearch",
				},
			}},
			err: multierror.NewPrefixed("deployment resource delete",
				multierror.NewPrefixed("deployment resource",
					apierror.ErrMissingAPI,
					errors.New("id \"\" is invalid"),
					errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
					errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
				),
				errors.New("resource kind \"elasticsearch\" is not supported"),
			),
		},
		{
			name: "fails due to API error",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					API:          api.NewMock(mock.SampleNotFoundError()),
					DeploymentID: mock.ValidClusterID,
					RefID:        "kibana",
					Kind:         "kibana",
				},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "succeeds on APM resource",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					RefID:        "kibana",
					Kind:         "kibana",
				},
			}},
		},
		{
			name: "fails due to API error on APM resource",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					API:          api.NewMock(mock.SampleInternalError()),
					DeploymentID: mock.ValidClusterID,
					RefID:        util.Apm,
					Kind:         util.Apm,
				},
			}},
			err: mock.MultierrorInternalError,
		},
		{
			name: "succeeds",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					RefID:        util.Apm,
					Kind:         util.Apm,
				},
			}},
		},
		{
			name: "succeeds with refID autodiscovery",
			args: args{params: DeleteStatelessParams{
				Params: Params{
					API: api.NewMock(
						mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
							Healthy: ec.Bool(true),
							ID:      ec.String(mock.ValidClusterID),
							Resources: &models.DeploymentResources{
								Apm: []*models.ApmResourceInfo{{
									ID:    ec.String(mock.ValidClusterID),
									RefID: ec.String(util.Apm),
								}},
							},
						})),
						mock.New200Response(mock.NewStringBody("")),
					),
					DeploymentID: mock.ValidClusterID,
					Kind:         util.Apm,
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteStateless(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("DeleteStateless() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}
