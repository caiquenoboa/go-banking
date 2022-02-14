package app

import (
	"encoding/json"
	"net/http"

	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (a TransactionHandler) newTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		response, appError := a.service.NewTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, response)
		}
	}

}
