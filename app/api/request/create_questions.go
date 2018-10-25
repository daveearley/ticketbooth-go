package request

type QuestionOption struct {
	Title string `json:"title" form:"title" binding:"required"`
}

type CreateQuestion struct {
	Title    string           `json:"title" form:"title" binding:"required"`
	Type     string           `json:"type"`
	Required bool             `json:"required" form:"required"`
	Options  []QuestionOption `json:"options"`
}
