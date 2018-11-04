package app

import (
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

// GetUserFromContext extracts the authenicated user from the gin context
func GetUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get("auth_user")

	if !exists {
		return nil
	}

	return user.(*models.User)
}

//IsUserAuthenticated checks if a user session exists
func IsUserAuthenticated(c *gin.Context) bool {
	_, exists := c.Get("auth_user")
	return exists
}
