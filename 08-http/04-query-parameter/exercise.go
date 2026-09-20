package koan

import (
	"net/http"
)

func Limit(r *http.Request) (int, error) { // TODO: queryを検証する
	return 20, nil
}
