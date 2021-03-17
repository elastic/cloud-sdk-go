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
	"fmt"
	"net/url"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/client/extensions"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// CreateParams is consumed by the Create function.
type CreateParams struct {
	*api.API

	Name        string
	Version     string
	Type        string
	DownloadURL string
	Description string
}

// Validate ensures the parameters are usable by Create.
func (params CreateParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid extension create params")

	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Type == "" {
		merr = merr.Append(errors.New("an extension type is required for this operation"))
	}

	if params.Version == "" {
		merr = merr.Append(errors.New("an extension version is required for this operation"))
	}

	if params.Name == "" {
		merr = merr.Append(errors.New("an extension name is required for this operation"))
	}

	if params.DownloadURL != "" {
		_, err := url.ParseRequestURI(params.DownloadURL)
		if err != nil {
			merr = merr.Append(fmt.Errorf("the provided URL is invalid: %v", err))
		}
	}

	return merr.ErrorOrNil()
}

// Create creates a new extension.
func Create(params CreateParams) (*models.Extension, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	body := models.CreateExtensionRequest{
		Description:   params.Description,
		DownloadURL:   params.DownloadURL,
		ExtensionType: &params.Type,
		Name:          &params.Name,
		Version:       &params.Version,
	}

	res, err := params.V1API.Extensions.CreateExtension(
		extensions.NewCreateExtensionParams().
			WithBody(&body),
		params.AuthWriter,
	)
	if err != nil {
		return nil, apierror.Wrap(err)
	}

	return res.Payload, nil
}
