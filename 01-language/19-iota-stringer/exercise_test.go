package koan

import (
	"fmt"
	"testing"
)

func TestStatusValues(t *testing.T) {
	equal(t, int(StatusPending), 0)
	equal(t, int(StatusActive), 1)
	equal(t, int(StatusClosed), 2)
}

func TestStatusString(t *testing.T) {
	for _, tt := range []struct {
		status Status
		want   string
	}{
		{StatusPending, "pending"},
		{StatusActive, "active"},
		{StatusClosed, "closed"},
		{Status(7), "Status(7)"},
	} {
		equal(t, fmt.Sprint(tt.status), tt.want)
	}
}
func equal[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v; want %v", got, want)
	}
}
