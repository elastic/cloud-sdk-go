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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopAllocatorMaintenanceModeReader is a Reader for the StopAllocatorMaintenanceMode structure.
type StopAllocatorMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAllocatorMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopAllocatorMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopAllocatorMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopAllocatorMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopAllocatorMaintenanceModeAccepted creates a StopAllocatorMaintenanceModeAccepted with default headers values
func NewStopAllocatorMaintenanceModeAccepted() *StopAllocatorMaintenanceModeAccepted {
	return &StopAllocatorMaintenanceModeAccepted{}
}

/*
StopAllocatorMaintenanceModeAccepted describes a response with status code 202, with default header values.

The stop maintenance mode command was issued successfully
*/
type StopAllocatorMaintenanceModeAccepted struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this stop allocator maintenance mode accepted response has a 2xx status code
func (o *StopAllocatorMaintenanceModeAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop allocator maintenance mode accepted response has a 3xx status code
func (o *StopAllocatorMaintenanceModeAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop allocator maintenance mode accepted response has a 4xx status code
func (o *StopAllocatorMaintenanceModeAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop allocator maintenance mode accepted response has a 5xx status code
func (o *StopAllocatorMaintenanceModeAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this stop allocator maintenance mode accepted response a status code equal to that given
func (o *StopAllocatorMaintenanceModeAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the stop allocator maintenance mode accepted response
func (o *StopAllocatorMaintenanceModeAccepted) Code() int {
	return 202
}

func (o *StopAllocatorMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StopAllocatorMaintenanceModeAccepted) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StopAllocatorMaintenanceModeAccepted) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *StopAllocatorMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopAllocatorMaintenanceModeForbidden creates a StopAllocatorMaintenanceModeForbidden with default headers values
func NewStopAllocatorMaintenanceModeForbidden() *StopAllocatorMaintenanceModeForbidden {
	return &StopAllocatorMaintenanceModeForbidden{}
}

/*
StopAllocatorMaintenanceModeForbidden describes a response with status code 403, with default header values.

The stop maintenance mode command was prohibited for the given allocator. (code: `root.unauthorized.rbac`)
*/
type StopAllocatorMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop allocator maintenance mode forbidden response has a 2xx status code
func (o *StopAllocatorMaintenanceModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop allocator maintenance mode forbidden response has a 3xx status code
func (o *StopAllocatorMaintenanceModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop allocator maintenance mode forbidden response has a 4xx status code
func (o *StopAllocatorMaintenanceModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop allocator maintenance mode forbidden response has a 5xx status code
func (o *StopAllocatorMaintenanceModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop allocator maintenance mode forbidden response a status code equal to that given
func (o *StopAllocatorMaintenanceModeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop allocator maintenance mode forbidden response
func (o *StopAllocatorMaintenanceModeForbidden) Code() int {
	return 403
}

func (o *StopAllocatorMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StopAllocatorMaintenanceModeForbidden) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StopAllocatorMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopAllocatorMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopAllocatorMaintenanceModeNotFound creates a StopAllocatorMaintenanceModeNotFound with default headers values
func NewStopAllocatorMaintenanceModeNotFound() *StopAllocatorMaintenanceModeNotFound {
	return &StopAllocatorMaintenanceModeNotFound{}
}

/*
StopAllocatorMaintenanceModeNotFound describes a response with status code 404, with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type StopAllocatorMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this stop allocator maintenance mode not found response has a 2xx status code
func (o *StopAllocatorMaintenanceModeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop allocator maintenance mode not found response has a 3xx status code
func (o *StopAllocatorMaintenanceModeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop allocator maintenance mode not found response has a 4xx status code
func (o *StopAllocatorMaintenanceModeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop allocator maintenance mode not found response has a 5xx status code
func (o *StopAllocatorMaintenanceModeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop allocator maintenance mode not found response a status code equal to that given
func (o *StopAllocatorMaintenanceModeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop allocator maintenance mode not found response
func (o *StopAllocatorMaintenanceModeNotFound) Code() int {
	return 404
}

func (o *StopAllocatorMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StopAllocatorMaintenanceModeNotFound) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/maintenance-mode/_stop][%d] stopAllocatorMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StopAllocatorMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopAllocatorMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
