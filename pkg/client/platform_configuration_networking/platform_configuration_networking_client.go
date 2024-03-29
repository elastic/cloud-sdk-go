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

package platform_configuration_networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform configuration networking API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform configuration networking API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDefaultDeploymentDomainName(params *GetDefaultDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultDeploymentDomainNameOK, error)

	GetResourceKindDeploymentDomainName(params *GetResourceKindDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceKindDeploymentDomainNameOK, error)

	SetDefaultDeploymentDomainName(params *SetDefaultDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetDefaultDeploymentDomainNameOK, error)

	SetResourceKindDeploymentDomainName(params *SetResourceKindDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetResourceKindDeploymentDomainNameOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetDefaultDeploymentDomainName gets default deployment domain name

Gets the default Deployment Domain Name configuration value.
*/
func (a *Client) GetDefaultDeploymentDomainName(params *GetDefaultDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDefaultDeploymentDomainNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultDeploymentDomainNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-default-deployment-domain-name",
		Method:             "GET",
		PathPattern:        "/platform/configuration/networking/deployment_domain_name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultDeploymentDomainNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDefaultDeploymentDomainNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-default-deployment-domain-name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetResourceKindDeploymentDomainName gets deployment domain name for a resource kind

Gets Deployment Domain Name configuration value for a resource kind.
*/
func (a *Client) GetResourceKindDeploymentDomainName(params *GetResourceKindDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetResourceKindDeploymentDomainNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetResourceKindDeploymentDomainNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-resource-kind-deployment-domain-name",
		Method:             "GET",
		PathPattern:        "/platform/configuration/networking/deployment_domain_name/{resource_kind}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResourceKindDeploymentDomainNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetResourceKindDeploymentDomainNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-resource-kind-deployment-domain-name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetDefaultDeploymentDomainName sets default deployment domain name

Saves the default Deployment Domain Name configuration value.
*/
func (a *Client) SetDefaultDeploymentDomainName(params *SetDefaultDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetDefaultDeploymentDomainNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetDefaultDeploymentDomainNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set-default-deployment-domain-name",
		Method:             "PUT",
		PathPattern:        "/platform/configuration/networking/deployment_domain_name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetDefaultDeploymentDomainNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetDefaultDeploymentDomainNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set-default-deployment-domain-name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SetResourceKindDeploymentDomainName sets deployment domain name for a resource kind

Saves the Deployment Domain Name configuration value for a resource kind.
*/
func (a *Client) SetResourceKindDeploymentDomainName(params *SetResourceKindDeploymentDomainNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetResourceKindDeploymentDomainNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetResourceKindDeploymentDomainNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "set-resource-kind-deployment-domain-name",
		Method:             "PUT",
		PathPattern:        "/platform/configuration/networking/deployment_domain_name/{resource_kind}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetResourceKindDeploymentDomainNameReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetResourceKindDeploymentDomainNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for set-resource-kind-deployment-domain-name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
