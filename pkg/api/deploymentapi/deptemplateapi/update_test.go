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

func TestUpdate(t *testing.T) {
	updateRawDef, err := ioutil.ReadFile("./testdata/update.json")
	if err != nil {
		t.Fatal(err)
	}

	var updateReq *models.DeploymentTemplateRequestBody
	if err := json.Unmarshal(updateRawDef, &updateReq); err != nil {
		t.Fatal(err)
	}

	updateRawBody, err := json.Marshal(updateReq)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment template update params",
				errors.New("api reference is required for the operation"),
				errors.New("required template request definition not provided"),
				errors.New("required template ID not provided"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "succeeds",
			args: args{params: UpdateParams{
				TemplateID: "my-preset-id",
				Region:     "us-east-1",
				Request:    updateReq,
				API: api.NewMock(mock.New201ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/my-preset-id",
						Query: url.Values{
							"create_only": []string{"false"},
							"region":      []string{"us-east-1"},
						},
						Body: mock.NewStringBody(string(updateRawBody) + "\n"),
					},
					mock.NewStringBody(`{"id": "my-preset-id"}`),
				)),
			}},
		},
		{
			name: "fails on API error",
			args: args{params: UpdateParams{
				TemplateID: "some",
				Region:     "us-east-1",
				Request:    &models.DeploymentTemplateRequestBody{},
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PUT",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/some",
						Query: url.Values{
							"create_only": []string{"false"},
							"region":      []string{"us-east-1"},
						},
						Body: mock.NewStructBody(&models.DeploymentTemplateRequestBody{}),
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Update(tt.args.params); !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
		})
	}
}
