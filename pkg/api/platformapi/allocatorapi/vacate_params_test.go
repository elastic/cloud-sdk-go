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
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/output"
	"github.com/elastic/cloud-sdk-go/pkg/sync/pool"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestVacateParamsValidate(t *testing.T) {
	type fields struct {
		Allocators          []string
		PreferredAllocators []string
		ClusterFilter       []string
		KindFilter          string
		Region              string
		PoolTimeout         pool.Timeout
		API                 *api.API
		Output              *output.Device
		TrackFrequency      time.Duration
		AllocatorDown       *bool
		Concurrency         uint16
		MaxPollRetries      uint8
	}
	tests := []struct {
		name   string
		fields fields
		err    string
	}{
		{
			name: "Accepts a correct set of parameters",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name: "Accepts a correct set of parameters with an elasticsearch kind filter",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				KindFilter:  "elasticsearch",
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name: "Accepts a correct set of parameters with a kibana kind filter",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				KindFilter:  "kibana",
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name: "Accepts a correct set of parameters with an apm kind filter",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				KindFilter:  util.Apm,
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name: "Accepts a correct set of parameters with an appsearch kind filter",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				KindFilter:  util.Appsearch,
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name: "Accepts a correct set of parameters with an enterprise_search kind filter",
			fields: fields{
				API:         new(api.API),
				Allocators:  []string{"an allocator"},
				KindFilter:  util.EnterpriseSearch,
				Concurrency: 1,
				Output:      new(output.Device),
				Region:      "us-east-1",
			},
		},
		{
			name:   "Empty parameters are not accepted",
			fields: fields{},
			err: multierror.NewPrefixed("invalid allocator vacate params",
				errAPIMustNotBeNil,
				errMustSpecifyAtLeast1Allocator,
				errConcurrencyCannotBeZero,
				errOutputDeviceCannotBeNil,
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
		{
			name: "Cluster filter is invalid",
			fields: fields{
				API:           new(api.API),
				Allocators:    []string{"an allocator"},
				ClusterFilter: []string{"something"},
				Concurrency:   1,
				Output:        new(output.Device),
				Region:        "us-east-1",
			},
			err: multierror.NewPrefixed("invalid allocator vacate params",
				errors.New(`cluster filter: id "something" is invalid, must be 32 characters long`),
			).Error(),
		},
		{
			name: "Invalid combination of cluster filter and kind filter",
			fields: fields{
				API:           new(api.API),
				Allocators:    []string{"an allocator"},
				ClusterFilter: []string{"63d765d37613423e97b1040257cf20c8"},
				KindFilter:    "elasticsearch",
				Concurrency:   1,
				Output:        new(output.Device),
				Region:        "us-east-1",
			},
			err: multierror.NewPrefixed("invalid allocator vacate params",
				errors.New(`only one of "clusters" or "kind" can be specified`),
			).Error(),
		},
		{
			name: "Invalid combination of allocatorDown and multiple allocators",
			fields: fields{
				API:           new(api.API),
				Allocators:    []string{"an allocator", "another allocator"},
				AllocatorDown: ec.Bool(true),
				Concurrency:   1,
				Output:        new(output.Device),
				Region:        "us-east-1",
			},
			err: multierror.NewPrefixed("invalid allocator vacate params",
				errors.New(`cannot set the AllocatorDown when multiple allocators are specified`),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := VacateParams{
				API:                 tt.fields.API,
				Allocators:          tt.fields.Allocators,
				PreferredAllocators: tt.fields.PreferredAllocators,
				ClusterFilter:       tt.fields.ClusterFilter,
				Region:              tt.fields.Region,
				KindFilter:          tt.fields.KindFilter,
				Concurrency:         tt.fields.Concurrency,
				Output:              tt.fields.Output,
				MaxPollRetries:      tt.fields.MaxPollRetries,
				TrackFrequency:      tt.fields.TrackFrequency,
				PoolTimeout:         tt.fields.PoolTimeout,
				AllocatorDown:       tt.fields.AllocatorDown,
			}
			err := params.Validate()
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("VacateParams.Validate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestVacateClusterParamsValidate(t *testing.T) {
	type fields struct {
		PreferredAllocators []string
		ClusterFilter       []string
		PlanOverrides
		Region         string
		ID             string
		ClusterID      string
		Kind           string
		API            *api.API
		TrackFrequency time.Duration
		AllocatorDown  *bool
		MoveOnly       *bool
		Output         *output.Device
		OutputFormat   string
		MaxPollRetries uint8
		SkipTracking   bool
		Moves          *models.MoveClustersDetails
	}
	tests := []struct {
		name   string
		fields fields
		err    string
	}{
		{
			name: "Accepts a correct set of parameters",
			fields: fields{
				API:       new(api.API),
				ID:        "i-abc",
				ClusterID: "63d765d37613423e97b1040257cf20c8",
				Kind:      "elasticsearch",
				Output:    new(output.Device),
				Region:    "us-east-1",
				Moves:     new(models.MoveClustersDetails),
			},
		},
		{
			name:   "Empty parameters are not accepted",
			fields: fields{},
			err: multierror.NewPrefixed("invalid allocator vacate params",
				apierror.ErrMissingAPI,
				errors.New("invalid allocator ID "),
				errors.New("invalid cluster ID "),
				errors.New("invalid kind "),
				errOutputDeviceCannotBeNil,
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := VacateClusterParams{
				PreferredAllocators: tt.fields.PreferredAllocators,
				ClusterFilter:       tt.fields.ClusterFilter,
				PlanOverrides:       tt.fields.PlanOverrides,
				Region:              tt.fields.Region,
				ID:                  tt.fields.ID,
				ClusterID:           tt.fields.ClusterID,
				Kind:                tt.fields.Kind,
				API:                 tt.fields.API,
				TrackFrequency:      tt.fields.TrackFrequency,
				AllocatorDown:       tt.fields.AllocatorDown,
				Output:              tt.fields.Output,
				OutputFormat:        tt.fields.OutputFormat,
				MaxPollRetries:      tt.fields.MaxPollRetries,
				SkipTracking:        tt.fields.SkipTracking,
			}
			err := params.Validate()
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("VacateParams.Validate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
