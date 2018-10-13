package event

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv Service
}

func NewController(srv Service) *controller {
	return &controller{srv}
}

func (ec *controller) GetById(c *gin.Context) {
	event, err := ec.srv.Find(utils.Str2int(c.Param("id")))

	if err != nil {
		app.NotFoundResponse(c)
		return
	}

	app.JsonResponse(c, event)
}

func (ec *controller) CreateEvent(c *gin.Context) {
	createRequest := request.CreateEvent{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	event, err := ec.srv.CreateEvent(createRequest, app.GetUserFromContext(c))

	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	app.CreatedResponse(c, event)
}
