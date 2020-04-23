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

// GetDeploymentEnterpriseSearchResourceInfoReader is a Reader for the GetDeploymentEnterpriseSearchResourceInfo structure.
type GetDeploymentEnterpriseSearchResourceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEnterpriseSearchResourceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEnterpriseSearchResourceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentEnterpriseSearchResourceInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentEnterpriseSearchResourceInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentEnterpriseSearchResourceInfoOK creates a GetDeploymentEnterpriseSearchResourceInfoOK with default headers values
func NewGetDeploymentEnterpriseSearchResourceInfoOK() *GetDeploymentEnterpriseSearchResourceInfoOK {
	return &GetDeploymentEnterpriseSearchResourceInfoOK{}
}

/*GetDeploymentEnterpriseSearchResourceInfoOK handles this case with default header values.

Standard response.
*/
type GetDeploymentEnterpriseSearchResourceInfoOK struct {
	Payload *models.EnterpriseSearchResourceInfo
}

func (o *GetDeploymentEnterpriseSearchResourceInfoOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/enterprise_search/{ref_id}][%d] getDeploymentEnterpriseSearchResourceInfoOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEnterpriseSearchResourceInfoOK) GetPayload() *models.EnterpriseSearchResourceInfo {
	return o.Payload
}

func (o *GetDeploymentEnterpriseSearchResourceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnterpriseSearchResourceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEnterpriseSearchResourceInfoNotFound creates a GetDeploymentEnterpriseSearchResourceInfoNotFound with default headers values
func NewGetDeploymentEnterpriseSearchResourceInfoNotFound() *GetDeploymentEnterpriseSearchResourceInfoNotFound {
	return &GetDeploymentEnterpriseSearchResourceInfoNotFound{}
}

/*GetDeploymentEnterpriseSearchResourceInfoNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type GetDeploymentEnterpriseSearchResourceInfoNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentEnterpriseSearchResourceInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/enterprise_search/{ref_id}][%d] getDeploymentEnterpriseSearchResourceInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEnterpriseSearchResourceInfoNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEnterpriseSearchResourceInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEnterpriseSearchResourceInfoInternalServerError creates a GetDeploymentEnterpriseSearchResourceInfoInternalServerError with default headers values
func NewGetDeploymentEnterpriseSearchResourceInfoInternalServerError() *GetDeploymentEnterpriseSearchResourceInfoInternalServerError {
	return &GetDeploymentEnterpriseSearchResourceInfoInternalServerError{}
}

/*GetDeploymentEnterpriseSearchResourceInfoInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type GetDeploymentEnterpriseSearchResourceInfoInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentEnterpriseSearchResourceInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/enterprise_search/{ref_id}][%d] getDeploymentEnterpriseSearchResourceInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentEnterpriseSearchResourceInfoInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEnterpriseSearchResourceInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
