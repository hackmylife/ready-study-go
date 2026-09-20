package koan

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) { return 0, io.ErrClosedPipe }
func TestBatch(t *testing.T) {
	var b bytes.Buffer
	noError(t, Run(strings.NewReader("name,quantity\n TEA ,2\ncoffee,3\n"), &b, map[string]int64{"tea": 100, "coffee": 150}))
	equal(t, b.String(), "name,quantity,total\ntea,2,200\ncoffee,3,450\n")
}
func TestBatchInvalid(t *testing.T) {
	for _, input := range []string{"bad,header\na,1", "name,quantity\ntea,0", "name,quantity\ntea,x", "name,quantity\nunknown,1", "name,quantity\ntea,9223372036854775807"} {
		wantError(t, Run(strings.NewReader(input), io.Discard, map[string]int64{"tea": 100}))
	}
	if !errors.Is(Run(strings.NewReader("name,quantity\ntea,1"), failingWriter{}, map[string]int64{"tea": 100}), io.ErrClosedPipe) {
		t.Fatal("flush error lost")
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
func wantError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error")
	}
}
