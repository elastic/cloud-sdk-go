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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestResetElasticsearchPassword(t *testing.T) {
	type args struct {
		params ResetElasticsearchPasswordParams
	}
	tests := []struct {
		name string
		args args
		want *models.ElasticsearchElasticUserPasswordResetResponse
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment elasticsearch password reset params",
				errors.New("api reference is required for the operation"),
				errors.New(`id "" is invalid`),
			),
		},
		{
			name: "succeeds with an explicit RefID",
			args: args{params: ResetElasticsearchPasswordParams{
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/_reset-password",
						Method: "POST",
					},
					mock.NewStringBody(`{"username": "elastic", "password": "my-password"}`),
				)),
				ID:    mock.ValidClusterID,
				RefID: "main-elasticsearch",
			}},
			want: &models.ElasticsearchElasticUserPasswordResetResponse{
				Username: ec.String("elastic"),
				Password: ec.String("my-password"),
			},
		},
		{
			name: "succeeds without a RefID",
			args: args{params: ResetElasticsearchPasswordParams{
				API: api.NewMock(
					mock.New200ResponseAssertion(
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
					),
					mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/elasticsearch/_reset-password",
							Method: "POST",
						},
						mock.NewStringBody(`{"username": "elastic", "password": "some-password"}`),
					),
				),
				ID: mock.ValidClusterID,
			}},
			want: &models.ElasticsearchElasticUserPasswordResetResponse{
				Username: ec.String("elastic"),
				Password: ec.String("some-password"),
			},
		},
		{
			name: "fails discovering a RefID",
			args: args{params: ResetElasticsearchPasswordParams{
				ID: mock.ValidClusterID,
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
		{
			name: "fails with an explicit RefID",
			args: args{params: ResetElasticsearchPasswordParams{
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/_reset-password",
						Method: "POST",
					},
					mock.SampleInternalError().Response.Body,
				)),
				ID:    mock.ValidClusterID,
				RefID: "main-elasticsearch",
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResetElasticsearchPassword(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err.Error()) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
