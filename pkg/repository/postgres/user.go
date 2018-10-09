package repository

import (
	"database/sql"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetById(id int) (*models.User, error) {
	user, err := models.FindUser(r.db, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Store(a *models.User) (*models.User, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return a, err
	}

	return a, nil
}
