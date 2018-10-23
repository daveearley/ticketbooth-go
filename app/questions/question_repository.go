package questions

import (
	"database/sql"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/models/generated"
)

type Repository interface {
	GetById(id int) (*models.Question, error)
	StoreTicketQuestion(question *models.Question, ticket *models.Ticket) (*models.Question, error)
	StoreEventQuestion(event *models.Question) (*models.Question, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Question, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *Repository) GetByID(id int) (*models.Question, error) {
	question := models.FindQ

}
