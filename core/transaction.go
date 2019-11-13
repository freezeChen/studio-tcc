package core

type TransactionType int

const (
	TransactionRoot   TransactionType = 1
	TransactionBranch TransactionType = 2
)

type Transaction struct {
	Id     string
	Type   TransactionType
	Status int
}


