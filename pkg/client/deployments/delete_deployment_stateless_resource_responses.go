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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteDeploymentStatelessResourceReader is a Reader for the DeleteDeploymentStatelessResource structure.
type DeleteDeploymentStatelessResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeploymentStatelessResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDeploymentStatelessResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteDeploymentStatelessResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteDeploymentStatelessResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 449:
		result := NewDeleteDeploymentStatelessResourceRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteDeploymentStatelessResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteDeploymentStatelessResourceOK creates a DeleteDeploymentStatelessResourceOK with default headers values
func NewDeleteDeploymentStatelessResourceOK() *DeleteDeploymentStatelessResourceOK {
	return &DeleteDeploymentStatelessResourceOK{}
}

/*DeleteDeploymentStatelessResourceOK handles this case with default header values.

Standard Deployment Resource Crud Response
*/
type DeleteDeploymentStatelessResourceOK struct {
	Payload *models.DeploymentResourceCrudResponse
}

func (o *DeleteDeploymentStatelessResourceOK) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}][%d] deleteDeploymentStatelessResourceOK  %+v", 200, o.Payload)
}

func (o *DeleteDeploymentStatelessResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeploymentStatelessResourceBadRequest creates a DeleteDeploymentStatelessResourceBadRequest with default headers values
func NewDeleteDeploymentStatelessResourceBadRequest() *DeleteDeploymentStatelessResourceBadRequest {
	return &DeleteDeploymentStatelessResourceBadRequest{}
}

/*DeleteDeploymentStatelessResourceBadRequest handles this case with default header values.

Resource has still instances. (code: `deployments.resource_has_instances`)
*/
type DeleteDeploymentStatelessResourceBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteDeploymentStatelessResourceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}][%d] deleteDeploymentStatelessResourceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteDeploymentStatelessResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeploymentStatelessResourceNotFound creates a DeleteDeploymentStatelessResourceNotFound with default headers values
func NewDeleteDeploymentStatelessResourceNotFound() *DeleteDeploymentStatelessResourceNotFound {
	return &DeleteDeploymentStatelessResourceNotFound{}
}

/*DeleteDeploymentStatelessResourceNotFound handles this case with default header values.

A resource with the given refId cannot be found in the deployment. (code: `deployments.deployment_resource_not_found`)
*/
type DeleteDeploymentStatelessResourceNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteDeploymentStatelessResourceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}][%d] deleteDeploymentStatelessResourceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteDeploymentStatelessResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeploymentStatelessResourceRetryWith creates a DeleteDeploymentStatelessResourceRetryWith with default headers values
func NewDeleteDeploymentStatelessResourceRetryWith() *DeleteDeploymentStatelessResourceRetryWith {
	return &DeleteDeploymentStatelessResourceRetryWith{}
}

/*DeleteDeploymentStatelessResourceRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteDeploymentStatelessResourceRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteDeploymentStatelessResourceRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}][%d] deleteDeploymentStatelessResourceRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteDeploymentStatelessResourceRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeploymentStatelessResourceInternalServerError creates a DeleteDeploymentStatelessResourceInternalServerError with default headers values
func NewDeleteDeploymentStatelessResourceInternalServerError() *DeleteDeploymentStatelessResourceInternalServerError {
	return &DeleteDeploymentStatelessResourceInternalServerError{}
}

/*DeleteDeploymentStatelessResourceInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type DeleteDeploymentStatelessResourceInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteDeploymentStatelessResourceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /deployments/{deployment_id}/{stateless_resource_kind}/{ref_id}][%d] deleteDeploymentStatelessResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteDeploymentStatelessResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
