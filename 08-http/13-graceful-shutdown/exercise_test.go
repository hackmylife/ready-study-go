package koan

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestStop(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	var once sync.Once
	defer once.Do(func() { close(release) })
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		<-release
		_, _ = io.WriteString(w, "done")
	}))
	defer ts.Close()
	clientDone := make(chan struct{})
	go func() {
		defer close(clientDone)
		resp, err := ts.Client().Get(ts.URL)
		if err == nil {
			_, _ = io.Copy(io.Discard, resp.Body)
			resp.Body.Close()
		}
	}()
	<-started
	err := Stop(ts.Config, 0)
	once.Do(func() { close(release) })
	<-clientDone
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("want deadline error, got %v", err)
	}
}
func TestStopIdle(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer ts.Close()
	noError(t, Stop(ts.Config, time.Second))
	resp, err := ts.Client().Get(ts.URL)
	if err == nil {
		resp.Body.Close()
		t.Fatal("server still accepts requests")
	}
}

func noError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
