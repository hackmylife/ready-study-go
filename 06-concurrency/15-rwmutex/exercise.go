package koan

type Scores struct { // TODO: 読み取り同士が同時に進めるように保護する
}

func (s *Scores) Set(name string, score int) { // TODO: 得点を登録する
}

func (s *Scores) Get(name string) (int, bool) { // TODO: 得点を読む
	return 0, false
}

func (s *Scores) View(fn func(scores map[string]int)) { // TODO: 読み取り中の内容をfnへ渡す
}
