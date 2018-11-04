package handler

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/api/transformer"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type eventHandlers struct {
	srv service.EventService
}

func NewEventHandlers(srv service.EventService) *eventHandlers {
	return &eventHandlers{srv}
}

func (ec *eventHandlers) GetById(c *gin.Context) {
	event, exists := c.Get("event")

	if !exists {
		response.NotFoundResponse(c)
		return
	}

	response.JSON(c, transformer.TransformEvent(c, event.(*models.Event)))
}

func (ec *eventHandlers) DeleteEvent(c *gin.Context) {
	event, exists := c.Get("event")

	if !exists {
		response.NotFoundResponse(c)
		return
	}

	if err := ec.srv.Delete(event.(*models.Event).ID); err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.NoContent(c)
}

func (ec *eventHandlers) CreateEvent(c *gin.Context) {
	createRequest := request.CreateEvent{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	event, err := ec.srv.Create(createRequest, app.GetUserFromContext(c))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Created(c, transformer.TransformEvent(c, event))
}

func (ec *eventHandlers) GetAll(c *gin.Context) {
	paginationParams := pagination.NewParams()

	if err := c.ShouldBindQuery(paginationParams); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	events, err := ec.srv.List(paginationParams, app.GetUserFromContext(c))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Paginated(c, paginationParams, transformer.TransformEvents(c, events))
}
