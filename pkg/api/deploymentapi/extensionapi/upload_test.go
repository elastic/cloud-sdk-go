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
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

func TestUpload(t *testing.T) {
	urlError := url.Error{
		Op:  "Put",
		URL: "https://mock.elastic.co/api/v1/deployments/extensions/123",
		Err: errors.New(`{"error": "some error"}`),
	}
	type args struct {
		params UploadParams
	}
	tests := []struct {
		name string
		args args
		want *models.Extension
		err  string
	}{
		{
			name: "Upload fails due to empty parameters",
			args: args{params: UploadParams{}},
			err: multierror.NewPrefixed("invalid extension upload params",
				errors.New("an extension ID is required for this operation"),
				errors.New("an extension file is required for this operation"),
				errors.New("api reference is required for the operation"),
			).Error(),
		},
		{
			name: "Upload fails due to API error",
			args: args{params: UploadParams{
				ExtensionID: "123",
				File:        strings.NewReader("extension.zip"),
				API: api.NewMock(mock.Response{
					Error: errors.New(`{"error": "some error"}`),
				}),
			}},
			err: urlError.Error(),
		},
		{
			name: "Upload Succeeds",
			args: args{params: UploadParams{
				ExtensionID: "123",
				File:        strings.NewReader("extension.zip"),
				API: api.NewMock(mock.Response{
					Response: http.Response{
						StatusCode: http.StatusOK,
						Status:     http.StatusText(http.StatusOK),
						Body:       mock.NewStringBody("{}"),
					},
				}),
				// Not testing for assertion as the content type is multipart/form-data
				// with a boundary that is a randomly generated string which changes every time.
			}},
			want: &models.Extension{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Upload(tt.args.params)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
