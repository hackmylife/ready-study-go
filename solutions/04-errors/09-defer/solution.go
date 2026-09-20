//go:build ignore

package koan

func Use(work func() error, close func()) error { defer close(); return work() }
