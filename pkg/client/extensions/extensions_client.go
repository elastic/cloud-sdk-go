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

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new extensions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for extensions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateExtension(params *CreateExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExtensionCreated, error)

	DeleteExtension(params *DeleteExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExtensionOK, error)

	GetExtension(params *GetExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionOK, error)

	ListExtensions(params *ListExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListExtensionsOK, error)

	UpdateExtension(params *UpdateExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExtensionOK, error)

	UploadExtension(params *UploadExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadExtensionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateExtension creates an extension

  Creates the extension.
 The extensions API supports two types of usage patterns.
 A). Specify a `download_url`, http or https URL,  where the extension is currently hosted. This will result in extension being copied to elastic repository.
 B). Create only the extension metadata using the `POST` endpoint and then use `PUT` to upload the extension file. Leave the `download_url` unspecified in this case.

*/
func (a *Client) CreateExtension(params *CreateExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateExtensionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-extension",
		Method:             "POST",
		PathPattern:        "/deployments/extensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateExtensionReader{formats: a.formats},
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
	success, ok := result.(*CreateExtensionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-extension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteExtension deletes extension

  Deletes a Extension.
*/
func (a *Client) DeleteExtension(params *DeleteExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-extension",
		Method:             "DELETE",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteExtensionReader{formats: a.formats},
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
	success, ok := result.(*DeleteExtensionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-extension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetExtension gets extension

  Retrieves an extension.
*/
func (a *Client) GetExtension(params *GetExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-extension",
		Method:             "GET",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetExtensionReader{formats: a.formats},
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
	success, ok := result.(*GetExtensionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-extension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListExtensions lists extensions

  Retrieves all of the available extensions.
*/
func (a *Client) ListExtensions(params *ListExtensionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExtensionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list-extensions",
		Method:             "GET",
		PathPattern:        "/deployments/extensions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListExtensionsReader{formats: a.formats},
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
	success, ok := result.(*ListExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-extensions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateExtension updates extension

  Updates an extension.
*/
func (a *Client) UpdateExtension(params *UpdateExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "update-extension",
		Method:             "POST",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateExtensionReader{formats: a.formats},
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
	success, ok := result.(*UpdateExtensionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-extension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadExtension uploads the extension

  Uploads archive for an extension.
*/
func (a *Client) UploadExtension(params *UploadExtensionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UploadExtensionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadExtensionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "upload-extension",
		Method:             "PUT",
		PathPattern:        "/deployments/extensions/{extension_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadExtensionReader{formats: a.formats},
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
	success, ok := result.(*UploadExtensionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for upload-extension: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
