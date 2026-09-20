package koan

import (
	"context"
	"net/http"
)

type Page struct {
	Items []string `json:"items"`
	Next  int      `json:"next"`
}

func List(ctx context.Context, client *http.Client, baseURL string, maxPages int) ([]string, error) { // TODO: 全ページを取得する
	return nil, nil
}
