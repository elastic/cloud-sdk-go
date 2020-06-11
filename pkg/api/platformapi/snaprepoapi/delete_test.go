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

func TestDelete(t *testing.T) {
	type args struct {
		params DeleteParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Delete succeeds",
			args: args{
				params: DeleteParams{
					Name:   "my_repo",
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							StatusCode: 200,
							Body:       mock.NewStringBody(`{}`),
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultWriteMockHeaders,
							Method: "DELETE",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/configuration/snapshots/repositories/my_repo",
							Query: url.Values{
								"cleanup_deployments": []string{"false"},
							},
						},
					}),
				},
			},
		},
		{
			name: "Delete fails on 404",
			args: args{
				params: DeleteParams{
					Name:   "my_repo",
					Region: "us-east-1",
					API:    api.NewMock(mock.New404Response(mock.NewStringBody(`{"error": "some error"}`))),
				},
			},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "Delete fails on invalid params",
			args: args{
				params: DeleteParams{},
			},
			err: multierror.NewPrefixed("invalid snapshot repository delete params",
				errors.New("api reference is required for the operation"),
				errors.New("name not specified and is required for this operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Delete(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
