package repository

import "github.com/daveearley/product/pkg/model"

type AccountRepositoryI interface {
	GetById(id string) (*model.Account, error)
	Store(a *model.Account) error
}
