package middleware

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/account"
	"github.com/daveearley/product/app/event"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// BindAndAuthorize handles checking authorization, 404s and binding IDs to models for use down the line
// todo move this logic into a service + tidy
func BindAndAuthorize(eventService event.Service, accountService account.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, v := range c.Params {
			switch v.Key {
			case "event_id":
				id := getID(&v)
				e, err := eventService.Find(id)

				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				if e.AccountID != app.GetUserFromContext(c).AccountID {
					c.AbortWithStatus(http.StatusUnauthorized)
				}

				c.Set("event", e)
				break
			case "account_id":
				id := getID(&v)
				a, err := accountService.Find(id)

				if err != nil {
					c.AbortWithStatus(http.StatusNotFound)
					return
				}

				if a.ID != app.GetUserFromContext(c).AccountID {
					c.AbortWithStatus(http.StatusForbidden)
					return
				}
				c.Set("account", a)
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
