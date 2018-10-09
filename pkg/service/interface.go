package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}

type AccountService interface {
	Find(id int) (*models.Account, error)
	CreateAccount(request *request.CreateAccount) (*models.Account, error)
	DeleteAccount(account *models.Account) error
}

type EventService interface {
	Find(id int) (*models.Event, error)
	CreateEvent(event request.CreateEvent) (*models.Event, error)
}
