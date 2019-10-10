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

package maps

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/elastic/cloud-sdk-go/pkg/util/slice"
)

func TestKeysToSlice(t *testing.T) {
	tests := []struct {
		name     string
		m        interface{}
		expSlice []string
		wantErr  bool
		err      error
	}{
		{
			m: map[string]string{
				"2": "two",
				"4": "four",
				"6": "six",
				"8": "eight",
			},
			name:     "KeysToSlice returns the expected string slice of string map",
			expSlice: []string{"2", "4", "6", "8"},
		},
		{
			m: map[string]bool{
				"2": false,
				"4": true,
				"6": true,
				"8": false,
			},
			name:     "KeysToSlice returns the expected string slice bool map",
			expSlice: []string{"2", "4", "6", "8"},
		},
		{
			m: map[string]int{
				"2": 2,
				"4": 4,
				"6": 6,
				"8": 8,
			},
			name:     "KeysToSlice returns the expected string slice int map",
			expSlice: []string{"2", "4", "6", "8"},
		},
		{
			m: map[string]interface{}{
				"2": nil,
				"4": nil,
				"6": nil,
				"8": nil,
			},
			name:     "KeysToSlice returns the expected string slice interface map",
			expSlice: []string{"2", "4", "6", "8"},
		},
		{
			m: map[int]string{
				2: "two",
				4: "four",
				6: "six",
				8: "eight",
			},
			name:    "KeysToSlice returns error",
			wantErr: true,
			err:     errors.New("unexpected key type map[int]string"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := KeysToStringSlice(tt.m)

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err == nil {
				t.Errorf("Validate() expected error = '%v' but no errors returned", tt.err)
			}

			if tt.wantErr && err.Error() != tt.err.Error() {
				t.Errorf("Validate() expected error = '%v' but got %v", tt.err, err)
			}

			if !tt.wantErr && !slice.ContainsAll(tt.expSlice, got) {
				t.Errorf("Validate() expected return = '%v' but got %v", tt.expSlice, got)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		m      map[string]interface{}
		filter string
	}
	tests := []struct {
		name    string
		args    args
		expKeys []string
		expSize int
	}{
		{
			name: "slice contains all elements of another slice",
			args: args{
				m: map[string]interface{}{
					"aws-east-1":   "two",
					"west-1":       "four",
					"gcp-euwest-1": "six",
					"ibm-1":        "eight",
				},
				filter: "gcp",
			},
			expKeys: []string{"gcp-euwest-1"},
			expSize: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filtered := Filter(tt.args.m, tt.args.filter)

			if len(filtered) != tt.expSize {
				t.Errorf("Filter() result should have a size of = %v, but it has %v", tt.expSize, len(filtered))
			}
			for _, e := range tt.expKeys {
				if !HasKey(filtered, e) {
					t.Errorf("Filter() result should have key = %v, but it hasn't", e)
				}
			}
		})
	}
}
