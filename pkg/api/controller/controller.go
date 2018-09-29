package controller

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func CreatedResponse(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, &model)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err,
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
