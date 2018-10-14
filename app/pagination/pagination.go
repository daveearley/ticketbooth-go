package pagination

import (
	"strings"
)

const defaultPaginationLimit = 20
const defaultOffset = 0

// Params stores pagination request information.
// todo: ideally switch to cursor pagination
type Params struct {
	PerPage          int    `form:"per_page" json:"per_page" binding:"max=100"`
	OrderBy          string `form:"order_by" json:"-"`
	OrderByDirection string `json:"-"`
	Page             int    `form:"page" json:"page,omitempty"`
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

func (p *Params) GetOffset() int {
	if p.Page == 0 || p.Page == 1 {
		return defaultOffset
	}

	return (p.Page - 1) * p.GetLimit()
}

func (p *Params) GetLimit() int {
	if p.PerPage == 0 {
		p.PerPage = defaultPaginationLimit
	}

	return p.PerPage
}
