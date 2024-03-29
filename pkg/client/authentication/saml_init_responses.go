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

// SamlInitReader is a Reader for the SamlInit structure.
type SamlInitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SamlInitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewSamlInitFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewSamlInitNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewSamlInitBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSamlInitFound creates a SamlInitFound with default headers values
func NewSamlInitFound() *SamlInitFound {
	return &SamlInitFound{}
}

/*
SamlInitFound describes a response with status code 302, with default header values.

Redirects the client to the identity provider with a SAML authentication request
*/
type SamlInitFound struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this saml init found response has a 2xx status code
func (o *SamlInitFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this saml init found response has a 3xx status code
func (o *SamlInitFound) IsRedirect() bool {
	return true
}

// IsClientError returns true when this saml init found response has a 4xx status code
func (o *SamlInitFound) IsClientError() bool {
	return false
}

// IsServerError returns true when this saml init found response has a 5xx status code
func (o *SamlInitFound) IsServerError() bool {
	return false
}

// IsCode returns true when this saml init found response a status code equal to that given
func (o *SamlInitFound) IsCode(code int) bool {
	return code == 302
}

// Code gets the status code for the saml init found response
func (o *SamlInitFound) Code() int {
	return 302
}

func (o *SamlInitFound) Error() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitFound  %+v", 302, o.Payload)
}

func (o *SamlInitFound) String() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitFound  %+v", 302, o.Payload)
}

func (o *SamlInitFound) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *SamlInitFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSamlInitNotImplemented creates a SamlInitNotImplemented with default headers values
func NewSamlInitNotImplemented() *SamlInitNotImplemented {
	return &SamlInitNotImplemented{}
}

/*
SamlInitNotImplemented describes a response with status code 501, with default header values.

The administrator needs to configure the authentication cluster. (code: `authc.no_authentication_cluster`)
*/
type SamlInitNotImplemented struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this saml init not implemented response has a 2xx status code
func (o *SamlInitNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this saml init not implemented response has a 3xx status code
func (o *SamlInitNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this saml init not implemented response has a 4xx status code
func (o *SamlInitNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this saml init not implemented response has a 5xx status code
func (o *SamlInitNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this saml init not implemented response a status code equal to that given
func (o *SamlInitNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the saml init not implemented response
func (o *SamlInitNotImplemented) Code() int {
	return 501
}

func (o *SamlInitNotImplemented) Error() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitNotImplemented  %+v", 501, o.Payload)
}

func (o *SamlInitNotImplemented) String() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitNotImplemented  %+v", 501, o.Payload)
}

func (o *SamlInitNotImplemented) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SamlInitNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSamlInitBadGateway creates a SamlInitBadGateway with default headers values
func NewSamlInitBadGateway() *SamlInitBadGateway {
	return &SamlInitBadGateway{}
}

/*
SamlInitBadGateway describes a response with status code 502, with default header values.

The authentication cluster failed to process the request. The response body contains details about the error. (code: `authc.authentication_cluster_error`)
*/
type SamlInitBadGateway struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this saml init bad gateway response has a 2xx status code
func (o *SamlInitBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this saml init bad gateway response has a 3xx status code
func (o *SamlInitBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this saml init bad gateway response has a 4xx status code
func (o *SamlInitBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this saml init bad gateway response has a 5xx status code
func (o *SamlInitBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this saml init bad gateway response a status code equal to that given
func (o *SamlInitBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the saml init bad gateway response
func (o *SamlInitBadGateway) Code() int {
	return 502
}

func (o *SamlInitBadGateway) Error() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitBadGateway  %+v", 502, o.Payload)
}

func (o *SamlInitBadGateway) String() string {
	return fmt.Sprintf("[GET /users/auth/saml/_init][%d] samlInitBadGateway  %+v", 502, o.Payload)
}

func (o *SamlInitBadGateway) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SamlInitBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
