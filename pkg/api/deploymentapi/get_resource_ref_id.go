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

package deploymentapi

import (
	"fmt"

	"github.com/elastic/cloud-sdk-go/pkg/api"
	"github.com/elastic/cloud-sdk-go/pkg/util"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

// PopulateRefIDParams is consumed by PopulateRefID.
type PopulateRefIDParams struct {
	*api.API

	DeploymentID string
	Kind         string

	// RefID string pointer where the value is persisted.
	RefID *string
}

// GetKindRefID obtains a resource kind RefID. If the kind is not supported
// an error is returned.
func GetKindRefID(params GetResourceParams) (string, error) {
	res, err := Get(params.GetParams)
	if err != nil {
		return "", err
	}

	var refID string
	switch params.Kind {
	case util.Apm:
		for _, resource := range res.Resources.Apm {
			refID = *resource.RefID
		}
	case util.Kibana:
		for _, resource := range res.Resources.Kibana {
			refID = *resource.RefID
		}
	case util.Elasticsearch:
		for _, resource := range res.Resources.Elasticsearch {
			refID = *resource.RefID
		}
	case util.Appsearch:
		for _, resource := range res.Resources.Appsearch {
			refID = *resource.RefID
		}
	case util.EnterpriseSearch:
		for _, resource := range res.Resources.EnterpriseSearch {
			refID = *resource.RefID
		}
	}

	if refID == "" {
		return "", fmt.Errorf("deployment get: resource kind %s is not available", params.Kind)
	}

	return refID, nil
}

// PopulateRefID takes in a the necessary parameters to discover a RefID from a
// resource kind, and populates the value in the params.RefID field in the form
// of a *string.
func PopulateRefID(params PopulateRefIDParams) error {
	// In the odd case RefID is nil, we don't want a panic to be thrown.
	if params.RefID == nil {
		params.RefID = ec.String("")
	}

	if *params.RefID != "" {
		return nil
	}

	refID, err := GetKindRefID(GetResourceParams{
		GetParams: GetParams{
			API:          params.API,
			DeploymentID: params.DeploymentID,
		},
		Kind: params.Kind,
	})
	if err != nil {
		return err
	}

	*params.RefID = refID

	return nil
}
