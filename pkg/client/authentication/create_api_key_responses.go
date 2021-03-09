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

// CreateAPIKeyReader is a Reader for the CreateAPIKey structure.
type CreateAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAPIKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAPIKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateAPIKeyRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateAPIKeyCreated creates a CreateAPIKeyCreated with default headers values
func NewCreateAPIKeyCreated() *CreateAPIKeyCreated {
	return &CreateAPIKeyCreated{}
}

/* CreateAPIKeyCreated describes a response with status code 201, with default header values.

The API key is created and returned in the body of the response.
*/
type CreateAPIKeyCreated struct {
	Payload *models.APIKeyResponse
}

func (o *CreateAPIKeyCreated) Error() string {
	return fmt.Sprintf("[POST /users/auth/keys][%d] createApiKeyCreated  %+v", 201, o.Payload)
}
func (o *CreateAPIKeyCreated) GetPayload() *models.APIKeyResponse {
	return o.Payload
}

func (o *CreateAPIKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIKeyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAPIKeyBadRequest creates a CreateAPIKeyBadRequest with default headers values
func NewCreateAPIKeyBadRequest() *CreateAPIKeyBadRequest {
	return &CreateAPIKeyBadRequest{}
}

/* CreateAPIKeyBadRequest describes a response with status code 400, with default header values.

The request is invalid. Specify a different request, then try again. (code: `api_keys.invalid_input`)
*/
type CreateAPIKeyBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateAPIKeyBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/auth/keys][%d] createApiKeyBadRequest  %+v", 400, o.Payload)
}
func (o *CreateAPIKeyBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateAPIKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateAPIKeyRetryWith creates a CreateAPIKeyRetryWith with default headers values
func NewCreateAPIKeyRetryWith() *CreateAPIKeyRetryWith {
	return &CreateAPIKeyRetryWith{}
}

/* CreateAPIKeyRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateAPIKeyRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateAPIKeyRetryWith) Error() string {
	return fmt.Sprintf("[POST /users/auth/keys][%d] createApiKeyRetryWith  %+v", 449, o.Payload)
}
func (o *CreateAPIKeyRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateAPIKeyRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
