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

package mock

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

func TestSampleInternalError(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		{
			name: "obtains the response",
			want: Response{Response: http.Response{
				StatusCode: 500,
				Status:     http.StatusText(500),
				Body: NewStructBody(models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{Code: ec.String(code500), Message: ec.String(message500)},
					},
				}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SampleInternalError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SampleInternalError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleNotFoundError(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		{
			name: "obtains the response",
			want: Response{Response: http.Response{
				StatusCode: 404,
				Status:     http.StatusText(404),
				Body: NewStructBody(models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{Code: ec.String(code404), Message: ec.String(message404)},
					},
				}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SampleNotFoundError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SampleNotFoundError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSampleBadRequestError(t *testing.T) {
	tests := []struct {
		name string
		want Response
	}{
		{
			name: "obtains the response",
			want: Response{Response: http.Response{
				StatusCode: 400,
				Status:     http.StatusText(400),
				Body: NewStructBody(models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{Code: ec.String(code400), Message: ec.String(message400)},
					},
				}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SampleBadRequestError(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SampleBadRequestError() = %v, want %v", got, tt.want)
			}
		})
	}
}
