package jsonmodel

import "github.com/daveearley/product/app/models/generated"

type Event struct {
	*models.Event
	Attributes *models.AttendeeSlice `json:"attributes"`
}
