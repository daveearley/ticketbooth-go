package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
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

type EventServiceI interface {
	Find(id int) (*models.Event, error)
	CreateEvent(event request.CreateEvent) (*models.Event, error)
}
