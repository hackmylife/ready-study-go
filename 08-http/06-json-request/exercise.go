package koan

import (
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) { // TODO: 本文を検証する
	w.WriteHeader(http.StatusNoContent)
}
