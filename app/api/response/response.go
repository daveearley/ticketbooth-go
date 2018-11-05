package response

import (
	"errors"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

const unauthorizedMessage string = "This action is unauthorized"

func NoContent(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}

func Created(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, data(&model))
}

func Error(c *gin.Context, statusCode int, err error) {
	c.Error(err)
	c.AbortWithStatusJSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func Unauthorized(c *gin.Context) {
	c.Error(errors.New(unauthorizedMessage))
	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"error": unauthorizedMessage,
	})
}

func NotFoundResponse(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func JSON(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, data(&json))
}

func Paginated(c *gin.Context, p *pagination.Params, json interface{}) {
	c.JSON(http.StatusOK, &gin.H{
		"data":       json,
		"pagination": p,
	})
}

func StringResponse(c *gin.Context, string string) {
	c.String(http.StatusOK, string)
}

func data(d *interface{}) *gin.H {
	return &gin.H{
		"data": d,
	}
}
