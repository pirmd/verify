package verify

import (
	"log"
	"testing"
)

type tbWriter struct {
	testing.TB
}

func (tbw tbWriter) Write(p []byte) (int, error) {
	tbw.Helper()
	tbw.Logf("%s", p)
	return len(p), nil
}

// NewLogger returns a new logger that logs to the provided testing.TB
func NewLogger(tb testing.TB) *log.Logger {
	tb.Helper()
	return log.New(tbWriter{TB: tb}, tb.Name()+" ", log.LstdFlags|log.Lshortfile|log.LUTC)
}
