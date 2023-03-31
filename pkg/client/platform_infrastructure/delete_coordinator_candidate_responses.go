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

// DeleteCoordinatorCandidateReader is a Reader for the DeleteCoordinatorCandidate structure.
type DeleteCoordinatorCandidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCoordinatorCandidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCoordinatorCandidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCoordinatorCandidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCoordinatorCandidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteCoordinatorCandidateRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCoordinatorCandidateOK creates a DeleteCoordinatorCandidateOK with default headers values
func NewDeleteCoordinatorCandidateOK() *DeleteCoordinatorCandidateOK {
	return &DeleteCoordinatorCandidateOK{}
}

/* DeleteCoordinatorCandidateOK describes a response with status code 200, with default header values.

The coordinator candidate specified by {coordinator_candidate_id} was deleted.
*/
type DeleteCoordinatorCandidateOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteCoordinatorCandidateOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] deleteCoordinatorCandidateOK  %+v", 200, o.Payload)
}
func (o *DeleteCoordinatorCandidateOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteCoordinatorCandidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCoordinatorCandidateBadRequest creates a DeleteCoordinatorCandidateBadRequest with default headers values
func NewDeleteCoordinatorCandidateBadRequest() *DeleteCoordinatorCandidateBadRequest {
	return &DeleteCoordinatorCandidateBadRequest{}
}

/* DeleteCoordinatorCandidateBadRequest describes a response with status code 400, with default header values.

The coordinator specified by {coordinator_candidate_id} could not be deleted. (code: `coordinators.candidate_deletion_failed`)
*/
type DeleteCoordinatorCandidateBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteCoordinatorCandidateBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] deleteCoordinatorCandidateBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteCoordinatorCandidateBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteCoordinatorCandidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCoordinatorCandidateNotFound creates a DeleteCoordinatorCandidateNotFound with default headers values
func NewDeleteCoordinatorCandidateNotFound() *DeleteCoordinatorCandidateNotFound {
	return &DeleteCoordinatorCandidateNotFound{}
}

/* DeleteCoordinatorCandidateNotFound describes a response with status code 404, with default header values.

The coordinator candidate specified by {coordinator_candidate_id} cannot be found. (code: `coordinators.candidate_not_found`)
*/
type DeleteCoordinatorCandidateNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteCoordinatorCandidateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] deleteCoordinatorCandidateNotFound  %+v", 404, o.Payload)
}
func (o *DeleteCoordinatorCandidateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteCoordinatorCandidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteCoordinatorCandidateRetryWith creates a DeleteCoordinatorCandidateRetryWith with default headers values
func NewDeleteCoordinatorCandidateRetryWith() *DeleteCoordinatorCandidateRetryWith {
	return &DeleteCoordinatorCandidateRetryWith{}
}

/* DeleteCoordinatorCandidateRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteCoordinatorCandidateRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteCoordinatorCandidateRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] deleteCoordinatorCandidateRetryWith  %+v", 449, o.Payload)
}
func (o *DeleteCoordinatorCandidateRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteCoordinatorCandidateRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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