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

// GetApmPlanReader is a Reader for the GetApmPlan structure.
type GetApmPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApmPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApmPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApmPlanNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetApmPlanPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetApmPlanOK creates a GetApmPlanOK with default headers values
func NewGetApmPlanOK() *GetApmPlanOK {
	return &GetApmPlanOK{}
}

/* GetApmPlanOK describes a response with status code 200, with default header values.

The pending plan is applied to the APM server.
*/
type GetApmPlanOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ApmPlan
}

func (o *GetApmPlanOK) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan][%d] getApmPlanOK  %+v", 200, o.Payload)
}
func (o *GetApmPlanOK) GetPayload() *models.ApmPlan {
	return o.Payload
}

func (o *GetApmPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-resource-created
	hdrXCloudResourceCreated := response.GetHeader("x-cloud-resource-created")

	if hdrXCloudResourceCreated != "" {
		o.XCloudResourceCreated = hdrXCloudResourceCreated
	}

	// hydrates response header x-cloud-resource-last-modified
	hdrXCloudResourceLastModified := response.GetHeader("x-cloud-resource-last-modified")

	if hdrXCloudResourceLastModified != "" {
		o.XCloudResourceLastModified = hdrXCloudResourceLastModified
	}

	// hydrates response header x-cloud-resource-version
	hdrXCloudResourceVersion := response.GetHeader("x-cloud-resource-version")

	if hdrXCloudResourceVersion != "" {
		o.XCloudResourceVersion = hdrXCloudResourceVersion
	}

	o.Payload = new(models.ApmPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmPlanNotFound creates a GetApmPlanNotFound with default headers values
func NewGetApmPlanNotFound() *GetApmPlanNotFound {
	return &GetApmPlanNotFound{}
}

/* GetApmPlanNotFound describes a response with status code 404, with default header values.

The {cluster_id} can't be found. (code: 'clusters.cluster_not_found')
*/
type GetApmPlanNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetApmPlanNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan][%d] getApmPlanNotFound  %+v", 404, o.Payload)
}
func (o *GetApmPlanNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmPlanNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmPlanPreconditionFailed creates a GetApmPlanPreconditionFailed with default headers values
func NewGetApmPlanPreconditionFailed() *GetApmPlanPreconditionFailed {
	return &GetApmPlanPreconditionFailed{}
}

/* GetApmPlanPreconditionFailed describes a response with status code 412, with default header values.

The APM server is unable to finish provisioning, or the provisioning failed. Apply a plan, then try again. (code: 'clusters.cluster_plan_state_error')
*/
type GetApmPlanPreconditionFailed struct {
	Payload *models.BasicFailedReply
}

func (o *GetApmPlanPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan][%d] getApmPlanPreconditionFailed  %+v", 412, o.Payload)
}
func (o *GetApmPlanPreconditionFailed) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmPlanPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
