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

// UpdateEsClusterSnapshotSettingsReader is a Reader for the UpdateEsClusterSnapshotSettings structure.
type UpdateEsClusterSnapshotSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEsClusterSnapshotSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEsClusterSnapshotSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateEsClusterSnapshotSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEsClusterSnapshotSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateEsClusterSnapshotSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEsClusterSnapshotSettingsOK creates a UpdateEsClusterSnapshotSettingsOK with default headers values
func NewUpdateEsClusterSnapshotSettingsOK() *UpdateEsClusterSnapshotSettingsOK {
	return &UpdateEsClusterSnapshotSettingsOK{}
}

/* UpdateEsClusterSnapshotSettingsOK describes a response with status code 200, with default header values.

The cluster snapshot settings were successfully updated
*/
type UpdateEsClusterSnapshotSettingsOK struct {

	/* The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string

	/* The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string

	/* The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ClusterSnapshotSettings
}

func (o *UpdateEsClusterSnapshotSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/snapshot/settings][%d] updateEsClusterSnapshotSettingsOK  %+v", 200, o.Payload)
}
func (o *UpdateEsClusterSnapshotSettingsOK) GetPayload() *models.ClusterSnapshotSettings {
	return o.Payload
}

func (o *UpdateEsClusterSnapshotSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ClusterSnapshotSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterSnapshotSettingsForbidden creates a UpdateEsClusterSnapshotSettingsForbidden with default headers values
func NewUpdateEsClusterSnapshotSettingsForbidden() *UpdateEsClusterSnapshotSettingsForbidden {
	return &UpdateEsClusterSnapshotSettingsForbidden{}
}

/* UpdateEsClusterSnapshotSettingsForbidden describes a response with status code 403, with default header values.

The provided action was prohibited for the given cluster.
*/
type UpdateEsClusterSnapshotSettingsForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterSnapshotSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/snapshot/settings][%d] updateEsClusterSnapshotSettingsForbidden  %+v", 403, o.Payload)
}
func (o *UpdateEsClusterSnapshotSettingsForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterSnapshotSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterSnapshotSettingsNotFound creates a UpdateEsClusterSnapshotSettingsNotFound with default headers values
func NewUpdateEsClusterSnapshotSettingsNotFound() *UpdateEsClusterSnapshotSettingsNotFound {
	return &UpdateEsClusterSnapshotSettingsNotFound{}
}

/* UpdateEsClusterSnapshotSettingsNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type UpdateEsClusterSnapshotSettingsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterSnapshotSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/snapshot/settings][%d] updateEsClusterSnapshotSettingsNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEsClusterSnapshotSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterSnapshotSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterSnapshotSettingsRetryWith creates a UpdateEsClusterSnapshotSettingsRetryWith with default headers values
func NewUpdateEsClusterSnapshotSettingsRetryWith() *UpdateEsClusterSnapshotSettingsRetryWith {
	return &UpdateEsClusterSnapshotSettingsRetryWith{}
}

/* UpdateEsClusterSnapshotSettingsRetryWith describes a response with status code 449, with default header values.

elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type UpdateEsClusterSnapshotSettingsRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterSnapshotSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/snapshot/settings][%d] updateEsClusterSnapshotSettingsRetryWith  %+v", 449, o.Payload)
}
func (o *UpdateEsClusterSnapshotSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterSnapshotSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
