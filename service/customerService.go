package service

import (
	"github.com/nitinkumar91296/banking/domain"
	"github.com/nitinkumar91296/banking/dto"
	errs "github.com/nitinkumar91296/banking/errors"
)

type CustomerService interface {
	GetAllCustomer(status string) ([]domain.Customer, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var resp []dto.CustomerResponse
	for _, c := range customers {
		resp = append(resp, c.ToDto())
	}

	return resp, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.FindCustomerById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return &resp, nil
}
