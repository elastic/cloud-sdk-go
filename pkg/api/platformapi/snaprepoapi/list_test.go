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

package snaprepoapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	var listSnapshotsSuccess = `
	{
		"configs": [
			{
				"config": {
					"access_key": "myaccesskey",
					"base_path": "apath",
					"bucket": "mybucket",
					"canned_acl": "private",
					"compress": true,
					"protocol": "http",
					"region": "us-east-1",
					"secret_key": "mysupersecretkey",
					"server_side_encryption": true,
					"storage_class": "standard"
				},
				"repository_name": "my_repo_1"
			},
			{
				"config": {
					"access_key": "myaccesskey",
					"bucket": "mybucket",
					"region": "us-east-1",
					"secret_key": "mysupersecretkey",
					"server_side_encryption": true,
					"storage_class": "standard"
				},
				"repository_name": "my_repo_2"
			}
		]
	}
`[1:]

	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.RepositoryConfigs
		err  error
	}{
		{
			name: "List Succeeds",
			args: args{
				params: ListParams{
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							StatusCode: 200,
							Body:       mock.NewStringBody(listSnapshotsSuccess),
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories",
						},
					}),
				},
			},
			want: &models.RepositoryConfigs{
				Configs: []*models.RepositoryConfig{
					{
						RepositoryName: ec.String("my_repo_1"),
						Config: map[string]interface{}{
							"region":                 "us-east-1",
							"bucket":                 "mybucket",
							"access_key":             "myaccesskey",
							"secret_key":             "mysupersecretkey",
							"base_path":              "apath",
							"compress":               true,
							"server_side_encryption": true,
							"canned_acl":             "private",
							"storage_class":          "standard",
							"protocol":               "http",
						},
					},
					{
						RepositoryName: ec.String("my_repo_2"),
						Config: map[string]interface{}{
							"region":                 "us-east-1",
							"bucket":                 "mybucket",
							"access_key":             "myaccesskey",
							"secret_key":             "mysupersecretkey",
							"server_side_encryption": true,
							"storage_class":          "standard",
						},
					},
				},
			},
		},
		{
			name: "List fails when the API returns an error",
			args: args{
				params: ListParams{
					Region: "us-east-1",
					API:    api.NewMock(mock.Response{Error: errors.New("ERROR")}),
				},
			},
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories",
				Err: errors.New("ERROR"),
			},
		},
		{
			name: "List fails when parameters are not valid",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid snapshot repository list params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			d := json.NewEncoder(os.Stdout)
			d.SetIndent("", "    ")
			d.Encode(tt.want)

			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
