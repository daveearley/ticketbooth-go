package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/daveearley/product/pkg/repository"
	"net/http"
	"github.com/daveearley/product/pkg/model"
)

type AccountController struct {
	rep repository.AccountRepository
}

func NewAccountController(rep repository.AccountRepository) *AccountController {
	return &AccountController{rep}
}

func (ac *AccountController) GetById(c *gin.Context) {
	account, err := ac.rep.GetById(1)

	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{
		"account": account.Id,
	})
}

func (ac *AccountController) Store(c *gin.Context) {
	account := model.Account{}
	if err := c.ShouldBindJSON(&account); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := ac.rep.Store(&account); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	CreatedResponse(c)
	return
}
