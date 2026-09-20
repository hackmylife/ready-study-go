//go:build ignore

package koan

func Forward(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v
		}
	}()
	return out
}
