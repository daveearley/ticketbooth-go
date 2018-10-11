package controller

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/service"
	"github.com/daveearley/product/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type eventController struct {
	srv service.EventService
}

func NewEventController(srv service.EventService) *eventController {
	return &eventController{srv}
}

func (ec *eventController) GetById(c *gin.Context) {
	event, err := ec.srv.Find(utils.Str2int(c.Param("id")))

	if err != nil {
		NotFoundResponse(c)
		return
	}

	JsonResponse(c, event)
}

func (ec *eventController) CreateEvent(c *gin.Context) {
	createRequest := request.CreateEvent{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	event, err := ec.srv.CreateEvent(createRequest, GetUserFromContext(c))

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	CreatedResponse(c, event)
}
