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

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteExtraCertificateReader is a Reader for the DeleteExtraCertificate structure.
type DeleteExtraCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExtraCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteExtraCertificateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteExtraCertificateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExtraCertificateOK creates a DeleteExtraCertificateOK with default headers values
func NewDeleteExtraCertificateOK() *DeleteExtraCertificateOK {
	return &DeleteExtraCertificateOK{}
}

/*DeleteExtraCertificateOK handles this case with default header values.

Extra certificate under the given id is deleted if exists
*/
type DeleteExtraCertificateOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteExtraCertificateOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/extra_certs/{cert_id}][%d] deleteExtraCertificateOK  %+v", 200, o.Payload)
}

func (o *DeleteExtraCertificateOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteExtraCertificateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExtraCertificateNotFound creates a DeleteExtraCertificateNotFound with default headers values
func NewDeleteExtraCertificateNotFound() *DeleteExtraCertificateNotFound {
	return &DeleteExtraCertificateNotFound{}
}

/*DeleteExtraCertificateNotFound handles this case with default header values.

The extra certificate cannot be found. (code: `cert.not_found`)
*/
type DeleteExtraCertificateNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteExtraCertificateNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/configuration/security/extra_certs/{cert_id}][%d] deleteExtraCertificateNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExtraCertificateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteExtraCertificateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
