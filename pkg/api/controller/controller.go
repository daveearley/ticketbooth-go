package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatedResponse(c *gin.Context) {
	c.String(http.StatusCreated, "")
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
