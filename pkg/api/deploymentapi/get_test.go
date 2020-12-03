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
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/deploymentapi/deputil"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGetParams_Validate(t *testing.T) {
	tests := []struct {
		name   string
		params GetParams
		err    string
	}{
		{
			name:   "validate should return all possible errors",
			params: GetParams{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				deputil.NewInvalidDeploymentIDError(""),
			).Error(),
		},
		{
			name: "validate should return error on missing api",
			params: GetParams{
				DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
			},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
			).Error(),
		},
		{
			name: "validate should return error on invalid ID",
			params: GetParams{
				API: &api.API{},
			},
			err: multierror.NewPrefixed("deployment get",
				deputil.NewInvalidDeploymentIDError(""),
			).Error(),
		},
		{
			name: "validate should pass if all params are properly set",
			params: GetParams{
				API:          &api.API{},
				DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.params.Validate()
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	const getResponse = `{
  "healthy": true,
  "id": "f1d329b0fb34470ba8b18361cabdd2bc"
}`
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentGetResponse
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{Response: http.Response{
						Body:       mock.NewStringBody("error"),
						StatusCode: 500,
					}}),
				},
			},
			err: "error",
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
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
					}),
				},
			},
			want: &models.DeploymentGetResponse{
				Healthy: ec.Bool(true),
				ID:      ec.String("f1d329b0fb34470ba8b18361cabdd2bc"),
			},
		},
		{
			name: "Get succeeds with ConvertLegacyPlans",
			args: args{
				params: GetParams{
					DeploymentID:       "f1d329b0fb34470ba8b18361cabdd2bc",
					ConvertLegacyPlans: true,
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc",
							Query: url.Values{
								"convert_legacy_plans": {"true"},
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
					}),
				},
			},
			want: &models.DeploymentGetResponse{
				Healthy: ec.Bool(true),
				ID:      ec.String("f1d329b0fb34470ba8b18361cabdd2bc"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetAppSearch(t *testing.T) {
	const getAppSearchResponse = `{
  "elasticsearch_cluster_ref_id": "main-elasticsearch",
  "id": "3531aaf988594efa87c1aabb7caed337",
  "ref_id": "main-appsearch"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.AppSearchResourceInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getAppSearchResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc/appsearch/",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
					}),
				},
			},
			want: &models.AppSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("main-appsearch"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAppSearch(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetApm(t *testing.T) {
	const getApmResponse = `{
  "elasticsearch_cluster_ref_id": "main-elasticsearch",
  "id": "3531aaf988594efa87c1aabb7caed337",
  "ref_id": "main-apm"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ApmResourceInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getApmResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc/apm/",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
					}),
				},
			},
			want: &models.ApmResourceInfo{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("main-apm"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetApm(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetElasticsearch(t *testing.T) {
	const getElasticsearchResponse = `{
  "id": "f1d329b0fb34470ba8b18361cabdd2bc",
  "ref_id": "main-elasticsearch"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ElasticsearchResourceInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getElasticsearchResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc/elasticsearch/",
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
					}),
				},
			},
			want: &models.ElasticsearchResourceInfo{
				ID:    ec.String("f1d329b0fb34470ba8b18361cabdd2bc"),
				RefID: ec.String("main-elasticsearch"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetElasticsearch(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetEnterpriseSearch(t *testing.T) {
	const getEnterpriseSearchResponse = `{
  "elasticsearch_cluster_ref_id": "main-elasticsearch",
  "id": "3531aaf988594efa87c1aabb7caed337",
  "ref_id": "main-enterprise_search"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.EnterpriseSearchResourceInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getEnterpriseSearchResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc/enterprise_search/",
							Query: url.Values{
								"show_metadata":      {"false"},
								"show_plan_defaults": {"false"},
								"show_plan_history":  {"false"},
								"show_plan_logs":     {"false"},
								"show_plans":         {"false"},
								"show_settings":      {"false"},
							},
						},
					}),
				},
			},
			want: &models.EnterpriseSearchResourceInfo{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("main-enterprise_search"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEnterpriseSearch(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetKibana(t *testing.T) {
	const getKibanaResponse = `{
  "elasticsearch_cluster_ref_id": "main-elasticsearch",
  "id": "3531aaf988594efa87c1aabb7caed337",
  "ref_id": "main-kibana"
}`

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.KibanaResourceInfo
		err  string
	}{
		{
			name: "Get fails due to parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("deployment get",
				apierror.ErrMissingAPI,
				errors.New(`id "" is invalid`),
			).Error(),
		},
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "f1d329b0fb34470ba8b18361cabdd2bc",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getKibanaResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/f1d329b0fb34470ba8b18361cabdd2bc/kibana/",
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
					}),
				},
			},
			want: &models.KibanaResourceInfo{
				ElasticsearchClusterRefID: ec.String("main-elasticsearch"),
				ID:                        ec.String("3531aaf988594efa87c1aabb7caed337"),
				RefID:                     ec.String("main-kibana"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKibana(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}

func TestGetElasticsearchID(t *testing.T) {
	const getResponse = `{
  "healthy": true,
  "id": "e3dac8bf3dc64c528c295a94d0f19a77",
  "resources": {
    "elasticsearch": [{
      "id": "418017cd1c7f402cbb7a981b2004ceeb",
      "ref_id": "main-elasticsearch",
      "region": "ece-region"
    }]
  }
}`
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want string
		err  string
	}{
		{
			name: "Get fails due to API failure",
			args: args{
				params: GetParams{
					DeploymentID: "e3dac8bf3dc64c528c295a94d0f19a77",
					API:          api.NewMock(mock.SampleInternalError()),
				},
			},
			err: mock.MultierrorInternalError.Error(),
		},
		{
			name: "Get succeeds",
			args: args{
				params: GetParams{
					DeploymentID: "e3dac8bf3dc64c528c295a94d0f19a77",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(getResponse),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/e3dac8bf3dc64c528c295a94d0f19a77",
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
					}),
				},
			},
			want: "418017cd1c7f402cbb7a981b2004ceeb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetElasticsearchID(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
