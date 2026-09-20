package koan

import (
	"errors"
	"net/http"
)

var ErrNotFound = errors.New("not found")

func WriteError(w http.ResponseWriter, err error) { // TODO: 公開用エラーへ変換する
}
