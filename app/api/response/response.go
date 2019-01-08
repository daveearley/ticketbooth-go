package response

import (
	"errors"
	"fmt"
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
)

var errNoData = errors.New("transformer returned nil")

const unauthorizedMessage string = "This action is unauthorized"

func NoContent(c *gin.Context) {
	c.String(http.StatusNoContent, "")
}

func Created(c *gin.Context, model interface{}) {
	c.JSON(http.StatusCreated, envelope(c, model))
}

func Error(c *gin.Context, err error) {

	if e, ok := err.(*app.Error); ok {
		fmt.Println(e)
		panic("HEYEYEYEY")
	}

	c.Error(err)
	c.AbortWithStatusJSON(200, gin.H{
		"error": err.Error(),
	})
}

func NotFoundResponse(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

func JSON(c *gin.Context, json interface{}) {
	c.JSON(http.StatusOK, envelope(c, json))
}

func Paginated(c *gin.Context, p *pagination.Params, json interface{}) {
	data := transform(c, json)

	if data == nil {
		Error(c, app.ServerError(errNoData))
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
	data := transform(c, d)

	if data == nil {
		log.Fatal(errNoData)
	}

	return &gin.H{
		"data": data,
	}
}

// transform takes one or more models and transforms them according to their type.
// todo - This solution needs work. Possibly use something other than reflection
func transform(c *gin.Context, data interface{}) interface{} {
	switch reflect.TypeOf(data) {
	// Multiple
	case reflect.TypeOf([]*models.Event{}):
		return TransformEvents(c, data.([]*models.Event))
	case reflect.TypeOf([]*models.Ticket{}):
		return TransformTickets(c, data.([]*models.Ticket))
	case reflect.TypeOf([]*models.Question{}):
		return TransformQuestions(c, data.([]*models.Question))
	case reflect.TypeOf([]*models.Attribute{}):
		return TransformAttributes(c, data.([]*models.Attribute))

		// Single
	case reflect.TypeOf(&models.Event{}):
		return TransformEvent(c, data.(*models.Event))
	case reflect.TypeOf(&models.Ticket{}):
		return TransformTicket(c, data.(*models.Ticket))
	case reflect.TypeOf(&models.Question{}):
		return TransformQuestion(c, data.(*models.Question))
	}

	return data
}
