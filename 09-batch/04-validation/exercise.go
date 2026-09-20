package koan

type Order struct {
	Name     string
	Quantity int
}

func Validate(order Order) error { // TODO: レコードを検証する
	return nil
}
