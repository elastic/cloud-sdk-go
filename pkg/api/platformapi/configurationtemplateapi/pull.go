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

package configurationtemplateapi

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/api/apierror"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var folderErrorMessage = "folder not specified and is required for the operation"

// PullToFolderParams is the parameter for deployment template pull to folder sub-command
type PullToFolderParams struct {
	*api.API
	Folder             string
	Region             string
	ShowInstanceConfig bool
}

// Validate is the implementation for the ecctl.Validator interface
func (params PullToFolderParams) Validate() error {
	var merr = multierror.NewPrefixed("invalid deployment template pull params")
	if params.API == nil {
		merr = merr.Append(apierror.ErrMissingAPI)
	}

	if params.Folder == "" {
		merr = merr.Append(errors.New(folderErrorMessage))
	}

	if err := ec.RequireRegionSet(params.Region); err != nil {
		merr = merr.Append(err)
	}

	return merr.ErrorOrNil()
}

// PullToFolder downloads deployment templates and save them in a local folder
func PullToFolder(params PullToFolderParams) error {
	if err := params.Validate(); err != nil {
		return err
	}

	res, err := ListTemplates(ListTemplateParams{
		API:                params.API,
		Region:             params.Region,
		ShowInstanceConfig: params.ShowInstanceConfig,
	})
	if err != nil {
		return err
	}

	return writeDeploymentTemplateToFolder(params.Folder, res)
}

// writeDeploymentTemplateToFolder this will write all the deployment template to a folder
// following this structure:
//   folder/
//   folder/id.json
func writeDeploymentTemplateToFolder(folder string, templates []*models.DeploymentTemplateInfo) error {
	p := folder
	if filepath.Ext(p) != "" {
		p = filepath.Dir(folder)
	}

	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		return err
	}

	var merr = multierror.NewPrefixed("failed persisting deployment templates")
	for _, template := range templates {
		template.Source = nil

		f, err := os.Create(filepath.Join(folder, template.ID+".json"))
		if err != nil {
			merr = merr.Append(err)
			continue
		}

		var enc = json.NewEncoder(f)
		enc.SetIndent("", "  ")
		if err := enc.Encode(template); err != nil {
			merr = merr.Append(err)
		}
	}

	return merr.ErrorOrNil()
}
