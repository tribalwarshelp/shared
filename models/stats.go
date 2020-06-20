package models

import (
	"context"
	"time"
)

type Stats struct {
	tableName struct{} `pg:"stats,alias:stats" json:"tableName" gqlgen:"tableName" xml:"tableName"`

	ActiveServers   int `json:"activeServers" gqlgen:"activeServers" xml:"activeServers"`
	InactiveServers int `json:"inactiveServers" gqlgen:"inactiveServers" xml:"inactiveServers"`
	Servers         int `json:"servers" gqlgen:"servers" xml:"servers"`
	ServerStats
}

func (s *Stats) BeforeInsert(ctx context.Context) (context.Context, error) {
	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	return ctx, nil
}

type StatsFilter struct {
	tableName struct{} `urlstruct:"stats"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`
}
