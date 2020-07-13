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

// SetAPIBaseURLReader is a Reader for the SetAPIBaseURL structure.
type SetAPIBaseURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAPIBaseURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAPIBaseURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetAPIBaseURLBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetAPIBaseURLNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetAPIBaseURLConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAPIBaseURLOK creates a SetAPIBaseURLOK with default headers values
func NewSetAPIBaseURLOK() *SetAPIBaseURLOK {
	return &SetAPIBaseURLOK{}
}

/*SetAPIBaseURLOK handles this case with default header values.

The API base Url was successfully saved.
*/
type SetAPIBaseURLOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.APIBaseURLData
}

func (o *SetAPIBaseURLOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/api_base_url][%d] setApiBaseUrlOK  %+v", 200, o.Payload)
}

func (o *SetAPIBaseURLOK) GetPayload() *models.APIBaseURLData {
	return o.Payload
}

func (o *SetAPIBaseURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.APIBaseURLData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAPIBaseURLBadRequest creates a SetAPIBaseURLBadRequest with default headers values
func NewSetAPIBaseURLBadRequest() *SetAPIBaseURLBadRequest {
	return &SetAPIBaseURLBadRequest{}
}

/*SetAPIBaseURLBadRequest handles this case with default header values.

The optimistic locking version format was wrong. (code: `adminconsole.base_url.bad_version_format`)
*/
type SetAPIBaseURLBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAPIBaseURLBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/api_base_url][%d] setApiBaseUrlBadRequest  %+v", 400, o.Payload)
}

func (o *SetAPIBaseURLBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAPIBaseURLBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAPIBaseURLNotFound creates a SetAPIBaseURLNotFound with default headers values
func NewSetAPIBaseURLNotFound() *SetAPIBaseURLNotFound {
	return &SetAPIBaseURLNotFound{}
}

/*SetAPIBaseURLNotFound handles this case with default header values.

There is no configured API base value but optimistic locking was sent. (code: `adminconsole.base_url.not_found`)
*/
type SetAPIBaseURLNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAPIBaseURLNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/api_base_url][%d] setApiBaseUrlNotFound  %+v", 404, o.Payload)
}

func (o *SetAPIBaseURLNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAPIBaseURLNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAPIBaseURLConflict creates a SetAPIBaseURLConflict with default headers values
func NewSetAPIBaseURLConflict() *SetAPIBaseURLConflict {
	return &SetAPIBaseURLConflict{}
}

/*SetAPIBaseURLConflict handles this case with default header values.

There was an optimistic locking version conflict. (code: `adminconsole.base_url.version_conflict`)
*/
type SetAPIBaseURLConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAPIBaseURLConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/api_base_url][%d] setApiBaseUrlConflict  %+v", 409, o.Payload)
}

func (o *SetAPIBaseURLConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAPIBaseURLConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
