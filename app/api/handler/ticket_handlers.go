package handler

import (
	"../../../app"
	"../../../app/api/pagination"
	"../../../app/api/request"
	"../../../app/api/response"
	"../../../app/models/generated"
	"../../../app/service"
	"github.com/gin-gonic/gin"
)

type ticketHandlers struct {
	srv      service.TicketService
	eventSrv service.EventService
}

func NewTicketHandlers(ticketSrv service.TicketService, eventSrv service.EventService) *ticketHandlers {
	return &ticketHandlers{ticketSrv, eventSrv}
}

func (ec *ticketHandlers) GetByID(c *gin.Context) {
	ticket, _ := c.Get(app.TicketResource)

	response.JSON(c, ticket)
}

func (ec *ticketHandlers) DeleteByID(c *gin.Context) {
	ticket, _ := c.Get(app.TicketResource)

	err := ec.srv.Delete(ticket.(*models.Ticket).ID)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.NoContent(c)
}

func (ec *ticketHandlers) CreateTicket(c *gin.Context) {
	createRequest := request.CreateTicket{}

	e, _ := c.Get(app.EventResource)

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, err)
		return
	}

	ticket, err := ec.srv.Create(createRequest, e.(*models.Event))

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, ticket)
}

func (ec *ticketHandlers) GetAll(c *gin.Context) {
	paginationParams := pagination.NewParams()

	if err := c.ShouldBindQuery(paginationParams); err != nil {
		response.Error(c, err)
		return
	}

	event, _ := c.Get(app.EventResource)
	tix, err := ec.srv.List(paginationParams, event.(*models.Event))

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Paginated(c, paginationParams, tix)
}

func (ec *ticketHandlers) AddQuestion(c *gin.Context) {
	createRequest := request.CreateQuestion{}

	ticket, _ := c.Get(app.TicketResource)

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, err)
		return
	}

	_, err := ec.srv.CreateQuestion(createRequest, ticket.(*models.Ticket))

	if err != nil {
		response.Error(c, err)
		return
	}

	response.Created(c, ticket)
}
