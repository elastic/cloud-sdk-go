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

package deployment_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetDeploymentTemplateV2Reader is a Reader for the SetDeploymentTemplateV2 structure.
type SetDeploymentTemplateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDeploymentTemplateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDeploymentTemplateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSetDeploymentTemplateV2Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDeploymentTemplateV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSetDeploymentTemplateV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetDeploymentTemplateV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetDeploymentTemplateV2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetDeploymentTemplateV2RetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetDeploymentTemplateV2OK creates a SetDeploymentTemplateV2OK with default headers values
func NewSetDeploymentTemplateV2OK() *SetDeploymentTemplateV2OK {
	return &SetDeploymentTemplateV2OK{}
}

/* SetDeploymentTemplateV2OK describes a response with status code 200, with default header values.

The deployment definition was valid and the template has been updated.
*/
type SetDeploymentTemplateV2OK struct {
	Payload *models.IDResponse
}

func (o *SetDeploymentTemplateV2OK) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2OK  %+v", 200, o.Payload)
}
func (o *SetDeploymentTemplateV2OK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *SetDeploymentTemplateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateV2Created creates a SetDeploymentTemplateV2Created with default headers values
func NewSetDeploymentTemplateV2Created() *SetDeploymentTemplateV2Created {
	return &SetDeploymentTemplateV2Created{}
}

/* SetDeploymentTemplateV2Created describes a response with status code 201, with default header values.

The deployment definition was valid and the template was created.
*/
type SetDeploymentTemplateV2Created struct {
	Payload *models.IDResponse
}

func (o *SetDeploymentTemplateV2Created) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2Created  %+v", 201, o.Payload)
}
func (o *SetDeploymentTemplateV2Created) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *SetDeploymentTemplateV2Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateV2BadRequest creates a SetDeploymentTemplateV2BadRequest with default headers values
func NewSetDeploymentTemplateV2BadRequest() *SetDeploymentTemplateV2BadRequest {
	return &SetDeploymentTemplateV2BadRequest{}
}

/* SetDeploymentTemplateV2BadRequest describes a response with status code 400, with default header values.

The requested region is not supported. (code: `templates.region_not_found`)
*/
type SetDeploymentTemplateV2BadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateV2BadRequest) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2BadRequest  %+v", 400, o.Payload)
}
func (o *SetDeploymentTemplateV2BadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentTemplateV2Unauthorized creates a SetDeploymentTemplateV2Unauthorized with default headers values
func NewSetDeploymentTemplateV2Unauthorized() *SetDeploymentTemplateV2Unauthorized {
	return &SetDeploymentTemplateV2Unauthorized{}
}

/* SetDeploymentTemplateV2Unauthorized describes a response with status code 401, with default header values.

The user is not authorized to access requested region. (code: `templates.region_not_allowed`)
*/
type SetDeploymentTemplateV2Unauthorized struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2Unauthorized  %+v", 401, o.Payload)
}
func (o *SetDeploymentTemplateV2Unauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentTemplateV2NotFound creates a SetDeploymentTemplateV2NotFound with default headers values
func NewSetDeploymentTemplateV2NotFound() *SetDeploymentTemplateV2NotFound {
	return &SetDeploymentTemplateV2NotFound{}
}

/* SetDeploymentTemplateV2NotFound describes a response with status code 404, with default header values.

The deployment template specified by {template_id} cannot be found. (code: `templates.template_not_found`)
*/
type SetDeploymentTemplateV2NotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateV2NotFound) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2NotFound  %+v", 404, o.Payload)
}
func (o *SetDeploymentTemplateV2NotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentTemplateV2Conflict creates a SetDeploymentTemplateV2Conflict with default headers values
func NewSetDeploymentTemplateV2Conflict() *SetDeploymentTemplateV2Conflict {
	return &SetDeploymentTemplateV2Conflict{}
}

/* SetDeploymentTemplateV2Conflict describes a response with status code 409, with default header values.

The version supplied in the request conflicted with the version found on the server. (code: `templates.version_conflict`)
*/
type SetDeploymentTemplateV2Conflict struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateV2Conflict) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2Conflict  %+v", 409, o.Payload)
}
func (o *SetDeploymentTemplateV2Conflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateV2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetDeploymentTemplateV2RetryWith creates a SetDeploymentTemplateV2RetryWith with default headers values
func NewSetDeploymentTemplateV2RetryWith() *SetDeploymentTemplateV2RetryWith {
	return &SetDeploymentTemplateV2RetryWith{}
}

/* SetDeploymentTemplateV2RetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetDeploymentTemplateV2RetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateV2RetryWith) Error() string {
	return fmt.Sprintf("[PUT /deployments/templates/{template_id}][%d] setDeploymentTemplateV2RetryWith  %+v", 449, o.Payload)
}
func (o *SetDeploymentTemplateV2RetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateV2RetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
