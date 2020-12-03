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

package deptemplateapi

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestCreate(t *testing.T) {
	createRawDef, err := ioutil.ReadFile("./testdata/create.json")
	if err != nil {
		t.Fatal(err)
	}

	var createReq *models.DeploymentTemplateRequestBody
	if err := json.Unmarshal(createRawDef, &createReq); err != nil {
		t.Fatal(err)
	}

	createRawBody, err := json.Marshal(createReq)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		params CreateParams
	}
	tests := []struct {
		name string
		args args
		want string
		err  string
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment template create params",
				errors.New("api reference is required for the operation"),
				errors.New("required template request definition not provided"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: CreateParams{
				Region:  "us-east-1",
				Request: createReq,
				API: api.NewMock(mock.New201ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"region": []string{"us-east-1"},
						},
						Body: mock.NewStringBody(string(createRawBody) + "\n"),
					},
					mock.NewStringBody(`{"id": "some-randomly-generated-id"}`),
				)),
			}},
			want: "some-randomly-generated-id",
		},
		{
			name: "succeeds with a preset ID",
			args: args{params: CreateParams{
				TemplateID: "my-preset-id",
				Region:     "us-east-1",
				Request:    createReq,
				API: api.NewMock(mock.New201ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/my-preset-id",
						Query: url.Values{
							"create_only": []string{"true"},
							"region":      []string{"us-east-1"},
						},
						Body: mock.NewStringBody(string(createRawBody) + "\n"),
					},
					mock.NewStringBody(`{"id": "my-preset-id"}`),
				)),
			}},
			want: "my-preset-id",
		},
		{
			name: "fails on API error",
			args: args{params: CreateParams{
				Region:  "us-east-1",
				Request: &models.DeploymentTemplateRequestBody{},
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"region": []string{"us-east-1"},
						},
						Body: mock.NewStructBody(&models.DeploymentTemplateRequestBody{}),
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
