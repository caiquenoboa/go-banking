package service

import (
	"time"

	"github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type TransactionService interface {
	NewTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo           domain.TransactionRepository
	accountService AccountService
}

func (d DefaultTransactionService) NewTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {

	err := request.Validate()

	if err != nil {
		return nil, err
	}

	transaction := domain.Transaction{
		TransactionId:   "",
		AccountId:       request.AccountId,
		Amount:          request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	accountAmountUpdated, err := d.accountService.UpdateBalanceAmount(request)
	if err != nil {
		return nil, err
	}

	transactionId, err := d.repo.NewTransaction(transaction)
	if err != nil {
		return nil, err
	}

	transactionResponse := dto.TransactionResponse{
		TransactionId:         transactionId.TransactionId,
		UpdatedBalanceAccount: accountAmountUpdated,
	}

	return &transactionResponse, nil
}

func NewDefaultTransactionService(repo domain.TransactionRepository, accountService AccountService) DefaultTransactionService {
	return DefaultTransactionService{
		repo:           repo,
		accountService: accountService,
	}
}
