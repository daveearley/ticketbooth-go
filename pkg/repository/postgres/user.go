package repository

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/daveearley/product/pkg/models/generated"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) GetById(id int) (*models.User, error) {
	user, err := models.FindUser(r.db, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Store(a *models.User) (*models.User, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return a, err
	}

	return a, nil
}
