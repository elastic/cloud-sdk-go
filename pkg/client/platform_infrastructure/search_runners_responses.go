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

// SearchRunnersReader is a Reader for the SearchRunners structure.
type SearchRunnersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchRunnersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchRunnersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchRunnersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchRunnersOK creates a SearchRunnersOK with default headers values
func NewSearchRunnersOK() *SearchRunnersOK {
	return &SearchRunnersOK{}
}

/*
SearchRunnersOK describes a response with status code 200, with default header values.

An overview of runners that matched the given search query.
*/
type SearchRunnersOK struct {
	Payload *models.RunnerOverview
}

// IsSuccess returns true when this search runners o k response has a 2xx status code
func (o *SearchRunnersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search runners o k response has a 3xx status code
func (o *SearchRunnersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search runners o k response has a 4xx status code
func (o *SearchRunnersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search runners o k response has a 5xx status code
func (o *SearchRunnersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search runners o k response a status code equal to that given
func (o *SearchRunnersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search runners o k response
func (o *SearchRunnersOK) Code() int {
	return 200
}

func (o *SearchRunnersOK) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/_search][%d] searchRunnersOK  %+v", 200, o.Payload)
}

func (o *SearchRunnersOK) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/_search][%d] searchRunnersOK  %+v", 200, o.Payload)
}

func (o *SearchRunnersOK) GetPayload() *models.RunnerOverview {
	return o.Payload
}

func (o *SearchRunnersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RunnerOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchRunnersBadRequest creates a SearchRunnersBadRequest with default headers values
func NewSearchRunnersBadRequest() *SearchRunnersBadRequest {
	return &SearchRunnersBadRequest{}
}

/*
SearchRunnersBadRequest describes a response with status code 400, with default header values.

The search request failed. (code: `runners.search_failed`)
*/
type SearchRunnersBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this search runners bad request response has a 2xx status code
func (o *SearchRunnersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search runners bad request response has a 3xx status code
func (o *SearchRunnersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search runners bad request response has a 4xx status code
func (o *SearchRunnersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search runners bad request response has a 5xx status code
func (o *SearchRunnersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search runners bad request response a status code equal to that given
func (o *SearchRunnersBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the search runners bad request response
func (o *SearchRunnersBadRequest) Code() int {
	return 400
}

func (o *SearchRunnersBadRequest) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/_search][%d] searchRunnersBadRequest  %+v", 400, o.Payload)
}

func (o *SearchRunnersBadRequest) String() string {
	return fmt.Sprintf("[POST /platform/infrastructure/runners/_search][%d] searchRunnersBadRequest  %+v", 400, o.Payload)
}

func (o *SearchRunnersBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SearchRunnersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
