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

package deploymentsize

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		strSize string
	}
	tests := []struct {
		name string
		args args
		want int32
		err  error
	}{
		{
			name: "parses a 0.5g (gigabyte notation)",
			args: args{strSize: "0.5g"}, want: 512,
		},
		{
			name: "parses a 8gb (gigabyte notation)",
			args: args{strSize: "8gb"}, want: 8 * 1024,
		},
		{
			name: "trying to parse 512m returns a failure",
			args: args{strSize: "512m"},
			err:  errors.New(`failed to convert "512m" to <size><g>`),
		},
		{
			name: "trying to parse 0.75g returns a failure",
			args: args{strSize: "0.75g"},
			err:  errors.New(`size "0.75g" is invalid: only increments of 0.5g are permitted`),
		},
		{
			name: "empty string returns an error",
			err:  errors.New(`failed to convert "" to <size><g>`),
		},
		{
			name: "unknown unit returns an error",
			args: args{strSize: "9999w"},
			err:  errors.New(`failed to convert "9999w" to <size><g>`),
		},
		{
			name: "invalid  prefix unit returns an error",
			args: args{strSize: "hellog"},
			err:  &strconv.NumError{Func: "ParseFloat", Num: "hello", Err: errors.New("invalid syntax")},
		},
		{
			name: "invalid size 0.25g (too small)",
			args: args{strSize: "0.25g"},
			err:  errors.New(`size "0.25g" is invalid: minimum size is 0.5g`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseGb(tt.args.strSize)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
