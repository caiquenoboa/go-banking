package dto

import "github.com/caiquenoboa/go-banking/errs"

type TransactionRequest struct {
	AccountId       string
	Amount          float64
	TransactionType string
}

func (t TransactionRequest) Validate() *errs.AppError {
	if t.Amount < 0 {
		return errs.NewValidationError("Amount cannot be negative")
	}
	return nil
}
