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

// contains the logic for the generator command the aim of this is to process
// the swagger definition that's obtained from Elastic Cloud and make a few
// changes in order for the SDK to be fully usable by Golang
//
// It will go over the swagger specification and set all of the boolean types
// to nullable using the vendor extension "x-nullable", this will cause any
// bool type to be converted to *bool in the Cloud SDK. This is required in
// order to fully use the V1 API.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const versionFmt = `package api

// Version contains the ECE API version compatibility for the current
// auto-generated client and models. This needs to be updated every time a new
// "api/apidocs.json" file added and the client and models are re-generated.
// Even though the Bugfix version is specified, the general support statement is
// on the minor version, i.e: 2.4.2 means that all the 2.4 branch is supported,
// expecting some potentially unfixed bugs when ECE version is the same feature
// version bug higher bugfix version.
const Version = "%s"
`

func Test_genVersionFile(t *testing.T) {
	type args struct {
		file    string
		version string
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		wantContents string
	}{
		{
			name: "Generates  file with 2.4.3 contents",
			args: args{
				version: "2.4.3",
			},
			wantContents: fmt.Sprintf(versionFmt, "2.4.3"),
		},
		{
			name: "Generates  file with 2.5.0 contents",
			args: args{
				version: "2.5.0",
			},
			wantContents: fmt.Sprintf(versionFmt, "2.5.0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := ioutil.TempFile("", tt.args.file)
			if err != nil {
				t.Error(err)
			}

			defer func() {
				f.Close()
				os.RemoveAll(f.Name())
			}()

			if err := genVersionFile(f.Name(), tt.args.version); (err != nil) != tt.wantErr {
				t.Errorf("genVersionFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Error(err)
			}
			if string(b) != tt.wantContents {
				t.Errorf("genVersionFile() contents = %v, wantContents %v", string(b), tt.wantContents)
			}
		})
	}
}
