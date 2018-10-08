package repository

import "github.com/daveearley/product/pkg/models/generated"

type AccountRepositoryI interface {
	GetById(id int) (*models.Account, error)
	Store(a *models.Account) (*models.Account, error)
}

type UserRepositoryI interface {
	GetById(id int) (*models.User, error)
	Store(a *models.User) (*models.User, error)
}

type EventRepositoryI interface {
	GetById(id int) (*models.Event, error)
	Store(event *models.Event) (*models.Event, error)
	SetAttributes(event *models.Event, attr []*models.Attribute) error
}
