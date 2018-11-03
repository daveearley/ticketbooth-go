package accounts

import (
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/users"
)

type Service interface {
	Find(id int) (*models.Account, error)
	Create(request *request.CreateAccount) (*models.Account, error)
	Delete(account *models.Account) error
}

type service struct {
	ur users.Repository
	ar Repository
}

func NewService(ar Repository, ur users.Repository) Service {
	return &service{ur, ar}
}

func (s *service) Find(id int) (*models.Account, error) {
	return s.ar.GetByID(id)
}

func (s *service) Create(request *request.CreateAccount) (*models.Account, error) {
	account, err := s.ar.Store(&models.Account{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	_, err = s.ur.Store(&models.User{
		AccountID: account.ID,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  request.Password,
	})

	if err != nil {
		return nil, err
	}

	return account, err
}

func (s *service) Delete(account *models.Account) error {
	return s.ar.DeleteByID(account.ID)
}
