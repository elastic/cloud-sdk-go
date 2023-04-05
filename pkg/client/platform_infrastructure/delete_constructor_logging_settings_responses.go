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

// DeleteConstructorLoggingSettingsReader is a Reader for the DeleteConstructorLoggingSettings structure.
type DeleteConstructorLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConstructorLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConstructorLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteConstructorLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConstructorLoggingSettingsOK creates a DeleteConstructorLoggingSettingsOK with default headers values
func NewDeleteConstructorLoggingSettingsOK() *DeleteConstructorLoggingSettingsOK {
	return &DeleteConstructorLoggingSettingsOK{}
}

/*
DeleteConstructorLoggingSettingsOK describes a response with status code 200, with default header values.

The updated logging settings for the constructor instance
*/
type DeleteConstructorLoggingSettingsOK struct {

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

// IsSuccess returns true when this delete constructor logging settings o k response has a 2xx status code
func (o *DeleteConstructorLoggingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete constructor logging settings o k response has a 3xx status code
func (o *DeleteConstructorLoggingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete constructor logging settings o k response has a 4xx status code
func (o *DeleteConstructorLoggingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete constructor logging settings o k response has a 5xx status code
func (o *DeleteConstructorLoggingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete constructor logging settings o k response a status code equal to that given
func (o *DeleteConstructorLoggingSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete constructor logging settings o k response
func (o *DeleteConstructorLoggingSettingsOK) Code() int {
	return 200
}

func (o *DeleteConstructorLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] deleteConstructorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteConstructorLoggingSettingsOK) String() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] deleteConstructorLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *DeleteConstructorLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *DeleteConstructorLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteConstructorLoggingSettingsNotFound creates a DeleteConstructorLoggingSettingsNotFound with default headers values
func NewDeleteConstructorLoggingSettingsNotFound() *DeleteConstructorLoggingSettingsNotFound {
	return &DeleteConstructorLoggingSettingsNotFound{}
}

/*
DeleteConstructorLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this constructor were not found. (code: `constructors.logging_settings.not_found`)
*/
type DeleteConstructorLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete constructor logging settings not found response has a 2xx status code
func (o *DeleteConstructorLoggingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete constructor logging settings not found response has a 3xx status code
func (o *DeleteConstructorLoggingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete constructor logging settings not found response has a 4xx status code
func (o *DeleteConstructorLoggingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete constructor logging settings not found response has a 5xx status code
func (o *DeleteConstructorLoggingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete constructor logging settings not found response a status code equal to that given
func (o *DeleteConstructorLoggingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete constructor logging settings not found response
func (o *DeleteConstructorLoggingSettingsNotFound) Code() int {
	return 404
}

func (o *DeleteConstructorLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] deleteConstructorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConstructorLoggingSettingsNotFound) String() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/constructors/{constructor_id}/logging_settings][%d] deleteConstructorLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConstructorLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteConstructorLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
