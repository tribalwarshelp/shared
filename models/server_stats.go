package models

import (
	"context"
	"time"
)

type ServerStats struct {
	tableName struct{} `pg:"?SERVER.stats,alias:stats"`

	ActivePlayers     int       `pg:",use_zero" json:"activePlayers" gqlgen:"activePlayers" xml:"activePlayers"`
	InactivePlayers   int       `pg:",use_zero" json:"inactivePlayers" gqlgen:"inactivePlayers" xml:"inactivePlayers"`
	Players           int       `pg:",use_zero" json:"players" gqlgen:"players" xml:"players"`
	ActiveTribes      int       `pg:",use_zero" json:"activeTribes" gqlgen:"activeTribes" xml:"activeTribes"`
	InactiveTribes    int       `pg:",use_zero" json:"inactiveTribes" gqlgen:"inactiveTribes" xml:"inactiveTribes"`
	Tribes            int       `pg:",use_zero" json:"tribes" gqlgen:"tribes" xml:"tribes"`
	Villages          int       `pg:",use_zero" json:"villages" gqlgen:"villages" xml:"villages"`
	BonusVillages     int       `pg:",use_zero" json:"bonusVillages" gqlgen:"bonusVillages" xml:"bonusVillages"`
	BarbarianVillages int       `pg:",use_zero" json:"barbarianVillages" gqlgen:"barbarianVillages" xml:"barbarianVillages"`
	PlayerVillages    int       `pg:",use_zero" json:"playerVillages" gqlgen:"playerVillages" xml:"playerVillages"`
	CreateDate        time.Time `pg:"default:now(),type:DATE,unique:group_1,use_zero" json:"createDate" gqlgen:"createDate" xml:"createDate"`
}

func (s *ServerStats) BeforeInsert(ctx context.Context) (context.Context, error) {
	if s.CreateDate.IsZero() {
		s.CreateDate = time.Now()
	}

	return ctx, nil
}

type ServerStatsFilter struct {
	tableName struct{} `urlstruct:"stats"`

	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
