package koan

import (
	"net/http"
	"time"
)

func Stop(server *http.Server, timeout time.Duration) error { // TODO: 期限付きでgraceful shutdownする
	return nil
}
