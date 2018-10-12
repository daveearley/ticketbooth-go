package routes

import (
	"database/sql"
	"github.com/daveearley/product/app"
	"github.com/daveearley/product/app/account"
	"github.com/daveearley/product/app/auth"
	"github.com/daveearley/product/app/event"
	"github.com/daveearley/product/app/middleware"
	"github.com/daveearley/product/app/user"
	"github.com/gin-gonic/gin"
)

// Register handles all DI and creation of routes
func Register(server *gin.Engine, db *sql.DB) {
	server.GET("/healthcheck", func(context *gin.Context) {
		app.StringResponse(context, "")
	})

	// Login routes
	authController := auth.NewController(
		auth.NewService(
			user.NewRepository(db),
		),
	)
	server.POST("/login", authController.Login)

	apiGroup := server.Group("/api")
	{
		userRepo := user.NewRepository(db)
		apiGroup.Use(api.JwtMiddleware(userRepo))

		// Account routes
		accountGroup := apiGroup.Group("account")
		{
			accountController := account.NewController(
				account.NewService(
					account.NewRepository(db),
					userRepo,
				),
			)
			accountGroup.POST("", accountController.CreateAccount)
			accountGroup.GET(":id", accountController.GetById)
		}

		// Event Routes
		eventGroup := apiGroup.Group("events")
		{
			eventController := event.NewController(
				event.NewService(
					event.NewRepository(db),
				),
			)
			eventGroup.POST("", eventController.CreateEvent)
			eventGroup.GET(":id", eventController.GetById)
		}
	}
}
