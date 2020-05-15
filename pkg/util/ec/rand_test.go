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

import "testing"

func TestRandomResourceID(t *testing.T) {
	someRandomString := RandomResourceID()
	tests := []struct {
		name    string
		length  int
		mustnot string
	}{
		{
			name:   "is of 32 characters length",
			length: 32,
		},
		{
			name:    "is actually random",
			length:  32,
			mustnot: someRandomString,
		},
		{
			name:    "is actually random",
			length:  32,
			mustnot: someRandomString,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomResourceID()
			if len(got) != tt.length {
				t.Errorf("RandomResourceID() = length %v, length is not %v", len(got), tt.length)
			}
			if got == tt.mustnot {
				t.Errorf("RandomResourceID() = %v, equals %v, so it's not random", got, tt.mustnot)
			}
		})
	}
}
