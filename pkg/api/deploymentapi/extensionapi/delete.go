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

package extensionapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/extensions"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// DeleteParams is consumed by the Delete function.
type DeleteParams struct {
	*api.API

	ExtensionID string
}

// Validate ensures the parameters are usable by Delete.
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid extension delete params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.ExtensionID == "" {
		merr = merr.Append(errors.New("an extension ID is required for this operation"))
	}

	return merr.ErrorOrNil()
}

// Delete deletes an extension.
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	_, err := params.V1API.Extensions.DeleteExtension(
		extensions.NewDeleteExtensionParams().
			WithExtensionID(params.ExtensionID),
		params.AuthWriter,
	)

	return apierror.Wrap(err)
}
