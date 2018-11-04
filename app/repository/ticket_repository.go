package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type TicketRepository interface {
	GetByID(id int) (*models.Ticket, error)
	DeleteByID(id int) error
	Store(event *models.Ticket) (*models.Ticket, error)
	SetAttributes(event *models.Ticket, attr []*models.Attribute) error
	SetQuestion(ticket *models.Ticket, question *models.Question) error
	List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error)
	ListQuestions(ticket *models.Ticket) ([]*models.Question, error)
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) GetByID(id int) (*models.Ticket, error) {
	event, err := models.FindTicket(r.db, id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *ticketRepository) DeleteByID(id int) error {
	ticket, err := models.FindTicket(r.db, id)

	if err != nil {
		return err
	}

	_, err = ticket.Delete(r.db)

	return err
}

func (r *ticketRepository) Store(ticket *models.Ticket) (*models.Ticket, error) {
	if err := ticket.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return ticket, nil
}

func (r *ticketRepository) SetAttributes(ticket *models.Ticket, attr []*models.Attribute) error {
	return ticket.SetAttributes(r.db, true, attr...)
}

func (r *ticketRepository) SetQuestion(ticket *models.Ticket, question *models.Question) error {
	return ticket.AddQuestions(r.db, true, question)
}

func (r *ticketRepository) List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error) {
	queryMods := pagination.QueryMods(p)
	queryMods = append(queryMods, qm.Load("Attributes"))
	queryMods = append(queryMods, qm.Where("event_id=?", event.ID))

	events, err := models.Tickets(queryMods...).All(r.db)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *ticketRepository) ListQuestions(ticket *models.Ticket) ([]*models.Question, error) {
	questions, err := ticket.Questions().All(r.db)

	if err != nil {
		return nil, err
	}

	return questions, nil
}