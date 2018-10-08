package service

import (
	"fmt"
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

	// todo Move attributes mapping to utils. Also handle conversion of map to json string
	if req.Attributes != nil {
		var attrs []*models.Attribute
		for k, v := range req.Attributes {

			//if reflect.TypeOf(v).Kind().String() == "map" {
			//	v, _ := json.Marshal(v)
			//}

			attrs = append(attrs, &models.Attribute{
				Name:  k,
				Value: fmt.Sprint(v),
			})
		}

		s.er.SetAttributes(event, attrs)
	}

	return event, nil
}
