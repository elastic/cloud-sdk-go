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

package cloudspec

import (
	"bytes"
	"testing"

	"github.com/go-openapi/spec"
)

var zeroCommands = `# ECE 2.4.3 Command Declaration

Total of 0 commands found

## Command list

`

var oneHeadCommand = `# ECE 2.4.3 Command Declaration

Total of 1 commands found

## Command list

- ` + "`someID`" + ` (**HEAD** /some/path)
`

var twoCommands = `# ECE 2.4.3 Command Declaration

Total of 2 commands found

## Command list

- ` + "`someGetID`" + ` (**GET** /some/path)
- ` + "`someID`" + ` (**HEAD** /some/path)
`

var sixCommands = `# ECE 2.4.3 Command Declaration

Total of 6 commands found

## Command list

- ` + "`someDeleteID`" + ` (**DELETE** /some/other/path/{id})
- ` + "`someGetID`" + ` (**GET** /some/path)
- ` + "`someID`" + ` (**HEAD** /some/path)
- ` + "`somePatchID`" + ` (**PATCH** /some/other/path)
- ` + "`somePostID`" + ` (**POST** /some/other/path)
- ` + "`somePutID`" + ` (**PUT** /some/other/path/{id})
`

func TestGetCommands(t *testing.T) {
	type args struct {
		cloudSpec *spec.Swagger
		version   string
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "has no commands",
			args: args{
				version: "2.4.3",
				cloudSpec: &spec.Swagger{
					SwaggerProps: spec.SwaggerProps{
						Paths: &spec.Paths{Paths: map[string]spec.PathItem{}},
					},
				},
			},
			wantW: zeroCommands,
		},
		{
			name: "has 1 command",
			args: args{
				version: "2.4.3",
				cloudSpec: &spec.Swagger{
					SwaggerProps: spec.SwaggerProps{
						Paths: &spec.Paths{Paths: map[string]spec.PathItem{
							"/some/path": {PathItemProps: spec.PathItemProps{
								Head: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someID",
								}},
							}},
						}},
					},
				},
			},
			wantW: oneHeadCommand,
		},
		{
			name: "has 6 commands",
			args: args{
				version: "2.4.3",
				cloudSpec: &spec.Swagger{
					SwaggerProps: spec.SwaggerProps{
						Paths: &spec.Paths{Paths: map[string]spec.PathItem{
							"/some/path": {PathItemProps: spec.PathItemProps{
								Head: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someID",
								}},
								Get: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someGetID",
								}},
							}},
							"/some/other/path": {PathItemProps: spec.PathItemProps{
								Patch: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "somePatchID",
								}},
								Post: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "somePostID",
								}},
							}},
							"/some/other/path/{id}": {PathItemProps: spec.PathItemProps{
								Delete: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someDeleteID",
								}},
								Put: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "somePutID",
								}},
							}},
						}},
					},
				},
			},
			wantW: sixCommands,
		},
		{
			name: "has 2 commands",
			args: args{
				version: "2.4.3",
				cloudSpec: &spec.Swagger{
					SwaggerProps: spec.SwaggerProps{
						Paths: &spec.Paths{Paths: map[string]spec.PathItem{
							"/some/path": {PathItemProps: spec.PathItemProps{
								Head: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someID",
								}},
								Get: &spec.Operation{OperationProps: spec.OperationProps{
									ID: "someGetID",
								}},
							}},
						}},
					},
				},
			},
			wantW: twoCommands,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			GetCommands(tt.args.cloudSpec, w, tt.args.version)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("GetCommands() = \n%v, want \n%v", gotW, tt.wantW)
			}
		})
	}
}
