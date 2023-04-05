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

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ListOrganizationsReader is a Reader for the ListOrganizations structure.
type ListOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListOrganizationsOK creates a ListOrganizationsOK with default headers values
func NewListOrganizationsOK() *ListOrganizationsOK {
	return &ListOrganizationsOK{}
}

/*
ListOrganizationsOK describes a response with status code 200, with default header values.

Organizations fetched successfully
*/
type ListOrganizationsOK struct {
	Payload *models.OrganizationList
}

// IsSuccess returns true when this list organizations o k response has a 2xx status code
func (o *ListOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list organizations o k response has a 3xx status code
func (o *ListOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations o k response has a 4xx status code
func (o *ListOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list organizations o k response has a 5xx status code
func (o *ListOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations o k response a status code equal to that given
func (o *ListOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list organizations o k response
func (o *ListOrganizationsOK) Code() int {
	return 200
}

func (o *ListOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] listOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationsOK) String() string {
	return fmt.Sprintf("[GET /organizations][%d] listOrganizationsOK  %+v", 200, o.Payload)
}

func (o *ListOrganizationsOK) GetPayload() *models.OrganizationList {
	return o.Payload
}

func (o *ListOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrganizationList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOrganizationsUnauthorized creates a ListOrganizationsUnauthorized with default headers values
func NewListOrganizationsUnauthorized() *ListOrganizationsUnauthorized {
	return &ListOrganizationsUnauthorized{}
}

/*
ListOrganizationsUnauthorized describes a response with status code 401, with default header values.

User not found. (code: `user.not_found`)
*/
type ListOrganizationsUnauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this list organizations unauthorized response has a 2xx status code
func (o *ListOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list organizations unauthorized response has a 3xx status code
func (o *ListOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list organizations unauthorized response has a 4xx status code
func (o *ListOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list organizations unauthorized response has a 5xx status code
func (o *ListOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list organizations unauthorized response a status code equal to that given
func (o *ListOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list organizations unauthorized response
func (o *ListOrganizationsUnauthorized) Code() int {
	return 401
}

func (o *ListOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] listOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /organizations][%d] listOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListOrganizationsUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ListOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
