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

// DeleteLdapConfigurationReader is a Reader for the DeleteLdapConfiguration structure.
type DeleteLdapConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLdapConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLdapConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteLdapConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLdapConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteLdapConfigurationOK creates a DeleteLdapConfigurationOK with default headers values
func NewDeleteLdapConfigurationOK() *DeleteLdapConfigurationOK {
	return &DeleteLdapConfigurationOK{}
}

/*
DeleteLdapConfigurationOK describes a response with status code 200, with default header values.

The LDAP configuration was successfully deleted
*/
type DeleteLdapConfigurationOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete ldap configuration o k response has a 2xx status code
func (o *DeleteLdapConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete ldap configuration o k response has a 3xx status code
func (o *DeleteLdapConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ldap configuration o k response has a 4xx status code
func (o *DeleteLdapConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete ldap configuration o k response has a 5xx status code
func (o *DeleteLdapConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete ldap configuration o k response a status code equal to that given
func (o *DeleteLdapConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete ldap configuration o k response
func (o *DeleteLdapConfigurationOK) Code() int {
	return 200
}

func (o *DeleteLdapConfigurationOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationOK  %+v", 200, o.Payload)
}

func (o *DeleteLdapConfigurationOK) String() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationOK  %+v", 200, o.Payload)
}

func (o *DeleteLdapConfigurationOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteLdapConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapConfigurationNotFound creates a DeleteLdapConfigurationNotFound with default headers values
func NewDeleteLdapConfigurationNotFound() *DeleteLdapConfigurationNotFound {
	return &DeleteLdapConfigurationNotFound{}
}

/*
DeleteLdapConfigurationNotFound describes a response with status code 404, with default header values.

The realm specified by {realm_id} cannot be found. (code: `security_realm.not_found`)
*/
type DeleteLdapConfigurationNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete ldap configuration not found response has a 2xx status code
func (o *DeleteLdapConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete ldap configuration not found response has a 3xx status code
func (o *DeleteLdapConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ldap configuration not found response has a 4xx status code
func (o *DeleteLdapConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete ldap configuration not found response has a 5xx status code
func (o *DeleteLdapConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete ldap configuration not found response a status code equal to that given
func (o *DeleteLdapConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete ldap configuration not found response
func (o *DeleteLdapConfigurationNotFound) Code() int {
	return 404
}

func (o *DeleteLdapConfigurationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLdapConfigurationNotFound) String() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLdapConfigurationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteLdapConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteLdapConfigurationConflict creates a DeleteLdapConfigurationConflict with default headers values
func NewDeleteLdapConfigurationConflict() *DeleteLdapConfigurationConflict {
	return &DeleteLdapConfigurationConflict{}
}

/*
DeleteLdapConfigurationConflict describes a response with status code 409, with default header values.

There is a version conflict. (code: `security_realm.version_conflict`)
*/
type DeleteLdapConfigurationConflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete ldap configuration conflict response has a 2xx status code
func (o *DeleteLdapConfigurationConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete ldap configuration conflict response has a 3xx status code
func (o *DeleteLdapConfigurationConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ldap configuration conflict response has a 4xx status code
func (o *DeleteLdapConfigurationConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete ldap configuration conflict response has a 5xx status code
func (o *DeleteLdapConfigurationConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete ldap configuration conflict response a status code equal to that given
func (o *DeleteLdapConfigurationConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete ldap configuration conflict response
func (o *DeleteLdapConfigurationConflict) Code() int {
	return 409
}

func (o *DeleteLdapConfigurationConflict) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationConflict  %+v", 409, o.Payload)
}

func (o *DeleteLdapConfigurationConflict) String() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/realms/ldap/{realm_id}][%d] deleteLdapConfigurationConflict  %+v", 409, o.Payload)
}

func (o *DeleteLdapConfigurationConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteLdapConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
