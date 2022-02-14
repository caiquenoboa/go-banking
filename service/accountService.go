package service

import (
	"time"

	"github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
	CheckBalance(dto.TransactionRequest, domain.Account) (bool, *errs.AppError)
	UpdateBalanceAmount(dto.TransactionRequest) (float64, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (d DefaultAccountService) NewAccount(request dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {

	err := request.Validate()

	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "",
	}

	newAccount, err := d.repo.Save(account)

	if err != nil {
		return nil, err
	}

	accountResponse := newAccount.NewAccountResponse()

	return &accountResponse, nil
}

func (d DefaultAccountService) CheckBalance(request dto.TransactionRequest, account domain.Account) (bool, *errs.AppError) {
	if account.Amount < request.Amount {
		return false, nil
	}

	return true, nil
}

func (d DefaultAccountService) UpdateBalanceAmount(request dto.TransactionRequest) (float64, *errs.AppError) {
	account, err := d.repo.GetAmountById(request.AccountId)
	if err != nil {
		return -1.0, err
	}

	var newAmount float64

	if request.TransactionType == "withdrawal" {
		hasAmount, err := d.CheckBalance(request, *account)
		if err != nil {
			return -1.0, err
		}
		if !hasAmount {
			return -1.0, errs.NewNotEnoughMoneyError()
		}
		newAmount = account.Amount - request.Amount
	}

	if request.TransactionType == "deposit" {
		newAmount = account.Amount + request.Amount
	}

	err = d.repo.UpdateBalanceAmountById(account.AccountId, newAmount)

	if err != nil {
		return -1.0, err
	}

	return newAmount, nil

}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
