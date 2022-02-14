package domain

import "github.com/caiquenoboa/go-banking/errs"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	NewTransaction(Transaction) (*Transaction, *errs.AppError)
}
