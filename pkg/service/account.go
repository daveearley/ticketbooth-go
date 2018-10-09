package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
	r "github.com/daveearley/product/pkg/repository"
)

type accountService struct {
	ur r.UserRepository
	ar r.AccountRepository
}

func NewAccountService(ar r.AccountRepository, ur r.UserRepository) *accountService {
	return &accountService{ur, ar}
}

func (s *accountService) Find(id int) (*models.Account, error) {
	return s.ar.GetById(id)
}

func (s *accountService) CreateAccount(request *request.CreateAccount) (*models.Account, error) {
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

func (s *accountService) DeleteAccount(account *models.Account) error {
	return nil
}
