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

package billing_costs_analysis

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetCostsOverviewReader is a Reader for the GetCostsOverview structure.
type GetCostsOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCostsOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCostsOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCostsOverviewBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCostsOverviewForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCostsOverviewNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCostsOverviewInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCostsOverviewOK creates a GetCostsOverviewOK with default headers values
func NewGetCostsOverviewOK() *GetCostsOverviewOK {
	return &GetCostsOverviewOK{}
}

/*
GetCostsOverviewOK describes a response with status code 200, with default header values.

Top-level cost overview for the organization
*/
type GetCostsOverviewOK struct {
	Payload *models.CostsOverview
}

// IsSuccess returns true when this get costs overview o k response has a 2xx status code
func (o *GetCostsOverviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get costs overview o k response has a 3xx status code
func (o *GetCostsOverviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get costs overview o k response has a 4xx status code
func (o *GetCostsOverviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get costs overview o k response has a 5xx status code
func (o *GetCostsOverviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get costs overview o k response a status code equal to that given
func (o *GetCostsOverviewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get costs overview o k response
func (o *GetCostsOverviewOK) Code() int {
	return 200
}

func (o *GetCostsOverviewOK) Error() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewOK  %+v", 200, o.Payload)
}

func (o *GetCostsOverviewOK) String() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewOK  %+v", 200, o.Payload)
}

func (o *GetCostsOverviewOK) GetPayload() *models.CostsOverview {
	return o.Payload
}

func (o *GetCostsOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CostsOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCostsOverviewBadRequest creates a GetCostsOverviewBadRequest with default headers values
func NewGetCostsOverviewBadRequest() *GetCostsOverviewBadRequest {
	return &GetCostsOverviewBadRequest{}
}

/*
GetCostsOverviewBadRequest describes a response with status code 400, with default header values.

The specified date range is invalid. (code: `costs.invalid_date_range`)
*/
type GetCostsOverviewBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get costs overview bad request response has a 2xx status code
func (o *GetCostsOverviewBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get costs overview bad request response has a 3xx status code
func (o *GetCostsOverviewBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get costs overview bad request response has a 4xx status code
func (o *GetCostsOverviewBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get costs overview bad request response has a 5xx status code
func (o *GetCostsOverviewBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get costs overview bad request response a status code equal to that given
func (o *GetCostsOverviewBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get costs overview bad request response
func (o *GetCostsOverviewBadRequest) Code() int {
	return 400
}

func (o *GetCostsOverviewBadRequest) Error() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetCostsOverviewBadRequest) String() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewBadRequest  %+v", 400, o.Payload)
}

func (o *GetCostsOverviewBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetCostsOverviewBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCostsOverviewForbidden creates a GetCostsOverviewForbidden with default headers values
func NewGetCostsOverviewForbidden() *GetCostsOverviewForbidden {
	return &GetCostsOverviewForbidden{}
}

/*
GetCostsOverviewForbidden describes a response with status code 403, with default header values.

The current user does not have access to the requested organization. (code: `organization.invalid_access`)
*/
type GetCostsOverviewForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get costs overview forbidden response has a 2xx status code
func (o *GetCostsOverviewForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get costs overview forbidden response has a 3xx status code
func (o *GetCostsOverviewForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get costs overview forbidden response has a 4xx status code
func (o *GetCostsOverviewForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get costs overview forbidden response has a 5xx status code
func (o *GetCostsOverviewForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get costs overview forbidden response a status code equal to that given
func (o *GetCostsOverviewForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get costs overview forbidden response
func (o *GetCostsOverviewForbidden) Code() int {
	return 403
}

func (o *GetCostsOverviewForbidden) Error() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewForbidden  %+v", 403, o.Payload)
}

func (o *GetCostsOverviewForbidden) String() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewForbidden  %+v", 403, o.Payload)
}

func (o *GetCostsOverviewForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetCostsOverviewForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCostsOverviewNotFound creates a GetCostsOverviewNotFound with default headers values
func NewGetCostsOverviewNotFound() *GetCostsOverviewNotFound {
	return &GetCostsOverviewNotFound{}
}

/*
GetCostsOverviewNotFound describes a response with status code 404, with default header values.

Organization not found. (code: `organization.not_found`)
*/
type GetCostsOverviewNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get costs overview not found response has a 2xx status code
func (o *GetCostsOverviewNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get costs overview not found response has a 3xx status code
func (o *GetCostsOverviewNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get costs overview not found response has a 4xx status code
func (o *GetCostsOverviewNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get costs overview not found response has a 5xx status code
func (o *GetCostsOverviewNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get costs overview not found response a status code equal to that given
func (o *GetCostsOverviewNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get costs overview not found response
func (o *GetCostsOverviewNotFound) Code() int {
	return 404
}

func (o *GetCostsOverviewNotFound) Error() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewNotFound  %+v", 404, o.Payload)
}

func (o *GetCostsOverviewNotFound) String() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewNotFound  %+v", 404, o.Payload)
}

func (o *GetCostsOverviewNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetCostsOverviewNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCostsOverviewInternalServerError creates a GetCostsOverviewInternalServerError with default headers values
func NewGetCostsOverviewInternalServerError() *GetCostsOverviewInternalServerError {
	return &GetCostsOverviewInternalServerError{}
}

/*
GetCostsOverviewInternalServerError describes a response with status code 500, with default header values.

An error occurred when fetching an overview of the costs for the organization. (code: `billing_service.failed_request`)
*/
type GetCostsOverviewInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get costs overview internal server error response has a 2xx status code
func (o *GetCostsOverviewInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get costs overview internal server error response has a 3xx status code
func (o *GetCostsOverviewInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get costs overview internal server error response has a 4xx status code
func (o *GetCostsOverviewInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get costs overview internal server error response has a 5xx status code
func (o *GetCostsOverviewInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get costs overview internal server error response a status code equal to that given
func (o *GetCostsOverviewInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get costs overview internal server error response
func (o *GetCostsOverviewInternalServerError) Code() int {
	return 500
}

func (o *GetCostsOverviewInternalServerError) Error() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCostsOverviewInternalServerError) String() string {
	return fmt.Sprintf("[GET /billing/costs/{organization_id}][%d] getCostsOverviewInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCostsOverviewInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetCostsOverviewInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
