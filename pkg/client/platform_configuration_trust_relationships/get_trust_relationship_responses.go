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

package platform_configuration_trust_relationships

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetTrustRelationshipReader is a Reader for the GetTrustRelationship structure.
type GetTrustRelationshipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTrustRelationshipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTrustRelationshipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetTrustRelationshipUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTrustRelationshipNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTrustRelationshipOK creates a GetTrustRelationshipOK with default headers values
func NewGetTrustRelationshipOK() *GetTrustRelationshipOK {
	return &GetTrustRelationshipOK{}
}

/*
GetTrustRelationshipOK describes a response with status code 200, with default header values.

The trust relationship info response
*/
type GetTrustRelationshipOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.TrustRelationshipGetResponse
}

// IsSuccess returns true when this get trust relationship o k response has a 2xx status code
func (o *GetTrustRelationshipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get trust relationship o k response has a 3xx status code
func (o *GetTrustRelationshipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trust relationship o k response has a 4xx status code
func (o *GetTrustRelationshipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get trust relationship o k response has a 5xx status code
func (o *GetTrustRelationshipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get trust relationship o k response a status code equal to that given
func (o *GetTrustRelationshipOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get trust relationship o k response
func (o *GetTrustRelationshipOK) Code() int {
	return 200
}

func (o *GetTrustRelationshipOK) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipOK  %+v", 200, o.Payload)
}

func (o *GetTrustRelationshipOK) String() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipOK  %+v", 200, o.Payload)
}

func (o *GetTrustRelationshipOK) GetPayload() *models.TrustRelationshipGetResponse {
	return o.Payload
}

func (o *GetTrustRelationshipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.TrustRelationshipGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTrustRelationshipUnauthorized creates a GetTrustRelationshipUnauthorized with default headers values
func NewGetTrustRelationshipUnauthorized() *GetTrustRelationshipUnauthorized {
	return &GetTrustRelationshipUnauthorized{}
}

/*
GetTrustRelationshipUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action.
*/
type GetTrustRelationshipUnauthorized struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get trust relationship unauthorized response has a 2xx status code
func (o *GetTrustRelationshipUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trust relationship unauthorized response has a 3xx status code
func (o *GetTrustRelationshipUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trust relationship unauthorized response has a 4xx status code
func (o *GetTrustRelationshipUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trust relationship unauthorized response has a 5xx status code
func (o *GetTrustRelationshipUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get trust relationship unauthorized response a status code equal to that given
func (o *GetTrustRelationshipUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get trust relationship unauthorized response
func (o *GetTrustRelationshipUnauthorized) Code() int {
	return 401
}

func (o *GetTrustRelationshipUnauthorized) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTrustRelationshipUnauthorized) String() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTrustRelationshipUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetTrustRelationshipUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTrustRelationshipNotFound creates a GetTrustRelationshipNotFound with default headers values
func NewGetTrustRelationshipNotFound() *GetTrustRelationshipNotFound {
	return &GetTrustRelationshipNotFound{}
}

/*
GetTrustRelationshipNotFound describes a response with status code 404, with default header values.

The trust relationship specified by {trust_relationship_id} cannot be found. (code: `trust_relationships.not_found`)
*/
type GetTrustRelationshipNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get trust relationship not found response has a 2xx status code
func (o *GetTrustRelationshipNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get trust relationship not found response has a 3xx status code
func (o *GetTrustRelationshipNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get trust relationship not found response has a 4xx status code
func (o *GetTrustRelationshipNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get trust relationship not found response has a 5xx status code
func (o *GetTrustRelationshipNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get trust relationship not found response a status code equal to that given
func (o *GetTrustRelationshipNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get trust relationship not found response
func (o *GetTrustRelationshipNotFound) Code() int {
	return 404
}

func (o *GetTrustRelationshipNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *GetTrustRelationshipNotFound) String() string {
	return fmt.Sprintf("[GET /platform/configuration/trust-relationships/{trust_relationship_id}][%d] getTrustRelationshipNotFound  %+v", 404, o.Payload)
}

func (o *GetTrustRelationshipNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetTrustRelationshipNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
