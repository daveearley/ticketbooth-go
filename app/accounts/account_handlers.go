package accounts

import (
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/api/response"
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
	account, _ := c.Get("account")

	response.JSON(c, account)
}

func (ac *controller) CreateAccount(c *gin.Context) {
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
	return
}
