package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
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
	GetRemainingTicketQuantity(ticket *models.Ticket) (int, error)
}

type ticketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) GetByID(id int) (*models.Ticket, error) {
	return models.FindTicket(r.db, id)
}

func (r *ticketRepository) DeleteByID(id int) error {
	ticket, err := models.FindTicket(r.db, id)

	if err != nil {
		return err
	}

	_, err = ticket.Delete(r.db)

	return err
}

func (r *ticketRepository) GetRemainingTicketQuantity(ticket *models.Ticket) (int, error) {
	type ResultCount struct {
		Count int `json:"count"`
	}

	var result ResultCount

	err := queries.Raw(`
  		SELECT COALESCE(sum(ticket_quantity), 0) AS count
		FROM ticket_reservations
		WHERE ticket_id = ?
  		AND current_timestamp  < reserved_until' `, ticket.ID,
	).Bind(nil, r.db, &result)

	if err != nil {
		return 0, err
	}

	return ticket.InititalQuantityAvailable - (ticket.QuantitySold + result.Count), nil
}

func (r *ticketRepository) Store(ticket *models.Ticket) (*models.Ticket, error) {
	err := ticket.Insert(r.db, boil.Infer())

	return ticket, err
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

	return models.Tickets(queryMods...).All(r.db)
}

func (r *ticketRepository) ListQuestions(ticket *models.Ticket) ([]*models.Question, error) {
	return ticket.Questions().All(r.db)
}
