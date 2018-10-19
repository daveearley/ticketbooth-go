package events

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/api/pagination"
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv Service
}

func NewController(srv Service) *controller {
	return &controller{srv}
}

func (ec *controller) GetById(c *gin.Context) {
	event, exists := c.Get("event")

	if !exists {
		response.NotFoundResponse(c)
		return
	}

	response.JSON(c, event)
}

func (ec *controller) CreateEvent(c *gin.Context) {
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

func (ec *controller) GetAll(c *gin.Context) {
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

	response.Paginated(c, paginationParams, TransformMany(events))
}
