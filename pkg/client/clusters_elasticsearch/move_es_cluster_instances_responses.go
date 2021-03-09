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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// MoveEsClusterInstancesReader is a Reader for the MoveEsClusterInstances structure.
type MoveEsClusterInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveEsClusterInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMoveEsClusterInstancesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMoveEsClusterInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveEsClusterInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMoveEsClusterInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewMoveEsClusterInstancesRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMoveEsClusterInstancesAccepted creates a MoveEsClusterInstancesAccepted with default headers values
func NewMoveEsClusterInstancesAccepted() *MoveEsClusterInstancesAccepted {
	return &MoveEsClusterInstancesAccepted{}
}

/* MoveEsClusterInstancesAccepted describes a response with status code 202, with default header values.

The move command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type MoveEsClusterInstancesAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *MoveEsClusterInstancesAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/_move][%d] moveEsClusterInstancesAccepted  %+v", 202, o.Payload)
}
func (o *MoveEsClusterInstancesAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *MoveEsClusterInstancesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveEsClusterInstancesBadRequest creates a MoveEsClusterInstancesBadRequest with default headers values
func NewMoveEsClusterInstancesBadRequest() *MoveEsClusterInstancesBadRequest {
	return &MoveEsClusterInstancesBadRequest{}
}

/* MoveEsClusterInstancesBadRequest describes a response with status code 400, with default header values.

 * The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)
* The cluster definition contained errors. (code: `clusters.plan_feature_not_implemented`)
*/
type MoveEsClusterInstancesBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/_move][%d] moveEsClusterInstancesBadRequest  %+v", 400, o.Payload)
}
func (o *MoveEsClusterInstancesBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveEsClusterInstancesForbidden creates a MoveEsClusterInstancesForbidden with default headers values
func NewMoveEsClusterInstancesForbidden() *MoveEsClusterInstancesForbidden {
	return &MoveEsClusterInstancesForbidden{}
}

/* MoveEsClusterInstancesForbidden describes a response with status code 403, with default header values.

The move command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type MoveEsClusterInstancesForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/_move][%d] moveEsClusterInstancesForbidden  %+v", 403, o.Payload)
}
func (o *MoveEsClusterInstancesForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveEsClusterInstancesNotFound creates a MoveEsClusterInstancesNotFound with default headers values
func NewMoveEsClusterInstancesNotFound() *MoveEsClusterInstancesNotFound {
	return &MoveEsClusterInstancesNotFound{}
}

/* MoveEsClusterInstancesNotFound describes a response with status code 404, with default header values.

 * The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
* One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
*/
type MoveEsClusterInstancesNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/_move][%d] moveEsClusterInstancesNotFound  %+v", 404, o.Payload)
}
func (o *MoveEsClusterInstancesNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveEsClusterInstancesRetryWith creates a MoveEsClusterInstancesRetryWith with default headers values
func NewMoveEsClusterInstancesRetryWith() *MoveEsClusterInstancesRetryWith {
	return &MoveEsClusterInstancesRetryWith{}
}

/* MoveEsClusterInstancesRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type MoveEsClusterInstancesRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/_move][%d] moveEsClusterInstancesRetryWith  %+v", 449, o.Payload)
}
func (o *MoveEsClusterInstancesRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
