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

// CreateDeploymentReader is a Reader for the CreateDeployment structure.
type CreateDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDeploymentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateDeploymentOK creates a CreateDeploymentOK with default headers values
func NewCreateDeploymentOK() *CreateDeploymentOK {
	return &CreateDeploymentOK{}
}

/*
CreateDeploymentOK describes a response with status code 200, with default header values.

The request was valid (used when validate_only is true).
*/
type CreateDeploymentOK struct {
	Payload *models.DeploymentCreateResponse
}

// IsSuccess returns true when this create deployment o k response has a 2xx status code
func (o *CreateDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create deployment o k response has a 3xx status code
func (o *CreateDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment o k response has a 4xx status code
func (o *CreateDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create deployment o k response has a 5xx status code
func (o *CreateDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment o k response a status code equal to that given
func (o *CreateDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create deployment o k response
func (o *CreateDeploymentOK) Code() int {
	return 200
}

func (o *CreateDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentOK  %+v", 200, o.Payload)
}

func (o *CreateDeploymentOK) String() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentOK  %+v", 200, o.Payload)
}

func (o *CreateDeploymentOK) GetPayload() *models.DeploymentCreateResponse {
	return o.Payload
}

func (o *CreateDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentCreated creates a CreateDeploymentCreated with default headers values
func NewCreateDeploymentCreated() *CreateDeploymentCreated {
	return &CreateDeploymentCreated{}
}

/*
CreateDeploymentCreated describes a response with status code 201, with default header values.

The request was valid and a new deployment was created
*/
type CreateDeploymentCreated struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.DeploymentCreateResponse
}

// IsSuccess returns true when this create deployment created response has a 2xx status code
func (o *CreateDeploymentCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create deployment created response has a 3xx status code
func (o *CreateDeploymentCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment created response has a 4xx status code
func (o *CreateDeploymentCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create deployment created response has a 5xx status code
func (o *CreateDeploymentCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment created response a status code equal to that given
func (o *CreateDeploymentCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create deployment created response
func (o *CreateDeploymentCreated) Code() int {
	return 201
}

func (o *CreateDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentCreated) String() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentCreated) GetPayload() *models.DeploymentCreateResponse {
	return o.Payload
}

func (o *CreateDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DeploymentCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentAccepted creates a CreateDeploymentAccepted with default headers values
func NewCreateDeploymentAccepted() *CreateDeploymentAccepted {
	return &CreateDeploymentAccepted{}
}

/*
CreateDeploymentAccepted describes a response with status code 202, with default header values.

The request was valid and deployment creation had already been started.
*/
type CreateDeploymentAccepted struct {
	Payload *models.DeploymentCreateResponse
}

// IsSuccess returns true when this create deployment accepted response has a 2xx status code
func (o *CreateDeploymentAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create deployment accepted response has a 3xx status code
func (o *CreateDeploymentAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment accepted response has a 4xx status code
func (o *CreateDeploymentAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this create deployment accepted response has a 5xx status code
func (o *CreateDeploymentAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment accepted response a status code equal to that given
func (o *CreateDeploymentAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the create deployment accepted response
func (o *CreateDeploymentAccepted) Code() int {
	return 202
}

func (o *CreateDeploymentAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentAccepted  %+v", 202, o.Payload)
}

func (o *CreateDeploymentAccepted) String() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentAccepted  %+v", 202, o.Payload)
}

func (o *CreateDeploymentAccepted) GetPayload() *models.DeploymentCreateResponse {
	return o.Payload
}

func (o *CreateDeploymentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentBadRequest creates a CreateDeploymentBadRequest with default headers values
func NewCreateDeploymentBadRequest() *CreateDeploymentBadRequest {
	return &CreateDeploymentBadRequest{}
}

/*
CreateDeploymentBadRequest describes a response with status code 400, with default header values.

The deployment request had errors.
*/
type CreateDeploymentBadRequest struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create deployment bad request response has a 2xx status code
func (o *CreateDeploymentBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create deployment bad request response has a 3xx status code
func (o *CreateDeploymentBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment bad request response has a 4xx status code
func (o *CreateDeploymentBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create deployment bad request response has a 5xx status code
func (o *CreateDeploymentBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment bad request response a status code equal to that given
func (o *CreateDeploymentBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create deployment bad request response
func (o *CreateDeploymentBadRequest) Code() int {
	return 400
}

func (o *CreateDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeploymentBadRequest) String() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeploymentBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentUnauthorized creates a CreateDeploymentUnauthorized with default headers values
func NewCreateDeploymentUnauthorized() *CreateDeploymentUnauthorized {
	return &CreateDeploymentUnauthorized{}
}

/*
CreateDeploymentUnauthorized describes a response with status code 401, with default header values.

You are not authorized to perform this action.
*/
type CreateDeploymentUnauthorized struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create deployment unauthorized response has a 2xx status code
func (o *CreateDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create deployment unauthorized response has a 3xx status code
func (o *CreateDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create deployment unauthorized response has a 4xx status code
func (o *CreateDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create deployment unauthorized response has a 5xx status code
func (o *CreateDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create deployment unauthorized response a status code equal to that given
func (o *CreateDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create deployment unauthorized response
func (o *CreateDeploymentUnauthorized) Code() int {
	return 401
}

func (o *CreateDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDeploymentUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
