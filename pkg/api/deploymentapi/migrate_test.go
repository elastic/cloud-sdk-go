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
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestMigrate(t *testing.T) {
	type args struct {
		params MigrateParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentTemplateMigrateResponse
		err  string
	}{
		{
			name: "fails on parameter validation",
			err: multierror.NewPrefixed("deployment migrate",
				apierror.ErrMissingAPI,
				apierror.ErrDeploymentID,
				errors.New("a target deployment template is necessary for this operation"),
			).Error(),
		},
		{
			name: "fails on API error",
			args: args{params: MigrateParams{
				API:          api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				DeploymentID: mock.ValidClusterID,
				TemplateID:   "aws-io-optimized",
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Succeeds",
			args: args{params: MigrateParams{
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/320b7b540dfc967a7a649c18e2fce4ed/_migrate",
						Query: url.Values{
							"template": {"aws-io-optimized"},
						},
					},
					mock.NewStructBody(models.DeploymentTemplateMigrateResponse{
						Resources: &models.DeploymentUpdateRequest{
							Resources: &models.DeploymentUpdateResources{
								Elasticsearch: []*models.ElasticsearchPayload{
									{
										Plan: &models.ElasticsearchClusterPlan{
											DeploymentTemplate: &models.DeploymentTemplateReference{
												ID: ec.String("aws-io-optimized"),
											},
										},
									},
								},
							},
						},
					}))),
				DeploymentID: mock.ValidClusterID,
				TemplateID:   "aws-io-optimized",
			}},
			want: &models.DeploymentTemplateMigrateResponse{
				Resources: &models.DeploymentUpdateRequest{
					Resources: &models.DeploymentUpdateResources{
						Elasticsearch: []*models.ElasticsearchPayload{
							{
								Plan: &models.ElasticsearchClusterPlan{
									DeploymentTemplate: &models.DeploymentTemplateReference{
										ID: ec.String("aws-io-optimized"),
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Migrate(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("Migrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
