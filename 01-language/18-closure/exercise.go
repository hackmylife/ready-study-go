package koan

func Counter() func() int { // TODO: 呼び出すたびに1ずつ増える値を返す関数を作る
	return func() int { return 0 }
}
