package koan

import (
	"reflect"
	"testing"
	"time"
)

func TestForward(t *testing.T) {
	in := make(chan int, 2)
	in <- 8
	in <- 9
	close(in)
	out := Forward(in)
	var got []int
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	for {
		select {
		case v, ok := <-out:
			if !ok {
				same(t, got, []int{8, 9})
				return
			}
			got = append(got, v)
		case <-timer.C:
			t.Fatal("output did not close")
		}
	}
}
func same(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v; want %#v", got, want)
	}
}
