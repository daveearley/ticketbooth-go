package response

import (
	"errors"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

const unauthorizedMessage string = "This action is unauthorized"

func CreatedResponse(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, &model)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.Error(err)
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func Unauthorized(c *gin.Context) {
	c.Error(errors.New(unauthorizedMessage))
	c.JSON(http.StatusForbidden, gin.H{
		"error": unauthorizedMessage,
	})
}

func NotFoundResponse(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func JsonResponse(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, json)
}

func Paginated(c *gin.Context, p *pagination.Params, json interface{}) {
	jsonResponse := make(map[string]interface{})
	jsonResponse["data"] = json
	jsonResponse["pagination"] = p

	c.JSON(http.StatusOK, jsonResponse)
}

func StringResponse(c *gin.Context, string string) {
	c.String(http.StatusOK, string)
}
