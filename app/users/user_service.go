package users

import "github.com/daveearley/ticketbooth/app/models/generated"

type Service interface {
	CreateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
}
