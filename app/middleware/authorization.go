package middleware

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/accounts"
	"github.com/daveearley/product/app/events"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BindAndAuthorize handles checking authorization, 404s and binding IDs to models for use down the line
// Todo move this logic into a service + tidy
func BindAndAuthorize(eventService events.Service, accountService accounts.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			switch v.Key {
			case "event_id":
				id := getID(&v)
				event, err := eventService.Find(id)

				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				if event.AccountID != app.GetUserFromContext(c).AccountID {
					c.AbortWithStatus(http.StatusUnauthorized)
				}

				c.Set("event", event)
				break
			case "account_id":
				id := getID(&v)
				account, err := accountService.Find(id)

				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				if account.ID != app.GetUserFromContext(c).AccountID {
					c.AbortWithStatus(http.StatusForbidden)
					return
				}
				c.Set("account", account)
				break
			}
		}

		c.Next()
	}
}

func getID(p *gin.Param) int {
	id, _ := strconv.Atoi(p.Value)

	return id
}
