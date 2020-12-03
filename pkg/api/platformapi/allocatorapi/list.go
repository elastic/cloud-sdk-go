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

package allocatorapi

import (
	"context"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// ListParams is used to list allocators
type ListParams struct {
	// Required API instance.
	*api.API

	// Required region on which to perform the allocator list.
	Region string

	// Optional Elasticsearch search query.
	Query string

	// Optional filter tags with expected format: key:value slice. i.e.
	// [key:val, key:value].
	FilterTags string

	// Optional toggle to show all allocators (By default connected:false) are
	// omitted.
	ShowAll bool

	// Optional number of allocators to return (Defaults to 100).
	Size int64
}

// Validate ensures that the parameters are correct
func (params ListParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator list params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// List obtains the full list of allocators
func List(params ListParams) (*models.AllocatorOverview, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	p := platform_infrastructure.NewGetAllocatorsParams().
		WithContext(api.WithRegion(context.Background(), params.Region)).
		WithQ(ec.String(params.Query))

	if params.Size > 0 {
		p.SetSize(&params.Size)
	}

	res, err := params.API.V1API.PlatformInfrastructure.GetAllocators(
		p, params.AuthWriter,
	)

	if err != nil {
		return nil, apierror.Wrap(err)
	}
	if !params.ShowAll {
		for _, z := range res.Payload.Zones {
			z.Allocators = FilterConnectedOrWithInstances(z.Allocators)
		}
	}

	for _, z := range res.Payload.Zones {
		z.Allocators = FilterByTag(tagsToMap(params.FilterTags), z.Allocators)
	}

	return res.Payload, nil
}
