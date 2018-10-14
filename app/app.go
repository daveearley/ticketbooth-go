package app

import (
	"fmt"
	"github.com/daveearley/product/app/models/generated"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
)

// GetUserFromContext extracts the authenicated user from the gin context
func GetUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get("auth_user")

	if !exists {
		log.Fatal("No authenticated user found in context.")
	}

	return user.(*models.User)
}

// IsUserAuthorized checks if a user is authorized to access an entity
func IsUserAuthorized(c *gin.Context, entity interface{}) bool {
	entityType := reflect.TypeOf(entity)
	authUser, _ := c.Get("auth_user")

	fmt.Println("name", entityType.String())

	switch entityType.String() {
	case "*models.Event":
		return entity.(*models.Event).AccountID == authUser.(*models.User).AccountID
	case "*models.Account":
		return entity.(*models.Account).ID == authUser.(*models.User).AccountID
	}

	return false
}
