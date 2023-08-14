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

package stack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteVersionStackReader is a Reader for the DeleteVersionStack structure.
type DeleteVersionStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVersionStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVersionStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteVersionStackNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVersionStackOK creates a DeleteVersionStackOK with default headers values
func NewDeleteVersionStackOK() *DeleteVersionStackOK {
	return &DeleteVersionStackOK{}
}

/*
DeleteVersionStackOK describes a response with status code 200, with default header values.

The `deleted` flag is applied to the specified Elastic Stack version.
*/
type DeleteVersionStackOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this delete version stack o k response has a 2xx status code
func (o *DeleteVersionStackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete version stack o k response has a 3xx status code
func (o *DeleteVersionStackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version stack o k response has a 4xx status code
func (o *DeleteVersionStackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete version stack o k response has a 5xx status code
func (o *DeleteVersionStackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version stack o k response a status code equal to that given
func (o *DeleteVersionStackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete version stack o k response
func (o *DeleteVersionStackOK) Code() int {
	return 200
}

func (o *DeleteVersionStackOK) Error() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionStackOK) String() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackOK  %+v", 200, o.Payload)
}

func (o *DeleteVersionStackOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteVersionStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVersionStackNotFound creates a DeleteVersionStackNotFound with default headers values
func NewDeleteVersionStackNotFound() *DeleteVersionStackNotFound {
	return &DeleteVersionStackNotFound{}
}

/*
DeleteVersionStackNotFound describes a response with status code 404, with default header values.

The Elastic Stack version can't be found. (code: `stackpack.version_not_found`)
*/
type DeleteVersionStackNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this delete version stack not found response has a 2xx status code
func (o *DeleteVersionStackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete version stack not found response has a 3xx status code
func (o *DeleteVersionStackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete version stack not found response has a 4xx status code
func (o *DeleteVersionStackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete version stack not found response has a 5xx status code
func (o *DeleteVersionStackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete version stack not found response a status code equal to that given
func (o *DeleteVersionStackNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete version stack not found response
func (o *DeleteVersionStackNotFound) Code() int {
	return 404
}

func (o *DeleteVersionStackNotFound) Error() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVersionStackNotFound) String() string {
	return fmt.Sprintf("[DELETE /stack/versions/{version}][%d] deleteVersionStackNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVersionStackNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteVersionStackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
