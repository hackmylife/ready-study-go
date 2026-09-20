package koan

func Use(work func() error, close func()) error { // TODO: 後処理を保証する
	return work()
}
