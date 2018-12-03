package handler

import (
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/api/response"
	"github.com/daveearley/ticketbooth/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"time"
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
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	event, _ := c.Get("event")

	trans, tickets, err := h.srv.CreateTransaction(&createRequest, event.(*models.Event))

	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.JSON(c, gin.H{
		"transaction": map[string]interface{}{
			"expiry": time.Now().Add(time.Minute * 10),
			"id": trans.UUID,
			"url": c.Request.Host + c.Request.RequestURI + "/" + trans.UUID.String,
		},
		"tickets": response.TransformTickets(c, tickets),
	})
}

func (h *transactionHandlers) PublicFinalizeTransaction(c *gin.Context) {

	transaction, _ := c.Get("transaction")

	response.JSON(c, gin.H{"ff":transaction})
}