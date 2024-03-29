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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetGlobalDeploymentTemplatesReader is a Reader for the GetGlobalDeploymentTemplates structure.
type GetGlobalDeploymentTemplatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGlobalDeploymentTemplatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGlobalDeploymentTemplatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGlobalDeploymentTemplatesOK creates a GetGlobalDeploymentTemplatesOK with default headers values
func NewGetGlobalDeploymentTemplatesOK() *GetGlobalDeploymentTemplatesOK {
	return &GetGlobalDeploymentTemplatesOK{}
}

/*
GetGlobalDeploymentTemplatesOK describes a response with status code 200, with default header values.

The deployment templates were returned successfully.
*/
type GetGlobalDeploymentTemplatesOK struct {
	Payload []*models.GlobalDeploymentTemplateInfo
}

// IsSuccess returns true when this get global deployment templates o k response has a 2xx status code
func (o *GetGlobalDeploymentTemplatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get global deployment templates o k response has a 3xx status code
func (o *GetGlobalDeploymentTemplatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get global deployment templates o k response has a 4xx status code
func (o *GetGlobalDeploymentTemplatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get global deployment templates o k response has a 5xx status code
func (o *GetGlobalDeploymentTemplatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get global deployment templates o k response a status code equal to that given
func (o *GetGlobalDeploymentTemplatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get global deployment templates o k response
func (o *GetGlobalDeploymentTemplatesOK) Code() int {
	return 200
}

func (o *GetGlobalDeploymentTemplatesOK) Error() string {
	return fmt.Sprintf("[GET /platform/configuration/templates/deployments/global][%d] getGlobalDeploymentTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetGlobalDeploymentTemplatesOK) String() string {
	return fmt.Sprintf("[GET /platform/configuration/templates/deployments/global][%d] getGlobalDeploymentTemplatesOK  %+v", 200, o.Payload)
}

func (o *GetGlobalDeploymentTemplatesOK) GetPayload() []*models.GlobalDeploymentTemplateInfo {
	return o.Payload
}

func (o *GetGlobalDeploymentTemplatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
