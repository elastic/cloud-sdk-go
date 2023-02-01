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
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-openapi/runtime"
	runtimeclient "github.com/go-openapi/runtime/client"

	"github.com/elastic/cloud-sdk-go/pkg/client"
)

const (
	// RegionBasePath is for /platform operations which require a Region to be
	// passed as of an API call context.Context. Previously used to create a
	// global client per API instance, now used on a per-operation basis.
	RegionBasePath = "/api/v1/regions/%s"

	rawMetadataDeploymentResourceTextProducer = "set-deployment-resource-raw-metadata"
	rawMetadataTextProducer                   = "set-es-cluster-metadata-raw"
	updateUserTextProducer                    = "update-user"
	updateCurrentUserTextProducer             = "update-current-user"
)

// DefaultBasePath is used as the base prefix for the API.
var DefaultBasePath = client.DefaultBasePath

// globalPath contains a mapping of the path prefixes which need/ don't need a
// region path interpolation, to work properly. If set to false, then the API
// for that path will require a custom context.Context containing the region
// value != "". Any of the prefixes set to `true` are global paths due to the
// region value already being integrated into the auto-generated parameters.
// Strictly, there's no need for a `"key":false` to be present in the map, but
// it does make it explicit and nicer to maintain.
var globalPath = map[string]bool{
	"clusters":    false,
	"comments":    false,
	"deployments": true,
	"phone-home":  true,
	"platform":    false,
	"stack":       false,
	"user":        true,
	"users":       true,
	"billing":     true,
}

type newRuntimeFunc func(region string) *runtimeclient.Runtime

// NewCloudClientRuntime creates a CloudClientRuntime from the config. Using
// the configured region (if any) to instantiate two different client.Runtime.
// If there's no region specified in the config then both are regionless.
func NewCloudClientRuntime(c Config) (*CloudClientRuntime, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, err
	}

	scheme := []string{u.Scheme}

	return &CloudClientRuntime{
		newRegionRuntime: func(r string) *runtimeclient.Runtime {
			return AddTypeConsumers(runtimeclient.NewWithClient(
				u.Host, fmt.Sprintf(RegionBasePath, r), scheme, c.Client,
			))
		},
		runtime: AddTypeConsumers(runtimeclient.NewWithClient(
			u.Host, DefaultBasePath, scheme, c.Client,
		)),
	}, nil
}

// CloudClientRuntime wraps runtimeclient.Runtime to allow operations to use a
// transport depending on the operation which is being performed.
type CloudClientRuntime struct {
	newRegionRuntime newRuntimeFunc
	runtime          *runtimeclient.Runtime
}

// Submit calls either the regionRuntime or the regionless runtime depending on
// which operation is being performed. Any API call to /deployments will use a
// regionless runtime while all others will use a region (if specified).
func (r *CloudClientRuntime) Submit(op *runtime.ClientOperation) (interface{}, error) {
	rTime, err := r.getRuntime(op)
	if err != nil {
		return nil, err
	}

	defer overrideJSONProducer(rTime, op.ID)()

	return rTime.Submit(op)
}

func (r *CloudClientRuntime) getRuntime(op *runtime.ClientOperation) (*runtimeclient.Runtime, error) {
	var notDeploymentNotes = !strings.Contains(op.PathPattern, "/note")
	regionless := globalPath[strings.Split(op.PathPattern, "/")[1]]
	if regionless && notDeploymentNotes {
		return r.runtime, nil
	}

	region, err := getRegion(op.Context)
	if err != nil {
		return nil, err
	}

	return r.newRegionRuntime(region), nil
}

// overrideJSONProducer will override the default JSON producer function for
// a Text producer which won't to serialize the data to JSON, and just send
// the body as is over the wire. This is useful in cases where a JSON body is
// being sent as a Go string value, not doing this will cause the payload json
// quotes to be escaped. See unit tests for examples.
// It returns a function which can be used as a callback to reset the producer
// to its original value.
func overrideJSONProducer(r *runtimeclient.Runtime, opID string) func() {
	if !(opID == updateUserTextProducer ||
		opID == rawMetadataTextProducer ||
		opID == updateCurrentUserTextProducer ||
		opID == rawMetadataDeploymentResourceTextProducer) {
		return func() {}
	}

	r.Producers[runtime.JSONMime] = runtime.TextProducer()
	return func() { r.Producers[runtime.JSONMime] = runtime.JSONProducer() }
}

func getRegion(ctx context.Context) (string, error) {
	if region, ok := GetContextRegion(ctx); ok {
		return region, nil
	}

	return "", errors.New(
		"the requested operation requires a region but none has been set",
	)
}
