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

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/ghodss/yaml"

	"github.com/elastic/cloud-sdk-go/internal/pkg/changelogger"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

var description = `
Generates a complete changelog aggregating individual changelog files from a specific folder.
The -changelog-dir flag's value is joined with -version, resulting: ${changelog-dir}/${version}.
`[1:]

type config struct {
	// flags
	dir      string
	version  string
	template string
	baseURL  string

	// device where the result will be written.
	out io.Writer
}

type appError struct {
	err  error
	code int
}

func (e *appError) Error() string {
	if e == nil || e.err == nil {
		return ""
	}
	return e.err.Error()
}

func main() {
	flagSet := flag.NewFlagSet("changelogger", flag.ContinueOnError)
	ogUsage := flagSet.Usage
	flagSet.Usage = func() {
		fmt.Fprintln(flagSet.Output(), description)
		ogUsage()
	}

	//  Flag Definition

	cfg := config{
		out: os.Stdout,
	}
	flagSet.StringVar(&cfg.dir, "changelog-dir", ".changelog", "path to the changelog directory")
	flagSet.StringVar(&cfg.version, "version", "", "version for the changelog being generated. Any 'v' prefix will be stripped")
	flagSet.StringVar(&cfg.template, "template", "", "template to generate the resulting changelog")
	flagSet.StringVar(&cfg.baseURL, "base-url", "", "base URL to use for each of the changes")

	//  Flag parsing

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Application wrapper

	if err := run(cfg); err != nil {
		code := 255
		fmt.Fprintln(os.Stderr, err)
		if e, ok := err.(*appError); ok {
			code = e.code
		}
		os.Exit(code)
	}
}

func run(cfg config) error {
	// config Validation
	if err := validateConfig(cfg); err != nil {
		return &appError{err: err, code: 1}
	}

	// Template parsing

	tplBytes, err := os.ReadFile(cfg.template)
	if err != nil {
		return &appError{
			err:  fmt.Errorf("failed opening template file: %w", err),
			code: 2,
		}
	}

	funcMap := template.FuncMap{
		"GitHubTracker": func(id string) string {
			baseURL := strings.TrimSuffix(cfg.baseURL, "/")
			return fmt.Sprintf("%s/issues/%s", baseURL, id)
		},
		"Version": func() string { return cfg.version },
		"Env":     os.Getenv,
	}
	tpl, err := template.New("changelog").Funcs(funcMap).Parse(string(tplBytes))
	if err != nil {
		return &appError{
			err:  fmt.Errorf("failed parsing template file contents: %w", err),
			code: 3,
		}
	}

	// Trim version prefix and walk the path.

	cleanVersion := cfg.version
	if strings.HasPrefix(cleanVersion, "v") {
		cleanVersion = strings.Replace(cleanVersion, "v", "", 1)
	}

	var changes changelogger.Changes
	changeValErr := multierror.NewPrefixed("invalid changelog entries")
	dir := filepath.Join(cfg.dir, cleanVersion)
	if err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed opening file %s: %w", path, err)
		}

		var change changelogger.Change
		if err := yaml.Unmarshal(b, &change); err != nil {
			return fmt.Errorf("failed decoding yaml file %s: %w", path, err)
		}

		// If there's no reference set in the file, use the file name.
		if change.Ref == "" {
			change.Ref = strings.Replace(info.Name(),
				filepath.Ext(info.Name()), "", 1,
			)
		}

		if !reflect.DeepEqual(change, changelogger.Change{}) {
			changes = append(changes, change)
		}

		if validateErr := change.Validate(info.Name()); validateErr != nil {
			changeValErr = changeValErr.Append(validateErr)
		}

		return nil
	}); err != nil {
		return &appError{
			err:  fmt.Errorf("failed walking the specified path: %w", err),
			code: 4,
		}
	}
	sort.Sort(changes)

	if len(changes) == 0 {
		return &appError{
			err:  fmt.Errorf("folder %s has no changelog files", dir),
			code: 5,
		}
	}

	if err := changeValErr.ErrorOrNil(); err != nil {
		return &appError{
			err:  err,
			code: 7,
		}
	}

	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, changes); err != nil {
		return &appError{
			err:  fmt.Errorf("failed executing the changelog template: %w", err),
			code: 8,
		}
	}

	if _, err := io.Copy(cfg.out, buf); err != nil {
		return &appError{
			err:  fmt.Errorf("failed copying the template output: %w", err),
			code: 9,
		}
	}

	return nil
}

func validateConfig(cfg config) error {
	merr := multierror.NewPrefixed("invalid flags")
	if cfg.version == "" {
		merr = merr.Append(errors.New("version cannot be empty"))
	}
	if cfg.baseURL == "" {
		merr = merr.Append(errors.New("base-url cannot be empty"))
	}
	if cfg.template == "" {
		merr = merr.Append(errors.New("template cannot be empty"))
	}
	if cfg.out == nil {
		merr = merr.Append(errors.New("out io.writer cannot be nil"))
	}

	return merr.ErrorOrNil()
}
