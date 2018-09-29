package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/model"
)

type UserServiceI interface {
	CreateUser(user *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
}

type AccountServiceI interface {
	Find(id uint64) (*model.Account, error)
	CreateAccount(request *request.CreateAccount) (*model.Account, error)
	DeleteAccount(account *model.Account) error
}
