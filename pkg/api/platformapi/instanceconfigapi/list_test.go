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

package instanceconfigapi

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
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
		want []*models.InstanceConfiguration
		err  error
	}{
		{
			name: "List succeeds",
			args: args{
				params: ListParams{
					Region: "us-east-1",
					API: api.NewMock(mock.Response{Response: http.Response{
						Body:       mock.NewStringBody(listInstanceConfigsSuccess),
						StatusCode: 200,
					}}),
				},
			},
			want: []*models.InstanceConfiguration{
				{
					ID:                "data.highstorage",
					Description:       "Instance configuration to be used for a higher disk/memory ratio",
					Name:              ec.String("data.highstorage"),
					InstanceType:      ec.String("elasticsearch"),
					StorageMultiplier: float64(32),
					NodeTypes:         []string{"data", "ingest", "master"},
					DiscreteSizes: &models.DiscreteSizes{
						DefaultSize: ec.Int32(1024),
						Resource:    ec.String("memory"),
						Sizes: []int32{
							1024,
							2048,
							4096,
							8192,
							16384,
							32768,
							65536,
							131072,
							262144,
						},
					},
				},
				{
					ID:                "kibana",
					Description:       "Instance configuration to be used for Kibana",
					Name:              ec.String("kibana"),
					InstanceType:      ec.String("kibana"),
					StorageMultiplier: float64(4),
					DiscreteSizes: &models.DiscreteSizes{
						DefaultSize: ec.Int32(1024),
						Resource:    ec.String("memory"),
						Sizes: []int32{
							1024,
							2048,
							4096,
							8192,
						},
					},
				},
			},
		},
		{
			name: "List fails on API error",
			args: args{params: ListParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "List fails on parameter validation failure",
			args: args{},
			err: multierror.NewPrefixed("invalid instance config list params",
				apierror.ErrMissingAPI,
				errors.New("region not specified and is required for this operation"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.params)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("List() error = %v, wantErr %v", err, tt.err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
