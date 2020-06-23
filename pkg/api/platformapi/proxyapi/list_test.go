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

package proxyapi

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

func TestList(t *testing.T) {
	var proxyList = `
	{
		"proxies": [
		  {
            "healthy": true,
            "host_ip": "",
            "metadata": {},
            "proxy_id": "87b2c433c761",
            "public_hostname": ""
		  }
		]
	}`[1:]
	type args struct {
		params ListParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxyOverview
		err  error
	}{
		{
			name: "Proxy list succeeds",
			args: args{params: ListParams{
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxyList)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies",
					},
				}),
			}},
			want: &models.ProxyOverview{
				Proxies: []*models.ProxyInfo{
					{
						Healthy:        ec.Bool(true),
						HostIP:         ec.String(""),
						Metadata:       make(map[string]interface{}),
						ProxyID:        ec.String("87b2c433c761"),
						PublicHostname: ec.String(""),
					},
				},
			},
		},
		{
			name: "Proxy list fails",
			args: args{params: ListParams{
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: errors.New(`{"error": "some error"}`),
		},
		{
			name: "Proxy list fails due to validation",
			args: args{params: ListParams{}},
			err: multierror.NewPrefixed("invalid proxy params",
				errors.New("api reference is required for the operation"),
				errors.New("region not specified and is required for this operation"),
			),
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
