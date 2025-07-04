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

package deploymentapi

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

func TestNewUpdateRequest(t *testing.T) {
	var apmGet, apmWant = getUpdateResponse(t,
		"./testdata/apm_get.json",
		"./testdata/apm_update.json",
	)

	// Add a resource with an empty plan so it gets ignored.
	apmGet.Resources.Elasticsearch = append(apmGet.Resources.Elasticsearch,
		&models.ElasticsearchResourceInfo{Info: &models.ElasticsearchClusterInfo{
			PlanInfo: &models.ElasticsearchClusterPlansInfo{},
		}},
	)
	apmGet.Resources.Kibana = append(apmGet.Resources.Kibana,
		&models.KibanaResourceInfo{Info: &models.KibanaClusterInfo{
			PlanInfo: &models.KibanaClusterPlansInfo{},
		}},
	)
	apmGet.Resources.Apm = append(apmGet.Resources.Apm,
		&models.ApmResourceInfo{Info: &models.ApmInfo{
			PlanInfo: &models.ApmPlansInfo{},
		}},
	)

	var appsearchGet, appsearchWant = getUpdateResponse(t,
		"./testdata/appsearch_get.json",
		"./testdata/appsearch_update.json",
	)
	appsearchGet.Resources.Appsearch = append(appsearchGet.Resources.Appsearch,
		&models.AppSearchResourceInfo{Info: &models.AppSearchInfo{
			PlanInfo: &models.AppSearchPlansInfo{},
		}},
	)

	var enterpriseSearchGet, enterpriseSearchWant = getUpdateResponse(t,
		"./testdata/enterprise_search_get.json",
		"./testdata/enterprise_search_update.json",
	)
	enterpriseSearchGet.Resources.EnterpriseSearch = append(enterpriseSearchGet.Resources.EnterpriseSearch,
		&models.EnterpriseSearchResourceInfo{Info: &models.EnterpriseSearchInfo{
			PlanInfo: &models.EnterpriseSearchPlansInfo{},
		}},
	)

	var observabilityGet, observabilityWant = getUpdateResponse(t,
		"./testdata/observability_get.json",
		"./testdata/observability_update.json",
	)

	var loggingMetricsGet, loggingMetricsWant = getUpdateResponse(t,
		"./testdata/logging_metrics_legacy_kibana_get.json",
		"./testdata/logging_metrics_legacy_kibana_update.json",
	)

	var integrationsServerGet, integrationsServerWant = getUpdateResponse(t,
		"./testdata/integrations_server_get.json",
		"./testdata/integrations_server_update.json",
	)

	type args struct {
		res *models.DeploymentGetResponse
	}
	tests := []struct {
		name string
		args args
		want *models.DeploymentUpdateRequest
	}{
		{
			name: "returns nil on empty response",
			args: args{res: nil},
		},
		{
			name: "parses a get response from a deployment with APM resources",
			args: args{res: apmGet},
			want: apmWant,
		},
		{
			name: "parses a get response from a deployment with App Search resources",
			args: args{res: appsearchGet},
			want: appsearchWant,
		},
		{
			name: "parses a get response from a deployment with Enterprise Search resources",
			args: args{res: enterpriseSearchGet},
			want: enterpriseSearchWant,
		},
		{
			name: "parses a get response from a deployment with observability settings and legacy topology",
			args: args{res: observabilityGet},
			want: observabilityWant,
		},
		{
			name: "parses a get response from a deployment with kibana legacy topology",
			args: args{res: loggingMetricsGet},
			want: loggingMetricsWant,
		},
		{
			name: "parses a get response from a deployment with integrations",
			args: args{res: integrationsServerGet},
			want: integrationsServerWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateRequest(tt.args.res); !assert.Equal(t, tt.want, got) {
				t.Errorf("NewUpdateRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getUpdateResponse(t *testing.T, get, update string) (*models.DeploymentGetResponse, *models.DeploymentUpdateRequest) {
	b, err := ioutil.ReadFile(get)
	if err != nil {
		t.Fatal(err)
	}
	var getResponse = new(models.DeploymentGetResponse)
	if err := getResponse.UnmarshalBinary(b); err != nil {
		t.Fatal(err)
	}

	updateRawB, err := ioutil.ReadFile(update)
	if err != nil {
		t.Fatal(err)
	}
	var updateWant = new(models.DeploymentUpdateRequest)
	if err := updateWant.UnmarshalBinary(updateRawB); err != nil {
		t.Fatal(err)
	}

	return getResponse, updateWant
}
