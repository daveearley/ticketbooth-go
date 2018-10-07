package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models"
)

type UserServiceI interface {
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}

type AccountServiceI interface {
	Find(id int) (*models.Account, error)
	CreateAccount(request *request.CreateAccount) (*models.Account, error)
	DeleteAccount(account *models.Account) error
}
