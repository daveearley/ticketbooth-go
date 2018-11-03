package middleware

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

// AuthorizeActions handles checking authorization for each route
// Todo move this logic into a service + tidy
func AuthorizeActions() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			switch v.Key {
			case "ticket_id":
				event, _ := c.Get("event")
				ticket, _ := c.Get("ticket")
				if ticket.(*models.Ticket).EventID != event.(*models.Event).ID {
					response.Unauthorized(c)
					return
				}
				break
			case "event_id":
				event, _ := c.Get("event")

				if event.(*models.Event).AccountID != app.GetUserFromContext(c).AccountID {
					response.Unauthorized(c)
				}
				break
			case "account_id":
				account, _ := c.Get("account")

				if account.(*models.Account).ID != app.GetUserFromContext(c).AccountID {
					response.Unauthorized(c)
					return
				}
				break
			}
		}

		c.Next()
	}
}
