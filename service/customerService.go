package service

import (
	"github.com/caiquenoboa/go-banking/domain"
	"github.com/caiquenoboa/go-banking/dto"
	"github.com/caiquenoboa/go-banking/errs"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetById(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	c, err := s.repo.FindAll(status)
	customerResponses := make([]dto.CustomerResponse, len(c))
	if err != nil {
		return nil, err
	}
	for _, r := range c {
		customerResponses = append(customerResponses, r.ToResponse())
	}
	return customerResponses, nil
}

func (s DefaultCustomerService) GetById(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	customerResponse := c.ToResponse()
	return &customerResponse, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
