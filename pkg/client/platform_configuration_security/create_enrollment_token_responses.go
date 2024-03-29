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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateEnrollmentTokenReader is a Reader for the CreateEnrollmentToken structure.
type CreateEnrollmentTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEnrollmentTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEnrollmentTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEnrollmentTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateEnrollmentTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEnrollmentTokenOK creates a CreateEnrollmentTokenOK with default headers values
func NewCreateEnrollmentTokenOK() *CreateEnrollmentTokenOK {
	return &CreateEnrollmentTokenOK{}
}

/*
CreateEnrollmentTokenOK describes a response with status code 200, with default header values.

A token has been generated that can be used to start new servers with the requested roles
*/
type CreateEnrollmentTokenOK struct {
	Payload *models.RequestEnrollmentTokenReply
}

// IsSuccess returns true when this create enrollment token o k response has a 2xx status code
func (o *CreateEnrollmentTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create enrollment token o k response has a 3xx status code
func (o *CreateEnrollmentTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create enrollment token o k response has a 4xx status code
func (o *CreateEnrollmentTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create enrollment token o k response has a 5xx status code
func (o *CreateEnrollmentTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create enrollment token o k response a status code equal to that given
func (o *CreateEnrollmentTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create enrollment token o k response
func (o *CreateEnrollmentTokenOK) Code() int {
	return 200
}

func (o *CreateEnrollmentTokenOK) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenOK  %+v", 200, o.Payload)
}

func (o *CreateEnrollmentTokenOK) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenOK  %+v", 200, o.Payload)
}

func (o *CreateEnrollmentTokenOK) GetPayload() *models.RequestEnrollmentTokenReply {
	return o.Payload
}

func (o *CreateEnrollmentTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestEnrollmentTokenReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentTokenBadRequest creates a CreateEnrollmentTokenBadRequest with default headers values
func NewCreateEnrollmentTokenBadRequest() *CreateEnrollmentTokenBadRequest {
	return &CreateEnrollmentTokenBadRequest{}
}

/*
CreateEnrollmentTokenBadRequest describes a response with status code 400, with default header values.

The token request format was invalid, details in the error (code: 'enrollment_tokens.invalid_token_request')
*/
type CreateEnrollmentTokenBadRequest struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create enrollment token bad request response has a 2xx status code
func (o *CreateEnrollmentTokenBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create enrollment token bad request response has a 3xx status code
func (o *CreateEnrollmentTokenBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create enrollment token bad request response has a 4xx status code
func (o *CreateEnrollmentTokenBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create enrollment token bad request response has a 5xx status code
func (o *CreateEnrollmentTokenBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create enrollment token bad request response a status code equal to that given
func (o *CreateEnrollmentTokenBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create enrollment token bad request response
func (o *CreateEnrollmentTokenBadRequest) Code() int {
	return 400
}

func (o *CreateEnrollmentTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEnrollmentTokenBadRequest) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateEnrollmentTokenBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateEnrollmentTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEnrollmentTokenForbidden creates a CreateEnrollmentTokenForbidden with default headers values
func NewCreateEnrollmentTokenForbidden() *CreateEnrollmentTokenForbidden {
	return &CreateEnrollmentTokenForbidden{}
}

/*
CreateEnrollmentTokenForbidden describes a response with status code 403, with default header values.

No signing key is available to generate a token (code: 'enrollment_tokens.signing_key_not_found')
*/
type CreateEnrollmentTokenForbidden struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create enrollment token forbidden response has a 2xx status code
func (o *CreateEnrollmentTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create enrollment token forbidden response has a 3xx status code
func (o *CreateEnrollmentTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create enrollment token forbidden response has a 4xx status code
func (o *CreateEnrollmentTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create enrollment token forbidden response has a 5xx status code
func (o *CreateEnrollmentTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create enrollment token forbidden response a status code equal to that given
func (o *CreateEnrollmentTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create enrollment token forbidden response
func (o *CreateEnrollmentTokenForbidden) Code() int {
	return 403
}

func (o *CreateEnrollmentTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenForbidden  %+v", 403, o.Payload)
}

func (o *CreateEnrollmentTokenForbidden) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/enrollment-tokens][%d] createEnrollmentTokenForbidden  %+v", 403, o.Payload)
}

func (o *CreateEnrollmentTokenForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateEnrollmentTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
