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

package user_role_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user role assignments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user role assignments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddRoleAssignments(params *AddRoleAssignmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddRoleAssignmentsOK, error)

	RemoveRoleAssignments(params *RemoveRoleAssignmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveRoleAssignmentsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddRoleAssignments adds role assignments

Adds a set of RoleAssignments to the specified User. Currently unavailable in self-hosted ECE.
*/
func (a *Client) AddRoleAssignments(params *AddRoleAssignmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddRoleAssignmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleAssignmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "add-role-assignments",
		Method:             "POST",
		PathPattern:        "/users/{user_id}/role_assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleAssignmentsReader{formats: a.formats},
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
	success, ok := result.(*AddRoleAssignmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add-role-assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RemoveRoleAssignments removes role assignments

Removes a set of RoleAssignments from the specified User. Currently unavailable in self-hosted ECE.
*/
func (a *Client) RemoveRoleAssignments(params *RemoveRoleAssignmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RemoveRoleAssignmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveRoleAssignmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "remove-role-assignments",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/role_assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemoveRoleAssignmentsReader{formats: a.formats},
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
	success, ok := result.(*RemoveRoleAssignmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for remove-role-assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}