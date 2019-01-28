package repository

import (
	"database/sql"
	"../../app"
	"../../app/models/generated"
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

//NewAccountRepository returns a new instance of accountRepository
func NewAccountRepository(conn *sql.DB) AccountRepository {
	return &accountRepository{conn}
}

//GetByID returns an account for a given ID
func (r *accountRepository) GetByID(id int) (*models.Account, error) {
	account, err := models.FindAccount(r.db, id)

	if err == nil {
		return account, nil
	}

	return nil, getErrorType(err, app.AccountResource, id)
}

//Store creates a new account
func (r *accountRepository) Store(a *models.Account) (*models.Account, error) {
	err := a.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, app.ServerError(err)
	}

	return a, nil
}

//DeleteByID deletes the account of the given ID
func (r *accountRepository) DeleteByID(id int) error {
	account, err := models.FindAccount(r.db, id)

	if err != nil {
		return getErrorType(err, app.AccountResource, id)
	}

	_, err = account.Delete(r.db)

	if err != nil {
		return app.ServerError(err)
	}

	return nil
}
