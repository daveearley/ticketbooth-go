package repository

import "github.com/daveearley/product/pkg/model"

type AccountRepository interface {
	GetById(id uint64) (*model.Account, error)
	Store(a *model.Account) error
}
