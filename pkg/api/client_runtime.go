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
	return r.getRuntime(op).Submit(op)
}

func (r *CloudClientRuntime) getRuntime(op *runtime.ClientOperation) *runtimeclient.Runtime {
	if strings.HasPrefix(op.PathPattern, "/deployments") {
		return r.runtime
	}
	return r.regionRuntime
}
