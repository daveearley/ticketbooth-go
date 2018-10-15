package routes

import (
	"database/sql"
	"github.com/daveearley/product/app/accounts"
	"github.com/daveearley/product/app/auth"
	"github.com/daveearley/product/app/events"
	"github.com/daveearley/product/app/middleware"
	"github.com/daveearley/product/app/response"
	"github.com/daveearley/product/app/tickets"
	"github.com/daveearley/product/app/users"
	"github.com/gin-gonic/gin"
)

// Register handles all DI and creation of routes
func Register(server *gin.Engine, db *sql.DB) {
	// Health Check
	server.GET("/healthcheck", func(context *gin.Context) {
		response.StringResponse(context, "")
	})

	// Repos
	userRepo := users.NewRepository(db)

	// Services
	authService := auth.NewService(users.NewRepository(db))
	eventService := events.NewService(events.NewRepository(db))
	ticketService := tickets.NewService(tickets.NewRepository(db))
	accountService := accounts.NewService(accounts.NewRepository(db), userRepo)

	// Controllers
	authController := auth.NewController(authService)
	ticketController := tickets.NewController(ticketService, eventService)
	eventController := events.NewController(eventService)
	accountController := accounts.NewController(accountService)

	server.POST("/login", authController.Login)

	apiGroup := server.Group("/v1")
	{
		apiGroup.Use(middleware.JwtMiddleware(userRepo))
		apiGroup.Use(middleware.BindAndAuthorize(eventService, accountService))

		// Account routes
		apiGroup.POST("/accounts", accountController.CreateAccount)
		apiGroup.GET("/accounts/:account_id", accountController.GetById)

		// Event routes
		apiGroup.POST("/events", eventController.CreateEvent)
		apiGroup.GET("/events/:event_id", eventController.GetById)
		apiGroup.GET("/events", eventController.GetEvents)
		apiGroup.POST("/events/:event_id/tickets", ticketController.CreateTicket)
	}
}
