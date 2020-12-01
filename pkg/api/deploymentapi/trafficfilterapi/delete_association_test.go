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

package trafficfilterapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestDeleteAssociation(t *testing.T) {
	type args struct {
		params DeleteAssociationParams
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid traffic filter association delete params",
				apierror.ErrMissingAPI,
				errors.New("rule set id is not specified and is required for the operation"),
				errors.New("entity id is not specified and is required for the operation"),
				errors.New("entity type is not specified and is required for the operation"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: DeleteAssociationParams{
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Host:   api.DefaultMockHost,
						Header: api.DefaultWriteMockHeaders,
						Path:   "/api/v1/deployments/traffic-filter/rulesets/some-id/associations/deployment/some-entity-id",
						Method: "DELETE",
					},
					mock.NewStringBody("{}"),
				)),
				ID:         "some-id",
				EntityID:   "some-entity-id",
				EntityType: "deployment",
			}},
		},
		{
			name: "fails",
			args: args{params: DeleteAssociationParams{
				API:        api.NewMock(mock.SampleInternalError()),
				ID:         "some-id",
				EntityID:   "some-entity-id",
				EntityType: "deployment",
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := DeleteAssociation(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("DeleteAssociation() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
