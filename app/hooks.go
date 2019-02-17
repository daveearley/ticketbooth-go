package app

import "github.com/daveearley/ticketbooth/app/models/generated"

func BeforeSaveTicket(ticket *models.Ticket) error {
	return nil
}
