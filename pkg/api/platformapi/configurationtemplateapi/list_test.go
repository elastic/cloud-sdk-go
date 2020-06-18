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

package configurationtemplateapi

import (
	"errors"
	"net/http"
	"net/url"
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

func TestListTemplates(t *testing.T) {
	var templateListSuccess = `
	[{
	  "description": "Test default Elasticsearch trial template",
	  "id": "84e0bd6d69bb44e294809d89cea88a7e",
	  "name": "(Trial) Default Elasticsearch",
	  "system_owned": false
	},
	{
	  "description": "Test default Elasticsearch template",
	  "id": "0efbab9c368849a59fc5622ec750ba47",
	  "name": "Default Elasticsearch",
	  "system_owned": true
	}
  ]`
	tests := []struct {
		name string
		args ListTemplateParams
		want []*models.DeploymentTemplateInfo
		err  error
	}{
		{
			name: "Platform deployment templates list succeeds",
			args: ListTemplateParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStringBody(templateListSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Query: url.Values{
							"format":                       {"cluster"},
							"show_hidden":                  {"false"},
							"show_instance_configurations": {"false"},
						},
						Path: "/api/v1/regions/us-east-1/platform/configuration/templates/deployments",
					},
				}),
			},
			want: []*models.DeploymentTemplateInfo{
				{
					ID:          "84e0bd6d69bb44e294809d89cea88a7e",
					Description: "Test default Elasticsearch trial template",
					Name:        ec.String("(Trial) Default Elasticsearch"),
					SystemOwned: ec.Bool(false),
				},
				{
					ID:          "0efbab9c368849a59fc5622ec750ba47",
					Description: "Test default Elasticsearch template",
					Name:        ec.String("Default Elasticsearch"),
					SystemOwned: ec.Bool(true),
				},
			},
		},
		{
			name: "Platform deployment templates list fails",
			args: ListTemplateParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			},
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/templates/deployments?format=cluster&show_hidden=false&show_instance_configurations=false",
				Err: errors.New("error"),
			},
		},
		{
			name: "Platform deployment templates fails with empty params",
			err: multierror.NewPrefixed("invalid deployment template list params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListTemplates(tt.args)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
