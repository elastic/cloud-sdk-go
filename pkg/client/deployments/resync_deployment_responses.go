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

// ResyncDeploymentReader is a Reader for the ResyncDeployment structure.
type ResyncDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResyncDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewResyncDeploymentRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewResyncDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResyncDeploymentOK creates a ResyncDeploymentOK with default headers values
func NewResyncDeploymentOK() *ResyncDeploymentOK {
	return &ResyncDeploymentOK{}
}

/*ResyncDeploymentOK handles this case with default header values.

The deployment resync operation executed successfully.
*/
type ResyncDeploymentOK struct {
	Payload *models.IndexSynchronizationResults
}

func (o *ResyncDeploymentOK) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_resync][%d] resyncDeploymentOK  %+v", 200, o.Payload)
}

func (o *ResyncDeploymentOK) GetPayload() *models.IndexSynchronizationResults {
	return o.Payload
}

func (o *ResyncDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IndexSynchronizationResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncDeploymentRetryWith creates a ResyncDeploymentRetryWith with default headers values
func NewResyncDeploymentRetryWith() *ResyncDeploymentRetryWith {
	return &ResyncDeploymentRetryWith{}
}

/*ResyncDeploymentRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncDeploymentRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncDeploymentRetryWith) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_resync][%d] resyncDeploymentRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncDeploymentRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncDeploymentRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncDeploymentInternalServerError creates a ResyncDeploymentInternalServerError with default headers values
func NewResyncDeploymentInternalServerError() *ResyncDeploymentInternalServerError {
	return &ResyncDeploymentInternalServerError{}
}

/*ResyncDeploymentInternalServerError handles this case with default header values.

The deployment resync operation failed for deployment {deployment_id}. (code: `deployments.resync_failed`)
*/
type ResyncDeploymentInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/{deployment_id}/_resync][%d] resyncDeploymentInternalServerError  %+v", 500, o.Payload)
}

func (o *ResyncDeploymentInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
