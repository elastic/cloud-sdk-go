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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteUserRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserOK creates a DeleteUserOK with default headers values
func NewDeleteUserOK() *DeleteUserOK {
	return &DeleteUserOK{}
}

/*DeleteUserOK handles this case with default header values.

User successfully deleted
*/
type DeleteUserOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_name}][%d] deleteUserOK  %+v", 200, o.Payload)
}

func (o *DeleteUserOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserBadRequest creates a DeleteUserBadRequest with default headers values
func NewDeleteUserBadRequest() *DeleteUserBadRequest {
	return &DeleteUserBadRequest{}
}

/*DeleteUserBadRequest handles this case with default header values.

* The user cannot be deleted. (code: `user.restricted_deletion`)
* External users cannot be modified. (code: `user.cannot_modify_external`)
* Built-in users cannot be modified. (code: `user.cannot_modify`)
 */
type DeleteUserBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteUserBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_name}][%d] deleteUserBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*DeleteUserNotFound handles this case with default header values.

User not found. (code: `user.not_found`)
*/
type DeleteUserNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_name}][%d] deleteUserNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRetryWith creates a DeleteUserRetryWith with default headers values
func NewDeleteUserRetryWith() *DeleteUserRetryWith {
	return &DeleteUserRetryWith{}
}

/*DeleteUserRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteUserRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteUserRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_name}][%d] deleteUserRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteUserRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteUserRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
