package handler

import (
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"github.com/daveearley/ticketbooth/app"
)

type transactionHandlers struct {
	srv service.TransactionService
}

func NewTransactionHandlers(transactionSrv service.TransactionService) *transactionHandlers {
	return &transactionHandlers{transactionSrv}
}

func (h *transactionHandlers) PublicCreateTransaction(c *gin.Context) {
	createRequest := request.CreateTransaction{}

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		response.Error(c, err)
		return
	}

	event, _ := c.Get(app.EventResource)

	transResponse, err := h.srv.CreateTransaction(&createRequest, event.(*models.Event))

	if err != nil {
		response.Error(c, err)
		return
	}

	response.JSON(c, gin.H{
		"transaction": map[string]interface{}{
			"expiry": transResponse.TransactionExpiry,
			"id":     transResponse.Transaction.UUID,
			"url":    c.Request.Host + c.Request.RequestURI + "/" + transResponse.Transaction.UUID.String,
		},
		"tickets": response.TransformTickets(c, transResponse.Tickets),
	})
}

func (h *transactionHandlers) PublicFinalizeTransaction(c *gin.Context) {
	transaction, _ := c.Get(app.TransactionResource)

	response.JSON(c, gin.H{"ff": transaction})
}
