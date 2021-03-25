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
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestNewUpdateRequestFromGet(t *testing.T) {
	type args struct {
		res *models.TrafficFilterRulesetInfo
	}
	tests := []struct {
		name string
		args args
		want *models.TrafficFilterRulesetRequest
		err  string
	}{
		{
			name: "succeeds with nil input",
			args: args{res: nil},
			want: nil,
		},
		{
			name: "succeeds with empty input",
			args: args{res: &models.TrafficFilterRulesetInfo{}},
			want: &models.TrafficFilterRulesetRequest{},
		},
		{
			name: "succeeds",
			args: args{res: &models.TrafficFilterRulesetInfo{
				Description:      "This is my traffic filter",
				IncludeByDefault: ec.Bool(true),
				Region:           ec.String("us-east1"),
				Name:             ec.String("my traffic filter"),
				Type:             ec.String("ip"),
				Rules: []*models.TrafficFilterRule{{
					Source: "0.0.0.0/0",
				}},
			}},
			want: &models.TrafficFilterRulesetRequest{
				Description:      "This is my traffic filter",
				IncludeByDefault: ec.Bool(true),
				Region:           ec.String("us-east1"),
				Name:             ec.String("my traffic filter"),
				Type:             ec.String("ip"),
				Rules: []*models.TrafficFilterRule{{
					Source: "0.0.0.0/0",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUpdateRequestFromGet(tt.args.res)
			assert.Equal(t, tt.want, got)
		})
	}
}
