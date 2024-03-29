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

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateSecurityDeploymentReader is a Reader for the CreateSecurityDeployment structure.
type CreateSecurityDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSecurityDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSecurityDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateSecurityDeploymentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSecurityDeploymentCreated creates a CreateSecurityDeploymentCreated with default headers values
func NewCreateSecurityDeploymentCreated() *CreateSecurityDeploymentCreated {
	return &CreateSecurityDeploymentCreated{}
}

/*
CreateSecurityDeploymentCreated describes a response with status code 201, with default header values.

The security deployment was successfully created
*/
type CreateSecurityDeploymentCreated struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.IDResponse
}

// IsSuccess returns true when this create security deployment created response has a 2xx status code
func (o *CreateSecurityDeploymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create security deployment created response has a 3xx status code
func (o *CreateSecurityDeploymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create security deployment created response has a 4xx status code
func (o *CreateSecurityDeploymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create security deployment created response has a 5xx status code
func (o *CreateSecurityDeploymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create security deployment created response a status code equal to that given
func (o *CreateSecurityDeploymentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create security deployment created response
func (o *CreateSecurityDeploymentCreated) Code() int {
	return 201
}

func (o *CreateSecurityDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment][%d] createSecurityDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateSecurityDeploymentCreated) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment][%d] createSecurityDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateSecurityDeploymentCreated) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *CreateSecurityDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSecurityDeploymentConflict creates a CreateSecurityDeploymentConflict with default headers values
func NewCreateSecurityDeploymentConflict() *CreateSecurityDeploymentConflict {
	return &CreateSecurityDeploymentConflict{}
}

/*
	CreateSecurityDeploymentConflict describes a response with status code 409, with default header values.

	* There is a version conflict. (code: `security_deployment.version_conflict`)

* There is a version conflict. (code: `security_deployment.already_exists`)
*/
type CreateSecurityDeploymentConflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create security deployment conflict response has a 2xx status code
func (o *CreateSecurityDeploymentConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create security deployment conflict response has a 3xx status code
func (o *CreateSecurityDeploymentConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create security deployment conflict response has a 4xx status code
func (o *CreateSecurityDeploymentConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create security deployment conflict response has a 5xx status code
func (o *CreateSecurityDeploymentConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create security deployment conflict response a status code equal to that given
func (o *CreateSecurityDeploymentConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create security deployment conflict response
func (o *CreateSecurityDeploymentConflict) Code() int {
	return 409
}

func (o *CreateSecurityDeploymentConflict) Error() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment][%d] createSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *CreateSecurityDeploymentConflict) String() string {
	return fmt.Sprintf("[POST /platform/configuration/security/deployment][%d] createSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *CreateSecurityDeploymentConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateSecurityDeploymentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
