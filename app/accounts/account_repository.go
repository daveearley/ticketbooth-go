package accounts

import (
	"database/sql"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type Repository interface {
	GetByID(id int) (*models.Account, error)
	Store(a *models.Account) (*models.Account, error)
	DeleteByID(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(conn *sql.DB) Repository {
	return &repository{conn}
}

func (r *repository) GetByID(id int) (*models.Account, error) {
	ac, err := models.FindAccount(r.db, id)

	if err != nil {
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

func (r *repository) DeleteByID(id int) error {
	account, err := models.FindAccount(r.db, id)

	if err != nil {
		return err
	}

	_, err = account.Delete(r.db)

	return err
}
