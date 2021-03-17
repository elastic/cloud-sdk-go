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
	"github.com/blang/semver/v4"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

var dataTiersMinVersion = semver.MustParse("7.10.0")

// PayloadOverrides represent the override settings to
type PayloadOverrides struct {
	// If set, it will override the deployment name.
	Name string

	// If set, it will override the region when not present in the
	// DeploymentCreateRequest.
	// Note this behavior is different from the rest of overrides
	// since this field tends to be populated by the global region
	// field which is implicit (by config) rather than explicit by flag.
	Region string

	// If set, it'll override all versions to match this one.
	Version string

	// ElasticsearchRefID used for the applications.
	ElasticcsearchRefID string

	// OverrideRefIDs when set, it'll override all the application's ref_id.
	OverrideRefIDs bool

	// Slice of built-in enabled plugins.
	ElasticcsearchBuiltinPlugins []string
}

// OverrideCreateOrUpdateRequest sets a series of overrides to either the Create or Update
// deployment request. See PayloadOverrides to understand which how each
// field of that struct affects the behavior of this function.
func OverrideCreateOrUpdateRequest(req interface{}, overrides *PayloadOverrides) error {
	if req == nil || overrides == nil {
		return nil
	}

	var apm []*models.ApmPayload
	var appsearch []*models.AppSearchPayload
	var elasticsearch []*models.ElasticsearchPayload
	var enterprisesearch []*models.EnterpriseSearchPayload
	var kibana []*models.KibanaPayload
	switch t := req.(type) {
	case *models.DeploymentUpdateRequest:
		if t.Resources == nil {
			return nil
		}
		apm, appsearch = t.Resources.Apm, t.Resources.Appsearch
		elasticsearch, kibana = t.Resources.Elasticsearch, t.Resources.Kibana
		enterprisesearch = t.Resources.EnterpriseSearch
	case *models.DeploymentCreateRequest:
		if overrides.Name != "" {
			t.Name = overrides.Name
		}
		if t.Resources == nil {
			return nil
		}
		apm, appsearch = t.Resources.Apm, t.Resources.Appsearch
		elasticsearch, kibana = t.Resources.Elasticsearch, t.Resources.Kibana
		enterprisesearch = t.Resources.EnterpriseSearch
	}

	return overrideByPayload(
		apm, appsearch, elasticsearch, kibana, enterprisesearch,
		overrides.Region, overrides.Version, overrides.ElasticcsearchRefID,
		overrides.ElasticcsearchBuiltinPlugins, overrides.OverrideRefIDs,
	)
}

// nolint
func overrideByPayload(apm []*models.ApmPayload, appsearch []*models.AppSearchPayload,
	elasticsearch []*models.ElasticsearchPayload, kibana []*models.KibanaPayload,
	enterprisesearch []*models.EnterpriseSearchPayload, region, version, refID string,
	plugins []string, overrideRefIDs bool) error {
	for _, resource := range apm {
		if resource.Region == nil && region != "" {
			resource.Region = &region
		}

		if overrideRefIDs {
			resource.RefID = ec.String("main-apm")
		}

		if refID != "" {
			resource.ElasticsearchClusterRefID = &refID
		}

		if version != "" {
			if resource.Plan != nil && resource.Plan.Apm != nil {
				resource.Plan.Apm.Version = version
			}
		}
	}

	for _, resource := range appsearch {
		if resource.Region == nil && region != "" {
			resource.Region = &region
		}

		if overrideRefIDs {
			resource.RefID = ec.String("main-appsearch")
		}

		if refID != "" {
			resource.ElasticsearchClusterRefID = &refID
		}

		if version != "" {
			if resource.Plan != nil && resource.Plan.Appsearch != nil {
				resource.Plan.Appsearch.Version = version
			}
		}
	}

	for _, resource := range elasticsearch {
		if resource.Region == nil && region != "" {
			resource.Region = &region
		}

		if refID != "" {
			resource.RefID = &refID
		}

		// If version is empty, then data tiers are supported by default.
		dataTierCompatible := true
		if version != "" {
			if resource.Plan != nil && resource.Plan.Elasticsearch != nil {
				resource.Plan.Elasticsearch.Version = version
				if len(plugins) > 0 {
					resource.Plan.Elasticsearch.EnabledBuiltInPlugins = append(
						resource.Plan.Elasticsearch.EnabledBuiltInPlugins, plugins...,
					)
				}

				esVersion, err := semver.Parse(version)
				if err != nil {
					return err
				}
				dataTierCompatible = esVersion.GE(dataTiersMinVersion)
			}
		}
		for _, top := range resource.Plan.ClusterTopology {
			if len(top.NodeRoles) > 0 && top.NodeType != nil {
				if dataTierCompatible {
					top.NodeType = nil
					continue
				}
				top.NodeRoles = nil
			}
		}
	}

	for _, resource := range enterprisesearch {
		if resource.Region == nil && region != "" {
			resource.Region = &region
		}

		if overrideRefIDs {
			resource.RefID = ec.String("main-enterprise_search")
		}

		if refID != "" {
			resource.ElasticsearchClusterRefID = &refID
		}

		if version != "" {
			if resource.Plan != nil && resource.Plan.EnterpriseSearch != nil {
				resource.Plan.EnterpriseSearch.Version = version
			}
		}
	}

	for _, resource := range kibana {
		if resource.Region == nil && region != "" {
			resource.Region = &region
		}

		if overrideRefIDs {
			resource.RefID = ec.String("main-kibana")
		}

		if refID != "" {
			resource.ElasticsearchClusterRefID = &refID
		}

		if version != "" {
			if resource.Plan != nil && resource.Plan.Kibana != nil {
				resource.Plan.Kibana.Version = version
			}
		}
	}

	return nil
}
