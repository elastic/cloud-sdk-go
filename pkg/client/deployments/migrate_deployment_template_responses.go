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

// MigrateDeploymentTemplateReader is a Reader for the MigrateDeploymentTemplate structure.
type MigrateDeploymentTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MigrateDeploymentTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMigrateDeploymentTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMigrateDeploymentTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMigrateDeploymentTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMigrateDeploymentTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMigrateDeploymentTemplateOK creates a MigrateDeploymentTemplateOK with default headers values
func NewMigrateDeploymentTemplateOK() *MigrateDeploymentTemplateOK {
	return &MigrateDeploymentTemplateOK{}
}

/*
MigrateDeploymentTemplateOK describes a response with status code 200, with default header values.

The request was valid and the deployment can be migrated to the template with the provided ID.
*/
type MigrateDeploymentTemplateOK struct {
	Payload *models.DeploymentUpdateRequest
}

// IsSuccess returns true when this migrate deployment template o k response has a 2xx status code
func (o *MigrateDeploymentTemplateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this migrate deployment template o k response has a 3xx status code
func (o *MigrateDeploymentTemplateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate deployment template o k response has a 4xx status code
func (o *MigrateDeploymentTemplateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this migrate deployment template o k response has a 5xx status code
func (o *MigrateDeploymentTemplateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate deployment template o k response a status code equal to that given
func (o *MigrateDeploymentTemplateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the migrate deployment template o k response
func (o *MigrateDeploymentTemplateOK) Code() int {
	return 200
}

func (o *MigrateDeploymentTemplateOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateOK  %+v", 200, o.Payload)
}

func (o *MigrateDeploymentTemplateOK) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateOK  %+v", 200, o.Payload)
}

func (o *MigrateDeploymentTemplateOK) GetPayload() *models.DeploymentUpdateRequest {
	return o.Payload
}

func (o *MigrateDeploymentTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentUpdateRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateDeploymentTemplateBadRequest creates a MigrateDeploymentTemplateBadRequest with default headers values
func NewMigrateDeploymentTemplateBadRequest() *MigrateDeploymentTemplateBadRequest {
	return &MigrateDeploymentTemplateBadRequest{}
}

/*
MigrateDeploymentTemplateBadRequest describes a response with status code 400, with default header values.

The deployment cannot be successfully migrated to the template with the provided ID.
*/
type MigrateDeploymentTemplateBadRequest struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this migrate deployment template bad request response has a 2xx status code
func (o *MigrateDeploymentTemplateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this migrate deployment template bad request response has a 3xx status code
func (o *MigrateDeploymentTemplateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate deployment template bad request response has a 4xx status code
func (o *MigrateDeploymentTemplateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this migrate deployment template bad request response has a 5xx status code
func (o *MigrateDeploymentTemplateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate deployment template bad request response a status code equal to that given
func (o *MigrateDeploymentTemplateBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the migrate deployment template bad request response
func (o *MigrateDeploymentTemplateBadRequest) Code() int {
	return 400
}

func (o *MigrateDeploymentTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *MigrateDeploymentTemplateBadRequest) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *MigrateDeploymentTemplateBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateDeploymentTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateDeploymentTemplateUnauthorized creates a MigrateDeploymentTemplateUnauthorized with default headers values
func NewMigrateDeploymentTemplateUnauthorized() *MigrateDeploymentTemplateUnauthorized {
	return &MigrateDeploymentTemplateUnauthorized{}
}

/*
MigrateDeploymentTemplateUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action.
*/
type MigrateDeploymentTemplateUnauthorized struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this migrate deployment template unauthorized response has a 2xx status code
func (o *MigrateDeploymentTemplateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this migrate deployment template unauthorized response has a 3xx status code
func (o *MigrateDeploymentTemplateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate deployment template unauthorized response has a 4xx status code
func (o *MigrateDeploymentTemplateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this migrate deployment template unauthorized response has a 5xx status code
func (o *MigrateDeploymentTemplateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate deployment template unauthorized response a status code equal to that given
func (o *MigrateDeploymentTemplateUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the migrate deployment template unauthorized response
func (o *MigrateDeploymentTemplateUnauthorized) Code() int {
	return 401
}

func (o *MigrateDeploymentTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *MigrateDeploymentTemplateUnauthorized) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *MigrateDeploymentTemplateUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateDeploymentTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMigrateDeploymentTemplateNotFound creates a MigrateDeploymentTemplateNotFound with default headers values
func NewMigrateDeploymentTemplateNotFound() *MigrateDeploymentTemplateNotFound {
	return &MigrateDeploymentTemplateNotFound{}
}

/*
MigrateDeploymentTemplateNotFound describes a response with status code 404, with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type MigrateDeploymentTemplateNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this migrate deployment template not found response has a 2xx status code
func (o *MigrateDeploymentTemplateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this migrate deployment template not found response has a 3xx status code
func (o *MigrateDeploymentTemplateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this migrate deployment template not found response has a 4xx status code
func (o *MigrateDeploymentTemplateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this migrate deployment template not found response has a 5xx status code
func (o *MigrateDeploymentTemplateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this migrate deployment template not found response a status code equal to that given
func (o *MigrateDeploymentTemplateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the migrate deployment template not found response
func (o *MigrateDeploymentTemplateNotFound) Code() int {
	return 404
}

func (o *MigrateDeploymentTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateNotFound  %+v", 404, o.Payload)
}

func (o *MigrateDeploymentTemplateNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/migrate_template][%d] migrateDeploymentTemplateNotFound  %+v", 404, o.Payload)
}

func (o *MigrateDeploymentTemplateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MigrateDeploymentTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
