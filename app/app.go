package app

import (
	"crypto/sha1"
	"fmt"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

const (
	EventResource       = "event"
	AttendeeResource    = "attendee"
	QuestionResource    = "question"
	TransactionResource = "transaction"
	AccountResource     = "account"
	UserResource        = "user"
	TicketResource      = "ticket"
)

// GetUserFromContext extracts the authenticated user from the gin context
func GetUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get(UserResource)

	if !exists {
		return nil
	}

	return user.(*models.User)
}

//IsUserAuthenticated checks if a user session exists
func IsUserAuthenticated(c *gin.Context) bool {
	_, exists := c.Get(UserResource)
	return exists
}

// GetUniqueUserID returns an MD5 hash of a user's user agent and IP address.
func GetUniqueUserID(c *gin.Context) string {
	h := sha1.New()
	h.Write([]byte(c.ClientIP() + c.Request.UserAgent()))
	s := h.Sum(nil)

	return fmt.Sprintf("%x\n", s)
}
