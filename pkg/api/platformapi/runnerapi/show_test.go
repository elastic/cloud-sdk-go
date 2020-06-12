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

package runnerapi

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

func TestShow(t *testing.T) {
	var runnerShow = `
{
  "connected": true,
  "runner_id": "192.168.44.10"
}`

	type args struct {
		params ShowParams
	}
	tests := []struct {
		name string
		args args
		want *models.RunnerInfo
		err  error
	}{
		{
			name: "Show runner succeeds",
			args: args{
				params: ShowParams{
					ID:     "192.168.44.10",
					Region: "us-east-1",
					API: api.NewMock(mock.Response{
						Response: http.Response{
							Body:       mock.NewStringBody(runnerShow),
							StatusCode: 200,
						},
						Assert: &mock.RequestAssertion{
							Header: api.DefaultReadMockHeaders,
							Method: "GET",
							Host:   api.DefaultMockHost,
							Path:   "/api/v1/regions/us-east-1/platform/infrastructure/runners/192.168.44.10",
						},
					}),
				},
			},
			want: &models.RunnerInfo{
				RunnerID:  ec.String("192.168.44.10"),
				Connected: ec.Bool(true),
			},
		},
		{
			name: "Show runner fails",
			args: args{
				params: ShowParams{
					ID:     "192.168.44.10",
					Region: "us-east-1",
					API:    api.NewMock(mock.Response{Error: errors.New("error")}),
				},
			},
			want: nil,
			err: &url.Error{
				Op:  "Get",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/infrastructure/runners/192.168.44.10",
				Err: errors.New("error"),
			},
		},
		{
			name: "Show runner fails due to validation",
			args: args{
				params: ShowParams{},
			},
			want: nil,
			err: multierror.NewPrefixed("invalid runner show params",
				errors.New("api reference is required for the operation"),
				errors.New("id not specified and is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Show(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
