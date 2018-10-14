package ticket

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/event"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/response"
	"github.com/daveearley/product/app/utils"
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

	e, err := ec.eventSrv.Find(utils.Str2int(c.Param("event_id")))

	if !app.IsUserAuthorized(c, e) {
		response.Unauthorized(c)
		return
	}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	ticket, err := ec.ticketSrv.Create(createRequest, e)

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, ticket)
}
