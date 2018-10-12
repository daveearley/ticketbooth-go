package auth

import (
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	srv Service
}

func NewController(srv Service) *authController {
	return &authController{srv}
}

func (ac *authController) Login(c *gin.Context) {
	var credentials *request.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		app.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := ac.srv.ValidateLoginAndReturnJwtToken(credentials)

	if err != nil {
		app.ErrorResponse(c, http.StatusForbidden, err)
		return
	}

	app.JsonResponse(c, gin.H{
		"token": token,
	})
}
