package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app"
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
	GetReservedTicketQuantity(ticket *models.Ticket) (int, error)
	CreateReservedTickets(resTickets []*models.TicketReservation) (err error)
	GetByEventID(ticketID int) ([]*models.Ticket, error)
}

type ticketRepository struct {
	db *sql.DB
}

//NewTicketRepository returns a new instance of ticketRepository
func NewTicketRepository(db *sql.DB) TicketRepository {
	return &ticketRepository{db}
}

//GetByID gets a ticket by a given ID
func (r *ticketRepository) GetByID(id int) (*models.Ticket, error) {
	return models.Tickets(qm.Where("id=?", id), qm.Load("Questions.QuestionOptions")).One(r.db)
}

//DeleteByID deletes a ticket with a given ID
func (r *ticketRepository) DeleteByID(id int) error {
	ticket, err := models.FindTicket(r.db, id)

	if err != nil {
		return err
	}

	_, err = ticket.Delete(r.db)

	return err
}

//GetReservedTicketQuantity returns the quantity currently reserved (in pending transactions)
func (r *ticketRepository) GetReservedTicketQuantity(ticket *models.Ticket) (int, error) {
	type ResultCount struct {
		Count int `json:"count"`
	}

	var result ResultCount
	err := models.NewQuery(
		qm.Select("COALESCE(sum(ticket_quantity), 0) AS count"),
		qm.From("ticket_reservations"),
		qm.Where("ticket_id=?", ticket.ID),
		qm.Where("current_timestamp  < reserved_until"),
	).Bind(nil, r.db, &result)

	if err != nil {
		return 0, app.ServerError(err)
	}

	return result.Count, nil
}

//Store creates a ticket
func (r *ticketRepository) Store(ticket *models.Ticket) (*models.Ticket, error) {
	err := ticket.Insert(r.db, boil.Infer())

	if err != nil {
		return nil, app.ServerError(err)
	}

	return ticket, err
}

//SetAttributes adds attributes to a ticket
func (r *ticketRepository) SetAttributes(ticket *models.Ticket, attr []*models.Attribute) error {
	err := ticket.SetAttributes(r.db, true, attr...)

	if err != nil {
		return app.ServerError(err)
	}

	return nil
}

//SetQuestion adds a question to a ticket
func (r *ticketRepository) SetQuestion(ticket *models.Ticket, question *models.Question) error {
	err := ticket.AddQuestions(r.db, true, question)

	if err != nil {
		return app.ServerError(err)
	}

	return nil
}

//List returns a slice of tickets
func (r *ticketRepository) List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error) {
	queryMods := pagination.QueryMods(p)
	queryMods = append(queryMods, qm.Load("Attributes"))
	queryMods = append(queryMods, qm.Where("event_id=?", event.ID))

	tickets, err := models.Tickets(queryMods...).All(r.db)

	if err != nil {
		return nil, app.ServerError(err)
	}

	return tickets, nil
}

//GetByEventID gets tickets related to an event
func (r *ticketRepository) GetByEventID(eventId int) ([]*models.Ticket, error) {
	tickets, err := models.Tickets(qm.Where("event_id=?", eventId)).All(r.db)

	if err != nil {
		return nil, app.ServerError(err)
	}

	return tickets, nil
}

//ListQuestions gets a ticket's questions
func (r *ticketRepository) ListQuestions(ticket *models.Ticket) ([]*models.Question, error) {
	questions, err := ticket.Questions().All(r.db)

	if err != nil {
		return nil, app.ServerError(err)
	}

	return questions, nil
}

//CreateReservedTickets creates a reserved ticket
func (r *ticketRepository) CreateReservedTickets(resTickets []*models.TicketReservation) (err error) {
	for _, res := range resTickets {
		err = res.Insert(r.db, boil.Infer())
	}

	if err != nil {
		return app.ServerError(err)
	}

	return err
}
