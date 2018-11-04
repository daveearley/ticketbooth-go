package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type AccountRepository interface {
	GetByID(id int) (*models.Account, error)
	Store(a *models.Account) (*models.Account, error)
	DeleteByID(id int) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(conn *sql.DB) AccountRepository {
	return &accountRepository{conn}
}

func (r *accountRepository) GetByID(id int) (*models.Account, error) {
	ac, err := models.FindAccount(r.db, id)

	if err != nil {
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

func (r *accountRepository) DeleteByID(id int) error {
	account, err := models.FindAccount(r.db, id)

	if err != nil {
		return err
	}

	_, err = account.Delete(r.db)

	return err
}
