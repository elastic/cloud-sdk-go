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
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGetResource(t *testing.T) {
	type args struct {
		params GetResourceParams
	}
	tests := []struct {
		name string
		args args
		want interface{}
		err  error
	}{
		{
			name: "fails due to param validation",
			args: args{params: GetResourceParams{}},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				deputil.NewInvalidDeploymentIDError(""),
			),
		},
		{
			name: "obtains a apm resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/apm/apm",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
						mock.NewStructBody(models.ApmResourceInfo{
							ElasticsearchClusterRefID: ec.String("elasticsearch"),
							ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
							RefID:                     ec.String(util.Apm),
						},
						))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        util.Apm,
				},
				Kind: util.Apm,
			}},
			want: &models.ApmResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String(util.Apm),
			},
		},
		{
			name: "obtains a apm resource without a RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337",
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
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Apm: []*models.ApmResourceInfo{{
									ElasticsearchClusterRefID: ec.String("elasticsearch"),
									ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID:                     ec.String(util.Apm),
								}},
							},
						})),
						mock.New200ResponseAssertion(
							&mock.RequestAssertion{
								Header: api.DefaultReadMockHeaders,
								Method: "GET",
								Host:   api.DefaultMockHost,
								Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/apm/apm",
								Query: url.Values{
									"show_metadata":      {"false"},
									"show_plan_defaults": {"false"},
									"show_plan_history":  {"false"},
									"show_plan_logs":     {"false"},
									"show_plans":         {"false"},
									"show_settings":      {"false"},
								},
							},
							mock.NewStructBody(models.ApmResourceInfo{
								ElasticsearchClusterRefID: ec.String("elasticsearch"),
								ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
								RefID:                     ec.String(util.Apm),
							})),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: util.Apm,
			}},
			want: &models.ApmResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String(util.Apm),
			},
		},
		{
			name: "obtains an elasticsearch resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/elasticsearch/elasticsearch",
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
						mock.NewStructBody(models.ElasticsearchResourceInfo{
							ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
							RefID: ec.String("elasticsearch"),
						},
						))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        "elasticsearch",
				},
				Kind: "elasticsearch",
			}},
			want: &models.ElasticsearchResourceInfo{
				ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID: ec.String("elasticsearch"),
			},
		},
		{
			name: "obtains an elasticsearch resource without a RefID",
			args: args{params: GetResourceParams{
				Kind: "elasticsearch",
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337",
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
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Elasticsearch: []*models.ElasticsearchResourceInfo{{
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID: ec.String("elasticsearch"),
								}},
							},
						})),
						mock.New200ResponseAssertion(
							&mock.RequestAssertion{
								Header: api.DefaultReadMockHeaders,
								Method: "GET",
								Host:   api.DefaultMockHost,
								Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/elasticsearch/elasticsearch",
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
							mock.NewStructBody(
								models.ElasticsearchResourceInfo{
									RefID: ec.String("elasticsearch"),
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
								},
							)),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
			}},
			want: &models.ElasticsearchResourceInfo{
				ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID: ec.String("elasticsearch"),
			},
		},
		{
			name: "obtains a kibana resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/kibana/kibana",
							Query: url.Values{
								"convert_legacy_plans": {"false"},
								"show_metadata":        {"false"},
								"show_plan_defaults":   {"false"},
								"show_plan_history":    {"false"},
								"show_plan_logs":       {"false"},
								"show_plans":           {"false"},
								"show_settings":        {"false"},
							},
						},
						mock.NewStructBody(models.KibanaResourceInfo{
							ElasticsearchClusterRefID: ec.String("elasticsearch"),
							ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
							RefID:                     ec.String("kibana"),
						},
						))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        "kibana",
				},
				Kind: "kibana",
			}},
			want: &models.KibanaResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("kibana"),
			},
		},
		{
			name: "obtains a kibana resource without a RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337",
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
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Kibana: []*models.KibanaResourceInfo{{
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID: ec.String("elasticsearch"),
								}},
							},
						})),
						mock.New200ResponseAssertion(
							&mock.RequestAssertion{
								Header: api.DefaultReadMockHeaders,
								Method: "GET",
								Host:   api.DefaultMockHost,
								Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/kibana/elasticsearch",
								Query: url.Values{
									"convert_legacy_plans": {"false"},
									"show_metadata":        {"false"},
									"show_plan_defaults":   {"false"},
									"show_plan_history":    {"false"},
									"show_plan_logs":       {"false"},
									"show_plans":           {"false"},
									"show_settings":        {"false"},
								},
							},
							mock.NewStructBody(
								models.KibanaResourceInfo{
									ElasticsearchClusterRefID: ec.String("elasticsearch"),
									ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID:                     ec.String("kibana"),
								},
							)),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: "kibana",
			}},
			want: &models.KibanaResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("kibana"),
			},
		},
		{
			name: "obtains a appsearch resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/appsearch/appsearch",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
						mock.NewStructBody(models.AppSearchResourceInfo{
							ElasticsearchClusterRefID: ec.String("elasticsearch"),
							ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
							RefID:                     ec.String("appsearch"),
						},
						))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        "appsearch",
				},
				Kind: "appsearch",
			}},
			want: &models.AppSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("appsearch"),
			},
		},
		{
			name: "obtains a appsearch resource without a RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337",
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
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Appsearch: []*models.AppSearchResourceInfo{{
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID: ec.String("appsearch"),
								}},
							},
						})),
						mock.New200ResponseAssertion(
							&mock.RequestAssertion{
								Header: api.DefaultReadMockHeaders,
								Method: "GET",
								Host:   api.DefaultMockHost,
								Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/appsearch/appsearch",
								Query: url.Values{
									"show_metadata":      {"false"},
									"show_plan_defaults": {"false"},
									"show_plan_history":  {"false"},
									"show_plan_logs":     {"false"},
									"show_plans":         {"false"},
									"show_settings":      {"false"},
								},
							},
							mock.NewStructBody(
								models.AppSearchResourceInfo{
									ElasticsearchClusterRefID: ec.String("elasticsearch"),
									ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID:                     ec.String("appsearch"),
								},
							)),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: "appsearch",
			}},
			want: &models.AppSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("appsearch"),
			},
		},
		{
			name: "obtains an enterprise search resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/enterprise_search/enterprise_search",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
						mock.NewStructBody(models.EnterpriseSearchResourceInfo{
							ElasticsearchClusterRefID: ec.String("elasticsearch"),
							ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
							RefID:                     ec.String(util.EnterpriseSearch),
						},
						))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        util.EnterpriseSearch,
				},
				Kind: util.EnterpriseSearch,
			}},
			want: &models.EnterpriseSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String(util.EnterpriseSearch),
			},
		},
		{
			name: "obtains an enterprise search resource without a RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337",
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
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								EnterpriseSearch: []*models.EnterpriseSearchResourceInfo{{
									ElasticsearchClusterRefID: ec.String("elasticsearch"),
									ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID:                     ec.String(util.EnterpriseSearch),
								}},
							},
						})),
						mock.New200ResponseAssertion(
							&mock.RequestAssertion{
								Header: api.DefaultReadMockHeaders,
								Method: "GET",
								Host:   api.DefaultMockHost,
								Path:   "/api/v1/deployments/3531aaf988594efa87c1aabb7caed337/enterprise_search/enterprise_search",
								Query: url.Values{
									"show_metadata":      {"false"},
									"show_plan_defaults": {"false"},
									"show_plan_history":  {"false"},
									"show_plan_logs":     {"false"},
									"show_plans":         {"false"},
									"show_settings":      {"false"},
								},
							},
							mock.NewStructBody(models.EnterpriseSearchResourceInfo{
								ElasticsearchClusterRefID: ec.String("elasticsearch"),
								ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
								RefID:                     ec.String(util.EnterpriseSearch),
							})),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: util.EnterpriseSearch,
			}},
			want: &models.EnterpriseSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String(util.EnterpriseSearch),
			},
		},
		{
			name: "obtains an invalid resource kind INVALID without a RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(
						mock.New200Response(mock.NewStructBody(models.DeploymentGetResponse{
							Healthy: ec.Bool(true),
							ID:      ec.String("3531aaf988594efa87c1aabb7caed337"),
							Resources: &models.DeploymentResources{
								Appsearch: []*models.AppSearchResourceInfo{{
									ID:    ec.String("3531aaf988594efa87c1aabb7caed337"),
									RefID: ec.String("appsearch"),
								}},
							},
						})),
					),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: "INVALID",
			}},
			err: errors.New("deployment get: resource kind INVALID is not available"),
		},
		{
			name: "returns an error when the RefID discovery fails",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New500Response(mock.NewStructBody(&models.BasicFailedReply{
						Errors: []*models.BasicFailedReplyElement{
							{Code: ec.String("deployment.missing")},
						},
					}))),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
				},
				Kind: "INVALID",
			}},
			err: multierror.NewPrefixed("api error", errors.New("deployment.missing: unknown")),
		},
		{
			name: "tries to obtain an INVALID resource with a set RefID",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API:          api.NewMock(),
					DeploymentID: "3531aaf988594efa87c1aabb7caed337",
					RefID:        "appsearch",
				},
				Kind: "INVALID",
			}},
			err: errors.New("deployment get: resource kind INVALID is not valid"),
		},
		{
			name: "obtains the whole deployment when Kind is empty",
			args: args{params: GetResourceParams{
				GetParams: GetParams{
					API: api.NewMock(mock.New200ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc",
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
							ID:      ec.String("f1d329b0fb34470ba8b18361cabdd2bc"),
						},
						))),
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
				},
			}},
			want: &models.DeploymentGetResponse{
				Healthy: ec.Bool(true),
				ID:      ec.String("f1d329b0fb34470ba8b18361cabdd2bc"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetResource(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
