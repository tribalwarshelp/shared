package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Ennoblement struct {
	tableName struct{} `pg:"?SERVER.ennoblements,alias:ennoblement"`

	ID              int       `json:"id" gqlgen:"id" xml:"id"`
	VillageID       int       `pg:",use_zero" gqlgen:"-" json:"villageID" xml:"villageID"`
	Village         *Village  `gqlgen:"-" json:"village,omitempty" xml:"village" pg:"rel:has-one"`
	NewOwnerID      int       `pg:",use_zero" gqlgen:"-" json:"newOwnerID" xml:"newOwnerID"`
	NewOwner        *Player   `gqlgen:"-" json:"newOwner,omitempty" xml:"newOwner" pg:"rel:has-one"`
	NewOwnerTribeID int       `pg:",use_zero" json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	NewOwnerTribe   *Tribe    `json:"newOwnerTribe,omitempty" gqlgen:"-" xml:"newOwnerTribe" pg:"rel:has-one"`
	OldOwnerID      int       `pg:",use_zero" gqlgen:"-" json:"oldOwnerID" xml:"oldOwnerID"`
	OldOwner        *Player   `gqlgen:"-" json:"oldOwner,omitempty" xml:"oldOwner" pg:"rel:has-one"`
	OldOwnerTribeID int       `pg:",use_zero" json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
	OldOwnerTribe   *Tribe    `json:"oldOwnerTribe,omitempty" gqlgen:"-" xml:"oldOwnerTribe" pg:"rel:has-one"`
	EnnobledAt      time.Time `pg:"default:now(),use_zero" json:"ennobledAt" gqlgen:"ennobledAt" xml:"ennobledAt"`
}

type EnnoblementFilterOr struct {
	NewOwnerID      []int `json:"newOwnerID" gqlgen:"newOwnerID" xml:"newOwnerID"`
	NewOwnerTribeID []int `json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	OldOwnerID      []int `json:"oldOwnerID" gqlgen:"oldOwnerID" xml:"oldOwnerID"`
	OldOwnerTribeID []int `json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
}

func (f *EnnoblementFilterOr) WhereWithAlias(q *orm.Query, alias string) *orm.Query {
	if f != nil {
		q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			if !isZero(f.NewOwnerID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_owner_id", alias), pg.Array(f.NewOwnerID))
			}
			if !isZero(f.NewOwnerTribeID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_owner_tribe_id", alias), pg.Array(f.NewOwnerTribeID))
			}
			if !isZero(f.OldOwnerID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_owner_id", alias), pg.Array(f.OldOwnerID))
			}
			if !isZero(f.OldOwnerTribeID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_owner_tribe_id", alias), pg.Array(f.OldOwnerTribeID))
			}
			return q, nil
		})
	}
	return q
}

type EnnoblementFilter struct {
	VillageID     []int          `json:"villageID" gqlgen:"villageID" xml:"villageID"`
	VillageIDNEQ  []int          `json:"villageIDNEQ" gqlgen:"villageIDNEQ" xml:"villageIDNEQ"`
	VillageFilter *VillageFilter `json:"villageFilter" xml:"villageFilter" gqlgen:"villageFilter"`

	NewOwnerID          []int         `json:"newOwnerID" gqlgen:"newOwnerID" xml:"newOwnerID"`
	NewOwnerIDNEQ       []int         `json:"newOwnerIDNEQ" gqlgen:"newOwnerIDNEQ" xml:"newOwnerIDNEQ"`
	NewOwnerFilter      *PlayerFilter `json:"newOwnerFilter" xml:"newOwnerFilter" gqlgen:"newOwnerFilter"`
	NewOwnerTribeID     []int         `json:"newOwnerTribeID" gqlgen:"newOwnerTribeID" xml:"newOwnerTribeID"`
	NewOwnerTribeIDNEQ  []int         `json:"newOwnerTribeIDNEQ" gqlgen:"newOwnerTribeIDNEQ" xml:"newOwnerTribeIDNEQ"`
	NewOwnerTribeFilter *TribeFilter  `json:"newOwnerTribeFilter" xml:"newOwnerTribeFilter" gqlgen:"newOwnerTribeFilter"`

	OldOwnerID          []int         `json:"oldOwnerID" gqlgen:"oldOwnerID" xml:"oldOwnerID"`
	OldOwnerIDNEQ       []int         `json:"oldOwnerIDNEQ" gqlgen:"oldOwnerIDNEQ" xml:"oldOwnerIDNEQ"`
	OldOwnerFilter      *PlayerFilter `json:"oldOwnerFilter" xml:"oldOwnerFilter" gqlgen:"oldOwnerFilter"`
	OldOwnerTribeID     []int         `json:"oldOwnerTribeID" gqlgen:"oldOwnerTribeID" xml:"oldOwnerTribeID"`
	OldOwnerTribeIDNEQ  []int         `json:"oldOwnerTribeIDNEQ" gqlgen:"oldOwnerTribeIDNEQ" xml:"oldOwnerTribeIDNEQ"`
	OldOwnerTribeFilter *TribeFilter  `json:"oldOwnerTribeFilter" xml:"oldOwnerTribeFilter" gqlgen:"oldOwnerTribeFilter"`

	EnnobledAt    time.Time `json:"ennobledAt" gqlgen:"ennobledAt" xml:"ennobledAt"`
	EnnobledAtGT  time.Time `json:"ennobledAtGT" gqlgen:"ennobledAtGT" xml:"ennobledAtGT"`
	EnnobledAtGTE time.Time `json:"ennobledAtGTE" gqlgen:"ennobledAtGTE" xml:"ennobledAtGTE"`
	EnnobledAtLT  time.Time `json:"ennobledAtLT" gqlgen:"ennobledAtLT" xml:"ennobledAtLT"`
	EnnobledAtLTE time.Time `json:"ennobledAtLTE" gqlgen:"ennobledAtLTE" xml:"ennobledAtLTE"`

	Or *EnnoblementFilterOr `json:"or" gqlgen:"or" xml:"or"`
}

func (f *EnnoblementFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	if !isZero(f.EnnobledAt) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("ennobled_at", alias), f.EnnobledAt)
	}
	if !isZero(f.EnnobledAtGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("ennobled_at", alias), f.EnnobledAtGT)
	}
	if !isZero(f.EnnobledAtGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("ennobled_at", alias), f.EnnobledAtGTE)
	}
	if !isZero(f.EnnobledAtLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("ennobled_at", alias), f.EnnobledAtLT)
	}
	if !isZero(f.EnnobledAtLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("ennobled_at", alias), f.EnnobledAtLTE)
	}

	if !isZero(f.VillageID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("village_id", alias), pg.Array(f.VillageID))
	}
	if !isZero(f.VillageIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("village_id", alias), pg.Array(f.VillageIDNEQ))
	}

	if !isZero(f.NewOwnerID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_owner_id", alias), pg.Array(f.NewOwnerID))
	}
	if !isZero(f.NewOwnerIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("new_owner_id", alias), pg.Array(f.NewOwnerIDNEQ))
	}
	if !isZero(f.NewOwnerTribeID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_owner_tribe_id", alias), pg.Array(f.NewOwnerTribeID))
	}
	if !isZero(f.NewOwnerTribeIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("new_owner_tribe_id", alias), pg.Array(f.NewOwnerTribeIDNEQ))
	}

	if !isZero(f.OldOwnerID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_owner_id", alias), pg.Array(f.OldOwnerID))
	}
	if !isZero(f.OldOwnerIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("old_owner_id", alias), pg.Array(f.OldOwnerIDNEQ))
	}
	if !isZero(f.OldOwnerTribeID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_owner_tribe_id", alias), pg.Array(f.OldOwnerTribeID))
	}
	if !isZero(f.OldOwnerTribeIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("old_owner_tribe_id", alias), pg.Array(f.OldOwnerTribeIDNEQ))
	}

	if f.Or != nil {
		q = f.Or.WhereWithAlias(q, alias)
	}

	var err error
	if f.NewOwnerFilter != nil {
		q, err = f.NewOwnerFilter.WhereWithAlias(q.Relation("NewOwner._"), "new_owner", "NewOwner.Tribe._", "new_owner__tribe")
		if err != nil {
			return q, err
		}
	}
	if f.NewOwnerTribeFilter != nil {
		q, err = f.NewOwnerTribeFilter.WhereWithAlias(q.Relation("NewOwnerTribe._"), "new_owner_tribe")
		if err != nil {
			return q, err
		}
	}
	if f.OldOwnerFilter != nil {
		q, err = f.OldOwnerFilter.WhereWithAlias(q.Relation("OldOwner._"), "old_owner", "OldOwner.Tribe._", "old_owner__tribe")
		if err != nil {
			return q, err
		}
	}
	if f.OldOwnerTribeFilter != nil {
		q, err = f.OldOwnerTribeFilter.WhereWithAlias(q.Relation("OldOwnerTribe._"), "old_owner_tribe")
		if err != nil {
			return q, err
		}
	}
	if f.VillageFilter != nil {
		q, err = f.
			VillageFilter.
			WhereWithAlias(
				q.Relation("Village._"),
				"village",
				"Village.Player._",
				"village__player",
				"Village.Player.Tribe._",
				"village__player__tribe",
			)
		if err != nil {
			return q, err
		}
	}

	return q, nil
}

func (f *EnnoblementFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "ennoblement")
}
