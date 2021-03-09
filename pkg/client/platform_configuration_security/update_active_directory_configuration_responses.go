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

// UpdateActiveDirectoryConfigurationReader is a Reader for the UpdateActiveDirectoryConfiguration structure.
type UpdateActiveDirectoryConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateActiveDirectoryConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateActiveDirectoryConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateActiveDirectoryConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateActiveDirectoryConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateActiveDirectoryConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateActiveDirectoryConfigurationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateActiveDirectoryConfigurationOK creates a UpdateActiveDirectoryConfigurationOK with default headers values
func NewUpdateActiveDirectoryConfigurationOK() *UpdateActiveDirectoryConfigurationOK {
	return &UpdateActiveDirectoryConfigurationOK{}
}

/* UpdateActiveDirectoryConfigurationOK describes a response with status code 200, with default header values.

The Active Directory configuration was successfully updated
*/
type UpdateActiveDirectoryConfigurationOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload models.EmptyResponse
}

func (o *UpdateActiveDirectoryConfigurationOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/active-directory/{realm_id}][%d] updateActiveDirectoryConfigurationOK  %+v", 200, o.Payload)
}
func (o *UpdateActiveDirectoryConfigurationOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *UpdateActiveDirectoryConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateActiveDirectoryConfigurationBadRequest creates a UpdateActiveDirectoryConfigurationBadRequest with default headers values
func NewUpdateActiveDirectoryConfigurationBadRequest() *UpdateActiveDirectoryConfigurationBadRequest {
	return &UpdateActiveDirectoryConfigurationBadRequest{}
}

/* UpdateActiveDirectoryConfigurationBadRequest describes a response with status code 400, with default header values.

 * The realm id is already in use. (code: `security_realm.id_conflict`)
* The selected id is not valid. (code: `security_realm.invalid_id`)
* Order must be greater than zero. (code: `security_realm.invalid_order`)
* Invalid Elasticsearch Security realm type. (code: `security_realm.invalid_type`)
* The realm order is already in use. (code: `security_realm.order_conflict`)
* Advanced YAML format is invalid. (code: `security_realm.invalid_yaml`)
* The url format is invalid. (code: `security_realm.invalid_url`)
* Invalid Active Directory URL. (code: `security_realm.active_directory.invalid_url`)
* Invalid certificate bundle URL. (code: `security_realm.invalid_bundle_url`)
*/
type UpdateActiveDirectoryConfigurationBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateActiveDirectoryConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/active-directory/{realm_id}][%d] updateActiveDirectoryConfigurationBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateActiveDirectoryConfigurationBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateActiveDirectoryConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateActiveDirectoryConfigurationNotFound creates a UpdateActiveDirectoryConfigurationNotFound with default headers values
func NewUpdateActiveDirectoryConfigurationNotFound() *UpdateActiveDirectoryConfigurationNotFound {
	return &UpdateActiveDirectoryConfigurationNotFound{}
}

/* UpdateActiveDirectoryConfigurationNotFound describes a response with status code 404, with default header values.

The realm specified by {realm_id} cannot be found. (code: `security_realm.not_found`)
*/
type UpdateActiveDirectoryConfigurationNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateActiveDirectoryConfigurationNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/active-directory/{realm_id}][%d] updateActiveDirectoryConfigurationNotFound  %+v", 404, o.Payload)
}
func (o *UpdateActiveDirectoryConfigurationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateActiveDirectoryConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateActiveDirectoryConfigurationConflict creates a UpdateActiveDirectoryConfigurationConflict with default headers values
func NewUpdateActiveDirectoryConfigurationConflict() *UpdateActiveDirectoryConfigurationConflict {
	return &UpdateActiveDirectoryConfigurationConflict{}
}

/* UpdateActiveDirectoryConfigurationConflict describes a response with status code 409, with default header values.

There is a version conflict. (code: `security_realm.version_conflict`)
*/
type UpdateActiveDirectoryConfigurationConflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateActiveDirectoryConfigurationConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/active-directory/{realm_id}][%d] updateActiveDirectoryConfigurationConflict  %+v", 409, o.Payload)
}
func (o *UpdateActiveDirectoryConfigurationConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateActiveDirectoryConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateActiveDirectoryConfigurationRetryWith creates a UpdateActiveDirectoryConfigurationRetryWith with default headers values
func NewUpdateActiveDirectoryConfigurationRetryWith() *UpdateActiveDirectoryConfigurationRetryWith {
	return &UpdateActiveDirectoryConfigurationRetryWith{}
}

/* UpdateActiveDirectoryConfigurationRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateActiveDirectoryConfigurationRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateActiveDirectoryConfigurationRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/security/realms/active-directory/{realm_id}][%d] updateActiveDirectoryConfigurationRetryWith  %+v", 449, o.Payload)
}
func (o *UpdateActiveDirectoryConfigurationRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateActiveDirectoryConfigurationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
