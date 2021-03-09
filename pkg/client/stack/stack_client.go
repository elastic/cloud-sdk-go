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

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new stack API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for stack API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteVersionStack(params *DeleteVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVersionStackOK, error)

	GetInstanceTypes(params *GetInstanceTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceTypesOK, error)

	GetVersionStack(params *GetVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionStackOK, error)

	GetVersionStacks(params *GetVersionStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionStacksOK, error)

	UpdateStackPacks(params *UpdateStackPacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStackPacksOK, error)

	UpdateVersionStack(params *UpdateVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVersionStackOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteVersionStack deletes stack version

  Uses the `deleted` flag, which removes only that version of the Elastic Stack from the list of available versions. To restore the version, send an update request. For more information, see the PUT request.
*/
func (a *Client) DeleteVersionStack(params *DeleteVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVersionStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-version-stack",
		Method:             "DELETE",
		PathPattern:        "/stack/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteVersionStackReader{formats: a.formats},
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
	success, ok := result.(*DeleteVersionStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-version-stack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetInstanceTypes gets instance types

  Retrieves a list of all instance types.
*/
func (a *Client) GetInstanceTypes(params *GetInstanceTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInstanceTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInstanceTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-instance-types",
		Method:             "GET",
		PathPattern:        "/stack/instance-types",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetInstanceTypesReader{formats: a.formats},
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
	success, ok := result.(*GetInstanceTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-instance-types: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVersionStack gets stack version

  Retrieves the Elastic Stack version and template.
*/
func (a *Client) GetVersionStack(params *GetVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-version-stack",
		Method:             "GET",
		PathPattern:        "/stack/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionStackReader{formats: a.formats},
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
	success, ok := result.(*GetVersionStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-version-stack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVersionStacks gets stack versions

  By default, retrieves only the available Elastic Stack versions. To retrieve all of the Elastic Stack versions, use the `show_deleted parameter`.
*/
func (a *Client) GetVersionStacks(params *GetVersionStacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVersionStacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionStacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-version-stacks",
		Method:             "GET",
		PathPattern:        "/stack/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetVersionStacksReader{formats: a.formats},
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
	success, ok := result.(*GetVersionStacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-version-stacks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateStackPacks uploads stack pack

  Creates or updates an Elastic Stack pack and template.
The endpoint supports `multipart/form-data` requests, as well as `application/zip` and `application/octet-stream` requests with a binary body. The maximum size of the payload is 1Mb.
When the archive contains an Elastic Stack configuration that is available through the API, the configuration and template are overwritten.
*/
func (a *Client) UpdateStackPacks(params *UpdateStackPacksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateStackPacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStackPacksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-stack-packs",
		Method:             "POST",
		PathPattern:        "/stack/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateStackPacksReader{formats: a.formats},
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
	success, ok := result.(*UpdateStackPacksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-stack-packs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVersionStack updates stack version

  Updates the Elastic Stack version configuration.
*/
func (a *Client) UpdateVersionStack(params *UpdateVersionStackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVersionStackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionStackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-version-stack",
		Method:             "PUT",
		PathPattern:        "/stack/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateVersionStackReader{formats: a.formats},
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
	success, ok := result.(*UpdateVersionStackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-version-stack: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
