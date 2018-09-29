package controller

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/service"
	"github.com/daveearley/product/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	srv service.AccountServiceI
}

func NewAccountController(as service.AccountServiceI) *AccountController {
	return &AccountController{as}
}

func (ac *AccountController) GetById(c *gin.Context) {
	account, err := ac.srv.Find(utils.Str2Uint(c.Param("id")))

	if err != nil {
		NotFoundResponse(c)
		return
	}

	JsonResponse(c, account)
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	createRequest := request.CreateAccount{}
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	account, err := ac.srv.CreateAccount(&createRequest)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	CreatedResponse(c, account)
	return
}
