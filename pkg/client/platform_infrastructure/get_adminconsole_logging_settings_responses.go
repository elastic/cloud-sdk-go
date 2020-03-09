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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetAdminconsoleLoggingSettingsReader is a Reader for the GetAdminconsoleLoggingSettings structure.
type GetAdminconsoleLoggingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdminconsoleLoggingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdminconsoleLoggingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAdminconsoleLoggingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAdminconsoleLoggingSettingsOK creates a GetAdminconsoleLoggingSettingsOK with default headers values
func NewGetAdminconsoleLoggingSettingsOK() *GetAdminconsoleLoggingSettingsOK {
	return &GetAdminconsoleLoggingSettingsOK{}
}

/*GetAdminconsoleLoggingSettingsOK handles this case with default header values.

The logging settings for the adminconsole instance
*/
type GetAdminconsoleLoggingSettingsOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.LoggingSettings
}

func (o *GetAdminconsoleLoggingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/adminconsoles/{adminconsole_id}/logging_settings][%d] getAdminconsoleLoggingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAdminconsoleLoggingSettingsOK) GetPayload() *models.LoggingSettings {
	return o.Payload
}

func (o *GetAdminconsoleLoggingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.LoggingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdminconsoleLoggingSettingsNotFound creates a GetAdminconsoleLoggingSettingsNotFound with default headers values
func NewGetAdminconsoleLoggingSettingsNotFound() *GetAdminconsoleLoggingSettingsNotFound {
	return &GetAdminconsoleLoggingSettingsNotFound{}
}

/*GetAdminconsoleLoggingSettingsNotFound handles this case with default header values.

The logging settings for this adminconsole were not found. (code: `adminconsoles.logging_settings.not_found`)
*/
type GetAdminconsoleLoggingSettingsNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetAdminconsoleLoggingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/adminconsoles/{adminconsole_id}/logging_settings][%d] getAdminconsoleLoggingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAdminconsoleLoggingSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetAdminconsoleLoggingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
