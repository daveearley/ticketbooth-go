package request

type CreateTransaction struct {
	Tickets []struct {
		ID       int `json:"id" form:"id" binding:"required"`
		Quantity int `json:"quantity" form:"quantity" binding:"required"`
	} `json:"tickets" form:"tickets" binding:"required"`
}
