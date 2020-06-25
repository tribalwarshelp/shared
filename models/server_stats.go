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
	CreatedAt         time.Time `pg:"default:now(),type:DATE,unique:group_1,use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

func (s *ServerStats) BeforeInsert(ctx context.Context) (context.Context, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return ctx, nil
}

type ServerStatsFilter struct {
	tableName struct{} `urlstruct:"stats"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
