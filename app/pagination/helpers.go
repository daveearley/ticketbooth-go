package pagination

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func ShouldSetCursor(p *Params, count int) bool {
	if count == 0 {
		return false
	}

	// If the number of results is less than the limit we are on the last page
	if count < p.Limit {
		return false
	}

	return true
}

func QueryMods(p *Params, authUser *models.User) []qm.QueryMod {
	var mods []qm.QueryMod

	if p.CursorID != 0 {
		var whereClause string

		if p.OrderByDirection == "asc" {
			whereClause = "id>?"
		} else {
			whereClause = "id<?"
		}

		mods = append(mods, qm.Where(whereClause, p.GetCursorID()))
	}

	mods = append(mods, []qm.QueryMod{
		qm.Where("account_id=?", authUser.AccountID),
		qm.OrderBy(p.GetOrderBy() + " " + p.OrderByDirection),
		qm.Limit(p.GetLimit()),
	}...)

	return mods
}
