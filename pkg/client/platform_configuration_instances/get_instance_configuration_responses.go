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

package platform_configuration_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetInstanceConfigurationReader is a Reader for the GetInstanceConfiguration structure.
type GetInstanceConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstanceConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetInstanceConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInstanceConfigurationOK creates a GetInstanceConfigurationOK with default headers values
func NewGetInstanceConfigurationOK() *GetInstanceConfigurationOK {
	return &GetInstanceConfigurationOK{}
}

/*
GetInstanceConfigurationOK describes a response with status code 200, with default header values.

An InstanceConfiguration
*/
type GetInstanceConfigurationOK struct {
	Payload *models.InstanceConfiguration
}

// IsSuccess returns true when this get instance configuration o k response has a 2xx status code
func (o *GetInstanceConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get instance configuration o k response has a 3xx status code
func (o *GetInstanceConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance configuration o k response has a 4xx status code
func (o *GetInstanceConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get instance configuration o k response has a 5xx status code
func (o *GetInstanceConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance configuration o k response a status code equal to that given
func (o *GetInstanceConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get instance configuration o k response
func (o *GetInstanceConfigurationOK) Code() int {
	return 200
}

func (o *GetInstanceConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/instances/{id}][%d] getInstanceConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetInstanceConfigurationOK) String() string {
	return fmt.Sprintf("[GET /platform/configuration/instances/{id}][%d] getInstanceConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetInstanceConfigurationOK) GetPayload() *models.InstanceConfiguration {
	return o.Payload
}

func (o *GetInstanceConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceConfigurationNotFound creates a GetInstanceConfigurationNotFound with default headers values
func NewGetInstanceConfigurationNotFound() *GetInstanceConfigurationNotFound {
	return &GetInstanceConfigurationNotFound{}
}

/*
GetInstanceConfigurationNotFound describes a response with status code 404, with default header values.

Instance configuration specified by {id} cannot be found (code: 'configuration.instance_configuration_not_found')
*/
type GetInstanceConfigurationNotFound struct {
	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this get instance configuration not found response has a 2xx status code
func (o *GetInstanceConfigurationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get instance configuration not found response has a 3xx status code
func (o *GetInstanceConfigurationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get instance configuration not found response has a 4xx status code
func (o *GetInstanceConfigurationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get instance configuration not found response has a 5xx status code
func (o *GetInstanceConfigurationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get instance configuration not found response a status code equal to that given
func (o *GetInstanceConfigurationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get instance configuration not found response
func (o *GetInstanceConfigurationNotFound) Code() int {
	return 404
}

func (o *GetInstanceConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/instances/{id}][%d] getInstanceConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceConfigurationNotFound) String() string {
	return fmt.Sprintf("[GET /platform/configuration/instances/{id}][%d] getInstanceConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceConfigurationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetInstanceConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
