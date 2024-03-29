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

// RestoreDeploymentReader is a Reader for the RestoreDeployment structure.
type RestoreDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestoreDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestoreDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestoreDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRestoreDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRestoreDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestoreDeploymentOK creates a RestoreDeploymentOK with default headers values
func NewRestoreDeploymentOK() *RestoreDeploymentOK {
	return &RestoreDeploymentOK{}
}

/*
RestoreDeploymentOK describes a response with status code 200, with default header values.

The request was valid and the resources of the deployment were restored.
*/
type RestoreDeploymentOK struct {
	Payload *models.DeploymentRestoreResponse
}

// IsSuccess returns true when this restore deployment o k response has a 2xx status code
func (o *RestoreDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this restore deployment o k response has a 3xx status code
func (o *RestoreDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deployment o k response has a 4xx status code
func (o *RestoreDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this restore deployment o k response has a 5xx status code
func (o *RestoreDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deployment o k response a status code equal to that given
func (o *RestoreDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the restore deployment o k response
func (o *RestoreDeploymentOK) Code() int {
	return 200
}

func (o *RestoreDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentOK  %+v", 200, o.Payload)
}

func (o *RestoreDeploymentOK) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentOK  %+v", 200, o.Payload)
}

func (o *RestoreDeploymentOK) GetPayload() *models.DeploymentRestoreResponse {
	return o.Payload
}

func (o *RestoreDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentRestoreResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentBadRequest creates a RestoreDeploymentBadRequest with default headers values
func NewRestoreDeploymentBadRequest() *RestoreDeploymentBadRequest {
	return &RestoreDeploymentBadRequest{}
}

/*
RestoreDeploymentBadRequest describes a response with status code 400, with default header values.

There are Elasticsearch resources in the deployment which are not shut down.
*/
type RestoreDeploymentBadRequest struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restore deployment bad request response has a 2xx status code
func (o *RestoreDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deployment bad request response has a 3xx status code
func (o *RestoreDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deployment bad request response has a 4xx status code
func (o *RestoreDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deployment bad request response has a 5xx status code
func (o *RestoreDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deployment bad request response a status code equal to that given
func (o *RestoreDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the restore deployment bad request response
func (o *RestoreDeploymentBadRequest) Code() int {
	return 400
}

func (o *RestoreDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *RestoreDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *RestoreDeploymentBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentUnauthorized creates a RestoreDeploymentUnauthorized with default headers values
func NewRestoreDeploymentUnauthorized() *RestoreDeploymentUnauthorized {
	return &RestoreDeploymentUnauthorized{}
}

/*
RestoreDeploymentUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action.
*/
type RestoreDeploymentUnauthorized struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restore deployment unauthorized response has a 2xx status code
func (o *RestoreDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deployment unauthorized response has a 3xx status code
func (o *RestoreDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deployment unauthorized response has a 4xx status code
func (o *RestoreDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deployment unauthorized response has a 5xx status code
func (o *RestoreDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deployment unauthorized response a status code equal to that given
func (o *RestoreDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the restore deployment unauthorized response
func (o *RestoreDeploymentUnauthorized) Code() int {
	return 401
}

func (o *RestoreDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *RestoreDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *RestoreDeploymentUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestoreDeploymentNotFound creates a RestoreDeploymentNotFound with default headers values
func NewRestoreDeploymentNotFound() *RestoreDeploymentNotFound {
	return &RestoreDeploymentNotFound{}
}

/*
RestoreDeploymentNotFound describes a response with status code 404, with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type RestoreDeploymentNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this restore deployment not found response has a 2xx status code
func (o *RestoreDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this restore deployment not found response has a 3xx status code
func (o *RestoreDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this restore deployment not found response has a 4xx status code
func (o *RestoreDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this restore deployment not found response has a 5xx status code
func (o *RestoreDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this restore deployment not found response a status code equal to that given
func (o *RestoreDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the restore deployment not found response
func (o *RestoreDeploymentNotFound) Code() int {
	return 404
}

func (o *RestoreDeploymentNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *RestoreDeploymentNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_restore][%d] restoreDeploymentNotFound  %+v", 404, o.Payload)
}

func (o *RestoreDeploymentNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestoreDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
