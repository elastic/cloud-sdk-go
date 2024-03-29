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

package extensions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ListExtensionsReader is a Reader for the ListExtensions structure.
type ListExtensionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListExtensionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListExtensionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListExtensionsOK creates a ListExtensionsOK with default headers values
func NewListExtensionsOK() *ListExtensionsOK {
	return &ListExtensionsOK{}
}

/*
ListExtensionsOK describes a response with status code 200, with default header values.

The extensions that are available
*/
type ListExtensionsOK struct {
	Payload *models.Extensions
}

// IsSuccess returns true when this list extensions o k response has a 2xx status code
func (o *ListExtensionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list extensions o k response has a 3xx status code
func (o *ListExtensionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list extensions o k response has a 4xx status code
func (o *ListExtensionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list extensions o k response has a 5xx status code
func (o *ListExtensionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list extensions o k response a status code equal to that given
func (o *ListExtensionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list extensions o k response
func (o *ListExtensionsOK) Code() int {
	return 200
}

func (o *ListExtensionsOK) Error() string {
	return fmt.Sprintf("[GET /deployments/extensions][%d] listExtensionsOK  %+v", 200, o.Payload)
}

func (o *ListExtensionsOK) String() string {
	return fmt.Sprintf("[GET /deployments/extensions][%d] listExtensionsOK  %+v", 200, o.Payload)
}

func (o *ListExtensionsOK) GetPayload() *models.Extensions {
	return o.Payload
}

func (o *ListExtensionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Extensions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
