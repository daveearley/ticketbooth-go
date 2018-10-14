package event

import "github.com/daveearley/product/app/models/generated"

type Response struct {
	*models.Event
	Attributes []*models.Attribute `json:"attributes"`
}

func TransformOne(e *models.Event) Response {
	return Response{e, e.R.Attributes}
}

func TransformMany(events []*models.Event) []Response {
	var transformedEvents []Response
	for _, v := range events {
		transformedEvents = append(transformedEvents, TransformOne(v))
	}

	return transformedEvents
}
