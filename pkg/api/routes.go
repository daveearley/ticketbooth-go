package api

import (
	"github.com/daveearley/product/pkg/api/controller"
	"github.com/daveearley/product/pkg/repository/postgres"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.GET("/healthcheck", func(context *gin.Context) {
		controller.StringResponse(context, "")
	})

	apiGroup := server.Group("/api")
	{
		accountGroup := apiGroup.Group("/account")
		{
			accountController := controller.NewAccountController(repository.NewAccountRepository(db), db)
			accountGroup.POST("", accountController.Store)
			accountGroup.GET("/:id", accountController.GetById)
		}
	}
}
