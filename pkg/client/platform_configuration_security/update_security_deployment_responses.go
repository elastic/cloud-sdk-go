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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateSecurityDeploymentReader is a Reader for the UpdateSecurityDeployment structure.
type UpdateSecurityDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSecurityDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSecurityDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateSecurityDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateSecurityDeploymentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateSecurityDeploymentRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateSecurityDeploymentOK creates a UpdateSecurityDeploymentOK with default headers values
func NewUpdateSecurityDeploymentOK() *UpdateSecurityDeploymentOK {
	return &UpdateSecurityDeploymentOK{}
}

/*UpdateSecurityDeploymentOK handles this case with default header values.

The security deployment was successfully updated
*/
type UpdateSecurityDeploymentOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.IDResponse
}

func (o *UpdateSecurityDeploymentOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentOK  %+v", 200, o.Payload)
}

func (o *UpdateSecurityDeploymentOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *UpdateSecurityDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityDeploymentNotFound creates a UpdateSecurityDeploymentNotFound with default headers values
func NewUpdateSecurityDeploymentNotFound() *UpdateSecurityDeploymentNotFound {
	return &UpdateSecurityDeploymentNotFound{}
}

/*UpdateSecurityDeploymentNotFound handles this case with default header values.

The security deployment was not found. (code: `security_deployment.not_found`)
*/
type UpdateSecurityDeploymentNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSecurityDeploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateSecurityDeploymentNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSecurityDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityDeploymentConflict creates a UpdateSecurityDeploymentConflict with default headers values
func NewUpdateSecurityDeploymentConflict() *UpdateSecurityDeploymentConflict {
	return &UpdateSecurityDeploymentConflict{}
}

/*UpdateSecurityDeploymentConflict handles this case with default header values.

* There is a version conflict. (code: `security_deployment.version_conflict`)
* There is a version conflict. (code: `security_deployment.already_exists`)
 */
type UpdateSecurityDeploymentConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSecurityDeploymentConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentConflict  %+v", 409, o.Payload)
}

func (o *UpdateSecurityDeploymentConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSecurityDeploymentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSecurityDeploymentRetryWith creates a UpdateSecurityDeploymentRetryWith with default headers values
func NewUpdateSecurityDeploymentRetryWith() *UpdateSecurityDeploymentRetryWith {
	return &UpdateSecurityDeploymentRetryWith{}
}

/*UpdateSecurityDeploymentRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateSecurityDeploymentRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateSecurityDeploymentRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/deployment][%d] updateSecurityDeploymentRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateSecurityDeploymentRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateSecurityDeploymentRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
