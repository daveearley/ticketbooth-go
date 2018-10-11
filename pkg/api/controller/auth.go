package controller

import (
	"github.com/daveearley/product/pkg/api/request"
	"github.com/daveearley/product/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	srv service.AuthService
}

func NewAuthController(srv service.AuthService) *authController {
	return &authController{srv}
}

func (ac *authController) Login(c *gin.Context) {
	var credentials *request.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := ac.srv.ValidateLoginAndReturnJwtToken(credentials)

	if err != nil {
		ErrorResponse(c, http.StatusForbidden, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
