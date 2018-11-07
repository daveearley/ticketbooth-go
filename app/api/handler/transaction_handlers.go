package handler

import (
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
)

type transactionHandlers struct {
	srv      service.TicketService
	eventSrv service.EventService
}

func NewTransactionHandlers(ticketSrv service.TicketService, eventSrv service.EventService) *transactionHandlers {
	return &transactionHandlers{ticketSrv, eventSrv}
}

func (h *transactionHandlers) PublicCreateTransaction(c *gin.Context) {
	// accept tickets IDs & quantity
	// return new transaction ID + expiry time
	// return url to complete transaction

	response.JSON(c, gin.H{"working": "ty"})
}
