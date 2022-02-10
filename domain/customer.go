package domain

import (
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToResponse() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.StatusAsText(),
	}
}
