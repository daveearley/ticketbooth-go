package transformer

import "github.com/daveearley/ticketbooth/app/models/generated"

type QuestionResponse struct {
	*models.Question
	QuestionOptions []*models.QuestionOption `json:"options"`
}

func TransformQuestion(q *models.Question) *QuestionResponse {
	return &QuestionResponse{q, q.R.QuestionOptions}
}

func TransformQuestions(questions []*models.Question) []*QuestionResponse {
	var transformed []*QuestionResponse
	for _, v := range questions {
		transformed = append(transformed, TransformQuestion(v))
	}

	return transformed
}
