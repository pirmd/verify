package verify

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// MockStdout is an helper to capture output to Stdout.
type MockStdout struct {
	orig     *os.File
	r, w     *os.File
	captured chan string
}

// NewMockStdout creates a new record session of os.Sdtout and start capture
// operation.
func NewMockStdout() *MockStdout {
	return &MockStdout{
		orig:     os.Stdout,
		captured: make(chan string),
	}
}

// StartMockStdout creates a new record session of os.Sdtout and start capture
// operation.
func StartMockStdout() (*MockStdout, error) {
	out := NewMockStdout()
	if err := out.Start(); err != nil {
		return nil, err
	}
	return out, nil
}

// Start starts capturing os.Stdout.
func (out *MockStdout) Start() error {
	var err error

	out.r, out.w, err = os.Pipe()
	if err != nil {
		return fmt.Errorf("fail to capture Stdout: %v", err)
	}

	os.Stdout = out.w

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, out.r)
		out.captured <- buf.String()
	}()

	return nil
}

// Stop terminates a os.Stdout's recording session and restore os.Stdout to its
// original state
func (out *MockStdout) Stop() {
	out.w.Close()
	os.Stdout = out.orig
}

// String stops the os.Stdout's recording and returns the content of os.Stdout
// capture session.
func (out *MockStdout) String() string {
	out.Stop()
	return <-out.captured
}

// EqualStdoutString verifies that captured os.Stdout is equal to 'want' and
// feedback a test error message with a line by line diff between them
func EqualStdoutString(got *MockStdout, want string) error {
	return Equal(got.String(), want)
}

// MatchStdoutGolden compares captured os.Stdout to the content of a 'golden'
// file If 'update' command flag is used, update the 'golden' file
func MatchStdoutGolden(name string, got *MockStdout) error {
	return MatchGolden(name, got.String())
}
