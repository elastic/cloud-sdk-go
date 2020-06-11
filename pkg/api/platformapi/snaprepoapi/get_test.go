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
	"errors"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	var getSnapshotSuccess = `
	{
		"repository_name": "my_snapshot_repo",
		"config": {
			"region":"us-east-1",
			"bucket":"mybucket",
			"access_key":"anaccesskey",
			"secret_key":"asecretkey"
		}
	}
`[1:]

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.RepositoryConfig
		err  error
	}{
		{
			name: "Getting a snapshot repository succeeds",
			args: args{
				params: GetParams{
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							StatusCode: 200,
							Body:       mock.NewStringBody(getSnapshotSuccess),
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories/my_snapshot_repo",
						},
					}),
					Name: "my_snapshot_repo",
				},
			},
			want: &models.RepositoryConfig{
				RepositoryName: ec.String("my_snapshot_repo"),
				Config: map[string]interface{}{
					"region":     "us-east-1",
					"bucket":     "mybucket",
					"access_key": "anaccesskey",
					"secret_key": "asecretkey",
				},
			},
		},
		{
			name: "Getting a snapshot repository fails when api returns an error",
			args: args{
				params: GetParams{
					Region: "us-east-1",
					API:    api.NewMock(mock.Response{Error: errors.New("ERROR")}),
					Name:   "my_snapshot_repo",
				},
			},
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories/my_snapshot_repo",
				Err: errors.New("ERROR"),
			},
		},
		{
			name: "Getting a snapshot repository fails when parameters are invalid",
			args: args{
				params: GetParams{},
			},
			err: multierror.NewPrefixed("invalid snapshot repository get params",
				errors.New("api reference is required for the operation"),
				errors.New("name not specified and is required for this operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)

			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
