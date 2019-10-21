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

package booleans

import (
	"errors"
	"reflect"
	"testing"
)

func TestCheckOnlyOneIsTrue(t *testing.T) {
	type args struct {
		bools  []bool
		errMsg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     error
	}{
		{
			name: "Returns an error if more than one element in the slice is true",
			args: args{
				bools:  []bool{true, true, true, false},
				errMsg: "more than one element is true",
			},
			wantErr: true,
			err:     errors.New("more than one element is true"),
		},
		{
			name: "Succeeds if all elements are false",
			args: args{
				bools:  []bool{false, false, false, false},
				errMsg: "more than one element is true",
			},
		},
		{
			name: "Succeeds if only one element is true",
			args: args{
				bools:  []bool{false, false, true, false},
				errMsg: "more than one element is true",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckOnlyOneIsTrue(tt.args.bools, tt.args.errMsg)

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && tt.err == nil {
				t.Errorf("Validate() expected errors = '%v' but no errors returned", tt.err)
			}

			if tt.wantErr && !reflect.DeepEqual(err, tt.err) {
				t.Errorf("Validate() expected errors = '%v' but got %v", tt.err, err)
			}
		})
	}
}
