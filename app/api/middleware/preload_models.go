package middleware

import (
	"fmt"
	"../../../app"
	"../../../app/api/response"
	"../../../app/repository"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PreloadModels binds parameter IDs to their models and sets them in context
// Todo move this logic into a service + tidy
func PreloadModels(
	eventRepo repository.EventRepository,
	accountRepo repository.AccountRepository,
	ticketRepo repository.TicketRepository,
	tranRepo repository.TransactionRepository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			id := getID(&v)
			fmt.Println(v.Key)
			switch v.Key {
			case "transaction_uuid":
				// todo - limit this to only recent/unfinalized transactions
				transaction, err := tranRepo.GetByUUID(v.Value)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}

				c.Set(app.TransactionResource, transaction)
			case "ticket_id":
				ticket, err := ticketRepo.GetByID(id)

				if err != nil {
					e := errors.WithStack(err)
					fmt.Println(e)
					response.NotFoundResponse(c)
					return
				}

				event, err := eventRepo.GetByID(ticket.EventID)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}

				c.Set(app.TicketResource, ticket)
				c.Set(app.EventResource, event)
				break
			case "event_id":

				if _, exists := c.Get(app.EventResource); exists {
					continue
				}

				event, err := eventRepo.GetByID(id)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}
				c.Set(app.EventResource, event)
				break
			case "account_id":
				account, err := accountRepo.GetByID(id)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}
				c.Set(app.AccountResource, account)
				break
			default:
				response.Error(c, app.InvalidValueError(v.Key, "Unknown parameter"))
			}
		}

		c.Next()
	}
}
