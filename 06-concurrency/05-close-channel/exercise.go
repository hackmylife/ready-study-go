package koan

func Forward(in <-chan int) <-chan int { // TODO: 所有する出力をcloseする
	out := make(chan int)
	close(out)
	return out
}
