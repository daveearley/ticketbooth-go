package questions

import "github.com/daveearley/product/app/models/generated"

type Response struct {
	*models.Question
	QuestionOptions []*models.QuestionOption `json:"options"`
}

func TransformOne(q *models.Question) *Response {
	return &Response{q, q.R.QuestionOptions}
}

func TransformMany(questions []*models.Question) []*Response {
	var transformed []*Response
	for _, v := range questions {
		transformed = append(transformed, TransformOne(v))
	}

	return transformed
}
