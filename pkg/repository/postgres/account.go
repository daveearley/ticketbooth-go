package repository

import (
	"database/sql"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(conn *sql.DB) *AccountRepository {
	return &AccountRepository{conn}
}

func (r *AccountRepository) GetById(id int) (*models.Account, error) {
	ac, err := models.FindAccount(r.db, id)

	if err != nil {
		return nil, err
	}

	return ac, nil
}

func (r *AccountRepository) Store(a *models.Account) (*models.Account, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return a, nil
}
