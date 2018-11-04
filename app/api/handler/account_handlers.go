package handler

import (
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type accountHandlers struct {
	srv service.AccountService
}

func NewAccountHandlers(as service.AccountService) *accountHandlers {
	return &accountHandlers{as}
}

func (ac *accountHandlers) GetById(c *gin.Context) {
	account, _ := c.Get("account")

	response.JSON(c, account)
}

func (ac *accountHandlers) Delete(c *gin.Context) {
	account, _ := c.Get("account")

	err := ac.srv.Delete(account.(*models.Account))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.NoContent(c)
}

func (ac *accountHandlers) CreateAccount(c *gin.Context) {
	createRequest := request.CreateAccount{}
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	account, err := ac.srv.Create(&createRequest)

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Created(c, account)
}
