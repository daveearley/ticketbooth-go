package controller

import (
	"github.com/daveearley/product/pkg/model"
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

func (ac *AccountController) Store(c *gin.Context) {
	account := model.Account{}
	if err := c.ShouldBindJSON(&account); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if _, err := ac.srv.CreateAccount(&account); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	CreatedResponse(c, account)
	return
}
