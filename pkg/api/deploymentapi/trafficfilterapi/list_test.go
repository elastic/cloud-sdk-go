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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.TrafficFilterRulesets
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid traffic filter list params",
				apierror.ErrMissingAPI,
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: ListParams{
				Region: "some-region",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Host:   api.DefaultMockHost,
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Path:   "/api/v1/deployments/traffic-filter/rulesets",
						Query: url.Values{
							"region":               []string{"some-region"},
							"include_associations": []string{"false"},
						},
					},
					mock.NewStringBody(`{"rulesets": [{"id": "some-id"}]}`),
				)),
			}},
			want: &models.TrafficFilterRulesets{Rulesets: []*models.TrafficFilterRulesetInfo{{
				ID: ec.String("some-id"),
			}}},
		},
		{
			name: "succeeds with IncludeAssociations and no region",
			args: args{params: ListParams{
				IncludeAssociations: true,
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Host:   api.DefaultMockHost,
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Path:   "/api/v1/deployments/traffic-filter/rulesets",
						Query: url.Values{
							"include_associations": []string{"true"},
						},
					},
					mock.NewStringBody(`{"rulesets": [{"id": "some-id", "associations": [{"id": "some-id", "entity_type": "deployment"}]}]}`),
				)),
			}},
			want: &models.TrafficFilterRulesets{Rulesets: []*models.TrafficFilterRulesetInfo{{
				ID: ec.String("some-id"),
				Associations: []*models.FilterAssociation{
					{ID: ec.String("some-id"), EntityType: ec.String("deployment")},
				},
			}}},
		},
		{
			name: "fails",
			args: args{params: ListParams{
				API:    api.NewMock(mock.SampleInternalError()),
				Region: "some-region",
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
