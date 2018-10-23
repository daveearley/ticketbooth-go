package tickets

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/attributes"
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/null"
)

type Service interface {
	Find(id int) (*models.Ticket, error)
	Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error)
}

type service struct {
	er Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Find(id int) (*models.Ticket, error) {
	ticket, err := s.er.GetById(id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *service) Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error) {
	ticket := &models.Ticket{
		Title:                    req.Title,
		SaleStartDate:            null.NewTime(req.SaleStartDate, true),
		SaleEndDate:              null.NewTime(req.SaleEndDate, true),
		IntitalQuantityAvailable: req.Quantity,
		EventID:                  event.ID,
	}

	app.BeforeSaveTicket(ticket)

	ticket, err := s.er.Store(ticket)

	if err != nil {
		return nil, err
	}

	if req.Attributes != nil {
		s.er.SetAttributes(ticket, attributes.MapToAttributes(&req.Attributes))
	}

	return ticket, nil
}

func (s *service) List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error) {
	tickets, err := s.er.List(p, event)

	if err != nil {
		return nil, err
	}

	return tickets, nil
}
