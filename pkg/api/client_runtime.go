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
	"fmt"
	"net/url"
	"strings"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"
)

const (
	// RegionPrefix is used when a Region is passed as part of the API config.
	RegionPrefix = "/api/v1/regions/%s"

	// RegionlessPrefix is used when no region is specified, assumed target is
	// most likely an ECE installation or a non federated one.
	RegionlessPrefix = "/api/v1"

	rawMetadataTextProducer = "set-es-cluster-metadata-raw"
)

// NewCloudClientRuntime creates a CloudClientRuntime from the config. Using
// the configured region (if any) to instantiate two different client.Runtime.
// If there's no region specified in the config then both are regionless.
func NewCloudClientRuntime(c Config) (*CloudClientRuntime, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, err
	}

	var basepath = RegionlessPrefix
	if c.Region != "" {
		basepath = fmt.Sprintf(RegionPrefix, c.Region)
	}

	// Additional consumers and producers that are needed for parts of the SDK
	// to work correctly
	rr := AddTypeConsumers(runtimeclient.NewWithClient(
		u.Host, basepath, []string{u.Scheme}, c.Client,
	))

	r := AddTypeConsumers(runtimeclient.NewWithClient(
		u.Host, RegionlessPrefix, []string{u.Scheme}, c.Client,
	))

	return &CloudClientRuntime{
		regionRuntime: rr,
		runtime:       r,
	}, nil
}

// CloudClientRuntime wraps runtimeclient.Runtime to allow operations to use a
// transport depending on the operation which is being performed.
type CloudClientRuntime struct {
	regionRuntime *runtimeclient.Runtime
	runtime       *runtimeclient.Runtime
}

// Submit calls either the regionRuntime or the regionless runtime depending on
// which operation is being performed. Any API call to /deployments will use a
// regionless runtime while all others will use a region (if specified).
func (r *CloudClientRuntime) Submit(op *runtime.ClientOperation) (interface{}, error) {
	rTime := r.getRuntime(op)

	defer overrideJSONProducer(rTime, op.ID)()
	return rTime.Submit(op)
}

func (r *CloudClientRuntime) getRuntime(op *runtime.ClientOperation) *runtimeclient.Runtime {
	var isDeployment = strings.HasPrefix(op.PathPattern, "/deployments")
	var notDeploymentNotes = !strings.Contains(op.PathPattern, "/note")
	if isDeployment && notDeploymentNotes {
		return r.runtime
	}
	return r.regionRuntime
}

// overrideJSONProducer will override the default JSON producer fucntion for
// a Text producer which won't to serialize the data to JSON, and just send
// the body as is over the wire. This is useful in cases where a JSON body is
// being sent as a Go string value, not doing this will cause the payload json
// quotes to be escaped. See unit tests for examples.
// It returns a function which can be used as a callback to reset the producer
// to its original value.
func overrideJSONProducer(r *runtimeclient.Runtime, opID string) func() {
	if opID != rawMetadataTextProducer {
		return func() {}
	}

	r.Producers[runtime.JSONMime] = runtime.TextProducer()
	return func() { r.Producers[runtime.JSONMime] = runtime.JSONProducer() }
}
