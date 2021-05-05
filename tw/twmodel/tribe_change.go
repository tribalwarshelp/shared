package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type TribeChange struct {
	tableName struct{} `pg:"?SERVER.tribe_changes,alias:tribe_change"`

	ID         int       `json:"id" gqlgen:"id" xml:"id"`
	PlayerID   int       `pg:",use_zero" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player     *Player   `json:"player,omitempty" gqlgen:"-" xml:"player" pg:"rel:has-one"`
	OldTribeID int       `pg:",use_zero" json:"oldTribeID" gqlgen:"oldTribeID" xml:"oldTribeID"`
	OldTribe   *Tribe    `json:"oldTribe,omitempty" gqlgen:"-" xml:"oldTribe" pg:"rel:has-one"`
	NewTribeID int       `pg:",use_zero" json:"newTribeID" gqlgen:"newTribeID" xml:"newTribeID"`
	NewTribe   *Tribe    `json:"newTribe,omitempty" gqlgen:"-" xml:"newTribe" pg:"rel:has-one"`
	CreatedAt  time.Time `pg:"default:now(),use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

type TribeChangeFilterOr struct {
	OldTribeID []int `json:"oldTribeID" gqlgen:"oldTribeID" xml:"oldTribeID"`
	NewTribeID []int `json:"newTribeID" gqlgen:"newTribeID" xml:"newTribeID"`
}

func (f *TribeChangeFilterOr) WhereWithAlias(q *orm.Query, alias string) *orm.Query {
	if f != nil {
		q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			if !isZero(f.OldTribeID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_tribe_id", alias), pg.Array(f.OldTribeID))
			}
			if !isZero(f.NewTribeID) {
				q = q.WhereOr(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_tribe_id", alias), pg.Array(f.NewTribeID))
			}
			return q, nil
		})
	}
	return q
}

type TribeChangeFilter struct {
	PlayerID     []int         `json:"id" gqlgen:"id" xml:"id"`
	PlayerIDNEQ  []int         `json:"idNEQ" gqlgen:"idNEQ" xml:"idNEQ"`
	PlayerFilter *PlayerFilter `json:"playerFilter" xml:"playerFilter" gqlgen:"playerFilter"`

	OldTribeID     []int        `json:"oldTribeID" gqlgen:"oldTribeID" xml:"oldTribeID"`
	OldTribeIDNEQ  []int        `json:"oldTribeIDNEQ" gqlgen:"oldTribeIDNEQ" xml:"oldTribeIDNEQ"`
	OldTribeFilter *TribeFilter `json:"oldTribeFilter" xml:"oldTribeFilter" gqlgen:"oldTribeFilter"`

	NewTribeID     []int        `json:"newTribeID" gqlgen:"newTribeID" xml:"newTribeID"`
	NewTribeIDNEQ  []int        `json:"newTribeIDNEQ" gqlgen:"newTribeIDNEQ" xml:"newTribeIDNEQ"`
	NewTribeFilter *TribeFilter `json:"newTribeFilter" xml:"newTribeFilter" gqlgen:"newTribeFilter"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Or *TribeChangeFilterOr `json:"or" gqlgen:"or" xml:"or"`
}

func (f *TribeChangeFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	if !isZero(f.PlayerID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerID))
	}
	if !isZero(f.PlayerIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerIDNEQ))
	}

	if !isZero(f.OldTribeID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("old_tribe_id", alias), pg.Array(f.OldTribeID))
	}
	if !isZero(f.OldTribeIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("old_tribe_id", alias), pg.Array(f.OldTribeIDNEQ))
	}

	if !isZero(f.NewTribeID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("new_tribe_id", alias), pg.Array(f.NewTribeID))
	}
	if !isZero(f.NewTribeIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("new_tribe_id", alias), pg.Array(f.NewTribeIDNEQ))
	}

	if !isZero(f.CreatedAt) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("created_at", alias), f.CreatedAt)
	}
	if !isZero(f.CreatedAtGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("created_at", alias), f.CreatedAtGT)
	}
	if !isZero(f.CreatedAtGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("created_at", alias), f.CreatedAtGTE)
	}
	if !isZero(f.CreatedAtLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("created_at", alias), f.CreatedAtLT)
	}
	if !isZero(f.CreatedAtLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("created_at", alias), f.CreatedAtLTE)
	}

	if f.Or != nil {
		q = f.Or.WhereWithAlias(q, alias)
	}

	return q, nil
}

func (f *TribeChangeFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "tribe_change")
}

func (f *TribeChangeFilter) WhereWithRelations(q *orm.Query) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	filtersToAppend := []filterToAppend{
		{
			filter: f,
			alias:  "tribe_change",
		},
	}
	if f.PlayerFilter != nil {
		filtersToAppend = append(filtersToAppend, filterToAppend{
			filter:       f.PlayerFilter,
			relationName: "Player",
		})
	}
	if f.OldTribeFilter != nil {
		filtersToAppend = append(filtersToAppend, filterToAppend{
			filter:       f.OldTribeFilter,
			relationName: "OldTribe",
		})
	}
	if f.NewTribeFilter != nil {
		filtersToAppend = append(filtersToAppend, filterToAppend{
			filter:       f.NewTribeFilter,
			relationName: "NewTribe",
		})
	}

	return appendFilters(q, filtersToAppend...)
}
