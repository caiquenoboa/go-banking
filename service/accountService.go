package service

import (
	"time"

	"github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
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

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
