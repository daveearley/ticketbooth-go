package service

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/api/request"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/daveearley/ticketbooth/app/repository"
	"github.com/daveearley/ticketbooth/app/utils"
	"strconv"
	"time"
)

type transactionSrv struct {
	transRepo repository.TransactionRepository
	ticketSrv TicketService
}

type TransactionService interface {
	GetTicketsFromRequest(req *request.CreateTransaction) ([]*models.Ticket, error)
	CreateTransaction(req *request.CreateTransaction, event *models.Event) (*CreateTransactionResponse, error)
}

type TransactionItem struct {
	TicketID int
	Title    string
	Price    float64
	Quantity int
	Total    float64
}

type CreateTransactionResponse struct {
	Transaction       *models.Transaction
	Tickets           []*models.Ticket
	TransactionExpiry time.Time
	Items             []TransactionItem
	Total             float64
	Tax               float64
}

func NewTransactionService(transRepo repository.TransactionRepository, ticServ TicketService) *transactionSrv {
	return &transactionSrv{transRepo, ticServ}
}

func (s *transactionSrv) CreateTransaction(req *request.CreateTransaction, event *models.Event) (*CreateTransactionResponse, error) {
	tickets, err := s.GetTicketsFromRequest(req)

	if err != nil {
		return nil, err
	}

	ticQtyMap := make([]TransactionItem, len(req.Tickets))
	for _, v := range req.Tickets {
		ticQtyMap = append(ticQtyMap, TransactionItem{
			TicketID:v.ID,
			Quantity:v.Quantity,
		})
	}

	trans, err := s.transRepo.Store(&models.Transaction{
		EventID: event.ID,
		// SqlBoiler panics if these defaults are not set.
		// Todo - investigate
		Total:         utils.IntToDecimal(0.00),
		TotalDiscount: utils.IntToDecimal(0.00),
		TotalTax:      utils.IntToDecimal(0.00),
		Status:        models.TransactionStatusPENDING,
	})

	if err != nil {
		return nil, err
	}

	transactionExpiry, err := s.ticketSrv.ReserveTickets(ticQtyMap, trans)

	if err != nil {
		return nil, err
	}

	return &CreateTransactionResponse{
		Transaction:       trans,
		Tickets:           tickets,
		TransactionExpiry: transactionExpiry,
	}, nil
}

func (s *transactionSrv) FinalizeTransaction() {

}

func (s *transactionSrv) GetLineItems() {
	// ticket qty * (price * tax)
	// ticket title
	//
}

//GetTicketsFromRequest loads tickets from the request and validates the quantity
func (s *transactionSrv) GetTicketsFromRequest(req *request.CreateTransaction) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	for i, v := range req.Tickets {
		tic, err := s.ticketSrv.Find(v.ID)

		if err != nil {
			return nil, err
		}

		if err = s.ValidateTicketQty(tic, i, v.Quantity); err != nil {
			return nil, err
		}

		tickets = append(tickets, tic)
	}

	return tickets, nil
}

//ValidateTicketQty checks if the user supplied ticket quantity is achievable
func (s *transactionSrv) ValidateTicketQty(ticket *models.Ticket, ticketIndex int,  qty int) error {
	remaining, err := s.ticketSrv.GetRemainingTicketQuantity(ticket)

	if err != nil {
		return err
	}

	if !ticket.MaxPerTransaction.IsZero() && qty > ticket.MaxPerTransaction.Int {
		return app.InvalidValueError("tickets."+strconv.Itoa(ticketIndex)+".quantity", "quantity is greater than allowed")
	}

	if qty > remaining {
		return app.InvalidValueError("tickets."+strconv.Itoa(ticketIndex)+".quantity", "quantity is greater than remaining")
	}

	return nil
}