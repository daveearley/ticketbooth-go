package service

import (
	"../../app"
	"../../app/api/request"
	"../../app/models/generated"
	"../../app/repository"
	"../../app/utils"
	"strconv"
)

type transactionSrv struct {
	transRepo repository.TransactionRepository
	ticketSrv TicketService
}

type TransactionService interface {
	ValidateTicketRequest(req *request.CreateTransaction) ([]*models.Ticket, error)
	CreateTransaction(req *request.CreateTransaction, event *models.Event) (*models.Transaction, []*models.Ticket, error)
	GetTicketsFromRequest(req *request.CreateTransaction) ([]*models.Ticket, error)
}

func NewTransactionService(transRepo repository.TransactionRepository, ticServ TicketService) *transactionSrv {
	return &transactionSrv{transRepo, ticServ}
}

func (s *transactionSrv) CreateTransaction(req *request.CreateTransaction, event *models.Event) (*models.Transaction, []*models.Ticket, error) {
	tickets, err := s.ValidateTicketRequest(req)

	if err != nil {
		return nil, nil, err
	}

	ticQtyMap := make(TicketQuantityMap)
	for _, v := range req.Tickets {
		ticQtyMap[v.ID] = v.Quantity
	}

	trans, err := s.transRepo.Store(&models.Transaction{
		EventID: event.ID,
		// SqlBoiler panics if these defaults are not set.
		// Likely a bug in the library related to decimal types.
		// Todo - investigate
		Total:         utils.IntToDecimal(0.00),
		TotalDiscount: utils.IntToDecimal(0.00),
		TotalTax:      utils.IntToDecimal(0.00),
	})

	if err != nil {
		return nil, nil, err
	}

	err = s.ticketSrv.ReserveTickets(ticQtyMap, trans)

	if err != nil {
		return nil, nil, err
	}

	return trans, tickets, nil
}

func (s *transactionSrv) FinalizeTransaction() {

}

func (s *transactionSrv) GetLineItems() {
	// ticket qty * (price * tax)
	// ticket title
	//
}

func (s *transactionSrv) ValidateTicketRequest(req *request.CreateTransaction) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	for i, v := range req.Tickets {
		tic, err := s.ticketSrv.Find(v.ID)

		if err != nil {
			return nil, err
		}

		remaining, err := s.ticketSrv.GetRemainingTicketQuantity(tic)

		if err != nil {
			return nil, err
		}

		if !tic.MaxPerTransaction.IsZero() && v.Quantity > tic.MaxPerTransaction.Int {
			return nil, app.InvalidValueError("tickets."+strconv.Itoa(i)+".quantity", "quantity is greater than allowed")
		}

		if v.Quantity > remaining {
			return nil, app.InvalidValueError("tickets."+strconv.Itoa(i)+".quantity", "quantity is greater than remaining")
		}

		tickets = append(tickets, tic)
	}

	return tickets, nil
}

//GetTicketsFromRequest
func (s *transactionSrv) GetTicketsFromRequest(req *request.CreateTransaction) ([]*models.Ticket, error) {
	var tickets []*models.Ticket
	for _, v := range req.Tickets {
		tic, err := s.ticketSrv.Find(v.ID)

		if err != nil {
			return nil, err
		}

		tickets = append(tickets, tic)
	}

	return tickets, nil
}
