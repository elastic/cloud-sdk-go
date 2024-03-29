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

// GetCoordinatorCandidateReader is a Reader for the GetCoordinatorCandidate structure.
type GetCoordinatorCandidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCoordinatorCandidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCoordinatorCandidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCoordinatorCandidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCoordinatorCandidateOK creates a GetCoordinatorCandidateOK with default headers values
func NewGetCoordinatorCandidateOK() *GetCoordinatorCandidateOK {
	return &GetCoordinatorCandidateOK{}
}

/*
GetCoordinatorCandidateOK describes a response with status code 200, with default header values.

Info about a coordinator candidate.
*/
type GetCoordinatorCandidateOK struct {
	Payload *models.CoordinatorCandidateInfo
}

// IsSuccess returns true when this get coordinator candidate o k response has a 2xx status code
func (o *GetCoordinatorCandidateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get coordinator candidate o k response has a 3xx status code
func (o *GetCoordinatorCandidateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coordinator candidate o k response has a 4xx status code
func (o *GetCoordinatorCandidateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get coordinator candidate o k response has a 5xx status code
func (o *GetCoordinatorCandidateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get coordinator candidate o k response a status code equal to that given
func (o *GetCoordinatorCandidateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get coordinator candidate o k response
func (o *GetCoordinatorCandidateOK) Code() int {
	return 200
}

func (o *GetCoordinatorCandidateOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] getCoordinatorCandidateOK  %+v", 200, o.Payload)
}

func (o *GetCoordinatorCandidateOK) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] getCoordinatorCandidateOK  %+v", 200, o.Payload)
}

func (o *GetCoordinatorCandidateOK) GetPayload() *models.CoordinatorCandidateInfo {
	return o.Payload
}

func (o *GetCoordinatorCandidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoordinatorCandidateInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCoordinatorCandidateNotFound creates a GetCoordinatorCandidateNotFound with default headers values
func NewGetCoordinatorCandidateNotFound() *GetCoordinatorCandidateNotFound {
	return &GetCoordinatorCandidateNotFound{}
}

/*
GetCoordinatorCandidateNotFound describes a response with status code 404, with default header values.

Unable to find coordinator candidate {coordinator_candidate_id}. Edit your request, then try again. (code: `coordinators.candidate_not_found`)
*/
type GetCoordinatorCandidateNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get coordinator candidate not found response has a 2xx status code
func (o *GetCoordinatorCandidateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get coordinator candidate not found response has a 3xx status code
func (o *GetCoordinatorCandidateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get coordinator candidate not found response has a 4xx status code
func (o *GetCoordinatorCandidateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get coordinator candidate not found response has a 5xx status code
func (o *GetCoordinatorCandidateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get coordinator candidate not found response a status code equal to that given
func (o *GetCoordinatorCandidateNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get coordinator candidate not found response
func (o *GetCoordinatorCandidateNotFound) Code() int {
	return 404
}

func (o *GetCoordinatorCandidateNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] getCoordinatorCandidateNotFound  %+v", 404, o.Payload)
}

func (o *GetCoordinatorCandidateNotFound) String() string {
	return fmt.Sprintf("[GET /platform/infrastructure/coordinators/candidates/{coordinator_candidate_id}][%d] getCoordinatorCandidateNotFound  %+v", 404, o.Payload)
}

func (o *GetCoordinatorCandidateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetCoordinatorCandidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
