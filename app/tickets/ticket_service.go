package tickets

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/attributes"
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/questions"
	"github.com/volatiletech/null"
)

type Service interface {
	Find(id int) (*models.Ticket, error)
	Delete(id int) error
	Create(req request.CreateTicket, event *models.Event) (*models.Ticket, error)
	CreateQuestion(req request.CreateQuestion, ticket *models.Ticket) (*models.Question, error)
	List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error)
}

type service struct {
	er Repository
	qr questions.Repository
}

func NewService(repository Repository, qRepository questions.Repository) Service {
	return &service{repository, qRepository}
}

func (s *service) Find(id int) (*models.Ticket, error) {
	ticket, err := s.er.GetByID(id)

	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *service) Delete(id int) error {
	err := s.er.DeleteByID(id)

	return err
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

func (s *service) CreateQuestion(req request.CreateQuestion, ticket *models.Ticket) (*models.Question, error) {
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

func (s *service) List(p *pagination.Params, event *models.Event) ([]*models.Ticket, error) {
	tickets, err := s.er.List(p, event)

	if err != nil {
		return nil, err
	}

	return tickets, nil
}
