package service

import (
	"testing"

	realdomain "github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
	"github.com/caiquenoboa/go-banking/mocks/domain"
	"github.com/golang/mock/gomock"
)

func Test_should_return_a_validation_error_reponse_when_the_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.AccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}
	service := NewAccountService(nil)

	//Act
	_, appError := service.NewAccount(request)

	//Assert
	if appError == nil {
		t.Error("failed while testing the new account validation")
	}
}

var (
	mockRepo *domain.MockAccountRepository
	service  AccountService
)

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_error_from_the_server_side_if_the_new_account_cannot_be_created(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	request := dto.AccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  request.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unexpected database error"))

	// Act
	_, appError := service.NewAccount(request)

	//Assert
	if appError == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func Test_should_return_new_account_response_when_a_new_account_is_saved_successfully(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	request := dto.AccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      6000,
	}
	account := realdomain.Account{
		CustomerId:  request.CustomerId,
		OpeningDate: dbTSLayout,
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      "",
	}
	accountWithId := account
	accountWithId.AccountId = "201"

	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	// Act
	newAccount, appError := service.NewAccount(request)

	//Assert
	if appError != nil {
		t.Error("Test failed while creating new account")
	}
	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Test failed while creating new account")
	}

}
