package koan

func Reduce[T, U any](values []T, initial U, f func(U, T) U) U {
	// TODO: 初期値から順番に値を畳み込む。
	return initial
}
