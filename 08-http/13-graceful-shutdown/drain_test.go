package koan

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

type observedListener struct {
	net.Listener
	closed chan struct{}
	once   sync.Once
}

func (l *observedListener) Close() error {
	err := l.Listener.Close()
	l.once.Do(func() { close(l.closed) })
	return err
}

func TestStopDrainsInFlightRequest(t *testing.T) {
	started := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	releaseHandler := func() { releaseOnce.Do(func() { close(release) }) }
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		close(started)
		<-release
		_, _ = io.WriteString(w, "completed")
	}))
	listener := &observedListener{Listener: server.Listener, closed: make(chan struct{})}
	server.Listener = listener
	server.Start()
	t.Cleanup(func() { releaseHandler(); server.Close() })
	type response struct {
		body string
		err  error
	}
	clientDone := make(chan response, 1)
	go func() {
		resp, err := server.Client().Get(server.URL)
		if err != nil {
			clientDone <- response{err: err}
			return
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		clientDone <- response{body: string(body), err: err}
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("request did not start")
	}
	stopDone := make(chan error, 1)
	go func() { stopDone <- Stop(server.Config, 3*time.Second) }()
	select {
	case <-listener.closed:
	case <-time.After(time.Second):
		t.Fatal("server did not stop accepting requests")
	}
	select {
	case err := <-stopDone:
		t.Fatalf("Stop returned while the handler was active: %v", err)
	default:
	}
	releaseHandler()
	select {
	case result := <-clientDone:
		noError(t, result.err)
		if result.body != "completed" {
			t.Fatalf("response = %q; want completed", result.body)
		}
	case <-time.After(time.Second):
		t.Fatal("in-flight response did not complete")
	}
	select {
	case err := <-stopDone:
		noError(t, err)
	case <-time.After(time.Second):
		t.Fatal("Stop did not return after the handler completed")
	}
}
