package repository

import (
	"github.com/jmoiron/sqlx"
)

type accountRePositoryDB struct {
	db *sqlx.DB
}

func NewAccountRePositoryDB(db *sqlx.DB) accountRePositoryDB {
	return accountRePositoryDB{db: db}
}

func (r accountRePositoryDB) Create(acc Account) (*Account, error) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.DB.Exec(query, acc.CustomerID, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != err {
		return nil, err
	}

	acc.AccountID = int(id)
	return &acc, nil
}

func (r accountRePositoryDB) GetAll(customerID int) ([]Account, error) {
	query := "SELECT account_id, customer_id, opening_date, account_type, amount, status FROM accounts where customer_id =?"
	account := []Account{}

	err := r.db.Select(&account, query, customerID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
