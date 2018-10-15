package ticket

import (
	"github.com/daveearley/product/app/event"
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	ticketSrv Service
	eventSrv  event.Service
}

func NewController(ticketSrv Service, eventSrv event.Service) *controller {
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
