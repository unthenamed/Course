package service

import (
	"api-laundy/model"
	"api-laundy/repo"
)

type transactionService struct {
	repo repo.TransactionRepo
}

func (t *transactionService) CreateNewTransaction(transaction model.Transaction) (model.Transaction, error) {
	return t.repo.CreateNewTransaction(transaction)
}

type TransactionService interface {
	CreateNewTransaction(model.Transaction) (model.Transaction, error)
}

func NewTransactionService(repo repo.TransactionRepo) TransactionService {
	return &transactionService{repo: repo}
}
