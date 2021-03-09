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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateEsClusterReader is a Reader for the CreateEsCluster structure.
type CreateEsClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEsClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEsClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateEsClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateEsClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEsClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewCreateEsClusterRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEsClusterOK creates a CreateEsClusterOK with default headers values
func NewCreateEsClusterOK() *CreateEsClusterOK {
	return &CreateEsClusterOK{}
}

/* CreateEsClusterOK describes a response with status code 200, with default header values.

The cluster definition was valid - no further action was requested. The return object contains an internal representation of the plan, for use in debugging
*/
type CreateEsClusterOK struct {
	Payload *models.ClusterCrudResponse
}

func (o *CreateEsClusterOK) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch][%d] createEsClusterOK  %+v", 200, o.Payload)
}
func (o *CreateEsClusterOK) GetPayload() *models.ClusterCrudResponse {
	return o.Payload
}

func (o *CreateEsClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEsClusterCreated creates a CreateEsClusterCreated with default headers values
func NewCreateEsClusterCreated() *CreateEsClusterCreated {
	return &CreateEsClusterCreated{}
}

/* CreateEsClusterCreated describes a response with status code 201, with default header values.

The cluster definition was valid and the cluster creation has started
*/
type CreateEsClusterCreated struct {
	Payload *models.ClusterCrudResponse
}

func (o *CreateEsClusterCreated) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch][%d] createEsClusterCreated  %+v", 201, o.Payload)
}
func (o *CreateEsClusterCreated) GetPayload() *models.ClusterCrudResponse {
	return o.Payload
}

func (o *CreateEsClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEsClusterAccepted creates a CreateEsClusterAccepted with default headers values
func NewCreateEsClusterAccepted() *CreateEsClusterAccepted {
	return &CreateEsClusterAccepted{}
}

/* CreateEsClusterAccepted describes a response with status code 202, with default header values.

The cluster definition was valid and the cluster creation has already started
*/
type CreateEsClusterAccepted struct {
	Payload *models.ClusterCrudResponse
}

func (o *CreateEsClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch][%d] createEsClusterAccepted  %+v", 202, o.Payload)
}
func (o *CreateEsClusterAccepted) GetPayload() *models.ClusterCrudResponse {
	return o.Payload
}

func (o *CreateEsClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCrudResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEsClusterBadRequest creates a CreateEsClusterBadRequest with default headers values
func NewCreateEsClusterBadRequest() *CreateEsClusterBadRequest {
	return &CreateEsClusterBadRequest{}
}

/* CreateEsClusterBadRequest describes a response with status code 400, with default header values.

 * The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)
* The features used in the cluster definition have not been implemented yet. (code: `clusters.plan_feature_not_implemented`)
*/
type CreateEsClusterBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateEsClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch][%d] createEsClusterBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEsClusterBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateEsClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateEsClusterRetryWith creates a CreateEsClusterRetryWith with default headers values
func NewCreateEsClusterRetryWith() *CreateEsClusterRetryWith {
	return &CreateEsClusterRetryWith{}
}

/* CreateEsClusterRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type CreateEsClusterRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *CreateEsClusterRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch][%d] createEsClusterRetryWith  %+v", 449, o.Payload)
}
func (o *CreateEsClusterRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateEsClusterRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
