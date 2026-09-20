package koan

func Start(work func()) <-chan struct{} { // TODO: 非同期に実行して完了を通知する
	done := make(chan struct{})
	close(done)
	return done
}
