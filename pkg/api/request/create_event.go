package request

import "time"

type CreateEvent struct {
	Title       string                 `json:"title" form:"title" binding:"required" `
	StartDate   time.Time              `json:"start_date" form:"start_date" time_format:"2013-01-02T15:04"`
	EndDate     time.Time              `json:"end_date" form:"end_date" time_format:"2013-01-02T15:04" binding:"gtefield=StartDate"`
	Description string                 `json:"description" form:"description"`
	Attributes  map[string]interface{} `json:"attributes" form:"attributes"`
}
