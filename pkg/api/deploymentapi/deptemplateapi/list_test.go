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

func TestList(t *testing.T) {
	listRawResp, err := ioutil.ReadFile("./testdata/list.json")
	if err != nil {
		t.Fatal(err)
	}

	var succeedResp []*models.DeploymentTemplateInfoV2
	if err := json.Unmarshal(listRawResp, &succeedResp); err != nil {
		t.Fatal(err)
	}
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want []*models.DeploymentTemplateInfoV2
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment template list params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "fails due to parameter invalid metadata filter",
			args: args{params: ListParams{
				MetadataFilter: "somewrongful value",
			}},
			err: multierror.NewPrefixed("invalid deployment template list params",
				errors.New("api reference is required for the operation"),
				errors.New(`invalid metadata filter "somewrongful value", must be formatted in the form of (key:value)`),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
							"show_hidden":                  []string{"false"},
						},
					},
					mock.NewByteBody(listRawResp),
				)),
			}},
			want: succeedResp,
		},
		{
			name: "succeeds with metadata filter",
			args: args{params: ListParams{
				Region:         "us-east-1",
				MetadataFilter: "parent_solution:stack",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"metadata":                     []string{"parent_solution:stack"},
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
							"show_hidden":                  []string{"false"},
						},
					},
					mock.NewByteBody(listRawResp),
				)),
			}},
			want: succeedResp,
		},
		{
			name: "succeeds with metadata and stack_version filter",
			args: args{params: ListParams{
				Region:         "us-east-1",
				MetadataFilter: "parent_solution:stack",
				StackVersion:   "6.2.1",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"metadata":                     []string{"parent_solution:stack"},
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
							"show_hidden":                  []string{"false"},
							"stack_version":                []string{"6.2.1"},
						},
					},
					mock.NewByteBody(listRawResp),
				)),
			}},
			want: succeedResp,
		},
		{
			name: "fails on API error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates",
						Query: url.Values{
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
							"show_hidden":                  []string{"false"},
						},
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
