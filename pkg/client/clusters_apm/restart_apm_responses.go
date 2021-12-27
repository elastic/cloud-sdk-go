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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// RestartApmReader is a Reader for the RestartApm structure.
type RestartApmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartApmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewRestartApmAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRestartApmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewRestartApmPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewRestartApmRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRestartApmAccepted creates a RestartApmAccepted with default headers values
func NewRestartApmAccepted() *RestartApmAccepted {
	return &RestartApmAccepted{}
}

/* RestartApmAccepted describes a response with status code 202, with default header values.

The stop command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type RestartApmAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *RestartApmAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_restart][%d] restartApmAccepted  %+v", 202, o.Payload)
}
func (o *RestartApmAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *RestartApmAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartApmNotFound creates a RestartApmNotFound with default headers values
func NewRestartApmNotFound() *RestartApmNotFound {
	return &RestartApmNotFound{}
}

/* RestartApmNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type RestartApmNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartApmNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_restart][%d] restartApmNotFound  %+v", 404, o.Payload)
}
func (o *RestartApmNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartApmNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartApmPreconditionFailed creates a RestartApmPreconditionFailed with default headers values
func NewRestartApmPreconditionFailed() *RestartApmPreconditionFailed {
	return &RestartApmPreconditionFailed{}
}

/* RestartApmPreconditionFailed describes a response with status code 412, with default header values.

The command sent to a cluster found the cluster in an illegal state, the error message gives more details. (code: `clusters.cluster_plan_state_error`)
*/
type RestartApmPreconditionFailed struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartApmPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_restart][%d] restartApmPreconditionFailed  %+v", 412, o.Payload)
}
func (o *RestartApmPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartApmPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRestartApmRetryWith creates a RestartApmRetryWith with default headers values
func NewRestartApmRetryWith() *RestartApmRetryWith {
	return &RestartApmRetryWith{}
}

/* RestartApmRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type RestartApmRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *RestartApmRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_restart][%d] restartApmRetryWith  %+v", 449, o.Payload)
}
func (o *RestartApmRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *RestartApmRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
