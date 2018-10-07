package controller

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EventController struct {
	srv service.EventServiceI
}

func NewEventController(srv service.EventServiceI) *EventController {
	return &EventController{srv}
}

func (ec *EventController) CreateEvent(c *gin.Context) {
	createRequest := request.CreateEvent{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	event, err := ec.srv.CreateEvent(createRequest)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	CreatedResponse(c, event)
}
