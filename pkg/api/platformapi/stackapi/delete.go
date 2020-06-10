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
	"strings"

	"github.com/blang/semver"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/stack"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// DeleteParams is consumed by Delete
type DeleteParams struct {
	*api.API
	Region  string
	Version string
}

// Validate ensures that the parameters are usable by the consuming
// function
func (params DeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("stack delete")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if _, e := semver.Parse(params.Version); e != nil {
		merr = merr.Append(errors.New(strings.ToLower(e.Error())))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// Delete deletes a stackpack
func Delete(params DeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	return api.ReturnErrOnly(
		params.API.V1API.Stack.DeleteVersionStack(
			stack.NewDeleteVersionStackParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithVersion(params.Version),
			params.AuthWriter,
		),
	)
}
