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

package apierror

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments_traffic_filter"
)

func TestError_Unwrap(t *testing.T) {
	type fields struct {
		Err error
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name:   "unwraps an traffic filter error (not found)",
			fields: fields{Err: deployments_traffic_filter.NewCreateTrafficFilterRulesetAssociationNotFound()},
			err:    &deployments_traffic_filter.CreateTrafficFilterRulesetAssociationNotFound{},
		},
		{
			name:   "unwraps an deployment creation error (bad request)",
			fields: fields{Err: deployments.NewCreateDeploymentBadRequest()},
			err:    &deployments.CreateDeploymentBadRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Err: tt.fields.Err,
			}

			if err := e.Unwrap(); !assert.Equal(t, tt.err, err) {
				t.Errorf("Unwrap() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "wraps no error",
		},
		{
			name: "wraps an error",
			args: args{err: deployments.NewCreateDeploymentBadRequest()},
			want: &Error{Err: &deployments.CreateDeploymentBadRequest{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.err); !assert.Equal(t, tt.want, got) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
