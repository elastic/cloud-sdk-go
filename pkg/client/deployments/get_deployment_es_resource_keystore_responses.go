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

// GetDeploymentEsResourceKeystoreReader is a Reader for the GetDeploymentEsResourceKeystore structure.
type GetDeploymentEsResourceKeystoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEsResourceKeystoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEsResourceKeystoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentEsResourceKeystoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentEsResourceKeystoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentEsResourceKeystoreOK creates a GetDeploymentEsResourceKeystoreOK with default headers values
func NewGetDeploymentEsResourceKeystoreOK() *GetDeploymentEsResourceKeystoreOK {
	return &GetDeploymentEsResourceKeystoreOK{}
}

/*
GetDeploymentEsResourceKeystoreOK describes a response with status code 200, with default header values.

The contents of the Elasticsearch keystore
*/
type GetDeploymentEsResourceKeystoreOK struct {
	Payload *models.KeystoreContents
}

// IsSuccess returns true when this get deployment es resource keystore o k response has a 2xx status code
func (o *GetDeploymentEsResourceKeystoreOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment es resource keystore o k response has a 3xx status code
func (o *GetDeploymentEsResourceKeystoreOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource keystore o k response has a 4xx status code
func (o *GetDeploymentEsResourceKeystoreOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment es resource keystore o k response has a 5xx status code
func (o *GetDeploymentEsResourceKeystoreOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource keystore o k response a status code equal to that given
func (o *GetDeploymentEsResourceKeystoreOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment es resource keystore o k response
func (o *GetDeploymentEsResourceKeystoreOK) Code() int {
	return 200
}

func (o *GetDeploymentEsResourceKeystoreOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreOK) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreOK) GetPayload() *models.KeystoreContents {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KeystoreContents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEsResourceKeystoreNotFound creates a GetDeploymentEsResourceKeystoreNotFound with default headers values
func NewGetDeploymentEsResourceKeystoreNotFound() *GetDeploymentEsResourceKeystoreNotFound {
	return &GetDeploymentEsResourceKeystoreNotFound{}
}

/*
	GetDeploymentEsResourceKeystoreNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type GetDeploymentEsResourceKeystoreNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource keystore not found response has a 2xx status code
func (o *GetDeploymentEsResourceKeystoreNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource keystore not found response has a 3xx status code
func (o *GetDeploymentEsResourceKeystoreNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource keystore not found response has a 4xx status code
func (o *GetDeploymentEsResourceKeystoreNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment es resource keystore not found response has a 5xx status code
func (o *GetDeploymentEsResourceKeystoreNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource keystore not found response a status code equal to that given
func (o *GetDeploymentEsResourceKeystoreNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment es resource keystore not found response
func (o *GetDeploymentEsResourceKeystoreNotFound) Code() int {
	return 404
}

func (o *GetDeploymentEsResourceKeystoreNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentEsResourceKeystoreInternalServerError creates a GetDeploymentEsResourceKeystoreInternalServerError with default headers values
func NewGetDeploymentEsResourceKeystoreInternalServerError() *GetDeploymentEsResourceKeystoreInternalServerError {
	return &GetDeploymentEsResourceKeystoreInternalServerError{}
}

/*
GetDeploymentEsResourceKeystoreInternalServerError describes a response with status code 500, with default header values.

We have failed you. (code: `deployments.metadata_internal_error`)
*/
type GetDeploymentEsResourceKeystoreInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource keystore internal server error response has a 2xx status code
func (o *GetDeploymentEsResourceKeystoreInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource keystore internal server error response has a 3xx status code
func (o *GetDeploymentEsResourceKeystoreInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource keystore internal server error response has a 4xx status code
func (o *GetDeploymentEsResourceKeystoreInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment es resource keystore internal server error response has a 5xx status code
func (o *GetDeploymentEsResourceKeystoreInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get deployment es resource keystore internal server error response a status code equal to that given
func (o *GetDeploymentEsResourceKeystoreInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get deployment es resource keystore internal server error response
func (o *GetDeploymentEsResourceKeystoreInternalServerError) Code() int {
	return 500
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/elasticsearch/{ref_id}/keystore][%d] getDeploymentEsResourceKeystoreInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceKeystoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
