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

package stackapi

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/stack"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
	"github.com/go-openapi/runtime"
)

// UploadParams is consumed by Upload
type UploadParams struct {
	*api.API
	Region    string
	StackPack io.Reader
}

// Validate ensures that the parameters are usable by the consuming
// function
func (params UploadParams) Validate() error {
	var merr = multierror.NewPrefixed("stack upload")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.StackPack == nil {
		merr = merr.Append(errors.New("stackpack cannot be empty"))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Upload uploads a stackpack from a location
func Upload(params UploadParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	res, err := params.V1API.Stack.UpdateStackPacks(
		stack.NewUpdateStackPacksParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithFile(runtime.NamedReader("StackPack", params.StackPack)),
		params.AuthWriter,
	)
	if err != nil {
		return api.UnwrapError(err)
	}

	var merr = multierror.NewPrefixed("stack upload")
	for _, e := range res.Payload.Errors {
		for _, ee := range e.Errors.Errors {
			// ECE stack packs seem to have a __MACOSX packed file which is
			// causing the command to return an error. Error thrown is:
			// This version cannot be parsed: [__MACOSX] because:
			// Unknown version string: [__MACOSX]
			if !strings.Contains(*ee.Message, "__MACOSX") {
				merr = merr.Append(fmt.Errorf("%s: %s", *ee.Code, *ee.Message))
			}
		}
	}
	return merr.ErrorOrNil()
}
