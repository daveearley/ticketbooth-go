package response

import (
	"github.com/daveearley/product/app/pagination"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatedResponse(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, &model)
}

func ErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}

func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": "This action is unauthorized",
	})
}

func NotFoundResponse(c *gin.Context) {
	c.String(http.StatusNotFound, "")
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
