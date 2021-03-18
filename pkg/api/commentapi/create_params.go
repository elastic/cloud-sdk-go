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

package commentapi

import (
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// CreateParams is consumed by the Create function.
type CreateParams struct {
	*api.API
	ResourceType, ResourceID, Message, Region string
}

// Validate ensures the parameters are usable by Create.
func (params CreateParams) Validate() error {
	errs := multierror.NewPrefixed("invalid comment create params",
		ec.RequireRegionSet(params.Region),
	)

	if params.API == nil {
		errs = errs.Append(apierror.ErrMissingAPI)
	}

	if params.ResourceType == "" {
		errs = errs.Append(errors.New("resource type is required for this operation"))
	}

	if params.ResourceID == "" {
		errs = errs.Append(errors.New("resource id is required for this operation"))
	}

	if params.Message == "" {
		errs = errs.Append(errors.New("message is required for this operation"))
	}
	return errs.ErrorOrNil()
}
