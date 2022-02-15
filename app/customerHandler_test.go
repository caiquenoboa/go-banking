package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
	"github.com/caiquenoboa/go-banking/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var (
	router      *mux.Router
	ch          CustomerHandlers
	mockService *service.MockCustomerService
)

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = service.NewMockCustomerService(ctrl)

	ch = CustomerHandlers{mockService}

	router = mux.NewRouter()

	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	//Arrange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Ashish", "New Delhi", "110011", "2000-01-01", "1"},
		{"1002", "Rob", "New Delhi", "110011", "2000-01-01", "1"},
		{"1003", "Caique", "Curitiba", "82630492", "1999-11-03", "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}

}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some data"))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	//Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}

}
