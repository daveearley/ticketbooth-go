package pagination

import (
	"strings"
)

const defaultPaginationLimit = 20
const defaultCursorValue = 0

type Params struct {
	Limit            int    `form:"limit" json:"limit"`
	OrderBy          string `form:"order_by" json:"-"`
	OrderByDirection string `json:"-"`
	CursorID         int    `form:"cursor" json:"cursor_id,omitempty"`
	NextCursorId     int    `json:"next_cursor"`
	ResultCount      int    `json:"result_count,omitempty"`
}

func (p *Params) GetOrderByDirection() string {
	dir := strings.Split(p.OrderBy, "_")[1]
	p.OrderByDirection = dir

	return dir
}

func (p *Params) GetOrderBy() string {
	if p.OrderBy == "" {
		p.OrderBy = "id_desc"
	}

	return strings.Split(p.OrderBy, "_")[0]
}

func (p *Params) GetCursorID() int {
	if p.CursorID == 0 {
		return defaultCursorValue
	}

	return p.CursorID
}

func (p *Params) GetLimit() int {
	if p.Limit == 0 {
		return defaultPaginationLimit
	}

	return p.Limit
}
