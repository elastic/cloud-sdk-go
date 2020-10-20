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

package api

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/elastic/cloud-sdk-go/pkg/api/mock"
	"github.com/elastic/cloud-sdk-go/pkg/models"
	"github.com/elastic/cloud-sdk-go/pkg/multierror"
	"github.com/elastic/cloud-sdk-go/pkg/util/ec"
)

const userAgentHeader = "User-Agent"

// DefaultUserAgent is used in UserAgentTransport when the agent is not set.
// It defaults to the project name + the current committed version.
var DefaultUserAgent = "cloud-sdk-go/" + Version

var defaultBackoff = time.Second * 1

// CustomTransport is the cloud-sdk-go custom transport which handles all the
// add-ons for http request and responses. It supports:
// * Adding a custom UserAgent header to all outgoing requests.
// * Writing a trail of the request / response flow to a device (verbose).
// * Adding support for request retries on timeout with a backoff period.
type CustomTransport struct {
	rt http.RoundTripper

	// UserAgent settings
	agent string

	// Retry settings
	retries int
	backoff time.Duration

	// Verbose settings
	verbose    bool
	count      int64
	redactAuth bool
	writer     io.Writer
}

// CustomTransportCfg is used to configure a CustomTransport.
type CustomTransportCfg struct {
	// RoundTripper to use for http calls.
	RoundTripper http.RoundTripper

	// UserAgent if specified, it sets the user agent on all outgoing requests.
	UserAgent string

	// Number of retries to perform on request timeout.
	Retries int

	// Cooldown time between retried requests.
	Backoff time.Duration

	// Verbose settings
	Verbose    bool
	RedactAuth bool
	Writer     io.Writer
}

func (cfg CustomTransportCfg) validate() error {
	var merr = multierror.NewPrefixed("invalid custom transport settings")
	if cfg.RoundTripper == nil {
		merr = merr.Append(errors.New("roundtripper cannot be nil"))
	}

	if cfg.Verbose && cfg.Writer == nil {
		merr = merr.Append(errors.New("verbose set to true, but no writer has been set"))
	}

	return merr.ErrorOrNil()
}

// NewCustomTransport creates a new CustomTransport. See the structure's Goddoc
// to learn about what it does.
func NewCustomTransport(cfg CustomTransportCfg) (*CustomTransport, error) {
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	if cfg.UserAgent == "" {
		cfg.UserAgent = DefaultUserAgent
	}

	if cfg.Backoff.Microseconds() <= 0 {
		cfg.Backoff = defaultBackoff
	}

	return &CustomTransport{
		rt:         cfg.RoundTripper,
		agent:      cfg.UserAgent,
		retries:    cfg.Retries,
		backoff:    cfg.Backoff,
		verbose:    cfg.Verbose,
		redactAuth: cfg.RedactAuth,
		writer:     cfg.Writer,
	}, nil
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// UserAgent header handling
	req.Header.Set(userAgentHeader, t.agent)

	return t.doRoundTrip(req, t.retries)
}

// doRoundTrip performs an http call with the specified request and retries the
// request when the returned error is context.DeadlineExceeded (timeout).
func (t *CustomTransport) doRoundTrip(req *http.Request, retries int) (*http.Response, error) {
	count := atomic.AddInt64(&t.count, 1)
	if t.verbose {
		handleVerboseRequest(t.writer, req, count, t.redactAuth)
	}

	res, err := t.rt.RoundTrip(req)

	if t.verbose && res != nil {
		handleVerboseResponse(t.writer, res, count)
	}

	if e, ok := err.(net.Error); ok && e.Timeout() {
		if t.verbose {
			msg := "request timed out, retrying..."
			if retries <= 0 {
				msg = "request timed out, giving up."
			}
			fmt.Fprintln(t.writer, msg)
		}

		if retries > 0 {
			retries--
			// Recursively do the roundtrip and return the result
			<-time.After(backoff(t.backoff))
			return t.doRoundTrip(req, retries)
		}
	}

	if res != nil && !t.verbose {
		_, _ = httputil.DumpResponse(res, res.Body != nil)

		// When the content type is "text/html", a bit of tweaking is required
		// for the response to be marshaled to JSON. Using the standard error
		// definition and populating it with parts of the request so the error
		// can be identified.
		if strings.Contains(res.Header.Get(contentType), textHTMLContentType) {
			res.Header.Set(contentType, jsonContentType)
			res.Body = newProxyBody(req, res.StatusCode)
		}
	}

	return res, err
}

func handleVerboseRequest(w io.Writer, req *http.Request, c int64, redact bool) {
	b, _ := httputil.DumpRequestOut(req, req.Body != nil)

	fmt.Fprintf(w, "==================== Start of Request #%d ====================\n", c)
	fmt.Fprintln(w, redactAuth(string(b), redact))
	fmt.Fprintf(w, "====================  End of Request #%d  ====================\n", c)
}

func handleVerboseResponse(w io.Writer, res *http.Response, c int64) {
	b, _ := httputil.DumpResponse(res, res.Body != nil)
	fmt.Fprintf(w, "==================== Start of Response #%d ====================\n", c)
	fmt.Fprintln(w, string(b))
	fmt.Fprintf(w, "====================  End of Response #%d  ====================\n", c)
}

func redactAuth(request string, obscure bool) string {
	if !obscure {
		return request
	}

	return regexp.MustCompile(`Authorization: .*\n`).ReplaceAllString(
		request, "Authorization: [REDACTED]\r\n",
	)
}

func newProxyBody(req *http.Request, code int) io.ReadCloser {
	return mock.NewStructBody(models.BasicFailedReply{
		Errors: []*models.BasicFailedReplyElement{
			{
				Code:    ec.String(strconv.Itoa(code)),
				Fields:  []string{fmt.Sprintf("%s %s", req.Method, req.URL.EscapedPath())},
				Message: ec.String(http.StatusText(code)),
			},
		},
	})
}

func backoff(d time.Duration) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return d / time.Duration(rand.Float32()*10+1)
}
