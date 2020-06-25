package models

import "time"

type PlayersDailyStats struct {
	tableName struct{} `pg:"?SERVER.players_daily_stats,alias:player_daily_stats"`

	PlayerID  int       `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player    *Player   `json:"player" gqlgen:"-" xml:"player"`
	Villages  int       `json:"villages" pg:",use_zero" gqlgen:"villages" xml:"villages"`
	Points    int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank      int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	CreatedAt time.Time `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`

	OpponentsDefeated
}

type PlayersDailyStatsFilter struct {
	tableName struct{} `urlstruct:"player_daily_stats"`

	PlayerID    []int `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	PlayerIdNEQ []int `json:"playerIDNEQ" gqlgen:"playerIDNEQ" xml:"playerIDNEQ"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
