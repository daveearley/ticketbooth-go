package ticket

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv *Service
}

func NewController(srv *Service) *controller {
	return &controller{srv}
}

func (ec *controller) CreateTicket(c *gin.Context) {
	createRequest := request.CreateTicket{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	event, err := ec.srv.Create(createRequest, app.GetUserFromContext(c))

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, event)
}
