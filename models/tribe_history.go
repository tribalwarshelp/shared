package models

import "time"

type TribeHistory struct {
	tableName struct{} `pg:"?SERVER.tribe_history,alias:tribe_history"`

	OpponentsDefeated

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
	tableName struct{} `urlstruct:"tribe_history"`

	TribeID    []int `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeIdNEQ []int `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
