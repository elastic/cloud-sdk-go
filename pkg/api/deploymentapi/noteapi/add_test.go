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

package noteapi

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestAdd(t *testing.T) {
	const getResponse = `{
  "healthy": true,
  "id": "e3dac8bf3dc64c528c295a94d0f19a77",
  "resources": {
    "elasticsearch": [{
      "id": "418017cd1c7f402cbb7a981b2004ceeb",
      "ref_id": "main-elasticsearch",
      "region": "ece-region"
    }]
  }
}`

	type args struct {
		params AddParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "Succeeds posting a deployment note",
			args: args{params: AddParams{
				Params: Params{
					API: api.NewMock(
						mock.Response{Response: http.Response{
							Body:       mock.NewStringBody(getResponse),
							StatusCode: 200,
						}},
						mock.Response{Response: http.Response{
							StatusCode: http.StatusCreated,
							Status:     http.StatusText(http.StatusCreated),
							Body:       mock.NewStringBody(`{}`),
						}},
					),
					ID: "e3dac8bf3dc64c528c295a94d0f19a77",
				},
				UserID:  "someid",
				Message: "note message",
			}},
		},
		{
			name: "Fails posting note (Fails to get deployment)",
			args: args{params: AddParams{
				Params: Params{
					API: api.NewMock(mock.SampleInternalError()),
					ID:  mock.ValidClusterID,
				},
				UserID:  "someid",
				Message: "note message",
			}},
			err: mock.MultierrorInternalError,
		},
		{
			name: "Fails due to parameter validation (empty params)",
			args: args{params: AddParams{}},
			err: multierror.NewPrefixed("deployment note add",
				errors.New("user id cannot be empty"),
				errors.New("note comment cannot be empty"),
				multierror.NewPrefixed("deployment note",
					errors.New("api reference is required for the operation"),
					errors.New(`id "" is invalid`),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Add(tt.args.params); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
