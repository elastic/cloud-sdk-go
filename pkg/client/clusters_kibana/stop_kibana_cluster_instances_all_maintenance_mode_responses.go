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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StopKibanaClusterInstancesAllMaintenanceModeReader is a Reader for the StopKibanaClusterInstancesAllMaintenanceMode structure.
type StopKibanaClusterInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopKibanaClusterInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStopKibanaClusterInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopKibanaClusterInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopKibanaClusterInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStopKibanaClusterInstancesAllMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopKibanaClusterInstancesAllMaintenanceModeAccepted creates a StopKibanaClusterInstancesAllMaintenanceModeAccepted with default headers values
func NewStopKibanaClusterInstancesAllMaintenanceModeAccepted() *StopKibanaClusterInstancesAllMaintenanceModeAccepted {
	return &StopKibanaClusterInstancesAllMaintenanceModeAccepted{}
}

/* StopKibanaClusterInstancesAllMaintenanceModeAccepted describes a response with status code 202, with default header values.

The stop maintenance mode command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type StopKibanaClusterInstancesAllMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_stop][%d] stopKibanaClusterInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StopKibanaClusterInstancesAllMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopKibanaClusterInstancesAllMaintenanceModeForbidden creates a StopKibanaClusterInstancesAllMaintenanceModeForbidden with default headers values
func NewStopKibanaClusterInstancesAllMaintenanceModeForbidden() *StopKibanaClusterInstancesAllMaintenanceModeForbidden {
	return &StopKibanaClusterInstancesAllMaintenanceModeForbidden{}
}

/* StopKibanaClusterInstancesAllMaintenanceModeForbidden describes a response with status code 403, with default header values.

The stop maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StopKibanaClusterInstancesAllMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_stop][%d] stopKibanaClusterInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StopKibanaClusterInstancesAllMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopKibanaClusterInstancesAllMaintenanceModeNotFound creates a StopKibanaClusterInstancesAllMaintenanceModeNotFound with default headers values
func NewStopKibanaClusterInstancesAllMaintenanceModeNotFound() *StopKibanaClusterInstancesAllMaintenanceModeNotFound {
	return &StopKibanaClusterInstancesAllMaintenanceModeNotFound{}
}

/* StopKibanaClusterInstancesAllMaintenanceModeNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StopKibanaClusterInstancesAllMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_stop][%d] stopKibanaClusterInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StopKibanaClusterInstancesAllMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStopKibanaClusterInstancesAllMaintenanceModeRetryWith creates a StopKibanaClusterInstancesAllMaintenanceModeRetryWith with default headers values
func NewStopKibanaClusterInstancesAllMaintenanceModeRetryWith() *StopKibanaClusterInstancesAllMaintenanceModeRetryWith {
	return &StopKibanaClusterInstancesAllMaintenanceModeRetryWith{}
}

/* StopKibanaClusterInstancesAllMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StopKibanaClusterInstancesAllMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_stop][%d] stopKibanaClusterInstancesAllMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StopKibanaClusterInstancesAllMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StopKibanaClusterInstancesAllMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
