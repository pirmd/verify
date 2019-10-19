package verify

import (
	"bytes"
	"io"
	"os"
	"testing"
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
func StartMockStdout(tb testing.TB) *MockStdout {
	out := NewMockStdout()
	out.Start(tb)
	return out
}

// Start starts capturing os.Stdout.
func (out *MockStdout) Start(tb testing.TB) {
	var err error

	out.r, out.w, err = os.Pipe()
	if err != nil {
		tb.Fatalf("fail to capture Stdout: %v", err)
	}

	os.Stdout = out.w

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, out.r)
		out.captured <- buf.String()
	}()
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
func EqualStdoutString(tb testing.TB, got *MockStdout, want string, message ...string) {
	EqualString(tb, got.String(), want, message...)
}

// MatchStdoutGolden compares captured os.Stdout to the content of a 'golden'
// file If 'update' command flag is used, update the 'golden' file
func MatchStdoutGolden(tb testing.TB, got *MockStdout, message ...string) {
	MatchGolden(tb, got.String(), message...)
}
