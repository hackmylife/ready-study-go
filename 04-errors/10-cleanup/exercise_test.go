package koan

import (
	"errors"
	"io"
	"testing"
)

type source struct {
	readErr, closeErr error
	closed            int
}

func (s *source) Read([]byte) (int, error) {
	if s.readErr != nil {
		return 0, s.readErr
	}
	return 0, io.EOF
}
func (s *source) Close() error { s.closed++; return s.closeErr }
func TestCleanup(t *testing.T) {
	a, b := errors.New("read"), errors.New("close")
	for _, tt := range []struct{ r, c error }{{nil, nil}, {a, nil}, {nil, b}, {a, b}} {
		s := &source{readErr: tt.r, closeErr: tt.c}
		_, err := ReadAndClose(s)
		equal(t, s.closed, 1)
		if tt.r == nil && tt.c == nil {
			noError(t, err)
		}
		for _, cause := range []error{tt.r, tt.c} {
			if cause != nil && !errors.Is(err, cause) {
				t.Fatalf("lost %v in %v", cause, err)
			}
		}
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
