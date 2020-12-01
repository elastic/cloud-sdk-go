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

func TestGet(t *testing.T) {
	var proxyGet = `
	{
            "healthy": true,
            "host_ip": "",
            "metadata": {},
            "proxy_id": "87b2c433c761",
            "public_hostname": ""
    }`[1:]

	type args struct {
		params GetParams
	}
	tests := []struct {
		name string
		args args
		want *models.ProxyInfo
		err  string
	}{
		{
			name: "Get proxy succeeds",
			args: args{params: GetParams{
				ID:     "87b2c433c761",
				Region: "us-east-1",
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       ioutil.NopCloser(strings.NewReader(proxyGet)),
					},
					Assert: &mock.RequestAssertion{
						Header: api.DefaultReadMockHeaders,
						Method: "GET",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/regions/us-east-1/platform/infrastructure/proxies/87b2c433c761",
					},
				}),
			},
			},
			want: &models.ProxyInfo{
				Healthy:        ec.Bool(true),
				HostIP:         ec.String(""),
				Metadata:       make(map[string]interface{}),
				ProxyID:        ec.String("87b2c433c761"),
				PublicHostname: ec.String(""),
			},
		},
		{
			name: "Proxy get fails",
			args: args{params: GetParams{
				ID:     "87b2c433c761",
				Region: "us-east-1",
				API:    api.NewMock(mock.New500Response(mock.NewStringBody(`{"error": "some error"}`))),
			}},
			err: `{"error": "some error"}`,
		},
		{
			name: "Get proxy fails due to validation",
			args: args{params: GetParams{}},
			err: multierror.NewPrefixed("invalid proxy params",
				errors.New("proxy id is not specified and is required for the operation"),
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
