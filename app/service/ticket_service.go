package service

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/repository"
	"github.com/volatiletech/null"
)

type TicketService interface {
	Find(id int) (*models.Ticket, error)
	Delete(id int) error
	Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error)
	CreateQuestion(req request.CreateQuestion, ticket *models.Ticket) (*models.Question, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error)
}

type ticketService struct {
	er repository.TicketRepository
	qr repository.QuestionRepository
}

func NewTicketService(repository repository.TicketRepository, qRepository repository.QuestionRepository) TicketService {
	return &ticketService{repository, qRepository}
}

func (s *ticketService) Find(id int) (*models.Ticket, error) {
	return s.er.GetByID(id)
}

func (s *ticketService) Delete(id int) error {
	err := s.er.DeleteByID(id)

	return err
}

func (s *ticketService) Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error) {
	ticket := &models.Ticket{
		Title:                     req.Title,
		SaleStartDate:             null.NewTime(req.SaleStartDate, true),
		SaleEndDate:               null.NewTime(req.SaleEndDate, true),
		InititalQuantityAvailable: req.Quantity,
		EventID:                   event.ID,
	}

	app.BeforeSaveTicket(ticket)

	ticket, err := s.er.Store(ticket)

	if err != nil {
		return nil, err
	}

	if req.Attributes != nil {
		s.er.SetAttributes(ticket, MapToAttributes(&req.Attributes))
	}

	return ticket, nil
}

func (s *ticketService) CreateQuestion(req request.CreateQuestion, ticket *models.Ticket) (*models.Question, error) {
	question := &models.Question{
		Title:    req.Title,
		Type:     req.Type,
		Required: req.Required,
	}

	err := s.er.SetQuestion(ticket, question)

	if err != nil {
		return nil, err
	}

	var opts []*models.QuestionOption
	for _, v := range req.Options {
		opts = append(opts, &models.QuestionOption{
			Title: v.Title,
		})
	}

	if err = s.qr.SetQuestionOptions(question, opts); err != nil {
		return nil, err
	}

	return question, nil
}

func (s *ticketService) List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error) {
	return s.er.List(p, event)
}
