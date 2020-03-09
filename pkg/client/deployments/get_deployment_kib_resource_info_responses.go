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

// GetDeploymentKibResourceInfoReader is a Reader for the GetDeploymentKibResourceInfo structure.
type GetDeploymentKibResourceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentKibResourceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentKibResourceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentKibResourceInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentKibResourceInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentKibResourceInfoOK creates a GetDeploymentKibResourceInfoOK with default headers values
func NewGetDeploymentKibResourceInfoOK() *GetDeploymentKibResourceInfoOK {
	return &GetDeploymentKibResourceInfoOK{}
}

/*GetDeploymentKibResourceInfoOK handles this case with default header values.

Standard response.
*/
type GetDeploymentKibResourceInfoOK struct {
	Payload *models.KibanaResourceInfo
}

func (o *GetDeploymentKibResourceInfoOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/kibana/{ref_id}][%d] getDeploymentKibResourceInfoOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentKibResourceInfoOK) GetPayload() *models.KibanaResourceInfo {
	return o.Payload
}

func (o *GetDeploymentKibResourceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KibanaResourceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentKibResourceInfoNotFound creates a GetDeploymentKibResourceInfoNotFound with default headers values
func NewGetDeploymentKibResourceInfoNotFound() *GetDeploymentKibResourceInfoNotFound {
	return &GetDeploymentKibResourceInfoNotFound{}
}

/*GetDeploymentKibResourceInfoNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type GetDeploymentKibResourceInfoNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentKibResourceInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/kibana/{ref_id}][%d] getDeploymentKibResourceInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentKibResourceInfoNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentKibResourceInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentKibResourceInfoInternalServerError creates a GetDeploymentKibResourceInfoInternalServerError with default headers values
func NewGetDeploymentKibResourceInfoInternalServerError() *GetDeploymentKibResourceInfoInternalServerError {
	return &GetDeploymentKibResourceInfoInternalServerError{}
}

/*GetDeploymentKibResourceInfoInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type GetDeploymentKibResourceInfoInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentKibResourceInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/kibana/{ref_id}][%d] getDeploymentKibResourceInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentKibResourceInfoInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentKibResourceInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
