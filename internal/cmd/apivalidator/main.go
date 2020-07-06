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

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/elastic/cloud-sdk-go/internal/cmd/apivalidator/requests"
)

const (
	defaultSourceFile = "https://raw.githubusercontent.com/elastic/cloud-sdk-go/master/api/apidocs-user.json"
	defaultPort       = "4010"
	host              = "http://127.0.0.1"
)

func main() {
	var sourceFile, port string
	flag.StringVar(&sourceFile, "source", defaultSourceFile, "source file to parse")
	flag.StringVar(&port, "port", defaultPort, "validation proxy port")
	flag.Parse()

	client := new(http.Client)

	config := requests.Config{
		Source: sourceFile,
		Host:   host,
		Port:   port,
		Client: client,
	}

	code, err := requests.Run(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(code)
	}
}
