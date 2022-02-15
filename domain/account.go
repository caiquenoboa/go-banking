package domain

import (
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/caiquenoboa/go-banking/domain AccountRepository
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

const dbTSLayout = "2006-01-02 15:04:05"

func NewAccountFromRequest(req dto.AccountRequest) Account {
	return Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "",
	}
}

func (a Account) NewAccountResponse() *dto.AccountResponse {
	return &dto.AccountResponse{
		AccountId: a.AccountId,
	}
}
