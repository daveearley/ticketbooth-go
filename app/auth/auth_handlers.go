package auth

import (
	"github.com/daveearley/product/app/api/request"
	"github.com/daveearley/product/app/api/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type controller struct {
	srv Service
}

func NewController(srv Service) *controller {
	return &controller{srv}
}

func (ac *controller) Login(c *gin.Context) {
	var credentials *request.Login
	if err := c.ShouldBindJSON(&credentials); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	token, err := ac.srv.ValidateLoginAndReturnJwtToken(credentials)

	if err != nil {
		response.ErrorResponse(c, http.StatusForbidden, err)
		return
	}

	response.JsonResponse(c, gin.H{
		"token": token,
	})
}
