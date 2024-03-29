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

// StartDeploymentResourceInstancesAllMaintenanceModeReader is a Reader for the StartDeploymentResourceInstancesAllMaintenanceMode structure.
type StartDeploymentResourceInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDeploymentResourceInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartDeploymentResourceInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartDeploymentResourceInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartDeploymentResourceInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartDeploymentResourceInstancesAllMaintenanceModeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartDeploymentResourceInstancesAllMaintenanceModeAccepted creates a StartDeploymentResourceInstancesAllMaintenanceModeAccepted with default headers values
func NewStartDeploymentResourceInstancesAllMaintenanceModeAccepted() *StartDeploymentResourceInstancesAllMaintenanceModeAccepted {
	return &StartDeploymentResourceInstancesAllMaintenanceModeAccepted{}
}

/*
StartDeploymentResourceInstancesAllMaintenanceModeAccepted describes a response with status code 202, with default header values.

The start maintenance command was issued successfully.
*/
type StartDeploymentResourceInstancesAllMaintenanceModeAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

// IsSuccess returns true when this start deployment resource instances all maintenance mode accepted response has a 2xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this start deployment resource instances all maintenance mode accepted response has a 3xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start deployment resource instances all maintenance mode accepted response has a 4xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this start deployment resource instances all maintenance mode accepted response has a 5xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this start deployment resource instances all maintenance mode accepted response a status code equal to that given
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the start deployment resource instances all maintenance mode accepted response
func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) Code() int {
	return 202
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllMaintenanceModeForbidden creates a StartDeploymentResourceInstancesAllMaintenanceModeForbidden with default headers values
func NewStartDeploymentResourceInstancesAllMaintenanceModeForbidden() *StartDeploymentResourceInstancesAllMaintenanceModeForbidden {
	return &StartDeploymentResourceInstancesAllMaintenanceModeForbidden{}
}

/*
StartDeploymentResourceInstancesAllMaintenanceModeForbidden describes a response with status code 403, with default header values.

The start maintenance mode command was prohibited for the given Resource. (code: `deployments.instance_update_prohibited_error`)
*/
type StartDeploymentResourceInstancesAllMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this start deployment resource instances all maintenance mode forbidden response has a 2xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start deployment resource instances all maintenance mode forbidden response has a 3xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start deployment resource instances all maintenance mode forbidden response has a 4xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this start deployment resource instances all maintenance mode forbidden response has a 5xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this start deployment resource instances all maintenance mode forbidden response a status code equal to that given
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the start deployment resource instances all maintenance mode forbidden response
func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) Code() int {
	return 403
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartDeploymentResourceInstancesAllMaintenanceModeNotFound creates a StartDeploymentResourceInstancesAllMaintenanceModeNotFound with default headers values
func NewStartDeploymentResourceInstancesAllMaintenanceModeNotFound() *StartDeploymentResourceInstancesAllMaintenanceModeNotFound {
	return &StartDeploymentResourceInstancesAllMaintenanceModeNotFound{}
}

/*
	StartDeploymentResourceInstancesAllMaintenanceModeNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* One or more instances of the given resource type are missing. (code: `deployments.instances_missing_on_update_error`)
*/
type StartDeploymentResourceInstancesAllMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this start deployment resource instances all maintenance mode not found response has a 2xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start deployment resource instances all maintenance mode not found response has a 3xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start deployment resource instances all maintenance mode not found response has a 4xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this start deployment resource instances all maintenance mode not found response has a 5xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this start deployment resource instances all maintenance mode not found response a status code equal to that given
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the start deployment resource instances all maintenance mode not found response
func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) Code() int {
	return 404
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartDeploymentResourceInstancesAllMaintenanceModeInternalServerError creates a StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError with default headers values
func NewStartDeploymentResourceInstancesAllMaintenanceModeInternalServerError() *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError {
	return &StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError{}
}

/*
StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError describes a response with status code 500, with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this start deployment resource instances all maintenance mode internal server error response has a 2xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this start deployment resource instances all maintenance mode internal server error response has a 3xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this start deployment resource instances all maintenance mode internal server error response has a 4xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this start deployment resource instances all maintenance mode internal server error response has a 5xx status code
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this start deployment resource instances all maintenance mode internal server error response a status code equal to that given
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the start deployment resource instances all maintenance mode internal server error response
func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) Code() int {
	return 500
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeInternalServerError  %+v", 500, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/maintenance-mode/_start][%d] startDeploymentResourceInstancesAllMaintenanceModeInternalServerError  %+v", 500, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllMaintenanceModeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
