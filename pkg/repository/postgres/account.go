package repository

import (
	"github.com/daveearley/product/pkg/model"
	"github.com/jinzhu/gorm"
)

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) *AccountRepository {
	return &AccountRepository{conn}
}

func (r *AccountRepository) GetById(id uint64) (*model.Account, error) {
	ac := &model.Account{}

	if err := r.Db.Preload("Users").First(ac, id).Error; err != nil {
		return nil, err
	}

	return ac, nil
}

func (r *AccountRepository) Store(a *model.Account) (*model.Account, error) {
	if err := r.Db.Create(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}
