package accounts

import (
	"database/sql"
	"fmt"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type Repository interface {
	GetById(id int) (*models.Account, error)
	Store(a *models.Account) (*models.Account, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn}
}

func (r *repository) GetById(id int) (*models.Account, error) {
	ac, err := models.FindAccount(r.db, id)

	if err != nil {
		fmt.Println("not found", err.Error(), id)
		return nil, err
	}

	return ac, nil
}

func (r *repository) Store(a *models.Account) (*models.Account, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return a, nil
}
