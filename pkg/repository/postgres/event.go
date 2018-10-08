package repository

import (
	"database/sql"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/volatiletech/sqlboiler/boil"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db}
}

func (r *EventRepository) GetById(id int) (*models.Event, error) {
	event, err := models.FindEvent(r.db, id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) Store(event *models.Event) (*models.Event, error) {
	if err := event.Insert(r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return event, nil
}

func (r *EventRepository) SetAttributes(event *models.Event, attr []*models.Attribute) error {
	return event.SetAttributes(r.db, true, attr...)
}
