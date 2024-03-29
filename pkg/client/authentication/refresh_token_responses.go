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

// RefreshTokenReader is a Reader for the RefreshToken structure.
type RefreshTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRefreshTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewRefreshTokenNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewRefreshTokenBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshTokenOK creates a RefreshTokenOK with default headers values
func NewRefreshTokenOK() *RefreshTokenOK {
	return &RefreshTokenOK{}
}

/*
RefreshTokenOK describes a response with status code 200, with default header values.

The token refreshed successfully and was returned in the body of the response.
*/
type RefreshTokenOK struct {
	Payload *models.TokenResponse
}

// IsSuccess returns true when this refresh token o k response has a 2xx status code
func (o *RefreshTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh token o k response has a 3xx status code
func (o *RefreshTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh token o k response has a 4xx status code
func (o *RefreshTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh token o k response has a 5xx status code
func (o *RefreshTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh token o k response a status code equal to that given
func (o *RefreshTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the refresh token o k response
func (o *RefreshTokenOK) Code() int {
	return 200
}

func (o *RefreshTokenOK) Error() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenOK  %+v", 200, o.Payload)
}

func (o *RefreshTokenOK) String() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenOK  %+v", 200, o.Payload)
}

func (o *RefreshTokenOK) GetPayload() *models.TokenResponse {
	return o.Payload
}

func (o *RefreshTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshTokenUnauthorized creates a RefreshTokenUnauthorized with default headers values
func NewRefreshTokenUnauthorized() *RefreshTokenUnauthorized {
	return &RefreshTokenUnauthorized{}
}

/*
RefreshTokenUnauthorized describes a response with status code 401, with default header values.

The authentication token is invalid or expired. (code: `root.unauthorized`)
*/
type RefreshTokenUnauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this refresh token unauthorized response has a 2xx status code
func (o *RefreshTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh token unauthorized response has a 3xx status code
func (o *RefreshTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh token unauthorized response has a 4xx status code
func (o *RefreshTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh token unauthorized response has a 5xx status code
func (o *RefreshTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh token unauthorized response a status code equal to that given
func (o *RefreshTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the refresh token unauthorized response
func (o *RefreshTokenUnauthorized) Code() int {
	return 401
}

func (o *RefreshTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshTokenUnauthorized) String() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshTokenUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RefreshTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshTokenNotImplemented creates a RefreshTokenNotImplemented with default headers values
func NewRefreshTokenNotImplemented() *RefreshTokenNotImplemented {
	return &RefreshTokenNotImplemented{}
}

/*
RefreshTokenNotImplemented describes a response with status code 501, with default header values.

The administrator needs to configure the authentication cluster. (code: `authc.no_authentication_cluster`)
*/
type RefreshTokenNotImplemented struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this refresh token not implemented response has a 2xx status code
func (o *RefreshTokenNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh token not implemented response has a 3xx status code
func (o *RefreshTokenNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh token not implemented response has a 4xx status code
func (o *RefreshTokenNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh token not implemented response has a 5xx status code
func (o *RefreshTokenNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh token not implemented response a status code equal to that given
func (o *RefreshTokenNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the refresh token not implemented response
func (o *RefreshTokenNotImplemented) Code() int {
	return 501
}

func (o *RefreshTokenNotImplemented) Error() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenNotImplemented  %+v", 501, o.Payload)
}

func (o *RefreshTokenNotImplemented) String() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenNotImplemented  %+v", 501, o.Payload)
}

func (o *RefreshTokenNotImplemented) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RefreshTokenNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshTokenBadGateway creates a RefreshTokenBadGateway with default headers values
func NewRefreshTokenBadGateway() *RefreshTokenBadGateway {
	return &RefreshTokenBadGateway{}
}

/*
RefreshTokenBadGateway describes a response with status code 502, with default header values.

The authentication cluster failed to process the request. The response body contains details about the error. (code: `authc.authentication_cluster_error`)
*/
type RefreshTokenBadGateway struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this refresh token bad gateway response has a 2xx status code
func (o *RefreshTokenBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh token bad gateway response has a 3xx status code
func (o *RefreshTokenBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh token bad gateway response has a 4xx status code
func (o *RefreshTokenBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh token bad gateway response has a 5xx status code
func (o *RefreshTokenBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh token bad gateway response a status code equal to that given
func (o *RefreshTokenBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the refresh token bad gateway response
func (o *RefreshTokenBadGateway) Code() int {
	return 502
}

func (o *RefreshTokenBadGateway) Error() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenBadGateway  %+v", 502, o.Payload)
}

func (o *RefreshTokenBadGateway) String() string {
	return fmt.Sprintf("[POST /users/auth/_refresh][%d] refreshTokenBadGateway  %+v", 502, o.Payload)
}

func (o *RefreshTokenBadGateway) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RefreshTokenBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
