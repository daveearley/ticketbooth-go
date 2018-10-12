package app

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetUserFromContext extracts the authenicated user from the gin context
func GetUserFromContext(c *gin.Context) *models.User {
	user, exists := c.Get("auth_user")

	if !exists {
		log.Fatal("No authenticated user found in context.")
	}

	return user.(*models.User)
}

func CreatedResponse(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, &model)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func NotFoundResponse(c *gin.Context) {
	c.String(http.StatusNotFound, "")
}

func JsonResponse(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, json)
}

func StringResponse(c *gin.Context, string string) {
	c.String(http.StatusOK, string)
}
