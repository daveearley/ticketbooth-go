package events

import (
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/attributes"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/null"
)

type Service interface {
	Find(id int) (*models.Event, error)
	Create(event request.CreateEvent, user *models.User) (*models.Event, error)
	List(p *pagination.Params, authUser *models.User) ([]*models.Event, error)
}

type service struct {
	er Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Find(id int) (*models.Event, error) {
	event, err := s.er.GetById(id)

	if err != nil {
		return nil, err
	}

	return event, nil
}

func (s *service) Create(req request.CreateEvent, user *models.User) (*models.Event, error) {
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
		s.er.SetAttributes(event, attributes.MapToAttributes(&req.Attributes))
	}

	return event, nil
}

func (s *service) List(p *pagination.Params, authUser *models.User) ([]*models.Event, error) {
	events, err := s.er.List(p, authUser)

	if err != nil {
		return nil, err
	}

	return events, nil
}
