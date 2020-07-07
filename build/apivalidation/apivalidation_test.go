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

// +build apivalidation

package apivalidation

import (
	"os"
	"testing"

	"github.com/elastic/cloud-sdk-go/internal/pkg/internalutil"
)

var (
	apiSpecSource = "../../" + os.Getenv("APISPEC_LOCATION")
	port          = os.Getenv("EXTERNAL_PORT")
)

func TestAPIValidation(t *testing.T) {
	type args struct {
		bin string
		arg []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Make API calls against API validation proxy",
			args: args{
				bin: "../../bin/apivalidator",
				arg: []string{
					"--source", apiSpecSource, "--port", port,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := internalutil.RunCommand(tt.args.bin, tt.args.arg...); (err != nil) != tt.wantErr {
				t.Error(err)
			}
		})
	}
}
