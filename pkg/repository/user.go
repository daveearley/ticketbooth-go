package repository

import "github.com/daveearley/product/pkg/model"

type UserRepositoryI interface {
	GetById(id string) (*model.User, error)
	Store(a *model.User) error
}
