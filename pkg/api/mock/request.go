package mock

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/elastic/cloud-sdk-go/pkg/multierror"
)

// RequestAssertion is used to assert the contents of the request passed to an
// http.RoundTripper.
type RequestAssertion struct {
	Body   io.ReadCloser
	Header http.Header
	Method string
	Path   string
	Host   string
}

// AssertRequest ensures that a RequestAssertion matches certain *http.Request
// fields. If they do not match, an error is return.
func AssertRequest(want *RequestAssertion, req *http.Request) error {
	var merr = multierror.NewPrefixed("request assertion")
	if req.Body != nil || want.Body != nil {
		if !reflect.DeepEqual(want.Body, req.Body) {
			var wantB []byte
			if want.Body != nil {
				wantB, _ = ioutil.ReadAll(
					io.TeeReader(want.Body, new(bytes.Buffer)),
				)
			}
			var gotB []byte
			if req.Body != nil {
				gotB, _ = ioutil.ReadAll(
					io.TeeReader(req.Body, new(bytes.Buffer)),
				)
			}
			if !reflect.DeepEqual(wantB, gotB) {
				merr = merr.Append(
					fmt.Errorf("got body %s, want %s", gotB, wantB),
				)
			}

		}
	}

	if !reflect.DeepEqual(want.Header, req.Header) {
		merr = merr.Append(fmt.Errorf(
			"headers do not match: %v != %v", want.Header, req.Header),
		)
	}

	if !reflect.DeepEqual(want.Method, req.Method) {
		merr = merr.Append(fmt.Errorf(
			"methods do not match: %s != %s", want.Method, req.Method),
		)
	}

	if req.URL != nil && !reflect.DeepEqual(want.Path, req.URL.Path) {
		merr = merr.Append(fmt.Errorf(
			"paths do not match: %s != %s", want.Path, req.URL.Path),
		)
	}

	if !reflect.DeepEqual(want.Host, req.Host) {
		merr = merr.Append(fmt.Errorf(
			"hosts do not match: %s != %s", want.Host, req.Host),
		)
	}

	return merr.ErrorOrNil()
}
