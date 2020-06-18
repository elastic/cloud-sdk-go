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

package roleapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestList(t *testing.T) {
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.RoleAggregates
		err  error
	}{
		{
			name: "fails on parameter validation",
			args: args{},
			err: multierror.NewPrefixed("invalid role list params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
		},
		{
			name: "fails on api error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.New500Response(mock.NewStringBody(
					`{"error": "failed listing roles"}`,
				))),
			}},
			err: errors.New(`{"error": "failed listing roles"}`),
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
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/blueprinter/roles",
					},
					mock.NewStructBody(
						models.RoleAggregates{Values: []*models.RoleAggregate{
							{ID: ec.String("one")},
							{ID: ec.String("two")},
							{ID: ec.String("three")},
						}},
					),
				)),
			}},
			want: &models.RoleAggregates{Values: []*models.RoleAggregate{
				{ID: ec.String("one")},
				{ID: ec.String("two")},
				{ID: ec.String("three")},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if !assert.Equal(t, tt.err, err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
