package transformer

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

type TicketResponse struct {
	*models.Ticket
	Questions *Envelope `json:"questions"`
}

type PublicTicketResponse struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	EventID           int       `json:"event_id"`
	QuantityAvailable int       `json:"quantity_available"`
	SaleStartDate     null.Time `json:"sale_start_date"`
	SaleEndDate       null.Time `json:"sale_end_date"`
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
	}
}

func TransformTickets(c *gin.Context, tickets []*models.Ticket) *Envelope {
	var transformed []interface{}
	for _, v := range tickets {
		transformed = append(transformed, TransformTicket(c, v))
	}

	return envelope(transformed)
}
