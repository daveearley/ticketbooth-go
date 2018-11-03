package jsonmodel

import "github.com/daveearley/ticketbooth/app/models/generated"

type Event struct {
	*models.Event
	Attributes *models.AttendeeSlice `json:"attributes"`
}
