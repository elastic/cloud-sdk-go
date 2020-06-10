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

package allocatorapi

import (
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestTagsToMap(t *testing.T) {
	tests := []struct {
		name string
		tags string
		want map[string]string
	}{
		{
			name: "Should return correct map with tags",
			tags: "[tag1:value1,tag2:value2]",
			want: map[string]string{
				"tag1": "value1",
				"tag2": "value2",
			},
		},
		{
			name: "Should return empty map if tags string is empty",
			tags: "",
			want: make(map[string]string),
		},
		{
			name: "Should ignore empty tags",
			tags: "tag1:value1,tag2",
			want: map[string]string{
				"tag1": "value1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tagsToMap(tt.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagsToMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterConnectedOrWithInstances(t *testing.T) {
	aConnected := &models.AllocatorInfo{
		Instances: []*models.AllocatedInstanceStatus{},
		Status: &models.AllocatorHealthStatus{
			Connected: ec.Bool(true),
		},
	}
	aWithInstances := &models.AllocatorInfo{
		Instances: []*models.AllocatedInstanceStatus{
			{ClusterID: ec.String("123")},
		},
		Status: &models.AllocatorHealthStatus{
			Connected: ec.Bool(false),
		},
	}
	aNotConnectedNoInstances := &models.AllocatorInfo{
		Instances: []*models.AllocatedInstanceStatus{},
		Status: &models.AllocatorHealthStatus{
			Connected: ec.Bool(false),
		},
	}
	allocators := []*models.AllocatorInfo{aConnected, aWithInstances, aNotConnectedNoInstances}

	tests := []struct {
		name   string
		input  []*models.AllocatorInfo
		wanted []*models.AllocatorInfo
	}{
		{
			name:   "Should return connected allocators or allocators with at least one instance",
			input:  allocators,
			wanted: []*models.AllocatorInfo{aConnected, aWithInstances},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := FilterConnectedOrWithInstances(tt.input)

			if len(tt.wanted) != len(filtered) {
				t.Errorf("FilterConnectedOrWithInstances() got = %d, want %d", len(tt.wanted), len(filtered))
			}

			for _, a := range tt.wanted {
				if !contains(filtered, a) {
					t.Errorf("FilterConnectedOrWithInstances() should contain = %v, but it doesn't", a)
				}
			}
		})
	}
}
