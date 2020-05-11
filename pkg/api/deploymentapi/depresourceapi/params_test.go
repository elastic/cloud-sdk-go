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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestResourceParams_Validate(t *testing.T) {
	type fields struct {
		API          *api.API
		DeploymentID string
		Kind         string
		RefID        string
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name:   "fails on empty parameters",
			fields: fields{},
			err: multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				deputil.NewInvalidDeploymentIDError(""),
				errors.New("resource kind cannot be empty"),
				multierror.NewPrefixed("failed auto-discovering the resource ref id",
					errors.New("deployment get: api reference is required for the operation"),
					errors.New(`deployment get: id "" is invalid`),
				),
			),
		},
		{
			name: "succeeds validation when refID is populated",
			fields: fields{
				API:          api.NewMock(),
				DeploymentID: mock.ValidClusterID,
				Kind:         "elasticsearch",
				RefID:        "main-elasticsearch",
			},
		},
		{
			name: "returns error when autodiscovery of ref-id fails",
			fields: fields{
				API:          api.NewMock(mock.SampleNotFoundError()),
				DeploymentID: mock.ValidClusterID,
				Kind:         "elasticsearch",
			},
			err: multierror.NewPrefixed("deployment resource", multierror.NewPrefixed(
				"failed auto-discovering the resource ref id", mock.MultierrorNotFound,
			)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &Params{
				API:          tt.fields.API,
				DeploymentID: tt.fields.DeploymentID,
				Kind:         tt.fields.Kind,
				RefID:        tt.fields.RefID,
			}

			if err := params.Validate(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("ResourceParams.Validate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
