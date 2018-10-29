package middleware

import (
	"fmt"
	"github.com/daveearley/product/app/accounts"
	"github.com/daveearley/product/app/api/response"
	"github.com/daveearley/product/app/events"
	"github.com/daveearley/product/app/tickets"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// PreloadModels binds parameter IDs to their models and sets them in context
// Todo move this logic into a service + tidy
func PreloadModels(
	eventRepo events.Repository,
	accountRepo accounts.Repository,
	ticketRepo tickets.Repository,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			id := getID(&v)
			fmt.Println(v.Key)
			switch v.Key {
			case "ticket_id":
				ticket, err := ticketRepo.GetByID(id)
				if err != nil {
					response.NotFoundResponse(c)
					return
				}

				event, err := eventRepo.GetByID(ticket.EventID)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}

				c.Set("ticket", ticket)
				c.Set("event", event)
				break
			case "event_id":

				if _, exists := c.Get("event"); exists {
					continue
				}

				event, err := eventRepo.GetByID(id)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}
				c.Set("event", event)
				break
			case "account_id":
				account, err := accountRepo.GetByID(id)

				if err != nil {
					response.NotFoundResponse(c)
					return
				}
				c.Set("account", account)
				break
			default:
				response.Error(c, http.StatusBadRequest, errors.New("Unknown parameter in URL"))
			}
		}

		c.Next()
	}
}
