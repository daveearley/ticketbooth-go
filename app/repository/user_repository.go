package repository

import (
	"database/sql"
	"../../app"
	"../../app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type UserRepository interface {
	GetByID(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Store(a *models.User) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

//NewUserRepository returns a new instance of userRepository
func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{conn}
}

//GetByID gets a user by ID
func (r *userRepository) GetByID(id int) (*models.User, error) {
	user, err := models.FindUser(r.db, id)

	if err != nil {
		return nil, getErrorType(err, app.UserResource, id)
	}

	return user, nil
}

//FindByEmail gets an user by email address
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email=?", email)).One(r.db)

	if err != nil {
		return nil, getErrorType(err, app.UserResource, email)
	}

	return user, nil
}

//Store creates a user
func (r *userRepository) Store(a *models.User) (*models.User, error) {
	err := a.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, app.ServerError(err)
	}

	return a, nil
}
