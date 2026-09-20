//go:build ignore

package koan

func Start(work func()) <-chan struct{} {
	done := make(chan struct{})
	go func() { defer close(done); work() }()
	return done
}
