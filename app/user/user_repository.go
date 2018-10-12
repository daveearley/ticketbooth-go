package user

import (
	"database/sql"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type Repository interface {
	GetById(id int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Store(a *models.User) (*models.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(conn *sql.DB) *repository {
	return &repository{conn}
}

func (r *repository) GetById(id int) (*models.User, error) {
	user, err := models.FindUser(r.db, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (*models.User, error) {
	user, err := models.Users(qm.Where("email=?", email)).One(r.db)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) Store(a *models.User) (*models.User, error) {
	if err := a.Insert(r.db, boil.Infer()); err != nil {
		return a, err
	}

	return a, nil
}
