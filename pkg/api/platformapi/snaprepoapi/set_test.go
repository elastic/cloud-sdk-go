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
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestSet(t *testing.T) {
	type args struct {
		params SetParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Set succeeds",
			args: args{
				params: SetParams{
					Name: "snapshot_repo_name",
					Type: "s3",
					Config: S3Config{
						Region:    "us-east-1",
						Bucket:    "mybucket",
						AccessKey: "myaccesskey",
						SecretKey: "mysecretkey",
					},
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							StatusCode: 200,
							Body:       mock.NewStringBody(`{}`),
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "PUT",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories/snapshot_repo_name",
							Body:   mock.NewStringBody(`{"settings":{"region":"us-east-1","bucket":"mybucket","access_key":"myaccesskey","secret_key":"mysecretkey"},"type":"s3"}`),
						},
					}),
				},
			},
		},
		{
			name: "Set fails due to API error",
			args: args{
				params: SetParams{
					Name: "snapshot_repo_name",
					Type: "s3",
					Config: S3Config{
						Region:    "us-east-1",
						Bucket:    "mybucket",
						AccessKey: "myaccesskey",
						SecretKey: "mysecretkey",
					},
					Region: "us-east-1",
					API:    api.NewMock(mock.Response{Error: errors.New("ERROR")}),
				},
			},
			err: &url.Error{
				Op:  "Put",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories/snapshot_repo_name",
				Err: errors.New("ERROR"),
			},
		},
		{
			name: "Set fails due to unset params",
			args: args{
				params: SetParams{},
			},
			err: multierror.NewPrefixed("invalid snapshot repository set params",
				errors.New("api reference is required for the operation"),
				errors.New("name not specified and is required for this operation"),
				errors.New("region not specified and is required for this operation"),
				errors.New("config not specified and is required for this operation"),
			),
		},
		{
			name: "Set fails due to invalid config",
			args: args{
				params: SetParams{
					Name:   "name",
					Type:   "s3",
					Region: "us-east-1",
					API: api.NewMock(mock.Response{Response: http.Response{
						StatusCode: 200,
						Body:       mock.NewStringBody(`{}`),
					}}),
					Config: new(S3Config),
				},
			},
			err: multierror.NewPrefixed("invalid snapshot repository set params",
				errors.New("s3 configuration: required setting: region cannot be empty"),
				errors.New("s3 configuration: required setting: bucket cannot be empty"),
				errors.New("s3 configuration: required setting: access key cannot be empty"),
				errors.New("s3 configuration: required setting: secret key cannot be empty"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Set(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
