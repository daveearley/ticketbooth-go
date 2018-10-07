package api

import (
	"database/sql"
	"github.com/daveearley/product/pkg/api/controller"
	"github.com/daveearley/product/pkg/repository/postgres"
	"github.com/daveearley/product/pkg/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *sql.DB) {
	server.GET("/healthcheck", func(context *gin.Context) {
		controller.StringResponse(context, "")
	})

	apiGroup := server.Group("/api")
	{
		// Account routes
		accountGroup := apiGroup.Group("account")
		{
			accountController := controller.NewAccountController(
				service.NewAccountService(
					repository.NewAccountRepository(db),
					repository.NewUserRepository(db),
				),
			)
			accountGroup.POST("", accountController.CreateAccount)
			accountGroup.GET(":id", accountController.GetById)
		}

		// Event Routes
		eventGroup := apiGroup.Group("events")
		{
			eventController := controller.NewEventController(
				service.NewEventService(
					repository.NewEventRepository(db),
				),
			)
			eventGroup.POST("", eventController.CreateEvent)
			eventGroup.GET(":id", eventController.CreateEvent)
		}
	}
}
