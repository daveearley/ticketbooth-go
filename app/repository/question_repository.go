package repository

import (
	"database/sql"
	"../../app"
	"../../app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type QuestionRepository interface {
	Store(question *models.Question) (*models.Question, error)
	SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error
}

type questionRepository struct {
	db *sql.DB
}

//NewQuestionRepository returns a new instance of questionRepository
func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db}
}

//Store creates a question
func (r *questionRepository) Store(question *models.Question) (*models.Question, error) {
	err := question.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, app.ServerError(err)
	}

	return question, nil
}

//SetQuestionOptions add options to a question
func (r *questionRepository) SetQuestionOptions(question *models.Question, opts []*models.QuestionOption) error {
	err := question.AddQuestionOptions(r.db, true, opts...)

	if err != nil {
		return app.ServerError(err)
	}

	return nil
}
