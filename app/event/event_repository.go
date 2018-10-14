package event

import (
	"database/sql"
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/pagination"
	"github.com/volatiletech/sqlboiler/boil"
)

type Repository interface {
	GetById(id int) (*models.Event, error)
	Store(event *models.Event) (*models.Event, error)
	SetAttributes(event *models.Event, attr []*models.Attribute) error
	List(p *pagination.Params, authUser *models.User) ([]*models.Event, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetById(id int) (*models.Event, error) {
	event, err := models.FindEvent(r.db, id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) Store(event *models.Event) (*models.Event, error) {
	if err := event.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *repository) SetAttributes(event *models.Event, attr []*models.Attribute) error {
	return event.SetAttributes(r.db, true, attr...)
}

func (r *repository) List(p *pagination.Params, authUser *models.User) ([]*models.Event, error) {
	boil.DebugMode = true

	queryMods := pagination.QueryMods(p, authUser)

	events, err := models.Events(queryMods...).All(r.db)

	if err != nil {
		return nil, err
	}

	if pagination.ShouldSetCursor(p, len(events)) {
		p.NextCursorId = events[len(events)-1].ID
	}

	return events, nil
}
