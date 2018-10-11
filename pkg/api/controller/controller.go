package controller

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/daveearley/product/pkg/models/generated"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
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

func PaginatedResponse(c *gin.Context, model interface{}, db *gorm.DB) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))

	paginator := pagination.Pagging(&pagination.Param{
		DB:      db,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &model)

	c.JSON(http.StatusOK, paginator)
}
