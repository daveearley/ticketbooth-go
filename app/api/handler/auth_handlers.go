package handler

import (
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
)

type authHandlers struct {
	srv service.AuthService
}

func NewAuthHandlers(srv service.AuthService) *authHandlers {
	return &authHandlers{srv}
}

func (ac *authHandlers) Login(c *gin.Context) {
	var credentials *request.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		response.Error(c, err)
		return
	}

	token, err := ac.srv.ValidateLoginAndReturnJwtToken(credentials)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.JSON(c, gin.H{
		"token": token,
	})
}
