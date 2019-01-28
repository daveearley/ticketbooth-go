package middleware

import (
	"../../../app"
	"../../../app/api/response"
	"../../../app/models/generated"
	"github.com/gin-gonic/gin"
)

// AuthorizeActions checks what resources are being requested and authorizes based om
// the authenticated user's permissions
// Todo this could get big over time - move this logic into a service + tidy
func AuthorizeActions() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			switch v.Key {
			case "ticket_id":
				event, _ := c.Get(app.EventResource)
				ticket, _ := c.Get(app.TicketResource)

				if ticket.(*models.Ticket).EventID != event.(*models.Event).ID {
					response.Unauthorized(c)
					return
				}
				break
			case "event_id":
				event, _ := c.Get(app.EventResource)

				if event.(*models.Event).AccountID != app.GetUserFromContext(c).AccountID {
					response.Unauthorized(c)
					return
				}
				break
			case "account_id":
				account, _ := c.Get(app.AccountResource)

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
