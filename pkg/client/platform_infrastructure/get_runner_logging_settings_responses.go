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

// GetRunnerLoggingSettingsReader is a Reader for the GetRunnerLoggingSettings structure.
type GetRunnerLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunnerLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRunnerLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetRunnerLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRunnerLoggingSettingsOK creates a GetRunnerLoggingSettingsOK with default headers values
func NewGetRunnerLoggingSettingsOK() *GetRunnerLoggingSettingsOK {
	return &GetRunnerLoggingSettingsOK{}
}

/*
GetRunnerLoggingSettingsOK describes a response with status code 200, with default header values.

The logging settings for the runner instance
*/
type GetRunnerLoggingSettingsOK struct {

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

// IsSuccess returns true when this get runner logging settings o k response has a 2xx status code
func (o *GetRunnerLoggingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runner logging settings o k response has a 3xx status code
func (o *GetRunnerLoggingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner logging settings o k response has a 4xx status code
func (o *GetRunnerLoggingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runner logging settings o k response has a 5xx status code
func (o *GetRunnerLoggingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner logging settings o k response a status code equal to that given
func (o *GetRunnerLoggingSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runner logging settings o k response
func (o *GetRunnerLoggingSettingsOK) Code() int {
	return 200
}

func (o *GetRunnerLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}/logging_settings][%d] getRunnerLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRunnerLoggingSettingsOK) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}/logging_settings][%d] getRunnerLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRunnerLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *GetRunnerLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRunnerLoggingSettingsNotFound creates a GetRunnerLoggingSettingsNotFound with default headers values
func NewGetRunnerLoggingSettingsNotFound() *GetRunnerLoggingSettingsNotFound {
	return &GetRunnerLoggingSettingsNotFound{}
}

/*
GetRunnerLoggingSettingsNotFound describes a response with status code 404, with default header values.

The logging settings for this runner were not found. (code: `runners.logging_settings.not_found`)
*/
type GetRunnerLoggingSettingsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get runner logging settings not found response has a 2xx status code
func (o *GetRunnerLoggingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get runner logging settings not found response has a 3xx status code
func (o *GetRunnerLoggingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runner logging settings not found response has a 4xx status code
func (o *GetRunnerLoggingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get runner logging settings not found response has a 5xx status code
func (o *GetRunnerLoggingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get runner logging settings not found response a status code equal to that given
func (o *GetRunnerLoggingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get runner logging settings not found response
func (o *GetRunnerLoggingSettingsNotFound) Code() int {
	return 404
}

func (o *GetRunnerLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}/logging_settings][%d] getRunnerLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerLoggingSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/runners/{runner_id}/logging_settings][%d] getRunnerLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRunnerLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetRunnerLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
