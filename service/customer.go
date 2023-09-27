package service

type CustomerResponse struct {
	CustomerID int    `json:"costomer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CostomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomerById(int) (*CustomerResponse, error)
}
