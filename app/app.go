package app

import (
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"log"
)

// GetUserFromContext extracts the authenicated user from the gin context
func GetUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get("auth_user")

	if !exists {
		log.Fatal("No authenticated user found in context.")
	}

	return user.(*models.User)
}
