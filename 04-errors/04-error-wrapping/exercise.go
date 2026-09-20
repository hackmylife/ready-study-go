package koan

func Load(path string, read func(string) ([]byte, error)) ([]byte, error) { // TODO: 原因を保持して返す
	return read(path)
}
