package api

import (
	"database/sql"
	"github.com/daveearley/product/pkg/api/controller"
	"github.com/daveearley/product/pkg/api/middleware"
	"github.com/daveearley/product/pkg/repository/postgres"
	"github.com/daveearley/product/pkg/service"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes handles all DI and creation of routes
func RegisterRoutes(server *gin.Engine, db *sql.DB) {
	server.GET("/healthcheck", func(context *gin.Context) {
		controller.StringResponse(context, "")
	})

	// Login routes
	authController := controller.NewAuthController(
		service.NewAuthService(
			repository.NewUserRepository(db),
		),
	)
	server.POST("/login", authController.Login)

	apiGroup := server.Group("/api")
	{
		userRepo := repository.NewUserRepository(db)
		apiGroup.Use(api.JwtMiddleware(userRepo))

		// Account routes
		accountGroup := apiGroup.Group("account")
		{
			accountController := controller.NewAccountController(
				service.NewAccountService(
					repository.NewAccountRepository(db),
					userRepo,
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
			eventGroup.GET(":id", eventController.GetById)
		}
	}
}
