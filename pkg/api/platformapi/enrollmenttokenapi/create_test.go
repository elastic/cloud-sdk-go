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

package enrollmenttokenapi

import (
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestCreate(t *testing.T) {
	type args struct {
		params CreateParams
	}
	tests := []struct {
		name string
		args args
		want *models.RequestEnrollmentTokenReply
		err  error
	}{
		{
			name: "Create fails due to incorrect duration",
			args: args{params: CreateParams{
				API:      new(api.API),
				Duration: time.Hour * 999999,
				Region:   "us-east-1",
			}},
			err: multierror.NewPrefixed("invalid enrollment-token create params",
				errors.New("validity value 3599996400 exceeds max allowed 2147483647 value in seconds"),
			),
		},
		{
			name: "Create fails due to missing API and region",
			args: args{params: CreateParams{}},
			err: multierror.NewPrefixed("invalid enrollment-token create params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "Create Succeeds with persistent token",
			args: args{params: CreateParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: 200,
						Body: mock.NewStructBody(models.RequestEnrollmentTokenReply{
							Token:   ec.String("some token"),
							TokenID: "some-token-id",
						}),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Host:   api.DefaultMockHost,
						Method: "POST",
						Body:   mock.NewStringBody(`{"persistent":true,"roles":null}` + "\n"),
						Path:   "/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens",
					},
				}),
			}},
			want: &models.RequestEnrollmentTokenReply{
				Token:   ec.String("some token"),
				TokenID: "some-token-id",
			},
		},
		{
			name: "Create fails due to API error",
			args: args{params: CreateParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Error: errors.New("error"),
				}),
			}},
			err: &url.Error{
				Op:  "Post",
				URL: "https://mock.elastic.co/api/v1/regions/us-east-1/platform/configuration/security/enrollment-tokens",
				Err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
