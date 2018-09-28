package service

import (
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

func (s *AccountService) CreateAccount(account *model.Account) (*model.Account, error) {
	// TODO create user

	return s.ar.Store(account)
}

func (s *AccountService) DeleteAccount(account *model.Account) error {
	return nil
}
