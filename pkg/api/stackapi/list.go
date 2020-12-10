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
	"sort"

	"github.com/blang/semver/v4"
	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/stack"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// ListParams is consumed by List
type ListParams struct {
	*api.API
	Region  string
	Deleted bool
}

// Validate ensures that the parameters are usable by the consuming
// function
func (params ListParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid stack list params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// List lists all stackpacks in the current installation
func List(params ListParams) (*models.StackVersionConfigs, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.Stack.GetVersionStacks(
		stack.NewGetVersionStacksParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithShowDeleted(ec.Bool(params.Deleted)),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	stacks := res.Payload.Stacks
	sort.Slice(stacks, func(i, j int) bool {
		return compareVersions(stacks[i].Version, stacks[j].Version)
	})

	return res.Payload, nil
}

func compareVersions(version1, version2 string) bool {
	a, err := semver.ParseTolerant(version1)
	if err != nil {
		return false
	}

	b, err := semver.ParseTolerant(version2)
	if err != nil {
		return false
	}

	if a.GT(b) {
		return true
	}

	return false
}
