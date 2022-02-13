package app

import (
	"encoding/json"
	"net/http"

	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/service"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (a AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.AccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		response, appError := a.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, response)
		}
	}

}

func (a AccountHandler) transaction(w http.ResponseWriter, r *http.Request) {

}
