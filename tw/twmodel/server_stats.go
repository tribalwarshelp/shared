package twmodel

import (
	"context"
	"github.com/Kichiyaki/gopgutil/v10"
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type ServerStats struct {
	tableName struct{} `pg:"?SERVER.stats,alias:stats"`

	ID                int       `json:"id" gqlgen:"id" xml:"id"`
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
	CreateDate    time.Time `json:"createDate" gqlgen:"createDate" xml:"createDate"`
	CreateDateGT  time.Time `json:"createDateGT" gqlgen:"createDateGT" xml:"createDateGT"`
	CreateDateGTE time.Time `json:"createDateGTE" gqlgen:"createDateGTE" xml:"createDateGTE"`
	CreateDateLT  time.Time `json:"createDateLT" gqlgen:"createDateLT" xml:"createDateLT"`
	CreateDateLTE time.Time `json:"createDateLTE" gqlgen:"createDateLTE" xml:"createDateLTE"`
}

func (f *ServerStatsFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	if !isZero(f.CreateDate) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDate)
	}
	if !isZero(f.CreateDateGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateGT)
	}
	if !isZero(f.CreateDateGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateGTE)
	}
	if !isZero(f.CreateDateLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateLT)
	}
	if !isZero(f.CreateDateLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("create_date", alias), f.CreateDateLTE)
	}

	return q, nil
}

func (f *ServerStatsFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "stats")
}
