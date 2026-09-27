package koan

type Status int

const ( // TODO: iotaで0から始まる連番にする
	StatusPending Status = 0
	StatusActive  Status = 0
	StatusClosed  Status = 0
)

func (s Status) String() string { // TODO: 状態名を返す
	return ""
}
