package modelx

import "github.com/daveearley/ticketbooth/app/models/generated"

type Ticket struct {
	*models.Ticket
	QuantityAvailable int
}
