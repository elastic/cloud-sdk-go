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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ResetApmSecretTokenReader is a Reader for the ResetApmSecretToken structure.
type ResetApmSecretTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetApmSecretTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetApmSecretTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewResetApmSecretTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewResetApmSecretTokenPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewResetApmSecretTokenRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetApmSecretTokenOK creates a ResetApmSecretTokenOK with default headers values
func NewResetApmSecretTokenOK() *ResetApmSecretTokenOK {
	return &ResetApmSecretTokenOK{}
}

/* ResetApmSecretTokenOK describes a response with status code 200, with default header values.

The token was successfully reset
*/
type ResetApmSecretTokenOK struct {
	Payload *models.ApmCrudResponse
}

func (o *ResetApmSecretTokenOK) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_reset-token][%d] resetApmSecretTokenOK  %+v", 200, o.Payload)
}
func (o *ResetApmSecretTokenOK) GetPayload() *models.ApmCrudResponse {
	return o.Payload
}

func (o *ResetApmSecretTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetApmSecretTokenNotFound creates a ResetApmSecretTokenNotFound with default headers values
func NewResetApmSecretTokenNotFound() *ResetApmSecretTokenNotFound {
	return &ResetApmSecretTokenNotFound{}
}

/* ResetApmSecretTokenNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type ResetApmSecretTokenNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *ResetApmSecretTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_reset-token][%d] resetApmSecretTokenNotFound  %+v", 404, o.Payload)
}
func (o *ResetApmSecretTokenNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetApmSecretTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetApmSecretTokenPreconditionFailed creates a ResetApmSecretTokenPreconditionFailed with default headers values
func NewResetApmSecretTokenPreconditionFailed() *ResetApmSecretTokenPreconditionFailed {
	return &ResetApmSecretTokenPreconditionFailed{}
}

/* ResetApmSecretTokenPreconditionFailed describes a response with status code 412, with default header values.

There is not currently applied plan - eg the cluster has not finished provisioning, or the provisioning failed (code: 'clusters.cluster_plan_state_error')
*/
type ResetApmSecretTokenPreconditionFailed struct {
	Payload *models.BasicFailedReply
}

func (o *ResetApmSecretTokenPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_reset-token][%d] resetApmSecretTokenPreconditionFailed  %+v", 412, o.Payload)
}
func (o *ResetApmSecretTokenPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetApmSecretTokenPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetApmSecretTokenRetryWith creates a ResetApmSecretTokenRetryWith with default headers values
func NewResetApmSecretTokenRetryWith() *ResetApmSecretTokenRetryWith {
	return &ResetApmSecretTokenRetryWith{}
}

/* ResetApmSecretTokenRetryWith describes a response with status code 449, with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type ResetApmSecretTokenRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *ResetApmSecretTokenRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/_reset-token][%d] resetApmSecretTokenRetryWith  %+v", 449, o.Payload)
}
func (o *ResetApmSecretTokenRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResetApmSecretTokenRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
