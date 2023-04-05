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

// GetDeploymentResourceProxyRequestsReader is a Reader for the GetDeploymentResourceProxyRequests structure.
type GetDeploymentResourceProxyRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentResourceProxyRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentResourceProxyRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentResourceProxyRequestsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewGetDeploymentResourceProxyRequestsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentResourceProxyRequestsOK creates a GetDeploymentResourceProxyRequestsOK with default headers values
func NewGetDeploymentResourceProxyRequestsOK() *GetDeploymentResourceProxyRequestsOK {
	return &GetDeploymentResourceProxyRequestsOK{}
}

/*
GetDeploymentResourceProxyRequestsOK describes a response with status code 200, with default header values.

The request has been processed successfully through the proxy.
*/
type GetDeploymentResourceProxyRequestsOK struct {
	Payload *models.GenericResponse
}

// IsSuccess returns true when this get deployment resource proxy requests o k response has a 2xx status code
func (o *GetDeploymentResourceProxyRequestsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment resource proxy requests o k response has a 3xx status code
func (o *GetDeploymentResourceProxyRequestsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resource proxy requests o k response has a 4xx status code
func (o *GetDeploymentResourceProxyRequestsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment resource proxy requests o k response has a 5xx status code
func (o *GetDeploymentResourceProxyRequestsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resource proxy requests o k response a status code equal to that given
func (o *GetDeploymentResourceProxyRequestsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment resource proxy requests o k response
func (o *GetDeploymentResourceProxyRequestsOK) Code() int {
	return 200
}

func (o *GetDeploymentResourceProxyRequestsOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsOK) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsOK) GetPayload() *models.GenericResponse {
	return o.Payload
}

func (o *GetDeploymentResourceProxyRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentResourceProxyRequestsNotFound creates a GetDeploymentResourceProxyRequestsNotFound with default headers values
func NewGetDeploymentResourceProxyRequestsNotFound() *GetDeploymentResourceProxyRequestsNotFound {
	return &GetDeploymentResourceProxyRequestsNotFound{}
}

/*
	GetDeploymentResourceProxyRequestsNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type GetDeploymentResourceProxyRequestsNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment resource proxy requests not found response has a 2xx status code
func (o *GetDeploymentResourceProxyRequestsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment resource proxy requests not found response has a 3xx status code
func (o *GetDeploymentResourceProxyRequestsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resource proxy requests not found response has a 4xx status code
func (o *GetDeploymentResourceProxyRequestsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment resource proxy requests not found response has a 5xx status code
func (o *GetDeploymentResourceProxyRequestsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resource proxy requests not found response a status code equal to that given
func (o *GetDeploymentResourceProxyRequestsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment resource proxy requests not found response
func (o *GetDeploymentResourceProxyRequestsNotFound) Code() int {
	return 404
}

func (o *GetDeploymentResourceProxyRequestsNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsNotFound) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentResourceProxyRequestsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentResourceProxyRequestsRetryWith creates a GetDeploymentResourceProxyRequestsRetryWith with default headers values
func NewGetDeploymentResourceProxyRequestsRetryWith() *GetDeploymentResourceProxyRequestsRetryWith {
	return &GetDeploymentResourceProxyRequestsRetryWith{}
}

/*
GetDeploymentResourceProxyRequestsRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type GetDeploymentResourceProxyRequestsRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment resource proxy requests retry with response has a 2xx status code
func (o *GetDeploymentResourceProxyRequestsRetryWith) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment resource proxy requests retry with response has a 3xx status code
func (o *GetDeploymentResourceProxyRequestsRetryWith) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment resource proxy requests retry with response has a 4xx status code
func (o *GetDeploymentResourceProxyRequestsRetryWith) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment resource proxy requests retry with response has a 5xx status code
func (o *GetDeploymentResourceProxyRequestsRetryWith) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment resource proxy requests retry with response a status code equal to that given
func (o *GetDeploymentResourceProxyRequestsRetryWith) IsCode(code int) bool {
	return code == 449
}

// Code gets the status code for the get deployment resource proxy requests retry with response
func (o *GetDeploymentResourceProxyRequestsRetryWith) Code() int {
	return 449
}

func (o *GetDeploymentResourceProxyRequestsRetryWith) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsRetryWith  %+v", 449, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsRetryWith) String() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/proxy/{proxy_path}][%d] getDeploymentResourceProxyRequestsRetryWith  %+v", 449, o.Payload)
}

func (o *GetDeploymentResourceProxyRequestsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentResourceProxyRequestsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
