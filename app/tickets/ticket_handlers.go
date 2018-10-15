package tickets

import (
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/api/response"
	"github.com/daveearley/product/app/events"
	"github.com/daveearley/product/app/models/generated"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	ticketSrv Service
	eventSrv  events.Service
}

func NewController(ticketSrv Service, eventSrv events.Service) *controller {
	return &controller{ticketSrv, eventSrv}
}

func (ec *controller) CreateTicket(c *gin.Context) {
	createRequest := request.CreateTicket{}

	e, _ := c.Get("event")

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	ticket, err := ec.ticketSrv.Create(createRequest, e.(*models.Event))

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ticket)
}
