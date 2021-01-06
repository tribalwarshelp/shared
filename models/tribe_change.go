package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type TribeChange struct {
	tableName struct{} `pg:"?SERVER.tribe_changes,alias:tribe_change"`

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
	q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
		if !isZero(f.OldTribeID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("old_tribe_id", alias)), pg.Array(f.OldTribeID))
		}
		if !isZero(f.NewTribeID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("new_tribe_id", alias)), pg.Array(f.NewTribeID))
		}
		return q, nil
	})
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

	Or *TribeChangeFilterOr `urlstruct:",nowhere" json:"or" gqlgen:"or" xml:"or"`
}

func (f *TribeChangeFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if !isZero(f.PlayerID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("player_id", alias)), pg.Array(f.PlayerID))
	}
	if !isZero(f.PlayerIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("player_id", alias)), pg.Array(f.PlayerIDNEQ))
	}

	if !isZero(f.OldTribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("old_tribe_id", alias)), pg.Array(f.OldTribeID))
	}
	if !isZero(f.OldTribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("old_tribe_id", alias)), pg.Array(f.OldTribeIDNEQ))
	}

	if !isZero(f.NewTribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("new_tribe_id", alias)), pg.Array(f.NewTribeID))
	}
	if !isZero(f.NewTribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("new_tribe_id", alias)), pg.Array(f.NewTribeIDNEQ))
	}

	if !isZero(f.CreatedAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("created_at", alias)), f.CreatedAt)
	}
	if !isZero(f.CreatedAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("created_at", alias)), f.CreatedAtGT)
	}
	if !isZero(f.CreatedAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("created_at", alias)), f.CreatedAtGTE)
	}
	if !isZero(f.CreatedAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("created_at", alias)), f.CreatedAtLT)
	}
	if !isZero(f.CreatedAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("created_at", alias)), f.CreatedAtLTE)
	}

	if f.Or != nil {
		q = f.Or.WhereWithAlias(q, alias)
	}

	return q, nil
}

func (f *TribeChangeFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "tribe_change")
}

type TribeChangeRelationshipAndSortAppender struct {
	Filter *TribeChangeFilter
	Sort   []string
}

func (a *TribeChangeRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
	var err error
	playerRequired := findStringWithPrefix(a.Sort, "player.") != ""
	if a.Filter.PlayerFilter != nil {
		q, err = a.Filter.PlayerFilter.WhereWithAlias(q, "player")
		if err != nil {
			return q, err
		}
		playerRequired = true
	}

	oldTribeRequired := findStringWithPrefix(a.Sort, "old_tribe.") != ""
	if a.Filter.OldTribeFilter != nil {
		q, err = a.Filter.OldTribeFilter.WhereWithAlias(q, "old_tribe")
		if err != nil {
			return q, err
		}
		oldTribeRequired = true
	}

	newTribeRequired := findStringWithPrefix(a.Sort, "new_tribe.") != ""
	if a.Filter.NewTribeFilter != nil {
		q, err = a.Filter.NewTribeFilter.WhereWithAlias(q, "new_tribe")
		if err != nil {
			return q, err
		}
		newTribeRequired = true
	}

	if !isZero(a.Sort) {
		q = q.Order(a.Sort...)
	}

	if playerRequired {
		q = q.Relation("Tribe._")
	}
	if oldTribeRequired {
		q = q.Relation("OldTribe._")
	}
	if newTribeRequired {
		q = q.Relation("NewTribe._")
	}

	return q, nil
}
