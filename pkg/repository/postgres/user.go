package repository

import (
	"github.com/daveearley/product/pkg/model"
	"database/sql"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) GetById(id uint64) (*model.User, error) {
	ac := &model.User{}

	if err := r.Db.Preload("Users").First(ac, id).Error; err != nil {
		return nil, err
	}

	return ac, nil
}

func (r *UserRepository) Store(a *model.User) (*model.User, error) {
	if err := r.Db.Create(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}
