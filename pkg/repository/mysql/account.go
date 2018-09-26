package repository

import (
	"github.com/daveearley/product/pkg/model"
	"github.com/daveearley/product/pkg/repository"
	"github.com/jinzhu/gorm"
)

type AccountRepository struct {
	Conn *gorm.DB
}

func NewAccountRepository(conn *gorm.DB) repository.AccountRepository {
	return &AccountRepository{conn}
}

func (r *AccountRepository) GetById(id uint64) (*model.Account, error)  {
	return &model.Account{
		Id: 123,
	}, nil
}

func (r *AccountRepository) Store(a *model.Account) error {
	if err := r.Conn.Create(&a).Error; err != nil {
		return err
	}

	return nil
}