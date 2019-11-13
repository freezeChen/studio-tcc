package core

import "time"

type TransactionRepository interface {
	Create(transaction Transaction) error
	Update(transaction Transaction) error
	Delete(transaction Transaction) error
	GetTransaction(id string) (Transaction, error)
	FindUnComplete(time time.Time) ([]Transaction, error)
}
