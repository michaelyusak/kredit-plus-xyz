package service

import "github.com/michaelyusak/kredit-plus-xyz/users/repository"

type transactionServiceImpl struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *transactionServiceImpl {
	return &transactionServiceImpl{
		transactionRepository: transactionRepository,
	}
}
