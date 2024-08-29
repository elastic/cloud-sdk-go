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

// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/client/accounts"
	"github.com/elastic/cloud-sdk-go/pkg/client/authentication"
	"github.com/elastic/cloud-sdk-go/pkg/client/billing_costs_analysis"
	"github.com/elastic/cloud-sdk-go/pkg/client/comments"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployment_templates"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments"
	"github.com/elastic/cloud-sdk-go/pkg/client/deployments_traffic_filter"
	"github.com/elastic/cloud-sdk-go/pkg/client/extensions"
	"github.com/elastic/cloud-sdk-go/pkg/client/organizations"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_instances"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_networking"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_security"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_snapshots"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_templates"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_configuration_trust_relationships"
	"github.com/elastic/cloud-sdk-go/pkg/client/platform_infrastructure"
	"github.com/elastic/cloud-sdk-go/pkg/client/stack"
	"github.com/elastic/cloud-sdk-go/pkg/client/telemetry"
	"github.com/elastic/cloud-sdk-go/pkg/client/trusted_environments"
	"github.com/elastic/cloud-sdk-go/pkg/client/user_role_assignments"
	"github.com/elastic/cloud-sdk-go/pkg/client/users"
)

// Default rest HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new rest HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Rest {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new rest HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Rest {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new rest client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Rest {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Rest)
	cli.Transport = transport
	cli.Accounts = accounts.New(transport, formats)
	cli.Authentication = authentication.New(transport, formats)
	cli.BillingCostsAnalysis = billing_costs_analysis.New(transport, formats)
	cli.Comments = comments.New(transport, formats)
	cli.DeploymentTemplates = deployment_templates.New(transport, formats)
	cli.Deployments = deployments.New(transport, formats)
	cli.DeploymentsTrafficFilter = deployments_traffic_filter.New(transport, formats)
	cli.Extensions = extensions.New(transport, formats)
	cli.Organizations = organizations.New(transport, formats)
	cli.Platform = platform.New(transport, formats)
	cli.PlatformConfigurationInstances = platform_configuration_instances.New(transport, formats)
	cli.PlatformConfigurationNetworking = platform_configuration_networking.New(transport, formats)
	cli.PlatformConfigurationSecurity = platform_configuration_security.New(transport, formats)
	cli.PlatformConfigurationSnapshots = platform_configuration_snapshots.New(transport, formats)
	cli.PlatformConfigurationTemplates = platform_configuration_templates.New(transport, formats)
	cli.PlatformConfigurationTrustRelationships = platform_configuration_trust_relationships.New(transport, formats)
	cli.PlatformInfrastructure = platform_infrastructure.New(transport, formats)
	cli.Stack = stack.New(transport, formats)
	cli.Telemetry = telemetry.New(transport, formats)
	cli.TrustedEnvironments = trusted_environments.New(transport, formats)
	cli.UserRoleAssignments = user_role_assignments.New(transport, formats)
	cli.Users = users.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Rest is a client for rest
type Rest struct {
	Accounts accounts.ClientService

	Authentication authentication.ClientService

	BillingCostsAnalysis billing_costs_analysis.ClientService

	Comments comments.ClientService

	DeploymentTemplates deployment_templates.ClientService

	Deployments deployments.ClientService

	DeploymentsTrafficFilter deployments_traffic_filter.ClientService

	Extensions extensions.ClientService

	Organizations organizations.ClientService

	Platform platform.ClientService

	PlatformConfigurationInstances platform_configuration_instances.ClientService

	PlatformConfigurationNetworking platform_configuration_networking.ClientService

	PlatformConfigurationSecurity platform_configuration_security.ClientService

	PlatformConfigurationSnapshots platform_configuration_snapshots.ClientService

	PlatformConfigurationTemplates platform_configuration_templates.ClientService

	PlatformConfigurationTrustRelationships platform_configuration_trust_relationships.ClientService

	PlatformInfrastructure platform_infrastructure.ClientService

	Stack stack.ClientService

	Telemetry telemetry.ClientService

	TrustedEnvironments trusted_environments.ClientService

	UserRoleAssignments user_role_assignments.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Rest) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Accounts.SetTransport(transport)
	c.Authentication.SetTransport(transport)
	c.BillingCostsAnalysis.SetTransport(transport)
	c.Comments.SetTransport(transport)
	c.DeploymentTemplates.SetTransport(transport)
	c.Deployments.SetTransport(transport)
	c.DeploymentsTrafficFilter.SetTransport(transport)
	c.Extensions.SetTransport(transport)
	c.Organizations.SetTransport(transport)
	c.Platform.SetTransport(transport)
	c.PlatformConfigurationInstances.SetTransport(transport)
	c.PlatformConfigurationNetworking.SetTransport(transport)
	c.PlatformConfigurationSecurity.SetTransport(transport)
	c.PlatformConfigurationSnapshots.SetTransport(transport)
	c.PlatformConfigurationTemplates.SetTransport(transport)
	c.PlatformConfigurationTrustRelationships.SetTransport(transport)
	c.PlatformInfrastructure.SetTransport(transport)
	c.Stack.SetTransport(transport)
	c.Telemetry.SetTransport(transport)
	c.TrustedEnvironments.SetTransport(transport)
	c.UserRoleAssignments.SetTransport(transport)
	c.Users.SetTransport(transport)
}
