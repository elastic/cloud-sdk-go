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

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCancelPlan(t *testing.T) {
	type args struct {
		params CancelPlanParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentResourceCrudResponse
		err  string
	}{
		{
			name: "fails due to parameter validation",
			args: args{},
			err: multierror.NewPrefixed("deployment resource",
				apierror.ErrMissingAPI,
				errors.New("id \"\" is invalid"),
				errors.New("resource kind cannot be empty"),
				errors.New(`failed auto-discovering the resource ref id: deployment get: api reference is required for the operation`),
				errors.New(`failed auto-discovering the resource ref id: deployment get: id "" is invalid`),
			).Error(),
		},
		{
			name: "fails due to API error",
			args: args{params: CancelPlanParams{
				Params: Params{
					API:          api.NewMock(mock.SampleInternalError()),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
					RefID:        "main-elasticsearch",
				},
			}},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "fails due to RefID discovery",
			args: args{params: CancelPlanParams{
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
			err: "deployment resource: 1 error occurred:\n\t* failed auto-discovering the resource ref id: api error: deployment.missing: unknown\n\n",
		},
		{
			name: "succeeds",
			args: args{params: CancelPlanParams{
				Params: Params{
					API:          api.NewMock(mock.New200Response(mock.NewStringBody(""))),
					DeploymentID: mock.ValidClusterID,
					Kind:         "elasticsearch",
					RefID:        "main-elasticsearch",
				},
			}},
			want: new(models.DeploymentResourceCrudResponse),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CancelPlan(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CancelPlan() = %v, want %v", got, tt.want)
			}
		})
	}
}
