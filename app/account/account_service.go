package account

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/user"
)

type Service interface {
	Find(id int) (*models.Account, error)
	CreateAccount(request *request.CreateAccount) (*models.Account, error)
	DeleteAccount(account *models.Account) error
}

type accountService struct {
	ur user.Repository
	ar Repository
}

func NewService(ar Repository, ur user.Repository) *accountService {
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
