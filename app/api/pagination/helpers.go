package pagination

import (
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// QueryMods returns SQLBoiler query mods required for pagination
func QueryMods(p *Params) []qm.QueryMod {
	var mods []qm.QueryMod

	mods = append(mods, []qm.QueryMod{
		qm.OrderBy(p.GetOrderBy() + " " + p.OrderByDirection),
		qm.Limit(p.GetLimit()),
		qm.Offset(p.GetOffset()),
	}...)

	return mods
}
