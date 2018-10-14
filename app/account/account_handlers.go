package account

import (
	"github.com/daveearley/product/app/request"
	"github.com/daveearley/product/app/response"
	"github.com/daveearley/product/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv Service
}

func NewController(as Service) *controller {
	return &controller{as}
}

func (ac *controller) GetById(c *gin.Context) {
	account, err := ac.srv.Find(utils.Str2int(c.Param("id")))

	if err != nil {
		response.NotFoundResponse(c)
		return
	}

	response.JsonResponse(c, account)
}

func (ac *controller) CreateAccount(c *gin.Context) {
	createRequest := request.CreateAccount{}
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	account, err := ac.srv.Create(&createRequest)

	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.CreatedResponse(c, account)
	return
}
