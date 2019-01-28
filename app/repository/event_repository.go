package repository

import (
	"database/sql"
	"../../app"
	"../../app/api/pagination"
	"../../app/models/generated"
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

//NewEventRepository returns a new instance of eventRepository
func NewEventRepository(db *sql.DB) *eventRepository {
	return &eventRepository{db}
}

//GetByID gets an event with a given ID
func (r *eventRepository) GetByID(id int) (*models.Event, error) {
	event, err := models.Events(
		qm.Load("Attributes"),
		qm.Where("id=?", id),
	).One(r.db)

	if err == nil {
		return event, nil
	}

	return nil, getErrorType(err, app.EventResource, id)
}

//DeleteByID deletes an event with a given ID
func (r *eventRepository) DeleteByID(id int) error {
	event, err := r.GetByID(id)

	if err != nil {
		return err
	}

	_, err = event.Delete(r.db)

	if err == nil {
		return nil
	}

	return app.ServerError(err, id)
}

//GetByTicketID gets an event using a ticket ID
func (r *eventRepository) GetByTicketID(id int) (*models.Event, error) {
	return models.Events(qm.Where("ticket_id=?", id)).One(r.db)
}

//Store saves an event
func (r *eventRepository) Store(event *models.Event) (*models.Event, error) {
	err := event.Insert(r.db, boil.Infer())

	if err == nil {
		return event, nil
	}

	return nil, app.ServerError(err)
}

//SetAttributes sets an events attributes
func (r *eventRepository) SetAttributes(event *models.Event, attr []*models.Attribute) error {
	err := event.SetAttributes(r.db, true, attr...)

	if err != nil {
		return app.ServerError(err)
	}

	return nil
}

// List returns a paginated slice of events
func (r *eventRepository) List(p *pagination.Params, authUser *models.User) ([]*models.Event, error) {
	queryMods := pagination.QueryMods(p)
	queryMods = append(queryMods, qm.Load("Attributes"))
	queryMods = append(queryMods, qm.Where("account_id=?", authUser.AccountID))

	events, err := models.Events(queryMods...).All(r.db)

	if err == nil {
		return events, nil
	}

	return nil, app.ServerError(err)
}
