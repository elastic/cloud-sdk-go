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
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestPullToFolder(t *testing.T) {
	urlError := url.Error{
		Op:  "Get",
		URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/templates/deployments?format=cluster&show_instance_configurations=false",
		Err: errors.New("error"),
	}
	var templateListSuccess = []*models.DeploymentTemplateInfo{
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
	}
	type args struct {
		params PullToFolderParams
	}
	tests := []struct {
		name string
		args args
		err  string
		want map[string]string
	}{
		{
			name: "fails due to param validation",
			err: multierror.NewPrefixed("invalid deployment template pull params",
				errors.New("api reference is required for the operation"),
				errors.New("folder not specified and is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "fails listing the templates due to API error",
			args: args{params: PullToFolderParams{
				Region: "us-east-1",
				Folder: "some",
				Format: "cluster",
				API:    api.NewMock(mock.Response{Error: errors.New("error")}),
			}},
			err: urlError.Error(),
		},
		{
			name: "pulls templates successfully",
			args: args{params: PullToFolderParams{
				Region: "us-east-1",
				Folder: "some-folder",
				Format: "cluster",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						Body:       mock.NewStructBody(templateListSuccess),
						StatusCode: 200,
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Query: url.Values{
							"format":                       {"cluster"},
							"show_instance_configurations": {"false"},
						},
						Path: "/api/v1/regions/us-east-1/platform/configuration/templates/deployments",
					},
				}),
			}},
			want: map[string]string{
				"some-folder/0efbab9c368849a59fc5622ec750ba47.json": `{
  "description": "Test default Elasticsearch template",
  "id": "0efbab9c368849a59fc5622ec750ba47",
  "kibana_deeplink": null,
  "metadata": null,
  "name": "Default Elasticsearch",
  "system_owned": true
}
`,
				"some-folder/84e0bd6d69bb44e294809d89cea88a7e.json": `{
  "description": "Test default Elasticsearch trial template",
  "id": "84e0bd6d69bb44e294809d89cea88a7e",
  "kibana_deeplink": null,
  "metadata": null,
  "name": "(Trial) Default Elasticsearch",
  "system_owned": false
}
`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PullToFolder(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}

			if p := tt.args.params.Folder; p != "" {
				matches, err := filepath.Glob(filepath.Join(p, "*.json"))
				if err != nil {
					t.Fatal(err)
				}

				for _, m := range matches {
					want, ok := tt.want[m]
					if !ok {
						t.Error("didn't find template", m, "in want")
					}

					gotV, err := ioutil.ReadFile(m)
					if err != nil {
						t.Error(err)
					}

					if got := string(gotV); got != want {
						t.Error("got", got, "!=", "want", want)
					}
				}

				if err := os.RemoveAll(p); err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}
