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

package deployments_traffic_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateTrafficFilterRulesetAssociationReader is a Reader for the CreateTrafficFilterRulesetAssociation structure.
type CreateTrafficFilterRulesetAssociationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTrafficFilterRulesetAssociationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateTrafficFilterRulesetAssociationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateTrafficFilterRulesetAssociationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewCreateTrafficFilterRulesetAssociationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateTrafficFilterRulesetAssociationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTrafficFilterRulesetAssociationOK creates a CreateTrafficFilterRulesetAssociationOK with default headers values
func NewCreateTrafficFilterRulesetAssociationOK() *CreateTrafficFilterRulesetAssociationOK {
	return &CreateTrafficFilterRulesetAssociationOK{}
}

/*
CreateTrafficFilterRulesetAssociationOK describes a response with status code 200, with default header values.

Create association request was valid and the association already exists
*/
type CreateTrafficFilterRulesetAssociationOK struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this create traffic filter ruleset association o k response has a 2xx status code
func (o *CreateTrafficFilterRulesetAssociationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create traffic filter ruleset association o k response has a 3xx status code
func (o *CreateTrafficFilterRulesetAssociationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset association o k response has a 4xx status code
func (o *CreateTrafficFilterRulesetAssociationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create traffic filter ruleset association o k response has a 5xx status code
func (o *CreateTrafficFilterRulesetAssociationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create traffic filter ruleset association o k response a status code equal to that given
func (o *CreateTrafficFilterRulesetAssociationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create traffic filter ruleset association o k response
func (o *CreateTrafficFilterRulesetAssociationOK) Code() int {
	return 200
}

func (o *CreateTrafficFilterRulesetAssociationOK) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationOK  %+v", 200, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationOK) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationOK  %+v", 200, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetAssociationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTrafficFilterRulesetAssociationCreated creates a CreateTrafficFilterRulesetAssociationCreated with default headers values
func NewCreateTrafficFilterRulesetAssociationCreated() *CreateTrafficFilterRulesetAssociationCreated {
	return &CreateTrafficFilterRulesetAssociationCreated{}
}

/*
CreateTrafficFilterRulesetAssociationCreated describes a response with status code 201, with default header values.

Create association request was valid and the association has been created
*/
type CreateTrafficFilterRulesetAssociationCreated struct {
	Payload models.EmptyResponse
}

// IsSuccess returns true when this create traffic filter ruleset association created response has a 2xx status code
func (o *CreateTrafficFilterRulesetAssociationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create traffic filter ruleset association created response has a 3xx status code
func (o *CreateTrafficFilterRulesetAssociationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset association created response has a 4xx status code
func (o *CreateTrafficFilterRulesetAssociationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create traffic filter ruleset association created response has a 5xx status code
func (o *CreateTrafficFilterRulesetAssociationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create traffic filter ruleset association created response a status code equal to that given
func (o *CreateTrafficFilterRulesetAssociationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create traffic filter ruleset association created response
func (o *CreateTrafficFilterRulesetAssociationCreated) Code() int {
	return 201
}

func (o *CreateTrafficFilterRulesetAssociationCreated) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationCreated  %+v", 201, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationCreated) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationCreated  %+v", 201, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationCreated) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetAssociationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTrafficFilterRulesetAssociationNotFound creates a CreateTrafficFilterRulesetAssociationNotFound with default headers values
func NewCreateTrafficFilterRulesetAssociationNotFound() *CreateTrafficFilterRulesetAssociationNotFound {
	return &CreateTrafficFilterRulesetAssociationNotFound{}
}

/*
CreateTrafficFilterRulesetAssociationNotFound describes a response with status code 404, with default header values.

The traffic filter ruleset specified by {ruleset_id} cannot be found. (code: `traffic_filter.not_found`)
*/
type CreateTrafficFilterRulesetAssociationNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create traffic filter ruleset association not found response has a 2xx status code
func (o *CreateTrafficFilterRulesetAssociationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create traffic filter ruleset association not found response has a 3xx status code
func (o *CreateTrafficFilterRulesetAssociationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset association not found response has a 4xx status code
func (o *CreateTrafficFilterRulesetAssociationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create traffic filter ruleset association not found response has a 5xx status code
func (o *CreateTrafficFilterRulesetAssociationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create traffic filter ruleset association not found response a status code equal to that given
func (o *CreateTrafficFilterRulesetAssociationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create traffic filter ruleset association not found response
func (o *CreateTrafficFilterRulesetAssociationNotFound) Code() int {
	return 404
}

func (o *CreateTrafficFilterRulesetAssociationNotFound) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationNotFound  %+v", 404, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationNotFound) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationNotFound  %+v", 404, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetAssociationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateTrafficFilterRulesetAssociationInternalServerError creates a CreateTrafficFilterRulesetAssociationInternalServerError with default headers values
func NewCreateTrafficFilterRulesetAssociationInternalServerError() *CreateTrafficFilterRulesetAssociationInternalServerError {
	return &CreateTrafficFilterRulesetAssociationInternalServerError{}
}

/*
CreateTrafficFilterRulesetAssociationInternalServerError describes a response with status code 500, with default header values.

Creation failed. (code: `traffic_filter.request_execution_failed`)
*/
type CreateTrafficFilterRulesetAssociationInternalServerError struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

// IsSuccess returns true when this create traffic filter ruleset association internal server error response has a 2xx status code
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create traffic filter ruleset association internal server error response has a 3xx status code
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create traffic filter ruleset association internal server error response has a 4xx status code
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create traffic filter ruleset association internal server error response has a 5xx status code
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create traffic filter ruleset association internal server error response a status code equal to that given
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create traffic filter ruleset association internal server error response
func (o *CreateTrafficFilterRulesetAssociationInternalServerError) Code() int {
	return 500
}

func (o *CreateTrafficFilterRulesetAssociationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationInternalServerError) String() string {
	return fmt.Sprintf("[POST /deployments/traffic-filter/rulesets/{ruleset_id}/associations][%d] createTrafficFilterRulesetAssociationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateTrafficFilterRulesetAssociationInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateTrafficFilterRulesetAssociationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
