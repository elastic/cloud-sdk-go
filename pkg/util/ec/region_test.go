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

package ec

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequireRegionSet(t *testing.T) {
	type args struct {
		region string
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "empty region returns error",
			err:  errors.New("region not specified and is required for this operation"),
		},
		{
			name: "spaced region returns error",
			args: args{region: "    "},
			err:  errors.New("region not specified and is required for this operation"),
		},
		{
			name: "non empty region returns no error",
			args: args{region: "some-region"},
		},
		{
			name: "non empty region with space prefix and suffix returns no error",
			args: args{region: "   some-region  "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RequireRegionSet(tt.args.region)
			assert.Equal(t, tt.err, err)
		})
	}
}
