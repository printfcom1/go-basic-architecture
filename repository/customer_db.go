package repository

import "github.com/jmoiron/sqlx"

type costomerRepositoryDB struct {
	db *sqlx.DB
}

func NewCostomerRepositoryDB(db *sqlx.DB) costomerRepositoryDB {
	return costomerRepositoryDB{db: db}
}

func (r costomerRepositoryDB) GetAll() ([]Customer, error) {
	customer := []Customer{}
	query := "select costomer_id, name ,date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customer, query)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (r costomerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select costomer_id, name ,date_of_birth, city, zipcode, status from customers where costomer_id =?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
