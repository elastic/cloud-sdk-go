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

// GetDeploymentEsResourceEligibleRemoteClustersReader is a Reader for the GetDeploymentEsResourceEligibleRemoteClusters structure.
type GetDeploymentEsResourceEligibleRemoteClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentEsResourceEligibleRemoteClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentEsResourceEligibleRemoteClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDeploymentEsResourceEligibleRemoteClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetDeploymentEsResourceEligibleRemoteClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentEsResourceEligibleRemoteClustersOK creates a GetDeploymentEsResourceEligibleRemoteClustersOK with default headers values
func NewGetDeploymentEsResourceEligibleRemoteClustersOK() *GetDeploymentEsResourceEligibleRemoteClustersOK {
	return &GetDeploymentEsResourceEligibleRemoteClustersOK{}
}

/*
GetDeploymentEsResourceEligibleRemoteClustersOK describes a response with status code 200, with default header values.

List of deployments which contains eligible remote clusters for the resource
*/
type GetDeploymentEsResourceEligibleRemoteClustersOK struct {
	Payload *models.DeploymentsSearchResponse
}

// IsSuccess returns true when this get deployment es resource eligible remote clusters o k response has a 2xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get deployment es resource eligible remote clusters o k response has a 3xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource eligible remote clusters o k response has a 4xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get deployment es resource eligible remote clusters o k response has a 5xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource eligible remote clusters o k response a status code equal to that given
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get deployment es resource eligible remote clusters o k response
func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) Code() int {
	return 200
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) GetPayload() *models.DeploymentsSearchResponse {
	return o.Payload
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentsSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentEsResourceEligibleRemoteClustersBadRequest creates a GetDeploymentEsResourceEligibleRemoteClustersBadRequest with default headers values
func NewGetDeploymentEsResourceEligibleRemoteClustersBadRequest() *GetDeploymentEsResourceEligibleRemoteClustersBadRequest {
	return &GetDeploymentEsResourceEligibleRemoteClustersBadRequest{}
}

/*
GetDeploymentEsResourceEligibleRemoteClustersBadRequest describes a response with status code 400, with default header values.

The resource specified doesn't yet have a valid version for the current plan or the pending plan. (code: `deployments.resource_does_not_have_a_valid_version`)
*/
type GetDeploymentEsResourceEligibleRemoteClustersBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource eligible remote clusters bad request response has a 2xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource eligible remote clusters bad request response has a 3xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource eligible remote clusters bad request response has a 4xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment es resource eligible remote clusters bad request response has a 5xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource eligible remote clusters bad request response a status code equal to that given
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get deployment es resource eligible remote clusters bad request response
func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) Code() int {
	return 400
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersBadRequest  %+v", 400, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDeploymentEsResourceEligibleRemoteClustersNotFound creates a GetDeploymentEsResourceEligibleRemoteClustersNotFound with default headers values
func NewGetDeploymentEsResourceEligibleRemoteClustersNotFound() *GetDeploymentEsResourceEligibleRemoteClustersNotFound {
	return &GetDeploymentEsResourceEligibleRemoteClustersNotFound{}
}

/*
	GetDeploymentEsResourceEligibleRemoteClustersNotFound describes a response with status code 404, with default header values.

	* The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)

* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type GetDeploymentEsResourceEligibleRemoteClustersNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get deployment es resource eligible remote clusters not found response has a 2xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get deployment es resource eligible remote clusters not found response has a 3xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get deployment es resource eligible remote clusters not found response has a 4xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get deployment es resource eligible remote clusters not found response has a 5xx status code
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get deployment es resource eligible remote clusters not found response a status code equal to that given
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get deployment es resource eligible remote clusters not found response
func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) Code() int {
	return 404
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/elasticsearch/{ref_id}/eligible-remote-clusters][%d] getDeploymentEsResourceEligibleRemoteClustersNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentEsResourceEligibleRemoteClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
