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

// RestartDeploymentStatelessResourceReader is a Reader for the RestartDeploymentStatelessResource structure.
type RestartDeploymentStatelessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartDeploymentStatelessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartDeploymentStatelessResourceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRestartDeploymentStatelessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRestartDeploymentStatelessResourceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRestartDeploymentStatelessResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartDeploymentStatelessResourceAccepted creates a RestartDeploymentStatelessResourceAccepted with default headers values
func NewRestartDeploymentStatelessResourceAccepted() *RestartDeploymentStatelessResourceAccepted {
	return &RestartDeploymentStatelessResourceAccepted{}
}

/*
RestartDeploymentStatelessResourceAccepted describes a response with status code 202, with default header values.

The restart command was issued successfully
*/
type RestartDeploymentStatelessResourceAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

// IsSuccess returns true when this restart deployment stateless resource accepted response has a 2xx status code
func (o *RestartDeploymentStatelessResourceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restart deployment stateless resource accepted response has a 3xx status code
func (o *RestartDeploymentStatelessResourceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart deployment stateless resource accepted response has a 4xx status code
func (o *RestartDeploymentStatelessResourceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this restart deployment stateless resource accepted response has a 5xx status code
func (o *RestartDeploymentStatelessResourceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this restart deployment stateless resource accepted response a status code equal to that given
func (o *RestartDeploymentStatelessResourceAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the restart deployment stateless resource accepted response
func (o *RestartDeploymentStatelessResourceAccepted) Code() int {
	return 202
}

func (o *RestartDeploymentStatelessResourceAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceAccepted  %+v", 202, o.Payload)
}

func (o *RestartDeploymentStatelessResourceAccepted) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceAccepted  %+v", 202, o.Payload)
}

func (o *RestartDeploymentStatelessResourceAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDeploymentStatelessResourceNotFound creates a RestartDeploymentStatelessResourceNotFound with default headers values
func NewRestartDeploymentStatelessResourceNotFound() *RestartDeploymentStatelessResourceNotFound {
	return &RestartDeploymentStatelessResourceNotFound{}
}

/*
	RestartDeploymentStatelessResourceNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type RestartDeploymentStatelessResourceNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restart deployment stateless resource not found response has a 2xx status code
func (o *RestartDeploymentStatelessResourceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restart deployment stateless resource not found response has a 3xx status code
func (o *RestartDeploymentStatelessResourceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart deployment stateless resource not found response has a 4xx status code
func (o *RestartDeploymentStatelessResourceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this restart deployment stateless resource not found response has a 5xx status code
func (o *RestartDeploymentStatelessResourceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this restart deployment stateless resource not found response a status code equal to that given
func (o *RestartDeploymentStatelessResourceNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the restart deployment stateless resource not found response
func (o *RestartDeploymentStatelessResourceNotFound) Code() int {
	return 404
}

func (o *RestartDeploymentStatelessResourceNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceNotFound  %+v", 404, o.Payload)
}

func (o *RestartDeploymentStatelessResourceNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceNotFound  %+v", 404, o.Payload)
}

func (o *RestartDeploymentStatelessResourceNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartDeploymentStatelessResourceUnprocessableEntity creates a RestartDeploymentStatelessResourceUnprocessableEntity with default headers values
func NewRestartDeploymentStatelessResourceUnprocessableEntity() *RestartDeploymentStatelessResourceUnprocessableEntity {
	return &RestartDeploymentStatelessResourceUnprocessableEntity{}
}

/*
RestartDeploymentStatelessResourceUnprocessableEntity describes a response with status code 422, with default header values.

The command sent to a Resource found the Resource in an illegal state, the error message gives more details. (code: `deployments.deployment_resource_plan_change_error`)
*/
type RestartDeploymentStatelessResourceUnprocessableEntity struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restart deployment stateless resource unprocessable entity response has a 2xx status code
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restart deployment stateless resource unprocessable entity response has a 3xx status code
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart deployment stateless resource unprocessable entity response has a 4xx status code
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this restart deployment stateless resource unprocessable entity response has a 5xx status code
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this restart deployment stateless resource unprocessable entity response a status code equal to that given
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the restart deployment stateless resource unprocessable entity response
func (o *RestartDeploymentStatelessResourceUnprocessableEntity) Code() int {
	return 422
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartDeploymentStatelessResourceInternalServerError creates a RestartDeploymentStatelessResourceInternalServerError with default headers values
func NewRestartDeploymentStatelessResourceInternalServerError() *RestartDeploymentStatelessResourceInternalServerError {
	return &RestartDeploymentStatelessResourceInternalServerError{}
}

/*
RestartDeploymentStatelessResourceInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type RestartDeploymentStatelessResourceInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restart deployment stateless resource internal server error response has a 2xx status code
func (o *RestartDeploymentStatelessResourceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restart deployment stateless resource internal server error response has a 3xx status code
func (o *RestartDeploymentStatelessResourceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restart deployment stateless resource internal server error response has a 4xx status code
func (o *RestartDeploymentStatelessResourceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this restart deployment stateless resource internal server error response has a 5xx status code
func (o *RestartDeploymentStatelessResourceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this restart deployment stateless resource internal server error response a status code equal to that given
func (o *RestartDeploymentStatelessResourceInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the restart deployment stateless resource internal server error response
func (o *RestartDeploymentStatelessResourceInternalServerError) Code() int {
	return 500
}

func (o *RestartDeploymentStatelessResourceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *RestartDeploymentStatelessResourceInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}/_restart][%d] restartDeploymentStatelessResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *RestartDeploymentStatelessResourceInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartDeploymentStatelessResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
