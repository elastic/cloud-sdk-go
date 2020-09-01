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

// StartDeploymentResourceInstancesAllReader is a Reader for the StartDeploymentResourceInstancesAll structure.
type StartDeploymentResourceInstancesAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartDeploymentResourceInstancesAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartDeploymentResourceInstancesAllAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartDeploymentResourceInstancesAllForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartDeploymentResourceInstancesAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartDeploymentResourceInstancesAllRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStartDeploymentResourceInstancesAllInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartDeploymentResourceInstancesAllAccepted creates a StartDeploymentResourceInstancesAllAccepted with default headers values
func NewStartDeploymentResourceInstancesAllAccepted() *StartDeploymentResourceInstancesAllAccepted {
	return &StartDeploymentResourceInstancesAllAccepted{}
}

/*StartDeploymentResourceInstancesAllAccepted handles this case with default header values.

The start command was issued successfully.
*/
type StartDeploymentResourceInstancesAllAccepted struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *StartDeploymentResourceInstancesAllAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllAccepted  %+v", 202, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllAccepted) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllForbidden creates a StartDeploymentResourceInstancesAllForbidden with default headers values
func NewStartDeploymentResourceInstancesAllForbidden() *StartDeploymentResourceInstancesAllForbidden {
	return &StartDeploymentResourceInstancesAllForbidden{}
}

/*StartDeploymentResourceInstancesAllForbidden handles this case with default header values.

The start maintenance mode command was prohibited for the given Resource. (code: `deployments.instance_update_prohibited_error`)
*/
type StartDeploymentResourceInstancesAllForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllForbidden) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllForbidden  %+v", 403, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllNotFound creates a StartDeploymentResourceInstancesAllNotFound with default headers values
func NewStartDeploymentResourceInstancesAllNotFound() *StartDeploymentResourceInstancesAllNotFound {
	return &StartDeploymentResourceInstancesAllNotFound{}
}

/*StartDeploymentResourceInstancesAllNotFound handles this case with default header values.

* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* One or more instances of the given resource type are missing. (code: `deployments.instances_missing_on_update_error`)
 */
type StartDeploymentResourceInstancesAllNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllNotFound  %+v", 404, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllRetryWith creates a StartDeploymentResourceInstancesAllRetryWith with default headers values
func NewStartDeploymentResourceInstancesAllRetryWith() *StartDeploymentResourceInstancesAllRetryWith {
	return &StartDeploymentResourceInstancesAllRetryWith{}
}

/*StartDeploymentResourceInstancesAllRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartDeploymentResourceInstancesAllRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllRetryWith  %+v", 449, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartDeploymentResourceInstancesAllInternalServerError creates a StartDeploymentResourceInstancesAllInternalServerError with default headers values
func NewStartDeploymentResourceInstancesAllInternalServerError() *StartDeploymentResourceInstancesAllInternalServerError {
	return &StartDeploymentResourceInstancesAllInternalServerError{}
}

/*StartDeploymentResourceInstancesAllInternalServerError handles this case with default header values.

A Resource that was previously stored no longer exists. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type StartDeploymentResourceInstancesAllInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartDeploymentResourceInstancesAllInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/{resource_kind}/{ref_id}/instances/_start][%d] startDeploymentResourceInstancesAllInternalServerError  %+v", 500, o.Payload)
}

func (o *StartDeploymentResourceInstancesAllInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartDeploymentResourceInstancesAllInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
