package tickets

import (
	"github.com/daveearley/ticketbooth/app/api/pagination"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/events"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv      Service
	eventSrv events.Service
}

func NewController(ticketSrv Service, eventSrv events.Service) *controller {
	return &controller{ticketSrv, eventSrv}
}

func (ec *controller) GetByID(c *gin.Context) {
	ticket, _ := c.Get("ticket")

	response.JSON(c, TransformOne(ticket.(*models.Ticket)))
}

func (ec *controller) DeleteByID(c *gin.Context) {
	ticket, _ := c.Get("ticket")

	err := ec.srv.Delete(ticket.(*models.Ticket).ID)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.NoContent(c)
}

func (ec *controller) CreateTicket(c *gin.Context) {
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

	response.Created(c, TransformOne(ticket))
}

func (ec *controller) GetAll(c *gin.Context) {
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

	response.Paginated(c, paginationParams, TransformMany(tix))
}

func (ec *controller) AddQuestion(c *gin.Context) {
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

	response.Created(c, TransformOne(ticket.(*models.Ticket)))
}
