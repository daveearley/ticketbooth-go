package controller

import (
	"github.com/daveearley/product/pkg/model"
	"github.com/daveearley/product/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type AccountController struct {
	rep repository.AccountRepositoryI
	Db  *gorm.DB
}

func NewAccountController(rep repository.AccountRepositoryI, db *gorm.DB) *AccountController {
	return &AccountController{rep, db}
}

func (ac *AccountController) GetById(c *gin.Context) {
	account, err := ac.rep.GetById(c.Param("id"))

	if err != nil {
		NotFoundResponse(c)
		return
	}

	PaginatedResponse(c, account, ac.Db)
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
