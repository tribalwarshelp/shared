package twmodel

import (
	"github.com/Kichiyaki/gopgutil/v10"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type DailyPlayerStats struct {
	tableName struct{} `pg:"?SERVER.daily_player_stats,alias:daily_player_stats"`

	ID         int       `json:"id" gqlgen:"id" xml:"id"`
	PlayerID   int       `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player     *Player   `json:"player" gqlgen:"-" xml:"player" pg:"rel:has-one"`
	Villages   int       `json:"villages" pg:",use_zero" gqlgen:"villages" xml:"villages"`
	Points     int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank       int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	CreateDate time.Time `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"createDate" gqlgen:"createDate" xml:"createDate"`

	OpponentsDefeated
}

type DailyPlayerStatsFilter struct {
	PlayerID     []int         `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	PlayerIDNEQ  []int         `json:"playerIDNEQ" gqlgen:"playerIDNEQ" xml:"playerIDNEQ"`
	PlayerFilter *PlayerFilter `json:"playerFilter" gqlgen:"playerFilter" xml:"playerFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}

func (f *DailyPlayerStatsFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
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

	if !isZero(f.PlayerID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerID))
	}
	if !isZero(f.PlayerIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerIDNEQ))
	}

	return q, nil
}

func (f *DailyPlayerStatsFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "daily_player_stats")
}

func (f *DailyPlayerStatsFilter) WhereWithRelations(q *orm.Query) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	filtersToAppend := []filterToAppend{
		{
			filter: f,
			alias:  "daily_player_stats",
		},
	}
	if f.PlayerFilter != nil {
		filtersToAppend = append(filtersToAppend, filterToAppend{
			filter:       f.PlayerFilter,
			relationName: "Player",
		})
	}

	return appendFilters(q, filtersToAppend...)
}
