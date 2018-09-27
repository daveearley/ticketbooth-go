package service

import (
	"fmt"
	r "github.com/daveearley/product/pkg/repository"
)

type CreateAccountService struct {
	ar *r.UserRepositoryI
	ur *r.AccountRepositoryI
}

func (s *CreateAccountService) NewCreateAccountService(ar *r.AccountRepositoryI, ur *r.UserRepositoryI) *CreateAccountService {
	return &CreateAccountService{ur, ar}
}

func (s *CreateAccountService) CreateNewAccount() {
	fmt.Print("Hey")
}
