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

package apierror

import (
	"errors"
	"testing"

	"github.com/go-openapi/runtime"
)

func TestIsRuntimeStatusCode(t *testing.T) {
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "returns false when the error is empty",
		},
		{
			name: "returns false when the error != runtime.APIError",
			args: args{err: errors.New("some")},
		},
		{
			name: "returns false when the error != runtime.APIError",
			args: args{
				err:  errors.New("some"),
				code: 400,
			},
		},
		{
			name: "returns false when the error is runtime.APIError, but the code doesn't match",
			args: args{
				err:  &runtime.APIError{Code: 401},
				code: 400,
			},
		},
		{
			name: "returns true when the error is runtime.APIError and the code match",
			args: args{
				err:  &runtime.APIError{Code: 403},
				code: 403,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRuntimeStatusCode(tt.args.err, tt.args.code); got != tt.want {
				t.Errorf("IsRuntimeStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
