package core

type transactionManager struct {
	repository TransactionRepository
}

func newTransactionManager(repository TransactionRepository) *transactionManager {
	return &transactionManager{repository: repository}
}
