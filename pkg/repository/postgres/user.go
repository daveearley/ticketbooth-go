package repository

import (
	"github.com/daveearley/product/pkg/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (r *UserRepository) GetById(id string) (*model.User, error) {
	ac := &model.User{}

	if err := r.Db.Preload("Users").First(ac, id).Error; err != nil {
		return nil, err
	}

	return ac, nil
}

func (r *UserRepository) Store(a *model.User) error {
	if err := r.Db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}
