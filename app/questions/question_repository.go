package questions

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type Repository interface {
	GetById(id int) (*models.Question, error)
	Store(question *models.Question) (*models.Question, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Question, error)
	SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) GetById(id int) (*models.Question, error) {
	panic("implement me")
}

func (r *repository) Store(question *models.Question) (*models.Question, error) {
	err := question.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (r *repository) List(p *pagination.Params, event *models.Event) ([]*models.Question, error) {
	panic("implement me")
}

func (r *repository) SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error {
	return question.AddQuestionOptions(r.db, true, opts...)
}
