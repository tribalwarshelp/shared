package models

import "time"

type TribeHistory struct {
	tableName struct{} `pg:"?SERVER.tribe_history,alias:tribe_history"`

	OpponentsDefeated

	TribeID       int       `pg:",unique:group_1" json:"-" pg:",use_zero" gqlgen:"tribeID" xml:"tribeID"`
	Tribe         *Tribe    `json:"tribe,omitempty" gqlgen:"-" xml:"tribe"`
	TotalMembers  int       `json:"totalMembers" gqlgen:"totalMembers" pg:",use_zero"`
	TotalVillages int       `json:"totalVillages" gqlgen:"totalVillages" pg:",use_zero"`
	Points        int       `json:"points" gqlgen:"points" pg:",use_zero"`
	AllPoints     int       `json:"allPoints" gqlgen:"allPoints" pg:",use_zero"`
	Rank          int       `json:"rank" gqlgen:"rank" pg:",use_zero"`
	Dominance     float64   `json:"dominance" gqlgen:"dominance" pg:",use_zero"`
	CreatedAt     time.Time `pg:"default:now(),type:DATE,unique:group_1,use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

type TribeHistoryFilter struct {
	tableName struct{} `urlstruct:"tribe_history"`

	TribeID    []int `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeIdNEQ []int `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
