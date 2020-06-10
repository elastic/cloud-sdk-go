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
	"errors"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// MetadataSetParams is is used to set a single allocator metadata key
type MetadataSetParams struct {
	*api.API
	ID, Key, Value string
	Region         string
}

// Validate ensures that the parameters are correct
func (params MetadataSetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator metadata set params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.ID == "" {
		merr = merr.Append(errors.New("id cannot be empty"))
	}
	if params.Key == "" {
		merr = merr.Append(errors.New("key cannot be empty"))
	}
	if params.Value == "" {
		merr = merr.Append(errors.New("key value cannot be empty"))
	}
	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// SetAllocatorMetadataItem sets a single metadata item to a given allocators metadata
func SetAllocatorMetadataItem(params MetadataSetParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.SetAllocatorMetadataItem(
			platform_infrastructure.NewSetAllocatorMetadataItemParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithAllocatorID(params.ID).
				WithKey(params.Key).
				WithBody(&models.MetadataItemValue{Value: &params.Value}),
			params.AuthWriter,
		),
	)
}

// MetadataDeleteParams is used to delete a single metadata key
type MetadataDeleteParams struct {
	*api.API
	ID, Key string
	Region  string
}

// Validate ensures that the parameters are correct
func (params MetadataDeleteParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator metadata delete params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.ID == "" {
		merr = merr.Append(errors.New("id cannot be empty"))
	}
	if params.Key == "" {
		merr = merr.Append(errors.New("key cannot be empty"))
	}
	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}
	return merr.ErrorOrNil()
}

// DeleteAllocatorMetadataItem delete a single metadata item to a given allocators metadata
func DeleteAllocatorMetadataItem(params MetadataDeleteParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	return api.ReturnErrOnly(
		params.API.V1API.PlatformInfrastructure.DeleteAllocatorMetadataItem(
			platform_infrastructure.NewDeleteAllocatorMetadataItemParams().
				WithContext(api.WithRegion(context.Background(), params.Region)).
				WithAllocatorID(params.ID).
				WithKey(params.Key),
			params.AuthWriter,
		),
	)
}

// MetadataGetParams is used to retrieve allocator metadata
type MetadataGetParams struct {
	*api.API
	ID     string
	Region string
}

// Validate ensures that the parameters are correct
func (params MetadataGetParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid allocator metadata get params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}
	if params.ID == "" {
		merr = merr.Append(errors.New("id cannot be empty"))
	}
	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}
	return merr.ErrorOrNil()
}

// GetAllocatorMetadata Retrieves the metadata for a given allocator
func GetAllocatorMetadata(params MetadataGetParams) ([]*models.MetadataItem, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	res, err := params.API.V1API.PlatformInfrastructure.GetAllocatorMetadata(
		platform_infrastructure.NewGetAllocatorMetadataParams().
			WithContext(api.WithRegion(context.Background(), params.Region)).
			WithAllocatorID(params.ID),
		params.AuthWriter,
	)
	if err != nil {
		return nil, api.UnwrapError(err)
	}

	return res.Payload, nil
}
