package api

import (
	"database/sql"
	"github.com/daveearley/product/app/accounts"
	"github.com/daveearley/product/app/api/middleware"
	"github.com/daveearley/product/app/api/response"
	"github.com/daveearley/product/app/auth"
	"github.com/daveearley/product/app/events"
	"github.com/daveearley/product/app/tickets"
	"github.com/daveearley/product/app/users"
	"github.com/daveearley/product/configs"
	"github.com/gin-gonic/gin"
)

// Register routes, handlers all DI
func BootstrapAndRegisterRoutes(server *gin.Engine, db *sql.DB, config *configs.Config) {
	// Error handing middleware
	server.Use(middleware.ErrorHandler())

	// Health Check
	server.GET("/healthcheck", func(context *gin.Context) {
		response.StringResponse(context, "")
	})

	// Repositories
	userRepo := users.NewRepository(db)
	eventRepo := events.NewRepository(db)
	ticketRepo := tickets.NewRepository(db)
	accountRepo := accounts.NewRepository(db)

	// Services
	authService := auth.NewService(userRepo, config)
	eventService := events.NewService(eventRepo)
	ticketService := tickets.NewService(ticketRepo)
	accountService := accounts.NewService(accountRepo, userRepo)

	// Controllers
	authController := auth.NewController(authService)
	ticketController := tickets.NewController(ticketService, eventService)
	eventController := events.NewController(eventService)
	accountController := accounts.NewController(accountService)

	server.POST("/login", authController.Login)

	apiGroup := server.Group("/v1")
	{
		apiGroup.Use(middleware.JwtMiddleware(userRepo, config))
		apiGroup.Use(middleware.PreloadModels(eventRepo, accountRepo, ticketRepo))
		apiGroup.Use(middleware.AuthorizeActions())

		// Account routes
		apiGroup.POST("/accounts", accountController.CreateAccount)
		apiGroup.GET("/accounts/:account_id", accountController.GetById)

		// Event routes
		apiGroup.POST("/events", eventController.CreateEvent)
		apiGroup.GET("/events/:event_id", eventController.GetById)
		apiGroup.GET("/events", eventController.GetAll)
		apiGroup.DELETE("/events/:event_id", eventController.DeleteEvent)

		apiGroup.POST("/events/:event_id/tickets", ticketController.CreateTicket)
		apiGroup.GET("/events/:event_id/tickets", ticketController.GetAll)
		apiGroup.GET("/events/:event_id/tickets/:ticket_id", ticketController.GetByID)
		apiGroup.DELETE("/events/:event_id/tickets/:ticket_id", ticketController.DeleteByID)

		apiGroup.POST("/events/:event_id/questions", ticketController.CreateTicket)             //CreateQuestion
		apiGroup.GET("/events/:event_id/questions", ticketController.CreateTicket)              //GetQuestion
		apiGroup.GET("/events/:event_id/questions/:question_id", ticketController.CreateTicket) //Paginate
	}
}
