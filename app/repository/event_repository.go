package repository

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type EventRepository interface {
	GetByID(id int) (*models.Event, error)
	GetByTicketID(id int) (*models.Event, error)
	Store(event *models.Event) (*models.Event, error)
	SetAttributes(event *models.Event, attr []*models.Attribute) error
	List(p *pagination.Params, authUser *models.User) ([]*models.Event, error)
	DeleteByID(id int) error
}

type eventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *eventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) GetByID(id int) (*models.Event, error) {
	event, err := models.Events(
		qm.Load("Tickets"),
		qm.Load("Attributes"),
		qm.Where("id=?", id),
	).One(r.db)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) DeleteByID(id int) error {
	event, err := r.GetByID(id)

	if err != nil {
		return err
	}

	_, err = event.Delete(r.db)

	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) GetByTicketID(id int) (*models.Event, error) {
	event, err := models.Events(qm.Where("ticket_id=?", id)).One(r.db)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) Store(event *models.Event) (*models.Event, error) {
	if err := event.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) SetAttributes(event *models.Event, attr []*models.Attribute) error {
	return event.SetAttributes(r.db, true, attr...)
}

func (r *eventRepository) List(p *pagination.Params, authUser *models.User) ([]*models.Event, error) {
	queryMods := pagination.QueryMods(p)
	queryMods = append(queryMods, qm.Load("Attributes"))
	queryMods = append(queryMods, qm.Where("account_id=?", authUser.AccountID))

	events, err := models.Events(queryMods...).All(r.db)

	if err != nil {
		return nil, err
	}

	return events, nil
}
