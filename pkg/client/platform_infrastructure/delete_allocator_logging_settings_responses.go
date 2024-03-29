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

// DeleteAllocatorLoggingSettingsReader is a Reader for the DeleteAllocatorLoggingSettings structure.
type DeleteAllocatorLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllocatorLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllocatorLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteAllocatorLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAllocatorLoggingSettingsOK creates a DeleteAllocatorLoggingSettingsOK with default headers values
func NewDeleteAllocatorLoggingSettingsOK() *DeleteAllocatorLoggingSettingsOK {
	return &DeleteAllocatorLoggingSettingsOK{}
}

/*
DeleteAllocatorLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the allocator instance
*/
type DeleteAllocatorLoggingSettingsOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.LoggingSettings
}

// IsSuccess returns true when this delete allocator logging settings o k response has a 2xx status code
func (o *DeleteAllocatorLoggingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete allocator logging settings o k response has a 3xx status code
func (o *DeleteAllocatorLoggingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete allocator logging settings o k response has a 4xx status code
func (o *DeleteAllocatorLoggingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete allocator logging settings o k response has a 5xx status code
func (o *DeleteAllocatorLoggingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete allocator logging settings o k response a status code equal to that given
func (o *DeleteAllocatorLoggingSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete allocator logging settings o k response
func (o *DeleteAllocatorLoggingSettingsOK) Code() int {
	return 200
}

func (o *DeleteAllocatorLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] deleteAllocatorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteAllocatorLoggingSettingsOK) String() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] deleteAllocatorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteAllocatorLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *DeleteAllocatorLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.LoggingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocatorLoggingSettingsNotFound creates a DeleteAllocatorLoggingSettingsNotFound with default headers values
func NewDeleteAllocatorLoggingSettingsNotFound() *DeleteAllocatorLoggingSettingsNotFound {
	return &DeleteAllocatorLoggingSettingsNotFound{}
}

/*
DeleteAllocatorLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this allocator were not found. (code: `allocators.logging_settings.not_found`)
*/
type DeleteAllocatorLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete allocator logging settings not found response has a 2xx status code
func (o *DeleteAllocatorLoggingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete allocator logging settings not found response has a 3xx status code
func (o *DeleteAllocatorLoggingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete allocator logging settings not found response has a 4xx status code
func (o *DeleteAllocatorLoggingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete allocator logging settings not found response has a 5xx status code
func (o *DeleteAllocatorLoggingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete allocator logging settings not found response a status code equal to that given
func (o *DeleteAllocatorLoggingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete allocator logging settings not found response
func (o *DeleteAllocatorLoggingSettingsNotFound) Code() int {
	return 404
}

func (o *DeleteAllocatorLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] deleteAllocatorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAllocatorLoggingSettingsNotFound) String() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}/logging_settings][%d] deleteAllocatorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAllocatorLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
