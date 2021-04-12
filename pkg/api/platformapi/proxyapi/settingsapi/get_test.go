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
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestGet(t *testing.T) {
	var proxiesSettings = `
{
   "expected_proxies_count" : 5,
   "http_settings" : {
      "cookie_secret" : "some-secret",
      "dashboards_base_url" : "some-url",
      "disconnected_cutoff" : 2,
      "minimum_proxy_services" : 1,
      "sso_settings" : {
         "cookie_name" : "some-cookie",
         "default_redirect_path" : "some-path",
         "dont_log_requests" : true,
         "maintenance_bypass_cookie_name" : "some-other-cookie",
         "max_age" : 10,
         "sso_secret" : "sso-secret"
      },
      "user_cookie_key" : "user-cookie"
   },
   "signature_secret" : "signature-secret",
   "signature_valid_for_millis" : 60
}`
	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxiesSettings
		err  string
	}{
		{
			name: "Proxies settings get succeeds",
			args: args{params: GetParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxiesSettings)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/settings",
					},
				}),
			}},
			want: &models.ProxiesSettings{
				ExpectedProxiesCount: ec.Int32(5),
				HTTPSettings: &models.ProxiesHTTPSettings{
					CookieSecret:         ec.String("some-secret"),
					DashboardsBaseURL:    ec.String("some-url"),
					DisconnectedCutoff:   ec.Int64(2),
					MinimumProxyServices: ec.Int32(1),
					SsoSettings: &models.ProxiesSSOSettings{
						CookieName:                  ec.String("some-cookie"),
						DefaultRedirectPath:         ec.String("some-path"),
						DontLogRequests:             ec.Bool(true),
						MaintenanceBypassCookieName: ec.String("some-other-cookie"),
						MaxAge:                      ec.Int64(10),
						SsoSecret:                   ec.String("sso-secret"),
					},
					UserCookieKey: ec.String("user-cookie"),
				},
				SignatureSecret:         ec.String("signature-secret"),
				SignatureValidForMillis: ec.Int64(60),
			},
		},
		{
			name: "Proxies settings get fails when api call fails",
			args: args{params: GetParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Proxies filtered group get fails due validation",
			args: args{params: GetParams{}},
			err: multierror.NewPrefixed("invalid get proxy settings params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			).Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			if !assert.Equal(t, tt.want, got) {
				t.Error(err)
			}
		})
	}
}
