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

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Contains(t *testing.T) {
	type args struct {
		slice StringSlice
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
			if got := tt.args.slice.Contains(tt.args.s); got != tt.want {
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

func TestStringSlice_IsEmpty(t *testing.T) {
	type args struct {
		slice StringSlice
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
			if got := tt.args.slice.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSlice_ToMap(t *testing.T) {
	type fields struct {
		slice StringSlice
	}
	type args struct {
		sep string
	}
	tests := []struct {
		name     string
		expected map[string]string
		fields   fields
		args     args
	}{
		{
			name:     "should return the expected map",
			fields:   fields{slice: StringSlice{"key:value", "anothervalue", ""}},
			expected: map[string]string{"key": "value"},
			args:     args{sep: ":"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.fields.slice.ToMap(tt.args.sep)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expect: \n%+v Got: \n%+v", tt.expected, actual)
			}
		})
	}
}

func TestStringSlice_ToCommaSeparatedString(t *testing.T) {
	type fields struct {
		slice StringSlice
	}
	tests := []struct {
		name     string
		expected string
		fields   fields
	}{
		{
			name:     "should return the expected string",
			fields:   fields{slice: StringSlice{"1", "2", "3"}},
			expected: "1,2,3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.fields.slice.ToCommaSeparatedString()
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Expect: \n%+v Got: \n%+v", tt.expected, actual)
			}
		})
	}
}

func TestDereference(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name     string
		args     args
		expected StringSlice
	}{
		{
			name: "should return an empty slice when input is nil",
			args: args{
				input: nil,
			},
			expected: StringSlice{},
		},
		{
			name: "should return an empty slice when input is not supported",
			args: args{
				input: []int{},
			},
			expected: StringSlice{},
		},
		{
			name: "should return the expected slice when input is a reference to string array",
			args: args{
				input: &[]string{"123", "321"},
			},
			expected: StringSlice{"123", "321"},
		},
		{
			name: "should return the expected slice when input is a reference to an interface array",
			args: args{
				input: &[]interface{}{"123", "321"},
			},
			expected: StringSlice{"123", "321"},
		},
		{
			name: "should return the expected slice when input is a reference to string map of json.RawMessage",
			args: args{
				input: &map[string]json.RawMessage{
					"123": []byte("doesn't matter"),
				},
			},
			expected: StringSlice{"123"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, Dereference(tt.args.input), tt.expected)
		})
	}
}
