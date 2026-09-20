package koan

import (
	"strings"
)

func Upper(s string) string { // TODO: 不要な非同期処理を削る
	ch := make(chan string)
	go func() { ch <- strings.ToUpper(s) }()
	return <-ch
}
