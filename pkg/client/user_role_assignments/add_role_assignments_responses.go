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
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// AddRoleAssignmentsReader is a Reader for the AddRoleAssignments structure.
type AddRoleAssignmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleAssignmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddRoleAssignmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddRoleAssignmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddRoleAssignmentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddRoleAssignmentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddRoleAssignmentsOK creates a AddRoleAssignmentsOK with default headers values
func NewAddRoleAssignmentsOK() *AddRoleAssignmentsOK {
	return &AddRoleAssignmentsOK{}
}

/*
AddRoleAssignmentsOK describes a response with status code 200, with default header values.

Role Assignments were successfully added to the target User
*/
type AddRoleAssignmentsOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this add role assignments o k response has a 2xx status code
func (o *AddRoleAssignmentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add role assignments o k response has a 3xx status code
func (o *AddRoleAssignmentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role assignments o k response has a 4xx status code
func (o *AddRoleAssignmentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add role assignments o k response has a 5xx status code
func (o *AddRoleAssignmentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add role assignments o k response a status code equal to that given
func (o *AddRoleAssignmentsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add role assignments o k response
func (o *AddRoleAssignmentsOK) Code() int {
	return 200
}

func (o *AddRoleAssignmentsOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsOK  %+v", 200, o.Payload)
}

func (o *AddRoleAssignmentsOK) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsOK  %+v", 200, o.Payload)
}

func (o *AddRoleAssignmentsOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *AddRoleAssignmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRoleAssignmentsBadRequest creates a AddRoleAssignmentsBadRequest with default headers values
func NewAddRoleAssignmentsBadRequest() *AddRoleAssignmentsBadRequest {
	return &AddRoleAssignmentsBadRequest{}
}

/*
AddRoleAssignmentsBadRequest describes a response with status code 400, with default header values.

The target user Id is invalid. (code: `role_assignments.invalid_target_user_id`)
*/
type AddRoleAssignmentsBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this add role assignments bad request response has a 2xx status code
func (o *AddRoleAssignmentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add role assignments bad request response has a 3xx status code
func (o *AddRoleAssignmentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role assignments bad request response has a 4xx status code
func (o *AddRoleAssignmentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add role assignments bad request response has a 5xx status code
func (o *AddRoleAssignmentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add role assignments bad request response a status code equal to that given
func (o *AddRoleAssignmentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add role assignments bad request response
func (o *AddRoleAssignmentsBadRequest) Code() int {
	return 400
}

func (o *AddRoleAssignmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsBadRequest  %+v", 400, o.Payload)
}

func (o *AddRoleAssignmentsBadRequest) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsBadRequest  %+v", 400, o.Payload)
}

func (o *AddRoleAssignmentsBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *AddRoleAssignmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRoleAssignmentsUnauthorized creates a AddRoleAssignmentsUnauthorized with default headers values
func NewAddRoleAssignmentsUnauthorized() *AddRoleAssignmentsUnauthorized {
	return &AddRoleAssignmentsUnauthorized{}
}

/*
AddRoleAssignmentsUnauthorized describes a response with status code 401, with default header values.

Credentials were invalid. (code: `root.unauthorized`)
*/
type AddRoleAssignmentsUnauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this add role assignments unauthorized response has a 2xx status code
func (o *AddRoleAssignmentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add role assignments unauthorized response has a 3xx status code
func (o *AddRoleAssignmentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role assignments unauthorized response has a 4xx status code
func (o *AddRoleAssignmentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add role assignments unauthorized response has a 5xx status code
func (o *AddRoleAssignmentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add role assignments unauthorized response a status code equal to that given
func (o *AddRoleAssignmentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add role assignments unauthorized response
func (o *AddRoleAssignmentsUnauthorized) Code() int {
	return 401
}

func (o *AddRoleAssignmentsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *AddRoleAssignmentsUnauthorized) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsUnauthorized  %+v", 401, o.Payload)
}

func (o *AddRoleAssignmentsUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *AddRoleAssignmentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddRoleAssignmentsForbidden creates a AddRoleAssignmentsForbidden with default headers values
func NewAddRoleAssignmentsForbidden() *AddRoleAssignmentsForbidden {
	return &AddRoleAssignmentsForbidden{}
}

/*
AddRoleAssignmentsForbidden describes a response with status code 403, with default header values.

You are not authorised to add the specified RoleAssignments. (code: `role_assignments.unauthorized_role_assignments`)
*/
type AddRoleAssignmentsForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this add role assignments forbidden response has a 2xx status code
func (o *AddRoleAssignmentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add role assignments forbidden response has a 3xx status code
func (o *AddRoleAssignmentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role assignments forbidden response has a 4xx status code
func (o *AddRoleAssignmentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add role assignments forbidden response has a 5xx status code
func (o *AddRoleAssignmentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add role assignments forbidden response a status code equal to that given
func (o *AddRoleAssignmentsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add role assignments forbidden response
func (o *AddRoleAssignmentsForbidden) Code() int {
	return 403
}

func (o *AddRoleAssignmentsForbidden) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsForbidden  %+v", 403, o.Payload)
}

func (o *AddRoleAssignmentsForbidden) String() string {
	return fmt.Sprintf("[POST /users/{user_id}/role_assignments][%d] addRoleAssignmentsForbidden  %+v", 403, o.Payload)
}

func (o *AddRoleAssignmentsForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *AddRoleAssignmentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}