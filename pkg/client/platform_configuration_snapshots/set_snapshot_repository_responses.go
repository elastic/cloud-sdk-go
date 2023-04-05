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

package platform_configuration_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetSnapshotRepositoryReader is a Reader for the SetSnapshotRepository structure.
type SetSnapshotRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetSnapshotRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetSnapshotRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetSnapshotRepositoryOK creates a SetSnapshotRepositoryOK with default headers values
func NewSetSnapshotRepositoryOK() *SetSnapshotRepositoryOK {
	return &SetSnapshotRepositoryOK{}
}

/*
SetSnapshotRepositoryOK describes a response with status code 200, with default header values.

New snapshot repository config
*/
type SetSnapshotRepositoryOK struct {
	Payload *models.RepositoryConfig
}

// IsSuccess returns true when this set snapshot repository o k response has a 2xx status code
func (o *SetSnapshotRepositoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set snapshot repository o k response has a 3xx status code
func (o *SetSnapshotRepositoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set snapshot repository o k response has a 4xx status code
func (o *SetSnapshotRepositoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set snapshot repository o k response has a 5xx status code
func (o *SetSnapshotRepositoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set snapshot repository o k response a status code equal to that given
func (o *SetSnapshotRepositoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the set snapshot repository o k response
func (o *SetSnapshotRepositoryOK) Code() int {
	return 200
}

func (o *SetSnapshotRepositoryOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/snapshots/repositories/{repository_name}][%d] setSnapshotRepositoryOK  %+v", 200, o.Payload)
}

func (o *SetSnapshotRepositoryOK) String() string {
	return fmt.Sprintf("[PUT /platform/configuration/snapshots/repositories/{repository_name}][%d] setSnapshotRepositoryOK  %+v", 200, o.Payload)
}

func (o *SetSnapshotRepositoryOK) GetPayload() *models.RepositoryConfig {
	return o.Payload
}

func (o *SetSnapshotRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepositoryConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
