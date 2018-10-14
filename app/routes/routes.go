package routes

import (
	"database/sql"
	"github.com/daveearley/product/app/account"
	"github.com/daveearley/product/app/auth"
	"github.com/daveearley/product/app/event"
	"github.com/daveearley/product/app/middleware"
	"github.com/daveearley/product/app/response"
	"github.com/daveearley/product/app/ticket"
	"github.com/daveearley/product/app/user"
	"github.com/gin-gonic/gin"
)

// Register handles all DI and creation of routes
func Register(server *gin.Engine, db *sql.DB) {
	// Health Check
	server.GET("/healthcheck", func(context *gin.Context) {
		response.StringResponse(context, "")
	})

	// Repos
	userRepo := user.NewRepository(db)

	// Services
	authService := auth.NewService(user.NewRepository(db))
	eventService := event.NewService(event.NewRepository(db))
	ticketService := ticket.NewService(ticket.NewRepository(db))
	accountService := account.NewService(account.NewRepository(db), userRepo)

	// Controllers
	authController := auth.NewController(authService)
	ticketController := ticket.NewController(ticketService, eventService)
	eventController := event.NewController(eventService)
	accountController := account.NewController(accountService)

	server.POST("/login", authController.Login)

	apiGroup := server.Group("/v1")
	{
		apiGroup.Use(api.JwtMiddleware(userRepo))

		// Account routes
		apiGroup.POST("/account", accountController.CreateAccount)
		apiGroup.GET("account/:id", accountController.GetById)

		// Event routes
		apiGroup.POST("/events", eventController.CreateEvent)
		apiGroup.GET("/events/:id", eventController.GetById)
		apiGroup.GET("/events", eventController.GetEvents)
		apiGroup.POST("/events/:event_id/tickets", ticketController.CreateTicket)
	}
}
