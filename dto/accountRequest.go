package dto

import (
	"strings"

	"github.com/caiquenoboa/go-banking/errs"
)

const SAVING = "saving"
const CHECKING = "checking"

type AccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (a AccountRequest) IsAccountTypeSaving() bool {
	return strings.ToLower(a.AccountType) == SAVING
}

func (a AccountRequest) IsAccountTypeChecking() bool {
	return strings.ToLower(a.AccountType) == CHECKING
}

func (a AccountRequest) Validate() *errs.AppError {
	if a.Amount < 5000 {
		return errs.NewValidationError("To open a new account, you need to deposite at least $5000.00")
	}
	if !a.IsAccountTypeSaving() && !a.IsAccountTypeChecking() {
		return errs.NewValidationError("Account Type should be saving or checking")
	}
	return nil
}
