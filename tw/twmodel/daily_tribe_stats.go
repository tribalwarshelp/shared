package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
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
	if f == nil {
		return q, nil
	}

	if !isZero(f.CreateDate) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDate)
	}
	if !isZero(f.CreateDateGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateGT)
	}
	if !isZero(f.CreateDateGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateGTE)
	}
	if !isZero(f.CreateDateLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateLT)
	}
	if !isZero(f.CreateDateLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateLTE)
	}

	if !isZero(f.TribeID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("tribe_id", alias), pg.Array(f.TribeID))
	}
	if !isZero(f.TribeIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("tribe_id", alias), pg.Array(f.TribeIDNEQ))
	}

	return q, nil
}

func (f *DailyTribeStatsFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "daily_tribe_stats")
}

func (f *DailyTribeStatsFilter) WhereWithRelations(q *orm.Query) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	filtersToAppend := []filterToAppend{
		{
			filter: f,
			alias:  "daily_tribe_stats",
		},
	}
	if f.TribeFilter != nil {
		filtersToAppend = append(filtersToAppend, filterToAppend{
			filter:       f.TribeFilter,
			relationName: "Tribe",
		})
	}

	return appendFilters(q, filtersToAppend...)
}
