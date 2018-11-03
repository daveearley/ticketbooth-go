package events

import (
	"github.com/daveearley/ticketbooth/app/attributes"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
	"time"
)

type Response struct {
	*models.Event
	Attributes []*models.Attribute `json:"attributes"`
}

type ResponsePublic struct {
	Title       string      `json:"title"`
	StartDate   time.Time   `json:"start_date"`
	EndDate     time.Time   `json:"end_date"`
	Description null.String `json:"description"`

	Attributes []*attributes.Response `json:"attributes"`
	Tickets    []*models.Ticket       `json:"tickets"`
}

func TransformOne(c *gin.Context, event *models.Event) interface{} {
	if _, exists := c.Get("auth_user"); exists {
		return &Response{event, event.R.Attributes}
	}

	return &ResponsePublic{
		Title:       event.Title,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Description: event.Description,
		Attributes:  attributes.TransformMany(c, event.R.Attributes),
		Tickets:     event.R.Tickets,
	}
}

func TransformMany(c *gin.Context, events []*models.Event) interface{} {
	var transformedEvents []interface{}

	for _, v := range events {
		transformedEvents = append(transformedEvents, TransformOne(c, v))
	}

	return &transformedEvents
}
