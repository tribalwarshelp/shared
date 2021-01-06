package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PlayerHistory struct {
	tableName struct{} `pg:"?SERVER.player_history,alias:player_history"`

	OpponentsDefeated

	PlayerID      int       `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player        *Player   `json:"player" gqlgen:"-" xml:"player" pg:"rel:has-one"`
	TotalVillages int       `json:"totalVillages" pg:",use_zero" gqlgen:"totalVillages" xml:"totalVillages"`
	Points        int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank          int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	TribeID       int       `json:"-" pg:",use_zero" gqlgen:"tribeID" xml:"tribeID"`
	Tribe         *Tribe    `json:"tribe,omitempty" gqlgen:"-" xml:"tribe" pg:"rel:has-one"`
	CreateDate    time.Time `pg:"default:CURRENT_DATE,type:DATE,unique:group_1,use_zero" json:"createDate" gqlgen:"createDate" xml:"createDate"`
}

type PlayerHistoryFilter struct {
	PlayerID     []int         `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	PlayerIDNEQ  []int         `json:"playerIDNEQ" gqlgen:"playerIDNEQ" xml:"playerIDNEQ"`
	PlayerFilter *PlayerFilter `json:"playerFilter" xml:"playerFilter" gqlgen:"playerFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}

func (f *PlayerHistoryFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
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

	if !isZero(f.PlayerID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("player_id", alias)), pg.Array(f.PlayerID))
	}
	if !isZero(f.PlayerIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("player_id", alias)), pg.Array(f.PlayerIDNEQ))
	}

	return q, nil
}

func (f *PlayerHistoryFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "player_history")
}

type PlayerHistoryRelationshipAndSortAppender struct {
	Filter *PlayerHistoryFilter
	Sort   []string
}

func (a *PlayerHistoryRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
	var err error
	playerRequired := findStringWithPrefix(a.Sort, "player.") != ""
	playerTribeRequired := findStringWithPrefix(a.Sort, "player.tribe.") != ""
	if a.Filter.PlayerFilter != nil {
		q, err = a.Filter.PlayerFilter.WhereWithAlias(q, "player")
		if err != nil {
			return q, err
		}
		playerRequired = true
		if a.Filter.PlayerFilter.TribeFilter != nil {
			q, err = a.Filter.PlayerFilter.WhereWithAlias(q, "tribe")
			if err != nil {
				return q, err
			}
			playerTribeRequired = true
		}
	}

	if !isZero(a.Sort) {
		q = q.Order(a.Sort...)
	}

	if playerRequired {
		q = q.Relation("Player._")
	}

	if playerTribeRequired {
		q = q.Join("LEFT JOIN ?SERVER.tribes AS tribe ON tribe.id = player.tribe_id")
	}

	return q, nil
}
