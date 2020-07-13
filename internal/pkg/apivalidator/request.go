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

package apivalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"

	"github.com/go-openapi/spec"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

type prismResponseBody struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Status int32  `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type apiRequests struct {
	reqs []*http.Request
}

// NewHTTPRequests creates multiple valid requests from a defined
// api specification and host.
func NewHTTPRequests(host string, client *http.Client, cloudSpec *spec.Swagger) error {
	if client == nil {
		return errors.New("requests to api: an http client must be specified")
	}

	var reqs []*http.Request
	requests := apiRequests{reqs}
	// nolint
	for path, pathItem := range cloudSpec.Paths.Paths {
		var request []*http.Request
		if pathItem.Head != nil {
			request = append(request, buildRequest(host, path, "HEAD"))
		}
		if pathItem.Get != nil {
			request = append(request, buildRequest(host, path, "GET"))
		}
		if pathItem.Post != nil {
			request = append(request, buildRequest(host, path, "POST"))
		}
		if pathItem.Put != nil {
			request = append(request, buildRequest(host, path, "PUT"))
		}
		if pathItem.Patch != nil {
			request = append(request, buildRequest(host, path, "PATCH"))
		}
		if pathItem.Delete != nil {
			request = append(request, buildRequest(host, path, "DELETE"))
		}

		requests.reqs = append(requests.reqs, request...)
	}

	sort.SliceStable(requests.reqs, func(i, j int) bool {
		return requests.reqs[i].Method < requests.reqs[j].Method
	})

	var merr = multierror.NewPrefixed("api spec validation")

	for i := range requests.reqs {
		req := requests.reqs[i]
		fmt.Printf("%v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		if err := validateRequest(req, client); err != nil {
			merr = merr.Append(err)
		}
	}

	return merr.ErrorOrNil()
}

func buildRequest(host, path, method string) *http.Request {
	// We are setting some defaults and populating empty bodies
	// in order to avoid validation errors from API
	path = strings.ReplaceAll(path, "{resource_kind}", "kibana")
	path = strings.ReplaceAll(path, "{stateless_resource_kind}", "apm")

	isPost := method == "POST"
	isPut := method == "PUT"
	isPatch := method == "PATCH"
	isUsersAuthKeys := strings.Contains(path, "/users/auth/keys")
	isDeploymentTemplates := strings.Contains(path, "deployments/templates")

	if isDeploymentTemplates {
		path += "?region=ece-region"
	}

	if isPost || isPut || isPatch || isUsersAuthKeys {
		r := strings.NewReader(`{}`)
		return &http.Request{
			Method: method,
			Body:   ioutil.NopCloser(r),
			URL: &url.URL{
				Host: host,
				Path: path,
			},
		}
	}

	return &http.Request{
		Method: method,
		URL: &url.URL{
			Host: host,
			Path: path,
		},
	}
}

func validateRequest(request *http.Request, client *http.Client) error {
	endpointURL := request.URL.Host + request.URL.Path
	req, err := http.NewRequest(request.Method, endpointURL, request.Body)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", os.ExpandEnv("ApiKey $EC_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var prismResponses *prismResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&prismResponses); err != nil {
		return err
	}

	// When a response from the prism validation proxy instead of an API response
	// is returned, it means we have discrepancies between our API spec and the live API.
	// These discrepancies are what we are looking for.
	if prismResponses.Type != "" {
		return fmt.Errorf("prism error: Type: %v, Title: %v, Status: %v, Detail: %v",
			prismResponses.Type,
			prismResponses.Title,
			prismResponses.Status,
			prismResponses.Detail)
	}

	return nil
}
