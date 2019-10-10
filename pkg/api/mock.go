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

package api

import (
	"io"

	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/auth"
	"github.com/elastic/cloud-sdk-go/pkg/client"
)

var (
	defaultMockSchema = []string{"https"}
	defaultMockHost   = "mock-host"
	defaultMockPath   = "mock-path"
)

// NewMock creates a new api.API from a list of Responses. Defaults to a dummy
// APIKey for authentication, which is not checked
func NewMock(res ...mock.Response) *API {
	transport := runtimeclient.NewWithClient(
		defaultMockHost,
		defaultMockPath,
		defaultMockSchema,
		mock.NewClient(res...),
	)
	return &API{
		V1API:      client.New(AddTypeConsumers(transport), nil),
		AuthWriter: auth.APIKey("dummy"),
	}
}

// NewDebugMock creates a new api.API from a list of Responses. Defaults to a
// dummy APIKey for authentication, which is not checked. Additionally adds the
// DebugTransport so that the responses go to the configured io.Writer.
func NewDebugMock(o io.Writer, res ...mock.Response) *API {
	c := mock.NewClient(res...)
	c.Transport = NewDebugTransport(c.Transport, o)
	transport := runtimeclient.NewWithClient(
		defaultMockHost,
		defaultMockPath,
		defaultMockSchema,
		c,
	)
	return &API{
		V1API:      client.New(AddTypeConsumers(transport), nil),
		AuthWriter: auth.APIKey("dummy"),
	}
}
