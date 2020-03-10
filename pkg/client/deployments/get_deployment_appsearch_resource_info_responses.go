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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentAppsearchResourceInfoReader is a Reader for the GetDeploymentAppsearchResourceInfo structure.
type GetDeploymentAppsearchResourceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentAppsearchResourceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentAppsearchResourceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentAppsearchResourceInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentAppsearchResourceInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentAppsearchResourceInfoOK creates a GetDeploymentAppsearchResourceInfoOK with default headers values
func NewGetDeploymentAppsearchResourceInfoOK() *GetDeploymentAppsearchResourceInfoOK {
	return &GetDeploymentAppsearchResourceInfoOK{}
}

/*GetDeploymentAppsearchResourceInfoOK handles this case with default header values.

Standard response.
*/
type GetDeploymentAppsearchResourceInfoOK struct {
	Payload *models.AppSearchResourceInfo
}

func (o *GetDeploymentAppsearchResourceInfoOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/appsearch/{ref_id}][%d] getDeploymentAppsearchResourceInfoOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentAppsearchResourceInfoOK) GetPayload() *models.AppSearchResourceInfo {
	return o.Payload
}

func (o *GetDeploymentAppsearchResourceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppSearchResourceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentAppsearchResourceInfoNotFound creates a GetDeploymentAppsearchResourceInfoNotFound with default headers values
func NewGetDeploymentAppsearchResourceInfoNotFound() *GetDeploymentAppsearchResourceInfoNotFound {
	return &GetDeploymentAppsearchResourceInfoNotFound{}
}

/*GetDeploymentAppsearchResourceInfoNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type GetDeploymentAppsearchResourceInfoNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentAppsearchResourceInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/appsearch/{ref_id}][%d] getDeploymentAppsearchResourceInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentAppsearchResourceInfoNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentAppsearchResourceInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentAppsearchResourceInfoInternalServerError creates a GetDeploymentAppsearchResourceInfoInternalServerError with default headers values
func NewGetDeploymentAppsearchResourceInfoInternalServerError() *GetDeploymentAppsearchResourceInfoInternalServerError {
	return &GetDeploymentAppsearchResourceInfoInternalServerError{}
}

/*GetDeploymentAppsearchResourceInfoInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type GetDeploymentAppsearchResourceInfoInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentAppsearchResourceInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/appsearch/{ref_id}][%d] getDeploymentAppsearchResourceInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentAppsearchResourceInfoInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentAppsearchResourceInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
