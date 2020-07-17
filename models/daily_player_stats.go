package models

import "time"

type DailyPlayerStats struct {
	tableName struct{} `pg:"?SERVER.daily_player_stats,alias:daily_player_stats"`

	PlayerID   int       `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player     *Player   `json:"player" gqlgen:"-" xml:"player"`
	Villages   int       `json:"villages" pg:",use_zero" gqlgen:"villages" xml:"villages"`
	Points     int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank       int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	CreateDate time.Time `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"createDate" gqlgen:"createDate" xml:"createDate"`

	OpponentsDefeated
}

type DailyPlayerStatsFilter struct {
	tableName struct{} `urlstruct:"daily_player_stats"`

	PlayerID     []int         `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	PlayerIdNEQ  []int         `json:"playerIDNEQ" gqlgen:"playerIDNEQ" xml:"playerIDNEQ"`
	PlayerFilter *PlayerFilter `urlstruct:",nowhere" json:"playerFilter" gqlgen:"playerFilter" xml:"playerFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
