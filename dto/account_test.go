package dto

import (
	"net/http"
	"testing"
)

func Test_shuld_return_error_when_account_amount_is_less_than_5000(t *testing.T) {
	// Arrange
	request := AccountRequest{
		Amount: 200,
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "To open a new account, you need to deposite at least $5000.00" {
		t.Error("Invalid message while testing account amount")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing account amount")
	}
}

func Test_shuld_return_error_when_account_type_is_not_saving_or_checking(t *testing.T) {
	// Arrange
	request := AccountRequest{
		Amount:      6000,
		AccountType: "invalid account type",
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Account Type should be saving or checking" {
		t.Error("Invalid message while testing account type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing account type")
	}
}

func Test_shuld_return_nil_when_account_is_correct(t *testing.T) {
	// Arrange
	request := AccountRequest{
		Amount:      6000,
		AccountType: SAVING,
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError != nil {
		t.Error("Invalid error while testing account")
	}
}
