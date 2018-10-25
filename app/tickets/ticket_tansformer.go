package tickets

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/daveearley/product/app/questions"
)

type Response struct {
	*models.Ticket
	Questions []*questions.Response `json:"questions"`
}

func TransformOne(t *models.Ticket) *Response {
	return &Response{t, questions.TransformMany(t.R.Questions)}
}

func TransformMany(tickets []*models.Ticket) []*Response {
	var transformed []*Response
	for _, v := range tickets {
		transformed = append(transformed, TransformOne(v))
	}

	return transformed
}
