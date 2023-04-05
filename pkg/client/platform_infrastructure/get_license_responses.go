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

// GetLicenseReader is a Reader for the GetLicense structure.
type GetLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLicenseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLicenseOK creates a GetLicenseOK with default headers values
func NewGetLicenseOK() *GetLicenseOK {
	return &GetLicenseOK{}
}

/*
GetLicenseOK describes a response with status code 200, with default header values.

The information for the license.
*/
type GetLicenseOK struct {
	Payload *models.LicenseObject
}

// IsSuccess returns true when this get license o k response has a 2xx status code
func (o *GetLicenseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get license o k response has a 3xx status code
func (o *GetLicenseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license o k response has a 4xx status code
func (o *GetLicenseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get license o k response has a 5xx status code
func (o *GetLicenseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get license o k response a status code equal to that given
func (o *GetLicenseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get license o k response
func (o *GetLicenseOK) Code() int {
	return 200
}

func (o *GetLicenseOK) Error() string {
	return fmt.Sprintf("[GET /platform/license][%d] getLicenseOK  %+v", 200, o.Payload)
}

func (o *GetLicenseOK) String() string {
	return fmt.Sprintf("[GET /platform/license][%d] getLicenseOK  %+v", 200, o.Payload)
}

func (o *GetLicenseOK) GetPayload() *models.LicenseObject {
	return o.Payload
}

func (o *GetLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseObject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLicenseNotFound creates a GetLicenseNotFound with default headers values
func NewGetLicenseNotFound() *GetLicenseNotFound {
	return &GetLicenseNotFound{}
}

/*
GetLicenseNotFound describes a response with status code 404, with default header values.

The license cannot be found. (code: `license.license_not_found`)
*/
type GetLicenseNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get license not found response has a 2xx status code
func (o *GetLicenseNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get license not found response has a 3xx status code
func (o *GetLicenseNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get license not found response has a 4xx status code
func (o *GetLicenseNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get license not found response has a 5xx status code
func (o *GetLicenseNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get license not found response a status code equal to that given
func (o *GetLicenseNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get license not found response
func (o *GetLicenseNotFound) Code() int {
	return 404
}

func (o *GetLicenseNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/license][%d] getLicenseNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseNotFound) String() string {
	return fmt.Sprintf("[GET /platform/license][%d] getLicenseNotFound  %+v", 404, o.Payload)
}

func (o *GetLicenseNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetLicenseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
