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

package mock

import (
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

const (
	code500    = "internal.server.error"
	message500 = "There was an internal server error"
)

const (
	code404    = "deployments.deployment_not_found"
	message404 = "No Deployment with id [62bed072dffd4509b84dfe7dc125bb12] could be found"
)

const (
	code400    = "root.invalid_json_request"
	message400 = "JSON request does not comply with schema: [String: DownField(region),DownArray,DownField(elasticsearch),DownField(resources): [String]]"
)

var (
	// MultierrorInternalError represents the multierror returned by apier.Unwrap().
	MultierrorInternalError = multierror.NewPrefixed("api error",
		fmt.Errorf("%s: %s", code500, message500),
	)

	// MultierrorNotFound represents the multierror returned by apier.Unwrap().
	MultierrorNotFound = multierror.NewPrefixed("api error",
		fmt.Errorf("%s: %s", code404, message404),
	)

	// MultierrorBadRequest represents the multierror returned by apier.Unwrap().
	MultierrorBadRequest = multierror.NewPrefixed("api error",
		fmt.Errorf("%s: %s", code400, message400),
	)
)

// SampleInternalError returns a response which encapsulates a 500 error.
func SampleInternalError() Response {
	return NewErrorResponse(500, APIError{Code: code500, Message: message500})
}

// SampleNotFoundError returns a response which encapsulates a 404 error.
func SampleNotFoundError() Response {
	return NewErrorResponse(404, APIError{Code: code404, Message: message404})
}

// SampleBadRequestError returns a response which encapsulates a 400 error.
func SampleBadRequestError() Response {
	return NewErrorResponse(400, APIError{Code: code400, Message: message400})
}
