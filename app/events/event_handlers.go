package events

import (
	"fmt"
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/pagination"
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/response"
	"github.com/daveearley/product/app/utils"
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
	event, err := ec.srv.Find(utils.Str2int(c.Param("id")))

	if !app.IsUserAuthorized(c, event) {
		response.Unauthorized(c)
		return
	}

	if err != nil {
		response.NotFoundResponse(c)
		return
	}

	response.JsonResponse(c, event)
}

func (ec *controller) CreateEvent(c *gin.Context) {
	createRequest := request.CreateEvent{}

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

func (ec *controller) GetEvents(c *gin.Context) {
	paginationParams := pagination.Params{}

	if err := c.ShouldBindQuery(&paginationParams); err != nil {
		fmt.Println(err.Error())
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	events, err := ec.srv.List(&paginationParams, app.GetUserFromContext(c))

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.Paginated(c, &paginationParams, TransformMany(events))
}
