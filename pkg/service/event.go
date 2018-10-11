package service

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/daveearley/product/pkg/repository"
	"github.com/volatiletech/null"
)

type eventService struct {
	er repository.EventRepository
}

func NewEventService(repository repository.EventRepository) *eventService {
	return &eventService{repository}
}

func (s *eventService) Find(id int) (*models.Event, error) {
	event, err := s.er.GetById(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *eventService) CreateEvent(req request.CreateEvent, user *models.User) (*models.Event, error) {
	event, err := s.er.Store(&models.Event{
		Title:       req.Title,
		Description: null.NewString(req.Description, true),
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		UserID:      user.ID,
		AccountID:   user.AccountID,
	})

	if err != nil {
		return nil, err
	}

	if req.Attributes != nil {
		s.er.SetAttributes(event, MapToAttributes(&req.Attributes))
	}

	return event, nil
}
