package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type QuestionRepository interface {
	GetById(id int) (*models.Question, error)
	Store(question *models.Question) (*models.Question, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Question, error)
	SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error
}

type questionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db}
}

func (r *questionRepository) GetById(id int) (*models.Question, error) {
	panic("implement me")
}

func (r *questionRepository) Store(question *models.Question) (*models.Question, error) {
	err := question.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (r *questionRepository) List(p *pagination.Params, event *models.Event) ([]*models.Question, error) {
	panic("implement me")
}

func (r *questionRepository) SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error {
	return question.AddQuestionOptions(r.db, true, opts...)
}
