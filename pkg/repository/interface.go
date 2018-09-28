package repository

import "github.com/daveearley/product/pkg/model"

type AccountRepositoryI interface {
	GetById(id uint64) (*model.Account, error)
	Store(a *model.Account) (*model.Account, error)
}

type UserRepositoryI interface {
	GetById(id uint64) (*model.User, error)
	Store(a *model.User) (*model.User, error)
}
