package service

import (
	"database/sql"

	"github.com/go-architecture/errs"
	"github.com/go-architecture/logs"
	"github.com/go-architecture/repository"
)

type customerService struct {
	custRepo repository.CostomerRepository
}

func NewCustomerService(custRopo repository.CostomerRepository) customerService {
	return customerService{custRepo: custRopo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}

func (s customerService) GetCustomerById(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("custormer not found")
		}
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}
	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
