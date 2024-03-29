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

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// LogoutReader is a Reader for the Logout structure.
type LogoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 501:
		result := NewLogoutNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewLogoutBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLogoutOK creates a LogoutOK with default headers values
func NewLogoutOK() *LogoutOK {
	return &LogoutOK{}
}

/*
LogoutOK describes a response with status code 200, with default header values.

The current session was successfully destroyed.
*/
type LogoutOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this logout o k response has a 2xx status code
func (o *LogoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this logout o k response has a 3xx status code
func (o *LogoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout o k response has a 4xx status code
func (o *LogoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this logout o k response has a 5xx status code
func (o *LogoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this logout o k response a status code equal to that given
func (o *LogoutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the logout o k response
func (o *LogoutOK) Code() int {
	return 200
}

func (o *LogoutOK) Error() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutOK  %+v", 200, o.Payload)
}

func (o *LogoutOK) String() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutOK  %+v", 200, o.Payload)
}

func (o *LogoutOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *LogoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogoutNotImplemented creates a LogoutNotImplemented with default headers values
func NewLogoutNotImplemented() *LogoutNotImplemented {
	return &LogoutNotImplemented{}
}

/*
LogoutNotImplemented describes a response with status code 501, with default header values.

The administrator needs to configure the authentication cluster. (code: `authc.no_authentication_cluster`)
*/
type LogoutNotImplemented struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this logout not implemented response has a 2xx status code
func (o *LogoutNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this logout not implemented response has a 3xx status code
func (o *LogoutNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout not implemented response has a 4xx status code
func (o *LogoutNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this logout not implemented response has a 5xx status code
func (o *LogoutNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this logout not implemented response a status code equal to that given
func (o *LogoutNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the logout not implemented response
func (o *LogoutNotImplemented) Code() int {
	return 501
}

func (o *LogoutNotImplemented) Error() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutNotImplemented  %+v", 501, o.Payload)
}

func (o *LogoutNotImplemented) String() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutNotImplemented  %+v", 501, o.Payload)
}

func (o *LogoutNotImplemented) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *LogoutNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLogoutBadGateway creates a LogoutBadGateway with default headers values
func NewLogoutBadGateway() *LogoutBadGateway {
	return &LogoutBadGateway{}
}

/*
LogoutBadGateway describes a response with status code 502, with default header values.

The authentication cluster failed to process the request. The response body contains details about the error. (code: `authc.authentication_cluster_error`)
*/
type LogoutBadGateway struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this logout bad gateway response has a 2xx status code
func (o *LogoutBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this logout bad gateway response has a 3xx status code
func (o *LogoutBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this logout bad gateway response has a 4xx status code
func (o *LogoutBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this logout bad gateway response has a 5xx status code
func (o *LogoutBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this logout bad gateway response a status code equal to that given
func (o *LogoutBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the logout bad gateway response
func (o *LogoutBadGateway) Code() int {
	return 502
}

func (o *LogoutBadGateway) Error() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutBadGateway  %+v", 502, o.Payload)
}

func (o *LogoutBadGateway) String() string {
	return fmt.Sprintf("[POST /users/auth/_logout][%d] logoutBadGateway  %+v", 502, o.Payload)
}

func (o *LogoutBadGateway) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *LogoutBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
