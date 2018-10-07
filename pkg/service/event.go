package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/daveearley/product/pkg/repository"
	"github.com/volatiletech/null"
)

type EventService struct {
	er repository.EventRepositoryI
}

func NewEventService(repository repository.EventRepositoryI) *EventService {
	return &EventService{repository}
}

func (s *EventService) Find(id int) (*models.Event, error) {
	event, err := s.er.GetById(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *EventService) CreateEvent(req request.CreateEvent) (*models.Event, error) {
	event, err := s.er.Store(&models.Event{
		Title:       req.Title,
		Description: null.NewString(req.Description, true),
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		UserID:      46,
		AccountID:   58,
	})

	if err != nil {
		return nil, err
	}

	return event, nil
}
