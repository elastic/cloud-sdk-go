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

package api

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
)

func TestReturnErrOnly(t *testing.T) {
	type args struct {
		in0 interface{}
		err error
	}
	tests := []struct {
		name string
		args args
		err  string
	}{
		{
			name: "returns nil when the passed error is nil",
			args: args{
				in0: "",
			},
		},
		{
			name: "returns the error when it receives one",
			args: args{
				in0: "with error",
				err: authentication.NewDeleteUserAPIKeysNotFound(),
			},
			err: authentication.NewDeleteUserAPIKeysNotFound().Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ReturnErrOnly(tt.args.in0, tt.args.err)
			if err != nil && !assert.EqualError(t, err, tt.err) {
				t.Errorf("ReturnErrOnly() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}
