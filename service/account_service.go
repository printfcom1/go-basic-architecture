package service

import (
	"strings"
	"time"

	"github.com/go-architecture/errs"
	"github.com/go-architecture/logs"
	"github.com/go-architecture/repository"
)

type accountService struct {
	accRepo repository.AccountRepositiry
}

func NewAccountService(accRepo repository.AccountRepositiry) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, newAcc NewAccountRequest) (*AccountResponse, error) {

	if newAcc.Amount < 5000 {
		return nil, errs.NewValidationError("amount at least 5,000")
	}

	if strings.ToLower(newAcc.AccountType) != "saving" && strings.ToLower(newAcc.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      1,
	}

	newAccount, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}

	response := AccountResponse{
		AccountID:   newAccount.AccountID,
		OpeningDate: newAccount.OpeningDate,
		AccountType: newAccount.AccountType,
		Amount:      newAccount.Amount,
		Status:      newAccount.Status,
	}
	return &response, nil
}

func (s accountService) GetAccount(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewInternalServerError()
	}

	response := []AccountResponse{}

	for _, account := range accounts {
		response = append(response, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return response, nil
}
