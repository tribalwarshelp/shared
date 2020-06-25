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
	CreatedAt     time.Time `pg:"default:now(),type:DATE,unique:group_1,use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

type PlayerHistoryFilter struct {
	tableName struct{} `urlstruct:"player_history"`

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
