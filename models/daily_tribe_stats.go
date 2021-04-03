package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DailyTribeStats struct {
	tableName struct{} `pg:"?SERVER.daily_tribe_stats,alias:daily_tribe_stats"`

	ID         int       `json:"id" gqlgen:"id" xml:"id"`
	TribeID    int       `pg:",unique:group_1" json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	Tribe      *Tribe    `json:"tribe" gqlgen:"-" xml:"tribe" pg:"rel:has-one"`
	Members    int       `json:"members" gqlgen:"members" pg:",use_zero" xml:"members"`
	Villages   int       `json:"villages" gqlgen:"villages" pg:",use_zero" xml:"villages"`
	Points     int       `json:"points" gqlgen:"points" pg:",use_zero" xml:"points"`
	AllPoints  int       `json:"allPoints" gqlgen:"allPoints" pg:",use_zero" xml:"allPoints"`
	Rank       int       `json:"rank" gqlgen:"rank" pg:",use_zero" xml:"rank"`
	Dominance  float64   `json:"dominance" gqlgen:"dominance" pg:",use_zero" xml:"dominance"`
	CreateDate time.Time `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"createDate" gqlgen:"createDate" xml:"createDate"`

	OpponentsDefeated
}

type DailyTribeStatsFilter struct {
	TribeID     []int        `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeIDNEQ  []int        `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`
	TribeFilter *TribeFilter `json:"tribeFilter" gqlgen:"tribeFilter" xml:"tribeFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}

func (f *DailyTribeStatsFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
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

func (f *DailyTribeStatsFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "daily_tribe_stats")
}

type DailyTribeStatsRelationshipAndSortAppender struct {
	Filter *DailyTribeStatsFilter
	Sort   []string
}

func (a *DailyTribeStatsRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
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
