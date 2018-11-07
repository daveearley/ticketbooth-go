package api

import (
	"database/sql"
	"github.com/daveearley/ticketbooth/app/api/handler"
	"github.com/daveearley/ticketbooth/app/api/middleware"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/repository"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/daveearley/ticketbooth/configs"
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
	userRepo := repository.NewUserRepository(db)
	eventRepo := repository.NewEventRepository(db)
	ticketRepo := repository.NewTicketRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	questionRepo := repository.NewQuestionRepository(db)

	// Services
	authService := service.NewAuthService(userRepo, config)
	eventService := service.NewEventService(eventRepo)
	ticketService := service.NewTicketService(ticketRepo, questionRepo)
	accountService := service.NewAccountService(accountRepo, userRepo)

	// Controllers
	authHandlers := handler.NewAuthHandlers(authService)
	ticketHandlers := handler.NewTicketHandlers(ticketService, eventService)
	eventHandlers := handler.NewEventHandlers(eventService, ticketService)
	accountHandlers := handler.NewAccountHandlers(accountService)
	transactionHandlers := handler.NewTransactionHandlers(ticketService, eventService)

	server.POST("/login", authHandlers.Login)

	server.Use(middleware.PreloadModels(eventRepo, accountRepo, ticketRepo))
	server.Use(middleware.DbTransaction(db))

	apiAuthGroup := server.Group("/v1")
	{
		apiAuthGroup.Use(middleware.JwtMiddleware(userRepo, config))
		apiAuthGroup.Use(middleware.AuthorizeActions())

		// Account routes
		apiAuthGroup.POST("/accounts", accountHandlers.CreateAccount)
		apiAuthGroup.GET("/accounts/:account_id", accountHandlers.GetById)
		apiAuthGroup.DELETE("/accounts/:account_id", accountHandlers.Delete)

		// Event routes
		apiAuthGroup.POST("/events", eventHandlers.CreateEvent)
		apiAuthGroup.GET("/events/:event_id", eventHandlers.GetById)
		apiAuthGroup.GET("/events", eventHandlers.GetAll)
		apiAuthGroup.DELETE("/events/:event_id", eventHandlers.DeleteEvent)

		// Attendees
		// Create
		// GetAll
		// GetByID
		// Delete
		//

		// Tickets
		apiAuthGroup.POST("/events/:event_id/tickets", ticketHandlers.CreateTicket)
		apiAuthGroup.GET("/events/:event_id/tickets", ticketHandlers.GetAll)
		apiAuthGroup.GET("/events/:event_id/tickets/:ticket_id", ticketHandlers.GetByID)
		apiAuthGroup.DELETE("/events/:event_id/tickets/:ticket_id", ticketHandlers.DeleteByID)
		apiAuthGroup.POST("/tickets/:ticket_id/questions", ticketHandlers.AddQuestion)

		// Transactions
		// GetAll
		// Delete
		// Create
		// GetOne

	}

	apiPublicGroup := server.Group("/v1/public")
	{
		apiPublicGroup.GET("/events/:event_id", eventHandlers.PublicGetByID)
		apiPublicGroup.POST("/transaction", transactionHandlers.PublicCreateTransaction)

		// 1. GET get event & tickets in single request
		// 2. POST reserve tickets & return transaction ID, ticket questions etc., expiry time
		// 3. POST transaction/:transaction_id with order/ticket/payment info
		// 4. done?
	}
}
