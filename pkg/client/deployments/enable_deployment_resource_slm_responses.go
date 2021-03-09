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

// EnableDeploymentResourceSlmReader is a Reader for the EnableDeploymentResourceSlm structure.
type EnableDeploymentResourceSlmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableDeploymentResourceSlmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnableDeploymentResourceSlmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewEnableDeploymentResourceSlmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewEnableDeploymentResourceSlmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEnableDeploymentResourceSlmInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnableDeploymentResourceSlmOK creates a EnableDeploymentResourceSlmOK with default headers values
func NewEnableDeploymentResourceSlmOK() *EnableDeploymentResourceSlmOK {
	return &EnableDeploymentResourceSlmOK{}
}

/* EnableDeploymentResourceSlmOK describes a response with status code 200, with default header values.

Standard response
*/
type EnableDeploymentResourceSlmOK struct {
	Payload *models.DeploymentResourceCommandResponse
}

func (o *EnableDeploymentResourceSlmOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-slm][%d] enableDeploymentResourceSlmOK  %+v", 200, o.Payload)
}
func (o *EnableDeploymentResourceSlmOK) GetPayload() *models.DeploymentResourceCommandResponse {
	return o.Payload
}

func (o *EnableDeploymentResourceSlmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResourceCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnableDeploymentResourceSlmNotFound creates a EnableDeploymentResourceSlmNotFound with default headers values
func NewEnableDeploymentResourceSlmNotFound() *EnableDeploymentResourceSlmNotFound {
	return &EnableDeploymentResourceSlmNotFound{}
}

/* EnableDeploymentResourceSlmNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type EnableDeploymentResourceSlmNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *EnableDeploymentResourceSlmNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-slm][%d] enableDeploymentResourceSlmNotFound  %+v", 404, o.Payload)
}
func (o *EnableDeploymentResourceSlmNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *EnableDeploymentResourceSlmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEnableDeploymentResourceSlmRetryWith creates a EnableDeploymentResourceSlmRetryWith with default headers values
func NewEnableDeploymentResourceSlmRetryWith() *EnableDeploymentResourceSlmRetryWith {
	return &EnableDeploymentResourceSlmRetryWith{}
}

/* EnableDeploymentResourceSlmRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type EnableDeploymentResourceSlmRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *EnableDeploymentResourceSlmRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-slm][%d] enableDeploymentResourceSlmRetryWith  %+v", 449, o.Payload)
}
func (o *EnableDeploymentResourceSlmRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *EnableDeploymentResourceSlmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewEnableDeploymentResourceSlmInternalServerError creates a EnableDeploymentResourceSlmInternalServerError with default headers values
func NewEnableDeploymentResourceSlmInternalServerError() *EnableDeploymentResourceSlmInternalServerError {
	return &EnableDeploymentResourceSlmInternalServerError{}
}

/* EnableDeploymentResourceSlmInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type EnableDeploymentResourceSlmInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *EnableDeploymentResourceSlmInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/_enable-slm][%d] enableDeploymentResourceSlmInternalServerError  %+v", 500, o.Payload)
}
func (o *EnableDeploymentResourceSlmInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *EnableDeploymentResourceSlmInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
