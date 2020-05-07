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

func TestNewErrorResponse(t *testing.T) {
	type args struct {
		code int
		errs []APIError
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		{
			name: "Creates a 400 error response",
			args: args{code: 400, errs: []APIError{
				{Code: "some code", Message: "Some message", Fields: []string{"f1", "f2"}},
			}},
			want: Response{Response: http.Response{
				StatusCode: 400,
				Status:     http.StatusText(400),
				Body: NewStructBody(models.BasicFailedReply{
					Errors: []*models.BasicFailedReplyElement{
						{
							Code:    ec.String("some code"),
							Message: ec.String("Some message"),
							Fields:  []string{"f1", "f2"},
						},
					},
				}),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewErrorResponse(tt.args.code, tt.args.errs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewErrorResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
