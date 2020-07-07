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

package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-openapi/spec"

	"github.com/elastic/cloud-sdk-go/internal/pkg/apivalidator"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

const (
	_ = iota
	codeFailedConfigValidation
	codeFailedDownloadingSpec
	codeCannotReadFile
	codeFailedUnmarshalingSpec
	codeFailedAPISpecValidation
)

var (
	defaultDestinationFile = fmt.Sprint(time.Now().Format(time.RFC3339), "-apidocs.json")

	errInvalidSource = errors.New("an api specification file must be specified")
	errInvalidHost   = errors.New("a host for connecting to the validation proxy must be specified")
	errInvalidPort   = errors.New("a port for connecting to the validation proxy must be specified")
	errInvalidClient = errors.New("an http client must be specified")
)

// Config contains the application configuration
type Config struct {
	Source string
	Host   string
	Port   string

	Client *http.Client
}

// Validate checks that the configuration is valid.
func (c Config) Validate() error {
	var err = multierror.NewPrefixed("api validation configuration")

	if c.Source == "" {
		err = err.Append(errInvalidSource)
	}

	if c.Host == "" {
		err = err.Append(errInvalidHost)
	}

	if c.Port == "" {
		err = err.Append(errInvalidPort)
	}

	if c.Client == nil {
		err = err.Append(errInvalidClient)
	}

	return err.ErrorOrNil()
}

// Run runs a series of validation requests based off a specified api spec against
// the specified validation proxy address.
func Run(config Config) (int, error) {
	if err := config.Validate(); err != nil {
		return codeFailedConfigValidation, err
	}

	destinationFile := defaultDestinationFile
	hostAddress := fmt.Sprint(config.Host, ":", config.Port)
	cloudSpec, code, err := decodeFile(config.Source, destinationFile, config.Client)
	if err != nil {
		return code, err
	}

	if err := apivalidator.NewHTTPRequests(hostAddress, config.Client, cloudSpec); err != nil {
		cleanupFile(destinationFile)
		return codeFailedAPISpecValidation, err
	}

	defer cleanupFile(destinationFile)

	return 0, nil
}

func cleanupFile(file string) {
	if fileExists(file) && file == defaultDestinationFile {
		os.Remove(file)
	}
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func downloadFile(address, file string, client *http.Client) error {
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := client.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func decodeFile(sourceFile, destinationFile string, client *http.Client) (*spec.Swagger, int, error) {
	// Attempts to parse sourceFile as a URL. If it fails, will treat it as a relative path.
	// Otherwise file is downloaded from its location.
	_, err := url.ParseRequestURI(sourceFile)
	if err != nil {
		destinationFile = sourceFile
	} else if err := downloadFile(sourceFile, destinationFile, client); err != nil {
		cleanupFile(destinationFile)
		return nil, codeFailedDownloadingSpec, err
	}

	f, err := os.Open(destinationFile)
	if err != nil {
		cleanupFile(destinationFile)
		return nil, codeCannotReadFile, err
	}
	defer f.Close()

	var cloudSpec *spec.Swagger
	if err := json.NewDecoder(f).Decode(&cloudSpec); err != nil {
		cleanupFile(destinationFile)
		return nil, codeFailedUnmarshalingSpec, err
	}

	return cloudSpec, 0, nil
}
