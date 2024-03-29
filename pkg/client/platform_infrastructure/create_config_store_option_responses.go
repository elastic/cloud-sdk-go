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

// CreateConfigStoreOptionReader is a Reader for the CreateConfigStoreOption structure.
type CreateConfigStoreOptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateConfigStoreOptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateConfigStoreOptionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateConfigStoreOptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateConfigStoreOptionCreated creates a CreateConfigStoreOptionCreated with default headers values
func NewCreateConfigStoreOptionCreated() *CreateConfigStoreOptionCreated {
	return &CreateConfigStoreOptionCreated{}
}

/*
CreateConfigStoreOptionCreated describes a response with status code 201, with default header values.

The Config Store Option was inserted successfully
*/
type CreateConfigStoreOptionCreated struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ConfigStoreOption
}

// IsSuccess returns true when this create config store option created response has a 2xx status code
func (o *CreateConfigStoreOptionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create config store option created response has a 3xx status code
func (o *CreateConfigStoreOptionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config store option created response has a 4xx status code
func (o *CreateConfigStoreOptionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create config store option created response has a 5xx status code
func (o *CreateConfigStoreOptionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create config store option created response a status code equal to that given
func (o *CreateConfigStoreOptionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create config store option created response
func (o *CreateConfigStoreOptionCreated) Code() int {
	return 201
}

func (o *CreateConfigStoreOptionCreated) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/store/{config_option_id}][%d] createConfigStoreOptionCreated  %+v", 201, o.Payload)
}

func (o *CreateConfigStoreOptionCreated) String() string {
	return fmt.Sprintf("[POST /platform/configuration/store/{config_option_id}][%d] createConfigStoreOptionCreated  %+v", 201, o.Payload)
}

func (o *CreateConfigStoreOptionCreated) GetPayload() *models.ConfigStoreOption {
	return o.Payload
}

func (o *CreateConfigStoreOptionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ConfigStoreOption)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateConfigStoreOptionBadRequest creates a CreateConfigStoreOptionBadRequest with default headers values
func NewCreateConfigStoreOptionBadRequest() *CreateConfigStoreOptionBadRequest {
	return &CreateConfigStoreOptionBadRequest{}
}

/*
CreateConfigStoreOptionBadRequest describes a response with status code 400, with default header values.

Config Store Option data already exists for the given name. (code: `platform.config.store.already_exists`)
*/
type CreateConfigStoreOptionBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create config store option bad request response has a 2xx status code
func (o *CreateConfigStoreOptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create config store option bad request response has a 3xx status code
func (o *CreateConfigStoreOptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create config store option bad request response has a 4xx status code
func (o *CreateConfigStoreOptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create config store option bad request response has a 5xx status code
func (o *CreateConfigStoreOptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create config store option bad request response a status code equal to that given
func (o *CreateConfigStoreOptionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create config store option bad request response
func (o *CreateConfigStoreOptionBadRequest) Code() int {
	return 400
}

func (o *CreateConfigStoreOptionBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/store/{config_option_id}][%d] createConfigStoreOptionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConfigStoreOptionBadRequest) String() string {
	return fmt.Sprintf("[POST /platform/configuration/store/{config_option_id}][%d] createConfigStoreOptionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateConfigStoreOptionBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateConfigStoreOptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
