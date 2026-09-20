package koan

import (
	"net/http"
)

func Routes() http.Handler { // TODO: path parameterを読む
	return http.NewServeMux()
}
