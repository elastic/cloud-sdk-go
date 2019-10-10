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

package auth

import (
	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/client"
)

var (
	defaultMockSchema = []string{"https"}
	defaultMockHost   = "mock-host"
	defaultMockPath   = "mock-path"
)

// newMock creates a new api.API from a list of Responses. Defaults to a dummy
// APIKey for authentication, which is not checked
func newMock(res ...mock.Response) *client.Rest {
	transport := runtimeclient.NewWithClient(
		defaultMockHost,
		defaultMockPath,
		defaultMockSchema,
		mock.NewClient(res...),
	)
	return client.New(addTypeConsumers(transport), nil)
}

// addTypeConsumers adds the missing consumers and producers to the
// client.Runtime. Even though a pointer is passed it is returned too.
func addTypeConsumers(rtime *runtimeclient.Runtime) *runtimeclient.Runtime {
	rtime.Consumers["application/zip"] = runtime.ByteStreamConsumer()
	rtime.Producers["application/zip"] = runtime.ByteStreamProducer()
	rtime.Producers["multipart/form-data"] = runtime.ByteStreamProducer()
	return rtime
}
