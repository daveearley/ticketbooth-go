package handler

import (
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/api/transformer"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ticketHandlers struct {
	srv      service.TicketService
	eventSrv service.EventService
}

func NewTicketHandlers(ticketSrv service.TicketService, eventSrv service.EventService) *ticketHandlers {
	return &ticketHandlers{ticketSrv, eventSrv}
}

func (ec *ticketHandlers) GetByID(c *gin.Context) {
	ticket, _ := c.Get("ticket")

	response.JSON(c, transformer.TransformTicket(c, ticket.(*models.Ticket)))
}

func (ec *ticketHandlers) DeleteByID(c *gin.Context) {
	ticket, _ := c.Get("ticket")

	err := ec.srv.Delete(ticket.(*models.Ticket).ID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.NoContent(c)
}

func (ec *ticketHandlers) CreateTicket(c *gin.Context) {
	createRequest := request.CreateTicket{}

	e, _ := c.Get("event")

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	ticket, err := ec.srv.Create(createRequest, e.(*models.Event))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Created(c, transformer.TransformTicket(c, ticket))
}

func (ec *ticketHandlers) GetAll(c *gin.Context) {
	paginationParams := pagination.NewParams()

	if err := c.ShouldBindQuery(paginationParams); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	event, _ := c.Get("event")
	tix, err := ec.srv.List(paginationParams, event.(*models.Event))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Paginated(c, paginationParams, transformer.TransformTickets(c, tix))
}

func (ec *ticketHandlers) AddQuestion(c *gin.Context) {
	createRequest := request.CreateQuestion{}

	ticket, _ := c.Get("ticket")

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	_, err := ec.srv.CreateQuestion(createRequest, ticket.(*models.Ticket))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Created(c, transformer.TransformTicket(c, ticket.(*models.Ticket)))
}
