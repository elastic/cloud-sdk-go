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

package response

import (
	"testing"

	"github.com/elastic/cloud-sdk-go/pkg/client"
)

func TestSetRawJSON(t *testing.T) {
	type args struct {
		params *client.Rest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "succeeds",
			args: args{params: &client.Rest{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { SetRawJSON(tt.args.params) })
	}
}

func TestUnsetRawJSON(t *testing.T) {
	type args struct {
		params *client.Rest
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "succeeds",
			args: args{params: &client.Rest{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { UnsetRawJSON(tt.args.params) })
	}
}
