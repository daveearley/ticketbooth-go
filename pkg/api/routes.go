package api

import (
	"github.com/gin-gonic/gin"
	"github.com/daveearley/product/pkg/api/controller"
	"github.com/daveearley/product/pkg/repository/mysql"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(server *gin.Engine, db *gorm.DB) {
	server.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, gin.H{})
	})

	apiGroup := server.Group("/api")
	{
		accountGroup := apiGroup.Group("/account")
		{
			accountController := controller.NewAccountController(repository.NewAccountRepository(db))
			accountGroup.POST("", accountController.Store)
			accountGroup.GET("/:id", accountController.GetById)
		}
	}
}