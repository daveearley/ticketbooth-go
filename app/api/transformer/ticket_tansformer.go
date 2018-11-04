package transformer

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

type TicketResponse struct {
	*models.Ticket
	Questions []*QuestionResponse `json:"questions"`
}

type PublicTicketResponse struct {
	ID                int
	Title             string
	EventID           int
	QuantityAvailable int
	SaleStartDate     null.Time
	SaleEndDate       null.Time

	Questions []*QuestionResponse `json:"questions"`
}

func TransformTicket(c *gin.Context, t *models.Ticket) interface{} {
	if app.IsUserAuthenticated(c) {
		return &TicketResponse{t, TransformQuestions(t.R.Questions)}
	}

	return &PublicTicketResponse{
		ID:                t.ID,
		Title:             t.Title,
		EventID:           t.EventID,
		QuantityAvailable: t.InititalQuantityAvailable,
		Questions:         TransformQuestions(t.R.Questions),
	}
}

func TransformTickets(c *gin.Context, tickets []*models.Ticket) interface{} {
	var transformed []interface{}
	for _, v := range tickets {
		transformed = append(transformed, TransformTicket(c, v))
	}

	return &transformed
}
