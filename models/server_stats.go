package models

import (
	"context"
	"time"
)

type ServerStats struct {
	tableName struct{} `pg:"?SERVER.stats,alias:stats"`

	ActivePlayers     int       `json:"activePlayers" gqlgen:"activePlayers" xml:"activePlayers"`
	InactivePlayers   int       `json:"inactivePlayers" gqlgen:"inactivePlayers" xml:"inactivePlayers"`
	Players           int       `json:"players" gqlgen:"players" xml:"players"`
	ActiveTribes      int       `json:"activeTribes" gqlgen:"activeTribes" xml:"activeTribes"`
	InactiveTribes    int       `json:"inactiveTribes" gqlgen:"inactiveTribes" xml:"inactiveTribes"`
	Tribes            int       `json:"tribes" gqlgen:"tribes" xml:"tribes"`
	Villages          int       `json:"villages" gqlgen:"villages" xml:"villages"`
	BonusVillages     int       `json:"bonusVillages" gqlgen:"bonusVillages" xml:"bonusVillages"`
	BarbarianVillages int       `json:"barbarianVillages" gqlgen:"barbarianVillages" xml:"barbarianVillages"`
	PlayerVillages    int       `json:"playerVillages" gqlgen:"playerVillages" xml:"playerVillages"`
	CreatedAt         time.Time `pg:"default:now(),use_zero" json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
}

func (s *ServerStats) BeforeInsert(ctx context.Context) (context.Context, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return ctx, nil
}
