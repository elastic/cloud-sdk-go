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

// GetInstanceTypesReader is a Reader for the GetInstanceTypes structure.
type GetInstanceTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceTypesOK creates a GetInstanceTypesOK with default headers values
func NewGetInstanceTypesOK() *GetInstanceTypesOK {
	return &GetInstanceTypesOK{}
}

/*
GetInstanceTypesOK describes a response with status code 200, with default header values.

List of node types
*/
type GetInstanceTypesOK struct {
	Payload []*models.InstanceTypeResource
}

// IsSuccess returns true when this get instance types o k response has a 2xx status code
func (o *GetInstanceTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get instance types o k response has a 3xx status code
func (o *GetInstanceTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance types o k response has a 4xx status code
func (o *GetInstanceTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get instance types o k response has a 5xx status code
func (o *GetInstanceTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance types o k response a status code equal to that given
func (o *GetInstanceTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get instance types o k response
func (o *GetInstanceTypesOK) Code() int {
	return 200
}

func (o *GetInstanceTypesOK) Error() string {
	return fmt.Sprintf("[GET /stack/instance-types][%d] getInstanceTypesOK  %+v", 200, o.Payload)
}

func (o *GetInstanceTypesOK) String() string {
	return fmt.Sprintf("[GET /stack/instance-types][%d] getInstanceTypesOK  %+v", 200, o.Payload)
}

func (o *GetInstanceTypesOK) GetPayload() []*models.InstanceTypeResource {
	return o.Payload
}

func (o *GetInstanceTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
