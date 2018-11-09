package service

import "github.com/daveearley/ticketbooth/app/repository"

type transactionSrv struct {
	transRepo repository.TransactionRepository
}

type TransactionService interface {
}

func NewTransactionService(transRepo repository.TransactionRepository) *transactionSrv {
	return &transactionSrv{transRepo}
}
