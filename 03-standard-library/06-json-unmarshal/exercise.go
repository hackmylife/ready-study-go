package koan

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Decode(data []byte) (User, error) { // TODO: JSONを読む
	return User{}, nil
}
