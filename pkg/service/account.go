package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/model"
	r "github.com/daveearley/product/pkg/repository"
)

type AccountService struct {
	ur r.UserRepositoryI
	ar r.AccountRepositoryI
}

func NewAccountService(ar r.AccountRepositoryI, ur r.UserRepositoryI) *AccountService {
	return &AccountService{ur, ar}
}

func (s *AccountService) Find(id uint64) (*model.Account, error) {
	return s.ar.GetById(id)
}

func (s *AccountService) CreateAccount(request *request.CreateAccount) (*model.Account, error) {
	account, err := s.ar.Store(&model.Account{
		Email: request.Email,
	})

	if err != nil {
		return nil, err
	}

	_, err = s.ur.Store(&model.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		AccountId: account.ID,
		Password:  request.Password,
	})

	return account, err
}

func (s *AccountService) DeleteAccount(account *model.Account) error {
	return nil
}
