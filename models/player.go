package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Player struct {
	tableName struct{} `pg:"?SERVER.players,alias:player" json:"tableName" xml:"tableName" gqlgen:"tableName"`

	ID             int       `json:"id" pg:"type:bigint,pk" gqlgen:"id" xml:"id"`
	Name           string    `json:"name" gqlgen:"name" xml:"name"`
	Exists         *bool     `json:"exists" pg:",use_zero" gqlgen:"exists" xml:"exists"`
	TotalVillages  int       `json:"totalVillages" pg:",use_zero" gqlgen:"totalVillages" xml:"totalVillages"`
	Points         int       `json:"points" pg:",use_zero" gqlgen:"points" xml:"points"`
	Rank           int       `json:"rank" pg:",use_zero" gqlgen:"rank" xml:"rank"`
	TribeID        int       `json:"-" pg:",use_zero" gqlgen:"tribeID" xml:"tribeID"`
	Tribe          *Tribe    `json:"tribe,omitempty" gqlgen:"-" pg:"rel:has-one" xml:"tribe"`
	DailyGrowth    int       `json:"dailyGrowth" pg:",use_zero" gqlgen:"dailyGrowth" xml:"dailyGrowth"`
	BestRank       int       `json:"bestRank" pg:",use_zero" gqlgen:"bestRank" xml:"bestRank"`
	BestRankAt     time.Time `json:"bestRankAt" pg:"default:now(),use_zero" gqlgen:"bestRankAt" xml:"bestRankAt"`
	MostPoints     int       `json:"mostPoints" pg:",use_zero" gqlgen:"mostPoints" xml:"mostPoints"`
	MostPointsAt   time.Time `json:"mostPointsAt" pg:"default:now(),use_zero" gqlgen:"mostPointsAt" xml:"mostPointsAt"`
	MostVillages   int       `json:"mostVillages" pg:",use_zero" gqlgen:"mostVillages" xml:"mostVillages"`
	MostVillagesAt time.Time `json:"mostVillagesAt" pg:"default:now(),use_zero" gqlgen:"mostVillagesAt" xml:"mostVillagesAt"`
	JoinedAt       time.Time `json:"joinedAt" pg:"default:now(),use_zero" gqlgen:"joinedAt" xml:"joinedAt"`
	LastActivityAt time.Time `pg:"default:now(),use_zero" json:"lastActivityAt" xml:"lastActivityAt" gqlgen:"lastActivityAt"`
	DeletedAt      time.Time `json:"deletedAt" gqlgen:"deletedAt" xml:"deletedAt"`

	OpponentsDefeated
}

type PlayerFilter struct {
	ID    []int `json:"id" gqlgen:"id" xml:"id"`
	IDNEQ []int `json:"idNEQ" gqlgen:"idNEQ" xml:"idNEQ"`

	Exists *bool ` json:"exists" gqlgen:"exists" xml:"exists"`

	Name      []string `json:"name" gqlgen:"name" xml:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ" xml:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH" xml:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ" xml:"nameIEQ"`

	TotalVillages    int `json:"totalVillages" gqlgen:"totalVillages" xml:"totalVillages"`
	TotalVillagesGT  int `json:"totalVillagesGT" gqlgen:"totalVillagesGT" xml:"totalVillagesGT"`
	TotalVillagesGTE int `json:"totalVillagesGTE" gqlgen:"totalVillagesGTE" xml:"totalVillagesGTE"`
	TotalVillagesLT  int `json:"totalVillagesLT" gqlgen:"totalVillagesLT" xml:"totalVillagesLT"`
	TotalVillagesLTE int `json:"totalVillagesLTE" gqlgen:"totalVillagesLTE" xml:"totalVillagesLTE"`

	Points    int `json:"points" gqlgen:"points" xml:"points"`
	PointsGT  int `json:"pointsGT" gqlgen:"pointsGT" xml:"pointsGT"`
	PointsGTE int `json:"pointsGTE" gqlgen:"pointsGTE" xml:"pointsGTE"`
	PointsLT  int `json:"pointsLT" gqlgen:"pointsLT" xml:"pointsLT"`
	PointsLTE int `json:"pointsLTE" gqlgen:"pointsLTE" xml:"pointsLTE"`

	Rank    int `json:"rank" gqlgen:"rank" xml:"rank"`
	RankGT  int `json:"rankGT" gqlgen:"rankGT" xml:"rankGT"`
	RankGTE int `json:"rankGTE" gqlgen:"rankGTE" xml:"rankGTE"`
	RankLT  int `json:"rankLT" gqlgen:"rankLT" xml:"rankLT"`
	RankLTE int `json:"rankLTE" gqlgen:"rankLTE" xml:"rankLTE"`

	DailyGrowth    int `json:"dailyGrowth" gqlgen:"dailyGrowth" xml:"dailyGrowth"`
	DailyGrowthGT  int `json:"dailyGrowthGT" gqlgen:"dailyGrowthGT" xml:"dailyGrowthGT"`
	DailyGrowthGTE int `json:"dailyGrowthGTE" gqlgen:"dailyGrowthGTE" xml:"dailyGrowthGTE"`
	DailyGrowthLT  int `json:"dailyGrowthLT" gqlgen:"dailyGrowthLT" xml:"dailyGrowthLT"`
	DailyGrowthLTE int `json:"dailyGrowthLTE" gqlgen:"dailyGrowthLTE" xml:"dailyGrowthLTE"`

	JoinedAt    time.Time `json:"joinedAt" gqlgen:"joinedAt" xml:"joinedAt"`
	JoinedAtGT  time.Time `json:"joinedAtGT" gqlgen:"joinedAtGT" xml:"joinedAtGT"`
	JoinedAtGTE time.Time `json:"joinedAtGTE" gqlgen:"joinedAtGTE" xml:"joinedAtGTE"`
	JoinedAtLT  time.Time `json:"joinedAtLT" gqlgen:"joinedAtLT" xml:"joinedAtLT"`
	JoinedAtLTE time.Time `json:"joinedAtLTE" gqlgen:"joinedAtLTE" xml:"joinedAtLTE"`

	LastActivityAt    time.Time `json:"lastActivityAt" gqlgen:"lastActivityAt" xml:"lastActivityAt"`
	LastActivityAtGT  time.Time `json:"lastActivityAtGT" gqlgen:"lastActivityAtGT" xml:"lastActivityAtGT"`
	LastActivityAtGTE time.Time `json:"lastActivityAtGTE" gqlgen:"lastActivityAtGTE" xml:"lastActivityAtGTE"`
	LastActivityAtLT  time.Time `json:"lastActivityAtLT" gqlgen:"lastActivityAtLT" xml:"lastActivityAtLT"`
	LastActivityAtLTE time.Time `json:"lastActivityAtLTE" gqlgen:"lastActivityAtLTE" xml:"lastActivityAtLTE"`

	DeletedAt    time.Time `json:"deletedAt" gqlgen:"deletedAt" xml:"deletedAt"`
	DeletedAtGT  time.Time `json:"deletedAtGT" gqlgen:"deletedAtGT" xml:"deletedAtGT"`
	DeletedAtGTE time.Time `json:"deletedAtGTE" gqlgen:"deletedAtGTE" xml:"deletedAtGTE"`
	DeletedAtLT  time.Time `json:"deletedAtLT" gqlgen:"deletedAtLT" xml:"deletedAtLT"`
	DeletedAtLTE time.Time `json:"deletedAtLTE" gqlgen:"deletedAtLTE" xml:"deletedAtLTE"`

	TribeID     []int        `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeIDNEQ  []int        `json:"tribeIDNEQ" gqlgen:"tribeIDNEQ" xml:"tribeIDNEQ"`
	TribeFilter *TribeFilter `json:"tribeFilter" gqlgen:"tribeFilter" xml:"tribeFilter"`

	OpponentsDefeatedFilter
}

func (f *PlayerFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if !isZero(f.ID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("id", alias)), pg.Array(f.ID))
	}
	if !isZero(f.IDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("id", alias)), pg.Array(f.IDNEQ))
	}

	if !isZero(f.Exists) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("exists", alias)), f.Exists)
	}

	if !isZero(f.Name) {
		q = q.Where(buildConditionArray(addAliasToColumnName("name", alias)), pg.Array(f.Name))
	}
	if !isZero(f.NameNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("name", alias)), pg.Array(f.NameNEQ))
	}
	if !isZero(f.NameMATCH) {
		q = q.Where(buildConditionMatch(addAliasToColumnName("name", alias)), f.NameMATCH)
	}
	if !isZero(f.NameIEQ) {
		q = q.Where(buildConditionIEQ(addAliasToColumnName("name", alias)), f.NameIEQ)
	}

	if !isZero(f.TotalVillages) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("total_villages", alias)), f.TotalVillages)
	}
	if !isZero(f.TotalVillagesGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("total_villages", alias)), f.TotalVillagesGT)
	}
	if !isZero(f.TotalVillagesGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("total_villages", alias)), f.TotalVillagesGTE)
	}
	if !isZero(f.TotalVillagesLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("total_villages", alias)), f.TotalVillagesLT)
	}
	if !isZero(f.TotalVillagesLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("total_villages", alias)), f.TotalVillagesLTE)
	}

	if !isZero(f.Points) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("points", alias)), f.Points)
	}
	if !isZero(f.PointsGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("points", alias)), f.PointsGT)
	}
	if !isZero(f.PointsGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("points", alias)), f.PointsGTE)
	}
	if !isZero(f.PointsLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("points", alias)), f.PointsLT)
	}
	if !isZero(f.PointsLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("points", alias)), f.PointsLTE)
	}

	if !isZero(f.Rank) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("rank", alias)), f.Rank)
	}
	if !isZero(f.RankGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("rank", alias)), f.RankGT)
	}
	if !isZero(f.RankGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("rank", alias)), f.RankGTE)
	}
	if !isZero(f.RankLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("rank", alias)), f.RankLT)
	}
	if !isZero(f.RankLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("rank", alias)), f.RankLTE)
	}

	if !isZero(f.DailyGrowth) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("daily_growth", alias)), f.DailyGrowth)
	}
	if !isZero(f.DailyGrowthGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("daily_growth", alias)), f.DailyGrowthGT)
	}
	if !isZero(f.DailyGrowthGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("daily_growth", alias)), f.DailyGrowthGTE)
	}
	if !isZero(f.DailyGrowthLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("daily_growth", alias)), f.DailyGrowthLT)
	}
	if !isZero(f.DailyGrowthLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("daily_growth", alias)), f.DailyGrowthLTE)
	}

	if !isZero(f.JoinedAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("joined_at", alias)), f.JoinedAt)
	}
	if !isZero(f.JoinedAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("joined_at", alias)), f.JoinedAtGT)
	}
	if !isZero(f.JoinedAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("joined_at", alias)), f.JoinedAtGTE)
	}
	if !isZero(f.JoinedAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("joined_at", alias)), f.JoinedAtLT)
	}
	if !isZero(f.JoinedAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("joined_at", alias)), f.JoinedAtLTE)
	}

	if !isZero(f.LastActivityAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("last_activity_at", alias)), f.LastActivityAt)
	}
	if !isZero(f.LastActivityAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("last_activity_at", alias)), f.LastActivityAtGT)
	}
	if !isZero(f.LastActivityAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("last_activity_at", alias)), f.LastActivityAtGTE)
	}
	if !isZero(f.LastActivityAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("last_activity_at", alias)), f.LastActivityAtLT)
	}
	if !isZero(f.LastActivityAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("last_activity_at", alias)), f.LastActivityAtLTE)
	}

	if !isZero(f.DeletedAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("deleted_at", alias)), f.DeletedAt)
	}
	if !isZero(f.DeletedAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("deleted_at", alias)), f.DeletedAtGT)
	}
	if !isZero(f.DeletedAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("deleted_at", alias)), f.DeletedAtGTE)
	}
	if !isZero(f.DeletedAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("deleted_at", alias)), f.DeletedAtLT)
	}
	if !isZero(f.DeletedAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("deleted_at", alias)), f.DeletedAtLTE)
	}

	if !isZero(f.TribeID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("tribe_id", alias)), pg.Array(f.TribeID))
	}
	if !isZero(f.TribeIDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("tribe_id", alias)), pg.Array(f.TribeIDNEQ))
	}

	return f.OpponentsDefeatedFilter.WhereWithAlias(q, alias)
}

func (f *PlayerFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "player")
}

type PlayerRelationshipAndSortAppender struct {
	Filter *PlayerFilter
	Sort   []string
}

func (a *PlayerRelationshipAndSortAppender) Append(q *orm.Query) (*orm.Query, error) {
	var err error
	tribeRequired := findStringWithPrefix(a.Sort, "tribe.") != ""
	if a.Filter.TribeFilter != nil {
		q, err = a.Filter.TribeFilter.WhereWithAlias(q, "tribe")
		if err != nil {
			return q, err
		}
		tribeRequired = true
	}

	if !isZero(a.Sort) {
		q = q.Order(a.Sort...)
	}

	if tribeRequired {
		q = q.Relation("Tribe._")
	}

	return q, nil
}

type FoundPlayer struct {
	Server       string `json:"server" xml:"server" gqlgen:"server"`
	ID           int    `json:"id" gqlgen:"id" xml:"id"`
	Name         string `json:"name" gqlgen:"name" xml:"name"`
	BestRank     int    `json:"bestRank" pg:",use_zero" gqlgen:"bestRank" xml:"bestRank"`
	MostPoints   int    `json:"mostPoints" pg:",use_zero" gqlgen:"mostPoints" xml:"mostPoints"`
	MostVillages int    `json:"mostVillages" pg:",use_zero" gqlgen:"mostVillages" xml:"mostVillages"`
	TribeID      int    `json:"tribeID" xml:"tribeID" gqlgen:"tribeID"`
	TribeTag     string `json:"tribeTag" xml:"tribeTag" gqlgen:"tribeTag"`
}
