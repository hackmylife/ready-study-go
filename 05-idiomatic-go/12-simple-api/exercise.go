package koan

import (
	"strings"
)

func Format(name string, loud bool) string {
	if loud {
		return strings.ToUpper("Hello, " + name)
	}
	return "Hello, " + name
}
func Greeting(name string) string { // TODO: 意図が分かるAPIにする
	return ""
}
func LoudGreeting(name string) string { return "" }
