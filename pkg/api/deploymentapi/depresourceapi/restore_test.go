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
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestRestore(t *testing.T) {
	type args struct {
		params RestoreParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails due to param validation",
			err: multierror.NewPrefixed("deployment resource", multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				errors.New("id \"\" is invalid"),
				errors.New("resource kind cannot be empty"),
				errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
				errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
			)),
		},
		{
			name: "fails due to restore Kibana due to API error",
			args: args{params: RestoreParams{
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
			name: "fails to restore APM due to API error",
			args: args{params: RestoreParams{
				Params: Params{
					API:          api.NewMock(mock.SampleNotFoundError()),
					DeploymentID: mock.ValidClusterID,
					RefID:        deputil.Apm,
					Kind:         deputil.Apm,
				},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "fails due to restore Elasticsearch due to API error",
			args: args{params: RestoreParams{
				Params: Params{
					API:          api.NewMock(mock.SampleNotFoundError()),
					DeploymentID: mock.ValidClusterID,
					RefID:        "elasticsearch",
					Kind:         "elasticsearch",
				},
			}},
			err: mock.MultierrorNotFound,
		},
		{
			name: "Succeeds restoring Kibana",
			args: args{params: RestoreParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					RefID:        "kibana",
					Kind:         "kibana",
				},
			}},
		},
		{
			name: "Succeeds restoring APM",
			args: args{params: RestoreParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					RefID:        deputil.Apm,
					Kind:         deputil.Apm,
				},
			}},
		},
		{
			name: "Succeeds restoring Elasticsearch",
			args: args{params: RestoreParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					RefID:        "elasticsearch",
					Kind:         "elasticsearch",
				},
			}},
		},
		{
			name: "Succeeds restoring Elasticsearch with refID autodiscovery",
			args: args{params: RestoreParams{
				Params: Params{
					API: api.NewMock(
						mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
							Healthy: ec.Bool(true),
							ID:      ec.String(mock.ValidClusterID),
							Resources: &models.DeploymentResources{
								Elasticsearch: []*models.ElasticsearchResourceInfo{{
									ID:    ec.String(mock.ValidClusterID),
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
			if err := Restore(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Restore() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
