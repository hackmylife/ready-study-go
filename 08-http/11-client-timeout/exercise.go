package koan

import (
	"net/http"
	"time"
)

func NewClient(timeout time.Duration) (*http.Client, error) { // TODO: 通信時間に上限を設ける
	return &http.Client{}, nil
}
