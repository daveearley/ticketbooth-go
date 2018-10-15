package pagination

import (
	"github.com/daveearley/product/app/models/generated"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

func QueryMods(p *Params, authUser *models.User) []qm.QueryMod {
	var mods []qm.QueryMod

	mods = append(mods, []qm.QueryMod{
		qm.Where("account_id=?", authUser.AccountID),
		qm.OrderBy(p.GetOrderBy() + " " + p.OrderByDirection),
		qm.Limit(p.GetLimit()),
		qm.Offset(p.GetOffset()),
	}...)

	return mods
}
