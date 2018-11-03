package events

import (
	"fmt"
	"github.com/daveearley/ticketbooth/app/models/generated"
)

type Response struct {
	*models.Event
	Attributes []*models.Attribute `json:"attributes"`
}

func TransformOne(e *models.Event) Response {

	fmt.Println(e)

	return Response{e, e.R.Attributes}
}

func TransformMany(events []*models.Event) []Response {
	var transformedEvents []Response
	for _, v := range events {
		transformedEvents = append(transformedEvents, TransformOne(v))
	}

	return transformedEvents
}
