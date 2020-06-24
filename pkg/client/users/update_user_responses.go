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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateUserRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserOK creates a UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

/*UpdateUserOK handles this case with default header values.

User successfully updated
*/
type UpdateUserOK struct {
	Payload *models.User
}

func (o *UpdateUserOK) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_name}][%d] updateUserOK  %+v", 200, o.Payload)
}

func (o *UpdateUserOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UpdateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/*UpdateUserBadRequest handles this case with default header values.

Invalid request
*/
type UpdateUserBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_name}][%d] updateUserBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/*UpdateUserForbidden handles this case with default header values.

Invalid permissions
*/
type UpdateUserForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_name}][%d] updateUserForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserNotFound creates a UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

/*UpdateUserNotFound handles this case with default header values.

User not found
*/
type UpdateUserNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_name}][%d] updateUserNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserRetryWith creates a UpdateUserRetryWith with default headers values
func NewUpdateUserRetryWith() *UpdateUserRetryWith {
	return &UpdateUserRetryWith{}
}

/*UpdateUserRetryWith handles this case with default header values.

Elevated permissions are required. (code: 'root.unauthorized.rbac.elevated_permissions_required')
*/
type UpdateUserRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateUserRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /users/{user_name}][%d] updateUserRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateUserRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateUserRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
