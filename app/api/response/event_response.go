package response

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"time"
)

type EventResponse struct {
	*models.Event
	Attributes interface{} `json:"attributes"`
}

type PublicEventResponse struct {
	Title       string      `json:"title"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Description null.String `json:"description"`

	Attributes interface{} `json:"attributes"`
	Tickets    interface{} `json:"tickets"`
}

func TransformEvent(c *gin.Context, event *models.Event) interface{} {
	if app.IsUserAuthenticated(c) {
		return &EventResponse{event, TransformAttributes(c, event.R.Attributes)}
	}

	return &PublicEventResponse{
		Title:       event.Title,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Description: event.Description,
		Attributes:  TransformAttributes(c, event.R.Attributes),
		Tickets:     TransformTickets(c, event.R.Tickets),
	}
}

func TransformEvents(c *gin.Context, events []*models.Event) interface{} {
	var transformedEvents []interface{}

	for _, v := range events {
		transformedEvents = append(transformedEvents, TransformEvent(c, v))
	}

	return &transformedEvents
}
