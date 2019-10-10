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

package plan

import (
	"errors"
	"reflect"
	"testing"
	"time"

	multierror "github.com/hashicorp/go-multierror"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_apm"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_elasticsearch"
	"github.com/elastic/cloud-sdk-go/pkg/client/clusters_kibana"
	"github.com/elastic/cloud-sdk-go/pkg/models"
)

func TestGetLogsFromPlanActivity(t *testing.T) {
	type args struct {
		res     interface{}
		pending bool
	}
	tests := []struct {
		name string
		args args
		want []*models.ClusterPlanStepInfo
		err  error
	}{
		{
			name: "Gets the pending plan logs from an Elasticsearch Plan Activity",
			args: args{
				res: &clusters_elasticsearch.GetEsClusterPlanActivityOK{
					Payload: &models.ElasticsearchClusterPlansInfo{
						Pending: &models.ElasticsearchClusterPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID Elasticsearch pending")},
							},
						},
					},
				},
				pending: true,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID Elasticsearch pending")},
			},
		},
		{
			name: "Gets the current plan logs from an Elasticsearch Plan Activity",
			args: args{
				res: &clusters_elasticsearch.GetEsClusterPlanActivityOK{
					Payload: &models.ElasticsearchClusterPlansInfo{
						Current: &models.ElasticsearchClusterPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID Elasticsearch current")},
							},
						},
					},
				},
				pending: false,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID Elasticsearch current")},
			},
		},
		{
			name: "Gets the pending plan logs from a Kibana Plan Activity",
			args: args{
				res: &clusters_kibana.GetKibanaClusterPlanActivityOK{
					Payload: &models.KibanaClusterPlansInfo{
						Current: &models.KibanaClusterPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID Kibana pending")},
							},
						},
					},
				},
				pending: false,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID Kibana pending")},
			},
		},
		{
			name: "Gets the current plan logs from a Kibana Plan Activity",
			args: args{
				res: &clusters_kibana.GetKibanaClusterPlanActivityOK{
					Payload: &models.KibanaClusterPlansInfo{
						Current: &models.KibanaClusterPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID kibana current")},
							},
						},
					},
				},
				pending: false,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID kibana current")},
			},
		},
		{
			name: "Gets the pending plan logs from an Apm Plan Activity",
			args: args{
				res: &clusters_apm.GetApmClusterPlanActivityOK{
					Payload: &models.ApmPlansInfo{
						Pending: &models.ApmPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID Apm pending")},
							},
						},
					},
				},
				pending: true,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID Apm pending")},
			},
		},
		{
			name: "Gets the current plan logs from an Apm Plan Activity",
			args: args{
				res: &clusters_apm.GetApmClusterPlanActivityOK{
					Payload: &models.ApmPlansInfo{
						Current: &models.ApmPlanInfo{
							PlanAttemptLog: []*models.ClusterPlanStepInfo{
								{StepID: newStringPointer("ID Apm current")},
							},
						},
					},
				},
				pending: false,
			},
			want: []*models.ClusterPlanStepInfo{
				{StepID: newStringPointer("ID Apm current")},
			},
		},
		{
			name: "Gets Error on missing Payload",
			args: args{
				res:     &clusters_elasticsearch.GetEsClusterPlanActivityOK{},
				pending: true,
			},
			err: errors.New("plan: failed to obtain Payload field from Response"),
		},
		{
			name: "Gets Error on missing Payload.Pending",
			args: args{
				res: &clusters_elasticsearch.GetEsClusterPlanActivityOK{
					Payload: &models.ElasticsearchClusterPlansInfo{},
				},
				pending: true,
			},
			err: errors.New("plan: failed to obtain Pending from Payload"),
		},
		{
			name: "Gets Error on missing Payload.Current",
			args: args{
				res: &clusters_elasticsearch.GetEsClusterPlanActivityOK{
					Payload: &models.ElasticsearchClusterPlansInfo{},
				},
				pending: false,
			},
			err: errors.New("plan: failed to obtain Current from Payload"),
		},
		{
			name: "Gets Error on missing Payload.Current.PlanAttemptLog",
			args: args{
				res: &clusters_elasticsearch.GetEsClusterPlanActivityOK{
					Payload: &models.ElasticsearchClusterPlansInfo{
						Current: &models.ElasticsearchClusterPlanInfo{},
					},
				},
				pending: false,
			},
			err: errors.New("plan: failed to obtain PlanAttemptLog field from Plan"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLogsFromPlanActivityResponse(tt.args.res, tt.args.pending)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("GetLogsFromPlanActivity() error = %v, wantErr %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLogsFromPlanActivity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetParams_Validate(t *testing.T) {
	type fields struct {
		ID         string
		Kind       string
		API        *api.API
		Cooldown   time.Duration
		Pending    bool
		MaxRetries uint8
		retries    uint8
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name:   "Validate fails due to empty params",
			fields: fields{},
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan get: API cannot be nil"),
					errors.New("plan get: invalid cluster id"),
					errors.New("plan get: invalid kind "),
				},
			},
		},
		{
			name: "Validate fails due to empty id and kind",
			fields: fields{
				API: new(api.API),
			},
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan get: invalid cluster id"),
					errors.New("plan get: invalid kind "),
				},
			},
		},
		{
			name: "Validate fails due to empty kind",
			fields: fields{
				API: new(api.API),
				ID:  "9e9c997ff4d0bfc273da17f549e45e76",
			},
			err: &multierror.Error{
				Errors: []error{
					errors.New("plan get: invalid kind "),
				},
			},
		},
		{
			name: "Validate succeeds",
			fields: fields{
				API:  new(api.API),
				ID:   "9e9c997ff4d0bfc273da17f549e45e76",
				Kind: "kibana",
			},
		},
		{
			name: "Validate succeeds",
			fields: fields{
				API:  new(api.API),
				ID:   "9e9c997ff4d0bfc273da17f549e45e76",
				Kind: "kibana",
			},
		},
		{
			name: "Validate fails due to invalid kind",
			fields: fields{
				API:  new(api.API),
				ID:   "9e9c997ff4d0bfc273da17f549e45e76",
				Kind: "something invalid",
			},
			err: &multierror.Error{
				Errors: []error{errors.New("plan get: invalid kind something invalid")},
			},
		},
		{
			name: "Validate fails due to missing cooldown",
			fields: fields{
				API:        new(api.API),
				ID:         "9e9c997ff4d0bfc273da17f549e45e76",
				Kind:       "elasticsearch",
				MaxRetries: 2,
			},
			err: &multierror.Error{
				Errors: []error{errors.New("plan get: both MaxRetries and Cooldown must be set")},
			},
		},
		{
			name: "Validate succeeds with retries",
			fields: fields{
				API:        new(api.API),
				ID:         "9e9c997ff4d0bfc273da17f549e45e76",
				Kind:       "elasticsearch",
				MaxRetries: 2,
				Cooldown:   time.Nanosecond,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := GetParams{
				API:        tt.fields.API,
				ID:         tt.fields.ID,
				Kind:       tt.fields.Kind,
				Pending:    tt.fields.Pending,
				MaxRetries: tt.fields.MaxRetries,
				Cooldown:   tt.fields.Cooldown,
				retries:    tt.fields.retries,
			}
			if err := params.Validate(); !reflect.DeepEqual(err, tt.err) {
				t.Errorf("GetParams.Validate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
