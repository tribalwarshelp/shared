package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type TribeHistory struct {
	tableName struct{} `pg:"?SERVER.tribe_history,alias:tribe_history"`

	OpponentsDefeated

	ID            int       `json:"id" gqlgen:"id" xml:"id"`
	TribeID       int       `pg:",unique:group_1" json:"-" pg:",use_zero" gqlgen:"tribeID" xml:"tribeID"`
	Tribe         *Tribe    `json:"tribe,omitempty" gqlgen:"-" xml:"tribe" pg:"rel:has-one"`
	TotalMembers  int       `json:"totalMembers" gqlgen:"totalMembers" pg:",use_zero"`
	TotalVillages int       `json:"totalVillages" gqlgen:"totalVillages" pg:",use_zero"`
	Points        int       `json:"points" gqlgen:"points" pg:",use_zero"`
	AllPoints     int       `json:"allPoints" gqlgen:"allPoints" pg:",use_zero"`
	Rank          int       `json:"rank" gqlgen:"rank" pg:",use_zero"`
	Dominance     float64   `json:"dominance" gqlgen:"dominance" pg:",use_zero"`
	CreateDate    time.Time `pg:"default:now(),type:DATE,unique:group_1,use_zero" json:"createDate" gqlgen:"createDate" xml:"createDate"`
}

type TribeHistoryFilter struct {
	TribeID     []int        `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeIDNEQ  []int        `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`
	TribeFilter *TribeFilter `json:"tribeFilter" xml:"tribeFilter" gqlgen:"tribeFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}

func (f *TribeHistoryFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if !isZero(f.CreateDate) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("create_date", alias)), f.CreateDate)
	}
	if !isZero(f.CreateDateGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("create_date", alias)), f.CreateDateGT)
	}
	if !isZero(f.CreateDateGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("create_date", alias)), f.CreateDateGTE)
	}
	if !isZero(f.CreateDateLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("create_date", alias)), f.CreateDateLT)
	}
	if !isZero(f.CreateDateLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("create_date", alias)), f.CreateDateLTE)
	}

	if !isZero(f.TribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("tribe_id", alias)), pg.Array(f.TribeID))
	}
	if !isZero(f.TribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("tribe_id", alias)), pg.Array(f.TribeIDNEQ))
	}

	return q, nil
}

func (f *TribeHistoryFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "tribe_history")
}

type TribeHistoryRelationshipAndSortAppender struct {
	Filter *TribeHistoryFilter
	Sort   []string
}

func (a *TribeHistoryRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
	var err error
	tribeRequired := findStringWithPrefix(a.Sort, "tribe.") != ""
	if a.Filter.TribeFilter != nil {
		q, err = a.Filter.TribeFilter.WhereWithAlias(q, "tribe")
		if err != nil {
			return q, err
		}
		tribeRequired = true
	}

	if !isZero(a.Sort) {
		q = q.Order(a.Sort...)
	}

	if tribeRequired {
		q = q.Relation("Tribe._")
	}

	return q, nil
}
