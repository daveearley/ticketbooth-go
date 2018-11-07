package handler

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type eventHandlers struct {
	srv    service.EventService
	ticSrv service.TicketService
}

func NewEventHandlers(srv service.EventService, ticSrv service.TicketService) *eventHandlers {
	return &eventHandlers{srv, ticSrv}
}

func (ec *eventHandlers) GetById(c *gin.Context) {
	event, exists := c.Get("event")

	if !exists {
		response.NotFoundResponse(c)
		return
	}

	response.JSON(c, event)
}

func (ec *eventHandlers) PublicGetByID(c *gin.Context) {
	event, _ := c.Get("event")
	tickets, err := ec.ticSrv.FindByEventID(event.(*models.Event).ID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
	}

	// todo - move to transformation layer
	var tics []*response.PublicTicketResponse
	for _, ticket := range tickets {
		qtyRemaining, _ := ec.ticSrv.GetRemainingTicketQuantity(ticket)

		tics = append(tics, &response.PublicTicketResponse{
			ID:                ticket.ID,
			Title:             ticket.Title,
			QuantityAvailable: qtyRemaining,
			MaxPerTransaction: ticket.MaxPerTransaction,
		})
	}

	resp := response.TransformEvent(c, event.(*models.Event)).(*response.PublicEventResponse)
	resp.Tickets = tics

	response.JSON(c, resp)
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

	response.Created(c, event)
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

	response.Paginated(c, paginationParams, events)
}
