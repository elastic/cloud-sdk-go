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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestAdd(t *testing.T) {
	type args struct {
		params AddParams
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "Succeeds posting a deployment note",
			args: args{params: AddParams{
				Params: Params{
					Region: "us-east-1",
					API: api.NewMock(
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
					Region: "us-east-1",
					API:    api.NewMock(mock.SampleInternalError()),
					ID:     mock.ValidClusterID,
				},
				UserID:  "someid",
				Message: "note message",
			}},
			err: mock.MultierrorInternalError.Error(),
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
					errors.New("region not specified and is required for this operation"),
				),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Add(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
