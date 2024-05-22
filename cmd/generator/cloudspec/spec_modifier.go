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

package cloudspec

import "github.com/go-openapi/spec"

const (
	boolType     = "boolean"
	nullableKey  = "x-nullable"
	omitEmptyKey = "x-omitempty"
)

// Modify iterates over the received spec and adds some extensions to specific
// fields that need to either be set to nullable (Pointer of typee) or be to be
// omitted when they're empty.
// nolint
func Modify(cloudSpec *spec.Swagger) {
	// Iterate over the spec's definitions.
	for k, def := range cloudSpec.Definitions {
		// This hack is required to remove all required fields from `ProxiesSettings` and allow clients
		// to use this model as structured json on Patch API request body.
		if k == "ProxiesSettings" {
			schema := cloudSpec.Definitions[k]
			schema.WithRequired([]string{}...)
			cloudSpec.Definitions[k] = schema
		}

		for kk, prop := range def.Properties {
			if prop.Type.Contains(boolType) {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}
			// Sets the deleted_on fields as a nullable since it should be.
			if kk == "deleted_on" && prop.Format == "date-time" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "DeploymentTemplateInfo" && kk == "instance_configurations" {
				prop.AddExtension(omitEmptyKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "DeploymentTemplateInfo" && kk == "order" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "DeploymentTemplateInfoV2" && kk == "order" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}
			if k == "DeploymentTemplateRequestBody" && kk == "order" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "ApmTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "AppSearchTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "ElasticsearchClusterTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "EnterpriseSearchTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "IntegrationsServerTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			if k == "KibanaClusterTopologyElement" && kk == "instance_configuration_version" {
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}

			addExtension := func(model string, field string, key string, value interface{}) {
				if k == model && kk == field {
					prop.AddExtension(key, value)
					cloudSpec.Definitions[k].Properties[kk] = prop
				}
			}
			addExtension("DeploymentGetResponse", "alias", nullableKey, true)
			addExtension("DeploymentSearchResponse", "alias", nullableKey, true)
			addExtension("DeploymentCreateRequest", "alias", nullableKey, true)
			addExtension("DeploymentCreateResponse", "alias", nullableKey, true)
			addExtension("DeploymentUpdateRequest", "alias", nullableKey, true)
			addExtension("DeploymentUpdateResponse", "alias", nullableKey, true)

			if k == "ElasticsearchConfiguration" {
				if kk == "enabled_built_in_plugins" ||
					kk == "user_bundles" ||
					kk == "user_plugins" {
					prop.AddExtension(omitEmptyKey, true)
					cloudSpec.Definitions[k].Properties[kk] = prop
				}
			}

			if k == "InstanceConfiguration" {
				if kk == "node_types" {
					prop.AddExtension(omitEmptyKey, true)
					cloudSpec.Definitions[k].Properties[kk] = prop
				}
			}

			if k == "UserSecurity" {
				if kk == "permissions" ||
					kk == "roles" {
					prop.AddExtension(omitEmptyKey, true)
					cloudSpec.Definitions[k].Properties[kk] = prop
				}
			}

			if k == "BoolQuery" {
				if kk == "filter" ||
					kk == "must" ||
					kk == "must_not" ||
					kk == "should" {
					prop.AddExtension(omitEmptyKey, true)
					cloudSpec.Definitions[k].Properties[kk] = prop
				}
			}

			// This hack will make json marshaling to omit all empty fields
			if k == "ProxiesSettings" {
				prop.AddExtension(omitEmptyKey, true)
				prop.AddExtension(nullableKey, true)
				cloudSpec.Definitions[k].Properties[kk] = prop
			}
		}
	}

	// Iterate over the paths and parameters as well.
	for k, path := range cloudSpec.Paths.Paths {
		// This hack changes the request body of the PUT container endpoint operation from string
		// to structured model of models.Container
		if k == "/platform/infrastructure/proxies/settings" {
			pathItemProps := path.PathItemProps
			pathItemProps.Patch.Parameters[1].ParamProps.Schema.SchemaProps.Type = nil
			pathItemProps.Patch.Parameters[1].ParamProps.Schema.SchemaProps.Ref = spec.MustCreateRef("#/definitions/ProxiesSettings")
		}
		for kk, param := range path.Parameters {
			if param.Type == boolType {
				param.AddExtension(nullableKey, true)
				cloudSpec.Paths.Paths[k].Parameters[kk] = param
			}
		}
	}
}
