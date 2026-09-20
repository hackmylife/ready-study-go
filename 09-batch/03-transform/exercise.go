package koan

type Order struct {
	Name     string
	Quantity int
}

func Normalize(order Order) Order { // TODO: レコードを正規化する
	return order
}
