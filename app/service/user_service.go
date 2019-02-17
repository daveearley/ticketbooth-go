package service

import "github.com/daveearley/ticketbooth/app/models/generated"

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}
