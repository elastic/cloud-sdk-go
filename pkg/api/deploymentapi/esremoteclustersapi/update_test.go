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

package esremoteclustersapi

import (
	"errors"
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

func TestUpdate(t *testing.T) {
	successContents := models.RemoteResources{Resources: []*models.RemoteResourceRef{
		{
			DeploymentID:       ec.String("first id"),
			ElasticsearchRefID: ec.String("main-elasticsearch"),
			Alias:              ec.String("caperucita"),
			SkipUnavailable:    ec.Bool(true),
		},
		{
			DeploymentID:       ec.String("second id"),
			ElasticsearchRefID: ec.String("elasticsearch"),
			Alias:              ec.String("bestest"),
		},
	}}

	rawBody, err := successContents.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	rawBody = append(rawBody, []byte("\n")...)
	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid elasticsearch remote clusters api params",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
				errors.New("required remote resources not provided"),
			),
		},
		{
			name: "succeeds",
			args: args{params: UpdateParams{
				DeploymentID:    mock.ValidClusterID,
				RefID:           "main-elasticsearch",
				RemoteResources: &successContents,
				API: api.NewMock(mock.New202ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/remote-clusters",
						Body:   mock.NewByteBody(rawBody),
					},
					mock.NewStringBody(`{}`),
				)),
			}},
		},
		{
			name: "succeeds with RefID discovery",
			args: args{params: UpdateParams{
				DeploymentID:    mock.ValidClusterID,
				RemoteResources: &successContents,
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
					mock.New202ResponseAssertion(
						&mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "PUT",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/elasticsearch/remote-clusters",
							Body:   mock.NewByteBody(rawBody),
						},
						mock.NewStringBody(`{}`),
					),
				),
			}},
		},
		{
			name: "fails on RefID discovery",
			args: args{params: UpdateParams{
				DeploymentID:    mock.ValidClusterID,
				RemoteResources: new(models.RemoteResources),
				API: api.NewMock(
					mock.New500ResponseAssertion(
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
					),
				),
			}},
			err: mock.MultierrorInternalError,
		},
		{
			name: "fails on API error",
			args: args{params: UpdateParams{
				DeploymentID:    mock.ValidClusterID,
				RefID:           "main-elasticsearch",
				RemoteResources: new(models.RemoteResources),
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/elasticsearch/main-elasticsearch/remote-clusters",
						Body:   mock.NewStringBody(`{"resources":null}` + "\n"),
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.params); !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
