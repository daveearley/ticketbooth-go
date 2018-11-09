package repository

type transactionRepository struct {
}

type TransactionRepository interface {
}

func NewTransactionRepository() *transactionRepository {
	return &transactionRepository{}
}
