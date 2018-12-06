package request

import "time"

type CreateTicket struct {
	Title             string                 `json:"title" form:"title" binding:"required" `
	SaleStartDate     time.Time              `json:"sale_start_date" form:"sale_start_date" time_format:"2013-01-02T15:04"`
	SaleEndDate       time.Time              `json:"sale_end_date" form:"sale_end_date" time_format:"2013-01-02T15:04" binding:"gtefield=SaleStartDate"`
	Quantity          int                    `json:"quantity" binding:"min=0"`
	MaxPerTransaction int                    `json:"max_per_transaction,omitempty" form:"max_per_transaction" binding:"omitempty,min=1"`
	Attributes        map[string]interface{} `json:"attributes" form:"attributes"`
}
