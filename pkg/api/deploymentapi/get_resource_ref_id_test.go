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

package deploymentapi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestPopulateRefID(t *testing.T) {
	type args struct {
		params PopulateRefIDParams
	}
	tests := []struct {
		name string
		args args
		err  error
		want string
	}{
		{
			name: "already set RefID returns it",
			args: args{params: PopulateRefIDParams{
				RefID: ec.String("some-ref-id"),
			}},
			want: "some-ref-id",
		},
		{
			name: "discovers a RefID",
			args: args{params: PopulateRefIDParams{
				DeploymentID: mock.ValidClusterID,
				Kind:         util.Elasticsearch,
				RefID:        ec.String(""),
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed",
						Query: url.Values{
							"convert_legacy_plans": {"false"},
							"enrich_with_template": {"true"},
							"show_metadata":        {"false"},
							"show_plan_defaults":   {"false"},
							"show_plan_history":    {"false"},
							"show_plan_logs":       {"false"},
							"show_plans":           {"false"},
							"show_security":        {"false"},
							"show_settings":        {"false"},
							"show_system_alerts":   {"5"},
						},
					},
					mock.NewStructBody(models.DeploymentGetResponse{
						Healthy: ec.Bool(true),
						ID:      ec.String(mock.ValidClusterID),
						Resources: &models.DeploymentResources{
							Elasticsearch: []*models.ElasticsearchResourceInfo{{
								ID:    ec.String(mock.ValidClusterID),
								RefID: ec.String("elasticsearch"),
							}},
						},
					}),
				)),
			}},
			want: "elasticsearch",
		},
		{
			name: "fails discovering a RefID",
			args: args{params: PopulateRefIDParams{
				DeploymentID: mock.ValidClusterID,
				Kind:         util.Elasticsearch,
				RefID:        ec.String(""),
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed",
						Query: url.Values{
							"convert_legacy_plans": {"false"},
							"enrich_with_template": {"true"},
							"show_metadata":        {"false"},
							"show_plan_defaults":   {"false"},
							"show_plan_history":    {"false"},
							"show_plan_logs":       {"false"},
							"show_plans":           {"false"},
							"show_security":        {"false"},
							"show_settings":        {"false"},
							"show_system_alerts":   {"5"},
						},
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PopulateRefID(tt.args.params); !assert.Equal(t, tt.err, err) {
				t.Errorf("PopulateRefID() error = %v, wantErr %v", err, tt.err)
			}
			assert.EqualValues(t, tt.want, *tt.args.params.RefID)
		})
	}
}
