package repository

import "errors"

type customerRepositoryMock struct {
	customer []Customer
}

func NewCostomerRepositoryMock() customerRepositoryMock {
	customers := []Customer{
		{CustomerID: 1, Name: "John Doe Dorver", DateOfBirth: "1990-05-15", City: "New York", ZipCode: "10001", Status: 1},
		{CustomerID: 2, Name: "Jane Smith", DateOfBirth: "1985-08-20", City: "Los Angeles", ZipCode: "90001", Status: 2},
		{CustomerID: 3, Name: "Bob Johnson", DateOfBirth: "1978-03-10", City: "Chicago", ZipCode: "60601", Status: 1},
		{CustomerID: 4, Name: "Alice Brown", DateOfBirth: "1995-11-25", City: "Houston", ZipCode: "77001", Status: 3},
		{CustomerID: 5, Name: "Charlie Wilson", DateOfBirth: "1982-07-02", City: "San Francisco", ZipCode: "94101", Status: 2},
	}
	return customerRepositoryMock{customer: customers}
}

func (m customerRepositoryMock) GetAll() ([]Customer, error) {
	return m.customer, nil
}

func (m customerRepositoryMock) GetById(id int) (*Customer, error) {
	for _, customer := range m.customer {
		if customer.CustomerID == id {
			return &customer, nil
		}
	}

	return nil, errors.New("customer not found")
}
