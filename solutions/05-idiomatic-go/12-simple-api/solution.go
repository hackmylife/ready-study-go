//go:build ignore

package koan

import (
	"strings"
)

func Greeting(name string) string {
	return "Hello, " + name
}
func LoudGreeting(name string) string { return strings.ToUpper(Greeting(name)) }
