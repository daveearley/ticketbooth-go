package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserRepository interface {
	GetById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Store(a *models.User) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{conn}
}

func (r *userRepository) GetById(id int) (*models.User, error) {
	user, err := models.FindUser(r.db, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email=?", email)).One(r.db)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Store(a *models.User) (*models.User, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return a, err
	}

	return a, nil
}