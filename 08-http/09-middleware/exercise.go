package koan

import (
	"net/http"
)

func RequestID(next http.Handler) http.Handler { // TODO: header処理を合成する
	return next
}
