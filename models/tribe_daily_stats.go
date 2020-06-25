package models

import "time"

type TribeDailyStats struct {
	tableName struct{} `pg:"?SERVER.tribes_daily_stats,alias:tribe_daily_stats"`

	TribeID   int       `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	Tribe     *Tribe    `json:"tribe" gqlgen:"-" xml:"tribe"`
	Members   int       `json:"members" gqlgen:"members" pg:",use_zero" xml:"members"`
	Villages  int       `json:"villages" gqlgen:"villages" pg:",use_zero" xml:"villages"`
	Points    int       `json:"points" gqlgen:"points" pg:",use_zero" xml:"points"`
	AllPoints int       `json:"allPoints" gqlgen:"allPoints" pg:",use_zero" xml:"allPoints"`
	Rank      int       `json:"rank" gqlgen:"rank" pg:",use_zero" xml:"rank"`
	Dominance float64   `json:"dominance" gqlgen:"dominance" pg:",use_zero" xml:"dominance"`
	CreatedAt time.Time `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`

	OpponentsDefeated
}

type TribesDailyStatsFilter struct {
	tableName struct{} `urlstruct:"tribe_daily_stats"`

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
