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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// MoveClustersReader is a Reader for the MoveClusters structure.
type MoveClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMoveClustersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMoveClustersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewMoveClustersRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMoveClustersAccepted creates a MoveClustersAccepted with default headers values
func NewMoveClustersAccepted() *MoveClustersAccepted {
	return &MoveClustersAccepted{}
}

/*
MoveClustersAccepted describes a response with status code 202, with default header values.

The move command was issued successfully, use the "GET" command on each /{cluster_id} resource to monitor progress
*/
type MoveClustersAccepted struct {
	Payload *models.MoveClustersCommandResponse
}

// IsSuccess returns true when this move clusters accepted response has a 2xx status code
func (o *MoveClustersAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this move clusters accepted response has a 3xx status code
func (o *MoveClustersAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move clusters accepted response has a 4xx status code
func (o *MoveClustersAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this move clusters accepted response has a 5xx status code
func (o *MoveClustersAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this move clusters accepted response a status code equal to that given
func (o *MoveClustersAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the move clusters accepted response
func (o *MoveClustersAccepted) Code() int {
	return 202
}

func (o *MoveClustersAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersAccepted  %+v", 202, o.Payload)
}

func (o *MoveClustersAccepted) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersAccepted  %+v", 202, o.Payload)
}

func (o *MoveClustersAccepted) GetPayload() *models.MoveClustersCommandResponse {
	return o.Payload
}

func (o *MoveClustersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MoveClustersCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveClustersBadRequest creates a MoveClustersBadRequest with default headers values
func NewMoveClustersBadRequest() *MoveClustersBadRequest {
	return &MoveClustersBadRequest{}
}

/*
	MoveClustersBadRequest describes a response with status code 400, with default header values.

	* The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)

* The cluster definition contained errors. (code: `clusters.plan_feature_not_implemented`)
*/
type MoveClustersBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this move clusters bad request response has a 2xx status code
func (o *MoveClustersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move clusters bad request response has a 3xx status code
func (o *MoveClustersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move clusters bad request response has a 4xx status code
func (o *MoveClustersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this move clusters bad request response has a 5xx status code
func (o *MoveClustersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this move clusters bad request response a status code equal to that given
func (o *MoveClustersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the move clusters bad request response
func (o *MoveClustersBadRequest) Code() int {
	return 400
}

func (o *MoveClustersBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersBadRequest  %+v", 400, o.Payload)
}

func (o *MoveClustersBadRequest) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersBadRequest  %+v", 400, o.Payload)
}

func (o *MoveClustersBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveClustersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveClustersForbidden creates a MoveClustersForbidden with default headers values
func NewMoveClustersForbidden() *MoveClustersForbidden {
	return &MoveClustersForbidden{}
}

/*
MoveClustersForbidden describes a response with status code 403, with default header values.

The move command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type MoveClustersForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this move clusters forbidden response has a 2xx status code
func (o *MoveClustersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move clusters forbidden response has a 3xx status code
func (o *MoveClustersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move clusters forbidden response has a 4xx status code
func (o *MoveClustersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this move clusters forbidden response has a 5xx status code
func (o *MoveClustersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this move clusters forbidden response a status code equal to that given
func (o *MoveClustersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the move clusters forbidden response
func (o *MoveClustersForbidden) Code() int {
	return 403
}

func (o *MoveClustersForbidden) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersForbidden  %+v", 403, o.Payload)
}

func (o *MoveClustersForbidden) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersForbidden  %+v", 403, o.Payload)
}

func (o *MoveClustersForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewMoveClustersRetryWith creates a MoveClustersRetryWith with default headers values
func NewMoveClustersRetryWith() *MoveClustersRetryWith {
	return &MoveClustersRetryWith{}
}

/*
MoveClustersRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type MoveClustersRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this move clusters retry with response has a 2xx status code
func (o *MoveClustersRetryWith) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this move clusters retry with response has a 3xx status code
func (o *MoveClustersRetryWith) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move clusters retry with response has a 4xx status code
func (o *MoveClustersRetryWith) IsClientError() bool {
	return true
}

// IsServerError returns true when this move clusters retry with response has a 5xx status code
func (o *MoveClustersRetryWith) IsServerError() bool {
	return false
}

// IsCode returns true when this move clusters retry with response a status code equal to that given
func (o *MoveClustersRetryWith) IsCode(code int) bool {
	return code == 449
}

// Code gets the status code for the move clusters retry with response
func (o *MoveClustersRetryWith) Code() int {
	return 449
}

func (o *MoveClustersRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersRetryWith  %+v", 449, o.Payload)
}

func (o *MoveClustersRetryWith) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/{allocator_id}/clusters/_move][%d] moveClustersRetryWith  %+v", 449, o.Payload)
}

func (o *MoveClustersRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveClustersRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
