package repository

import (
	"database/sql"
	"fmt"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(conn *sql.DB) *accountRepository {
	return &accountRepository{conn}
}

func (r *accountRepository) GetById(id int) (*models.Account, error) {
	ac, err := models.FindAccount(r.db, id)

	if err != nil {
		fmt.Println("not found", err.Error(), id)
		return nil, err
	}

	return ac, nil
}

func (r *accountRepository) Store(a *models.Account) (*models.Account, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return a, nil
}
