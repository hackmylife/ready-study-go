package koan

type User struct {
	Name   string
	Age    int
	Secret string
}

func Encode(u User) ([]byte, error) { // TODO: JSONに変換する
	return nil, nil
}
