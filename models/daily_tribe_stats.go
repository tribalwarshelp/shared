package models

import (
	"time"
)

type DailyTribeStats struct {
	tableName struct{} `pg:"?SERVER.daily_tribe_stats,alias:daily_tribe_stats"`

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
	TribeIdNEQ  []int        `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`
	TribeFilter *TribeFilter `urlstruct:",nowhere" json:"tribeFilter" gqlgen:"tribeFilter" xml:"tribeFilter"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}
