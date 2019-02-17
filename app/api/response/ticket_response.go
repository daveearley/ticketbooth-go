package response

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

type TicketResponse struct {
	*models.Ticket
	Questions interface{} `json:"questions"`
}

type PublicTicketResponse struct {
	ID                int         `json:"id"`
	Title             string      `json:"title"`
	QuantityAvailable int         `json:"quantity_available"`
	SaleStartDate     null.Time   `json:"sale_start_date"`
	SaleEndDate       null.Time   `json:"sale_end_date"`
	MaxPerTransaction null.Int    `json:"max_per_transaction"`
	Questions         interface{} `json:"questions"`
}

//TransformTicket transforms a ticket model for frontend use
func TransformTicket(c *gin.Context, t *models.Ticket) interface{} {
	if app.IsUserAuthenticated(c) {
		return &TicketResponse{t, getTicketQuestions(t)}
	}

	return &PublicTicketResponse{
		ID:                t.ID,
		Title:             t.Title,
		QuantityAvailable: t.InititalQuantityAvailable,
		Questions:         getTicketQuestions(t),
	}
}

//TransformTickets transforms a slice of tickets
func TransformTickets(c *gin.Context, tickets []*models.Ticket) interface{} {
	var transformed []interface{}
	for _, v := range tickets {
		transformed = append(transformed, TransformTicket(c, v))
	}

	return &transformed
}

//getTicketQuestions returns a tickets questions or nil
func getTicketQuestions(ticket *models.Ticket) models.QuestionSlice {
	if ticket.R != nil {
		return ticket.R.Questions
	}

	return nil
}
