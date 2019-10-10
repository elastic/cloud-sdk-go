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

package slice

import "testing"

func TestHasString(t *testing.T) {
	type args struct {
		slice []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice contains string",
			args: args{
				slice: []string{"1", "2", "3"},
				s:     "1",
			},
			want: true,
		},
		{
			name: "slice doesn't contain string",
			args: args{
				slice: []string{"1", "2", "3"},
				s:     "4",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasString(tt.args.slice, tt.args.s); got != tt.want {
				t.Errorf("HasString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsAll(t *testing.T) {
	type args struct {
		sliceA []string
		sliceB []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice contains all elements of another slice",
			args: args{
				sliceA: []string{"1", "2", "3"},
				sliceB: []string{"2", "3", "1"},
			},
			want: true,
		},
		{
			name: "slice doesn't contain all elements of another slice",
			args: args{
				sliceA: []string{"1", "2", "3"},
				sliceB: []string{"2", "3"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAll(tt.args.sliceB, tt.args.sliceA); got != tt.want {
				t.Errorf("ContainsAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice is empty",
			args: args{slice: nil},
			want: true,
		},
		{
			name: "slice is empty",
			args: args{slice: []string{}},
			want: true,
		},
		{
			name: "slice is empty",
			args: args{slice: []string{"", ""}},
			want: true,
		},
		{
			name: "slice is not empty",
			args: args{slice: []string{"", "something"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.slice); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
