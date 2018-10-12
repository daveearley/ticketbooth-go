package account

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/request"
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
		app.NotFoundResponse(c)
		return
	}

	app.JsonResponse(c, account)
}

func (ac *controller) CreateAccount(c *gin.Context) {
	createRequest := request.CreateAccount{}
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	account, err := ac.srv.CreateAccount(&createRequest)

	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	app.CreatedResponse(c, account)
	return
}
