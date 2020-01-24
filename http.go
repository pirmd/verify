package verify

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

var (
	// updateMockHTTP updates the http.Response files provided through MockHTTPResponse
	updateMockHTTP = flag.Bool("test.mockhttp-update", false, "update files served through the mock http transport")
	// where to find http.Response files
	mockHTTPDir = flag.String("test.mockhttpdir", "./testdata", "path to folder hosting golden files")
)

// MockHTTPResponse represents a mock http transport for testing purpose. It
// implements http.RoundTripper, which handles single http requests issued by a
// http.Client.
//
// MockHTTPResponse returns the content of a file from provided path whose
// filename is the requested url. If no file is found in path with requested
// url as filename, an empty http.Response if no corresponding file exists.
type MockHTTPResponse struct {
	origTransport http.RoundTripper
}

// NewMockHTTPResponse creates a new MockHTTPResponse.
func NewMockHTTPResponse() *MockHTTPResponse {
	return &MockHTTPResponse{}
}

// StartMockHTTPResponse creates a new MockHTTPResponse and starts it.
func StartMockHTTPResponse() *MockHTTPResponse {
	m := NewMockHTTPResponse()
	m.Start()
	return m
}

// Start actually starts the MockHTTP transport by replacing
// http.DefaultTransport by itself.
func (m *MockHTTPResponse) Start() {
	m.origTransport = http.DefaultTransport
	http.DefaultTransport = m
}

// Stop stops MockHTTPResponse and restore initial http.DefaultTransport.
func (m *MockHTTPResponse) Stop() {
	http.DefaultTransport = m.origTransport
}

// RoundTrip implements http.Roundtripper interface.
func (m *MockHTTPResponse) RoundTrip(req *http.Request) (*http.Response, error) {
	if *updateMockHTTP {
		if err := m.updateResponsesFiles(req); err != nil {
			return nil, fmt.Errorf("fail to update mock http response file (in %s): %v", m.pathFor(req), err)
		}
	}

	resp, err := m.respondFromFiles(req)
	if err != nil {
		return nil, fmt.Errorf("fail to mock http response (for %s): %v", req.URL, err)
	}

	return resp, err
}

func (m *MockHTTPResponse) respondFromFiles(req *http.Request) (*http.Response, error) {
	f, err := os.Open(m.pathFor(req))
	if err != nil {
		return nil, err
	}

	resp, err := http.ReadResponse(bufio.NewReader(f), nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *MockHTTPResponse) updateResponsesFiles(req *http.Request) error {
	m.Stop()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}

	path := m.pathFor(req)
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, b, 0666); err != nil {
		return err
	}

	m.Start()

	return nil
}

func (m *MockHTTPResponse) pathFor(req *http.Request) string {
	return filepath.Join(*mockHTTPDir, req.URL.Host, req.URL.Path, req.URL.RawQuery+".http")
}
