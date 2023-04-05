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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetDeploymentEsResourceKeystoreReader is a Reader for the SetDeploymentEsResourceKeystore structure.
type SetDeploymentEsResourceKeystoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDeploymentEsResourceKeystoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDeploymentEsResourceKeystoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetDeploymentEsResourceKeystoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetDeploymentEsResourceKeystoreRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetDeploymentEsResourceKeystoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetDeploymentEsResourceKeystoreOK creates a SetDeploymentEsResourceKeystoreOK with default headers values
func NewSetDeploymentEsResourceKeystoreOK() *SetDeploymentEsResourceKeystoreOK {
	return &SetDeploymentEsResourceKeystoreOK{}
}

/*
SetDeploymentEsResourceKeystoreOK describes a response with status code 200, with default header values.

The new contents of the Elasticsearch keystore
*/
type SetDeploymentEsResourceKeystoreOK struct {
	Payload *models.KeystoreContents
}

// IsSuccess returns true when this set deployment es resource keystore o k response has a 2xx status code
func (o *SetDeploymentEsResourceKeystoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set deployment es resource keystore o k response has a 3xx status code
func (o *SetDeploymentEsResourceKeystoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set deployment es resource keystore o k response has a 4xx status code
func (o *SetDeploymentEsResourceKeystoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set deployment es resource keystore o k response has a 5xx status code
func (o *SetDeploymentEsResourceKeystoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set deployment es resource keystore o k response a status code equal to that given
func (o *SetDeploymentEsResourceKeystoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set deployment es resource keystore o k response
func (o *SetDeploymentEsResourceKeystoreOK) Code() int {
	return 200
}

func (o *SetDeploymentEsResourceKeystoreOK) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreOK  %+v", 200, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreOK) String() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreOK  %+v", 200, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreOK) GetPayload() *models.KeystoreContents {
	return o.Payload
}

func (o *SetDeploymentEsResourceKeystoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeystoreContents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentEsResourceKeystoreNotFound creates a SetDeploymentEsResourceKeystoreNotFound with default headers values
func NewSetDeploymentEsResourceKeystoreNotFound() *SetDeploymentEsResourceKeystoreNotFound {
	return &SetDeploymentEsResourceKeystoreNotFound{}
}

/*
	SetDeploymentEsResourceKeystoreNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type SetDeploymentEsResourceKeystoreNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this set deployment es resource keystore not found response has a 2xx status code
func (o *SetDeploymentEsResourceKeystoreNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set deployment es resource keystore not found response has a 3xx status code
func (o *SetDeploymentEsResourceKeystoreNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set deployment es resource keystore not found response has a 4xx status code
func (o *SetDeploymentEsResourceKeystoreNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this set deployment es resource keystore not found response has a 5xx status code
func (o *SetDeploymentEsResourceKeystoreNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this set deployment es resource keystore not found response a status code equal to that given
func (o *SetDeploymentEsResourceKeystoreNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the set deployment es resource keystore not found response
func (o *SetDeploymentEsResourceKeystoreNotFound) Code() int {
	return 404
}

func (o *SetDeploymentEsResourceKeystoreNotFound) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreNotFound  %+v", 404, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreNotFound) String() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreNotFound  %+v", 404, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentEsResourceKeystoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentEsResourceKeystoreRetryWith creates a SetDeploymentEsResourceKeystoreRetryWith with default headers values
func NewSetDeploymentEsResourceKeystoreRetryWith() *SetDeploymentEsResourceKeystoreRetryWith {
	return &SetDeploymentEsResourceKeystoreRetryWith{}
}

/*
SetDeploymentEsResourceKeystoreRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetDeploymentEsResourceKeystoreRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this set deployment es resource keystore retry with response has a 2xx status code
func (o *SetDeploymentEsResourceKeystoreRetryWith) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set deployment es resource keystore retry with response has a 3xx status code
func (o *SetDeploymentEsResourceKeystoreRetryWith) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set deployment es resource keystore retry with response has a 4xx status code
func (o *SetDeploymentEsResourceKeystoreRetryWith) IsClientError() bool {
	return true
}

// IsServerError returns true when this set deployment es resource keystore retry with response has a 5xx status code
func (o *SetDeploymentEsResourceKeystoreRetryWith) IsServerError() bool {
	return false
}

// IsCode returns true when this set deployment es resource keystore retry with response a status code equal to that given
func (o *SetDeploymentEsResourceKeystoreRetryWith) IsCode(code int) bool {
	return code == 449
}

// Code gets the status code for the set deployment es resource keystore retry with response
func (o *SetDeploymentEsResourceKeystoreRetryWith) Code() int {
	return 449
}

func (o *SetDeploymentEsResourceKeystoreRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreRetryWith  %+v", 449, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreRetryWith) String() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreRetryWith  %+v", 449, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentEsResourceKeystoreRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentEsResourceKeystoreInternalServerError creates a SetDeploymentEsResourceKeystoreInternalServerError with default headers values
func NewSetDeploymentEsResourceKeystoreInternalServerError() *SetDeploymentEsResourceKeystoreInternalServerError {
	return &SetDeploymentEsResourceKeystoreInternalServerError{}
}

/*
SetDeploymentEsResourceKeystoreInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.metadata_internal_error`)
*/
type SetDeploymentEsResourceKeystoreInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this set deployment es resource keystore internal server error response has a 2xx status code
func (o *SetDeploymentEsResourceKeystoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this set deployment es resource keystore internal server error response has a 3xx status code
func (o *SetDeploymentEsResourceKeystoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set deployment es resource keystore internal server error response has a 4xx status code
func (o *SetDeploymentEsResourceKeystoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this set deployment es resource keystore internal server error response has a 5xx status code
func (o *SetDeploymentEsResourceKeystoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this set deployment es resource keystore internal server error response a status code equal to that given
func (o *SetDeploymentEsResourceKeystoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the set deployment es resource keystore internal server error response
func (o *SetDeploymentEsResourceKeystoreInternalServerError) Code() int {
	return 500
}

func (o *SetDeploymentEsResourceKeystoreInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreInternalServerError  %+v", 500, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] setDeploymentEsResourceKeystoreInternalServerError  %+v", 500, o.Payload)
}

func (o *SetDeploymentEsResourceKeystoreInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentEsResourceKeystoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
