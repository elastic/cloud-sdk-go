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

package trafficfilterapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
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
		want *models.TrafficFilterRulesetResponse
		err  error
	}{
		{
			name: "fails due to parameter validation",
			err: multierror.NewPrefixed("invalid traffic filter create params",
				apierror.ErrMissingAPI,
				errors.New("request payload is not specified and is required for the operation"),
			),
		},
		{
			name: "succeeds",
			args: args{params: CreateParams{
				API: api.NewMock(mock.New201ResponseAssertion(
					&mock.RequestAssertion{
						Host:   api.DefaultMockHost,
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Path:   "/api/v1/deployments/traffic-filter/rulesets",
						Body:   mock.NewStringBody(`{"include_by_default":false,"name":"some name","region":"us-east-1","rules":[{"source":"0.0.0.0/0"}],"type":"ip"}` + "\n"),
					},
					mock.NewStringBody(`{"id": "some-id"}`),
				)),
				Req: &models.TrafficFilterRulesetRequest{
					IncludeByDefault: ec.Bool(false),
					Name:             ec.String("some name"),
					Region:           ec.String("us-east-1"),
					Type:             ec.String("ip"),
					Rules: []*models.TrafficFilterRule{{
						Source: ec.String("0.0.0.0/0"),
					}},
				},
			}},
			want: &models.TrafficFilterRulesetResponse{
				ID: ec.String("some-id"),
			},
		},
		{
			name: "fails",
			args: args{params: CreateParams{
				API: api.NewMock(mock.SampleInternalError()),
				Req: &models.TrafficFilterRulesetRequest{
					IncludeByDefault: ec.Bool(false),
					Name:             ec.String("some name"),
					Region:           ec.String("us-east-1"),
					Type:             ec.String("ip"),
					Rules: []*models.TrafficFilterRule{{
						Source: ec.String("0.0.0.0/0"),
					}},
				},
			}},
			err: mock.MultierrorInternalError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
