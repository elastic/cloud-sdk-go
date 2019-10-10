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
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-openapi/spec"

	"github.com/elastic/cloud-sdk-go/cmd/generator/cloudspec"
)

var (
	sourceFile, destinationFile string
	defaultSourceFile           = "apidocs.json"
	defaultDestinationFile      = "swagger.json"
	errFileMustNotBeEmpty       = errors.New("file must not be empty")
	errFailedUnmarshalingSpec   = errors.New("failed unmarshaling spec")
)

const (
	_ = iota
	codeCannotOpenFile
	codeFileMustNotBeEmpty
	codeFailedUnmarshalingSpec
	codeFailedCreatingDestinationFile
	codeFailedEncodingSpec
)

func main() {
	flag.StringVar(&sourceFile, "source", defaultSourceFile, "source file to parse")
	flag.StringVar(&destinationFile, "destination", defaultDestinationFile, "destination file to save the result")
	flag.Parse()

	b, err := ioutil.ReadFile(sourceFile)
	exitOnError(err, codeCannotOpenFile)

	if len(b) == 0 {
		exitOnError(errFileMustNotBeEmpty, codeFileMustNotBeEmpty)
	}

	var cloudSpec *spec.Swagger
	if err := json.Unmarshal(b, &cloudSpec); err != nil {
		exitOnError(errFailedUnmarshalingSpec, codeFailedUnmarshalingSpec)
	}
	cloudSpec.Info.Version, cloudSpec.Info.Title = "v1", "rest"

	// Modifies the spec mainly to make some properties of the swagger spec
	// nullable or ommitable when they're empty to make it fully compatible
	// for Go programatic consumption.
	cloudspec.Modify(cloudSpec)

	f, err := os.Create(destinationFile)
	exitOnError(err, codeFailedCreatingDestinationFile)
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	exitOnError(enc.Encode(cloudSpec), codeFailedEncodingSpec)
}

func exitOnError(err error, code int) {
	if err != nil {
		fmt.Println(err)
		os.Exit(code)
	}
}
