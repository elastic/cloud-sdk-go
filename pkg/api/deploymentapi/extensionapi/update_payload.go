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
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// NewUpdateRequestFromGet populates UpdateParams from models.Extension
func NewUpdateRequestFromGet(res *models.Extension) *UpdateParams {
	if res == nil {
		return nil
	}

	if res.ExtensionType == nil {
		res.ExtensionType = ec.String("")
	}

	if res.Name == nil {
		res.Name = ec.String("")
	}

	if res.Version == nil {
		res.Version = ec.String("")
	}

	req := UpdateParams{
		Description: res.Description,
		DownloadURL: res.DownloadURL,
		Type:        *res.ExtensionType,
		Name:        *res.Name,
		Version:     *res.Version,
	}

	return &req
}
