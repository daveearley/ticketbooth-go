package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
	r "github.com/daveearley/product/pkg/repository"
)

type AccountService struct {
	ur r.UserRepositoryI
	ar r.AccountRepositoryI
}

func NewAccountService(ar r.AccountRepositoryI, ur r.UserRepositoryI) *AccountService {
	return &AccountService{ur, ar}
}

func (s *AccountService) Find(id int) (*models.Account, error) {
	return s.ar.GetById(id)
}

func (s *AccountService) CreateAccount(request *request.CreateAccount) (*models.Account, error) {
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

func (s *AccountService) DeleteAccount(account *models.Account) error {
	return nil
}
