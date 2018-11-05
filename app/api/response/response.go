package response

import (
	"errors"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

const unauthorizedMessage string = "This action is unauthorized"

func NoContent(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}

func Created(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, envelope(c, model))
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
	c.JSON(http.StatusOK, envelope(c, json))
}

func Paginated(c *gin.Context, p *pagination.Params, json interface{}) {

	data, err := transform(c, json)

	if err != nil {
		Error(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &gin.H{
		"envelope":   data,
		"pagination": p,
	})
}

func StringResponse(c *gin.Context, string string) {
	c.String(http.StatusOK, string)
}

// envelope wraps JSON responses in a 'data' object
func envelope(c *gin.Context, d interface{}) *gin.H {
	data, err := transform(c, d)

	if err != nil {
		log.Fatal(err)
	}

	return &gin.H{
		"data": data,
	}
}

// transform takes one or more models and transforms them according to their type.
// todo - This solution needs work. Possibly use something other than reflection
func transform(c *gin.Context, data interface{}) (interface{}, error) {
	switch reflect.TypeOf(data) {
	// Multiple
	case reflect.TypeOf([]*models.Event{}):
		return TransformEvents(c, data.([]*models.Event)), nil
	case reflect.TypeOf([]*models.Ticket{}):
		return TransformTickets(c, data.([]*models.Ticket)), nil
	case reflect.TypeOf([]*models.Question{}):
		return TransformQuestions(c, data.([]*models.Question)), nil
	case reflect.TypeOf([]*models.Attribute{}):
		return TransformAttributes(c, data.([]*models.Attribute)), nil

		// Single
	case reflect.TypeOf(&models.Event{}):
		return TransformEvent(c, data.(*models.Event)), nil
	case reflect.TypeOf(&models.Ticket{}):
		return TransformTicket(c, data.(*models.Ticket)), nil
	case reflect.TypeOf(&models.Question{}):
		return TransformQuestion(c, data.(*models.Question)), nil
	}

	return nil, errors.New("unable to transform: " + reflect.TypeOf(data).String())
}
