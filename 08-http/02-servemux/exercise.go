package koan

import (
	"net/http"
)

func Routes() http.Handler { // TODO: ルートを登録する
	return http.NewServeMux()
}
