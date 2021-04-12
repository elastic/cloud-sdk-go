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

package settingsapi

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestPatch(t *testing.T) {

	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxiesSettings
		err  string
	}{
		{
			name: "Proxies settings set succeeds",
			args: args{params: UpdateParams{
				Region:          "us-east-1",
				ProxiesSettings: proxySettings(),
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       mock.NewStructBody(proxySettings()),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "PATCH",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/settings",
						Body:   mock.NewStringBody(`{"expected_proxies_count":5,"http_settings":{"cookie_secret":"some-secret","dashboards_base_url":"some-url","disconnected_cutoff":2,"minimum_proxy_services":1,"sso_settings":{"cookie_name":"some-cookie","default_redirect_path":"some-path","dont_log_requests":true,"maintenance_bypass_cookie_name":"some-other-cookie","max_age":10,"sso_secret":"sso-secret"},"user_cookie_key":"user-cookie"},"signature_secret":"signature-secret","signature_valid_for_millis":60}` + "\n"),
					},
				}),
			}},
			want: proxySettings(),
		},
		{
			name: "Proxies settings set fails when api call fails",
			args: args{params: UpdateParams{
				Region:          "us-east-1",
				API:             api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
				ProxiesSettings: &models.ProxiesSettings{},
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Proxies settings set fails due validation",
			args: args{params: UpdateParams{}},
			err: multierror.NewPrefixed("invalid set proxy settings params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
				errors.New("a proxy settings object is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Patch(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
