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

package extensionapi

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestUpdate(t *testing.T) {
	type args struct {
		params UpdateParams
	}
	tests := []struct {
		name string
		args args
		want *models.Extension
		err  string
	}{
		{
			name: "fails due to parameter validation",
			args: args{params: UpdateParams{
				DownloadURL: "imaurl",
			}},
			err: multierror.NewPrefixed("invalid extension update params",
				apierror.ErrMissingAPI,
				errors.New("an extension ID is required for this operation"),
				errors.New("an extension type is required for this operation"),
				errors.New(`the provided URL is invalid: parse "imaurl": invalid URI for request`),
			).Error(),
		},
		{
			name: "succeeds",
			args: args{params: UpdateParams{
				ExtensionID: "someid",
				Name:        "Boop",
				Version:     "v1.0",
				Type:        "sometype",
				DownloadURL: "https://www.example.com",
				Description: "Why hello there",
				API: api.NewMock(mock.New200ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/extensions/someid",
						Body:   mock.NewStringBody(`{"description":"Why hello there","download_url":"https://www.example.com","extension_type":"sometype","name":"Boop","version":"v1.0"}` + "\n"),
					},
					mock.NewStructBody(models.Extension{
						ID:            ec.String("someid"),
						Name:          ec.String("Boop"),
						Version:       ec.String("v1.0"),
						ExtensionType: ec.String("sometype"),
						DownloadURL:   "https://www.example.com",
						Description:   "Why hello there",
					}),
				)),
			}},
			want: &models.Extension{
				ID:            ec.String("someid"),
				Name:          ec.String("Boop"),
				Version:       ec.String("v1.0"),
				ExtensionType: ec.String("sometype"),
				DownloadURL:   "https://www.example.com",
				Description:   "Why hello there",
			},
		},
		{
			name: "fails on API error",
			args: args{params: UpdateParams{
				ExtensionID: "someid",
				Type:        "sometype",
				API: api.NewMock(mock.New500ResponseAssertion(
					&mock.RequestAssertion{
						Header: api.DefaultWriteMockHeaders,
						Method: "POST",
						Host:   api.DefaultMockHost,
						Path:   "/api/v1/deployments/extensions/someid",
						Body:   mock.NewStringBody(`{"extension_type":"sometype","name":"","version":""}` + "\n"),
					},
					mock.SampleInternalError().Response.Body,
				)),
			}},
			err: mock.MultierrorInternalError.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Error(err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
