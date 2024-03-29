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

// PromoteCoordinatorCandidateReader is a Reader for the PromoteCoordinatorCandidate structure.
type PromoteCoordinatorCandidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteCoordinatorCandidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPromoteCoordinatorCandidateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPromoteCoordinatorCandidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPromoteCoordinatorCandidateAccepted creates a PromoteCoordinatorCandidateAccepted with default headers values
func NewPromoteCoordinatorCandidateAccepted() *PromoteCoordinatorCandidateAccepted {
	return &PromoteCoordinatorCandidateAccepted{}
}

/*
PromoteCoordinatorCandidateAccepted describes a response with status code 202, with default header values.

Accepted promote of coordinator candidate.
*/
type PromoteCoordinatorCandidateAccepted struct {
	Payload *models.CoordinatorCandidateInfo
}

// IsSuccess returns true when this promote coordinator candidate accepted response has a 2xx status code
func (o *PromoteCoordinatorCandidateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this promote coordinator candidate accepted response has a 3xx status code
func (o *PromoteCoordinatorCandidateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote coordinator candidate accepted response has a 4xx status code
func (o *PromoteCoordinatorCandidateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this promote coordinator candidate accepted response has a 5xx status code
func (o *PromoteCoordinatorCandidateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this promote coordinator candidate accepted response a status code equal to that given
func (o *PromoteCoordinatorCandidateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the promote coordinator candidate accepted response
func (o *PromoteCoordinatorCandidateAccepted) Code() int {
	return 202
}

func (o *PromoteCoordinatorCandidateAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}/_promote][%d] promoteCoordinatorCandidateAccepted  %+v", 202, o.Payload)
}

func (o *PromoteCoordinatorCandidateAccepted) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}/_promote][%d] promoteCoordinatorCandidateAccepted  %+v", 202, o.Payload)
}

func (o *PromoteCoordinatorCandidateAccepted) GetPayload() *models.CoordinatorCandidateInfo {
	return o.Payload
}

func (o *PromoteCoordinatorCandidateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoordinatorCandidateInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteCoordinatorCandidateNotFound creates a PromoteCoordinatorCandidateNotFound with default headers values
func NewPromoteCoordinatorCandidateNotFound() *PromoteCoordinatorCandidateNotFound {
	return &PromoteCoordinatorCandidateNotFound{}
}

/*
PromoteCoordinatorCandidateNotFound describes a response with status code 404, with default header values.

Unable to find coordinator candidate {coordinator_id}. Edit your request, then try again. (code: `coordinators.candidate_not_found`)
*/
type PromoteCoordinatorCandidateNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this promote coordinator candidate not found response has a 2xx status code
func (o *PromoteCoordinatorCandidateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this promote coordinator candidate not found response has a 3xx status code
func (o *PromoteCoordinatorCandidateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this promote coordinator candidate not found response has a 4xx status code
func (o *PromoteCoordinatorCandidateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this promote coordinator candidate not found response has a 5xx status code
func (o *PromoteCoordinatorCandidateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this promote coordinator candidate not found response a status code equal to that given
func (o *PromoteCoordinatorCandidateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the promote coordinator candidate not found response
func (o *PromoteCoordinatorCandidateNotFound) Code() int {
	return 404
}

func (o *PromoteCoordinatorCandidateNotFound) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}/_promote][%d] promoteCoordinatorCandidateNotFound  %+v", 404, o.Payload)
}

func (o *PromoteCoordinatorCandidateNotFound) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}/_promote][%d] promoteCoordinatorCandidateNotFound  %+v", 404, o.Payload)
}

func (o *PromoteCoordinatorCandidateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *PromoteCoordinatorCandidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
