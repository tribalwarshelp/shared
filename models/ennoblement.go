package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Ennoblement struct {
	tableName struct{} `pg:"?SERVER.ennoblements,alias:ennoblement"`

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
	q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
		if !isZero(f.NewOwnerID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("new_owner_id", alias)), pg.Array(f.NewOwnerID))
		}
		if !isZero(f.NewOwnerTribeID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("new_owner_tribe_id", alias)), pg.Array(f.NewOwnerTribeID))
		}
		if !isZero(f.OldOwnerID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("old_owner_id", alias)), pg.Array(f.OldOwnerID))
		}
		if !isZero(f.OldOwnerTribeID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("old_owner_tribe_id", alias)), pg.Array(f.OldOwnerTribeID))
		}
		return q, nil
	})
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
	if !isZero(f.EnnobledAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("ennobled_at", alias)), f.EnnobledAt)
	}
	if !isZero(f.EnnobledAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("ennobled_at", alias)), f.EnnobledAtGT)
	}
	if !isZero(f.EnnobledAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("ennobled_at", alias)), f.EnnobledAtGTE)
	}
	if !isZero(f.EnnobledAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("ennobled_at", alias)), f.EnnobledAtLT)
	}
	if !isZero(f.EnnobledAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("ennobled_at", alias)), f.EnnobledAtLTE)
	}

	if !isZero(f.VillageID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("village_id", alias)), pg.Array(f.VillageID))
	}
	if !isZero(f.VillageIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("village_id", alias)), pg.Array(f.VillageIDNEQ))
	}

	if !isZero(f.NewOwnerID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("new_owner_id", alias)), pg.Array(f.NewOwnerID))
	}
	if !isZero(f.NewOwnerIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("new_owner_id", alias)), pg.Array(f.NewOwnerIDNEQ))
	}
	if !isZero(f.NewOwnerTribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("new_owner_tribe_id", alias)), pg.Array(f.NewOwnerTribeID))
	}
	if !isZero(f.NewOwnerTribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("new_owner_tribe_id", alias)), pg.Array(f.NewOwnerTribeIDNEQ))
	}

	if !isZero(f.OldOwnerID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("old_owner_id", alias)), pg.Array(f.OldOwnerID))
	}
	if !isZero(f.OldOwnerIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("old_owner_id", alias)), pg.Array(f.OldOwnerIDNEQ))
	}
	if !isZero(f.OldOwnerTribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("old_owner_tribe_id", alias)), pg.Array(f.OldOwnerTribeID))
	}
	if !isZero(f.OldOwnerTribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("old_owner_tribe_id", alias)), pg.Array(f.OldOwnerTribeIDNEQ))
	}

	if f.Or != nil {
		q = f.Or.WhereWithAlias(q, alias)
	}

	return q, nil
}

func (f *EnnoblementFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "ennoblement")
}

type EnnoblementRelationshipAndSortAppender struct {
	Filter *EnnoblementFilter
	Sort   []string
}

func (a *EnnoblementRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
	var err error
	villageRequired := findStringWithPrefix(a.Sort, "village.") != ""
	if a.Filter.VillageFilter != nil {
		q, err = a.Filter.VillageFilter.WhereWithAlias(q, "village")
		if err != nil {
			return q, err
		}
		villageRequired = true
	}

	oldOwnerRequired := findStringWithPrefix(a.Sort, "old_owner.") != ""
	if a.Filter.OldOwnerFilter != nil {
		q, err = a.Filter.OldOwnerFilter.WhereWithAlias(q, "old_owner")
		if err != nil {
			return q, err
		}
		oldOwnerRequired = true
	}
	oldOwnerTribeRequired := findStringWithPrefix(a.Sort, "old_owner_tribe.") != ""
	if a.Filter.OldOwnerTribeFilter != nil {
		q, err = a.Filter.OldOwnerTribeFilter.WhereWithAlias(q, "old_owner_tribe")
		if err != nil {
			return q, err
		}
		oldOwnerTribeRequired = true
	}

	newOwnerRequired := findStringWithPrefix(a.Sort, "new_owner.") != ""
	if a.Filter.NewOwnerFilter != nil {
		q, err = a.Filter.NewOwnerFilter.WhereWithAlias(q, "new_owner")
		if err != nil {
			return q, err
		}
		newOwnerRequired = true
	}
	newOwnerTribeRequired := findStringWithPrefix(a.Sort, "new_owner_tribe.") != ""
	if a.Filter.NewOwnerTribeFilter != nil {
		q, err = a.Filter.NewOwnerTribeFilter.WhereWithAlias(q, "new_owner_tribe")
		if err != nil {
			return q, err
		}
		newOwnerTribeRequired = true
	}

	if !isZero(a.Sort) {
		q = q.Order(a.Sort...)
	}

	if villageRequired {
		q = q.Relation("Village._")
	}
	if oldOwnerRequired {
		q = q.Relation("OldOwner._")
	}
	if oldOwnerTribeRequired {
		q = q.Relation("OldOwnerTribe._")
	}
	if newOwnerRequired {
		q = q.Relation("NewOwner._")
	}
	if newOwnerTribeRequired {
		q = q.Relation("NewOwnerTribe._")
	}

	return q, nil
}
