package koan

import (
	"net/http"
)

type Note struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func NewHandler() http.Handler { // TODO: 独立した状態を持つJSON APIを作る
	return http.NewServeMux()
}
