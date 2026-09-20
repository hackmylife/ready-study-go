package koan

func Queue(values []int) <-chan int { // TODO: bufferへ格納する
	ch := make(chan int)
	close(ch)
	return ch
}
