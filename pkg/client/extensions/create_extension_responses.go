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

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateExtensionReader is a Reader for the CreateExtension structure.
type CreateExtensionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateExtensionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateExtensionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateExtensionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateExtensionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateExtensionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateExtensionCreated creates a CreateExtensionCreated with default headers values
func NewCreateExtensionCreated() *CreateExtensionCreated {
	return &CreateExtensionCreated{}
}

/*CreateExtensionCreated handles this case with default header values.

The extension that was just created.
*/
type CreateExtensionCreated struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.Extension
}

func (o *CreateExtensionCreated) Error() string {
	return fmt.Sprintf("[POST /deployments/extensions][%d] createExtensionCreated  %+v", 201, o.Payload)
}

func (o *CreateExtensionCreated) GetPayload() *models.Extension {
	return o.Payload
}

func (o *CreateExtensionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.Extension)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionBadRequest creates a CreateExtensionBadRequest with default headers values
func NewCreateExtensionBadRequest() *CreateExtensionBadRequest {
	return &CreateExtensionBadRequest{}
}

/*CreateExtensionBadRequest handles this case with default header values.

Could not download the extension from the specified URL. (code: `extensions.request_execution_failed`)
*/
type CreateExtensionBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateExtensionBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/extensions][%d] createExtensionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateExtensionBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateExtensionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionNotFound creates a CreateExtensionNotFound with default headers values
func NewCreateExtensionNotFound() *CreateExtensionNotFound {
	return &CreateExtensionNotFound{}
}

/*CreateExtensionNotFound handles this case with default header values.

Your current session does not have a user id associated with it. (code: `extensions.no_user_id`)
*/
type CreateExtensionNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateExtensionNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/extensions][%d] createExtensionNotFound  %+v", 404, o.Payload)
}

func (o *CreateExtensionNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateExtensionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateExtensionConflict creates a CreateExtensionConflict with default headers values
func NewCreateExtensionConflict() *CreateExtensionConflict {
	return &CreateExtensionConflict{}
}

/*CreateExtensionConflict handles this case with default header values.

An extension already exists with the generated id. Please try again. (code: `extensions.id_already_exists`)
*/
type CreateExtensionConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateExtensionConflict) Error() string {
	return fmt.Sprintf("[POST /deployments/extensions][%d] createExtensionConflict  %+v", 409, o.Payload)
}

func (o *CreateExtensionConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateExtensionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
