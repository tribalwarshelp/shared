package models

import "time"

type PlayerHistory struct {
	tableName struct{} `pg:"?SERVER.player_history,alias:player_history"`

	OpponentsDefeated

	PlayerID      int       `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	Player        *Player   `json:"player" gqlgen:"-" xml:"player"`
	TotalVillages int       `json:"totalVillages" pg:",use_zero" gqlgen:"totalVillages" xml:"totalVillages"`
	Points        int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank          int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	TribeID       int       `json:"-" pg:",use_zero" gqlgen:"tribeID" xml:"tribeID"`
	Tribe         *Tribe    `json:"tribe,omitempty" gqlgen:"-" xml:"tribe"`
	CreateDate    time.Time `pg:"default:CURRENT_DATE,type:DATE,unique:group_1,use_zero" json:"createDate" gqlgen:"createDate" xml:"createDate"`
}

type PlayerHistoryFilter struct {
	tableName struct{} `urlstruct:"player_history"`

	PlayerID    []int `json:"playerID" gqlgen:"playerID" xml:"playerID"`
	PlayerIdNEQ []int `json:"playerIDNEQ" gqlgen:"playerIDNEQ" xml:"playerIDNEQ"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
