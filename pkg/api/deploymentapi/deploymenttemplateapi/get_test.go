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

package deploymenttemplateapi

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

func TestGet(t *testing.T) {
	getRawResp, err := ioutil.ReadFile("./testdata/get.json")
	if err != nil {
		t.Fatal(err)
	}

	var succeedResp *models.DeploymentTemplateInfoV2
	if err := json.Unmarshal(getRawResp, &succeedResp); err != nil {
		t.Fatal(err)
	}
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentTemplateInfoV2
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid deployment template get params",
				errors.New("api reference is required for the operation"),
				errors.New("required template id not provided"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "succeeds",
			args: args{params: GetParams{
				Region:     "us-east-1",
				TemplateID: "default",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/default",
						Query: url.Values{
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
						},
					},
					mock.NewByteBody(getRawResp),
				)),
			}},
			want: succeedResp,
		},
		{
			name: "succeeds with stack filter",
			args: args{params: GetParams{
				Region:       "us-east-1",
				TemplateID:   "default",
				StackVersion: "6.8.0",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/default",
						Query: url.Values{
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
							"stack_version":                []string{"6.8.0"},
						},
					},
					mock.NewByteBody(getRawResp),
				)),
			}},
			want: succeedResp,
		},
		{
			name: "fails on API error",
			args: args{params: GetParams{
				Region:     "us-east-1",
				TemplateID: "some-id",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/templates/some-id",
						Query: url.Values{
							"region":                       []string{"us-east-1"},
							"show_instance_configurations": []string{"true"},
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
			got, err := Get(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
