package response

import (
	"github.com/daveearley/ticketbooth/app"
	"github.com/daveearley/ticketbooth/app/models/generated"
	"github.com/gin-gonic/gin"
)

type QuestionResponse struct {
	*models.Question
	QuestionOptions interface{} `json:"options"`
}

type PublicQuestionResponse struct {
	Title           string      `json:"title"`
	Type            string      `json:"type"`
	Required        bool        `json:"required"`
	QuestionOptions interface{} `json:"options"`
}

type QuestionOptionResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func TransformQuestion(c *gin.Context, q *models.Question) interface{} {
	if app.IsUserAuthenticated(c) {
		return &QuestionResponse{q, TransformQuestionOptions(c, q.R.QuestionOptions)}
	}

	return &PublicQuestionResponse{
		Title:           q.Title,
		QuestionOptions: TransformQuestionOptions(c, q.R.QuestionOptions),
		Type:            q.Type,
		Required:        q.Required,
	}
}

func TransformQuestions(c *gin.Context, questions []*models.Question) interface{} {
	var transformed []interface{}
	for _, v := range questions {
		transformed = append(transformed, TransformQuestion(c, v))
	}

	return &transformed
}

func TransformQuestionOptions(c *gin.Context, questionOptions []*models.QuestionOption) interface{} {
	var transformed []interface{}
	for _, v := range questionOptions {
		transformed = append(transformed, QuestionOptionResponse{
			ID:    v.ID,
			Title: v.Title,
		})
	}

	return &transformed
}
