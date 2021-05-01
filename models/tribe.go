package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Tribe struct {
	tableName struct{} `pg:"?SERVER.tribes,alias:tribe"`

	ID             int       `json:"id" pg:"type:bigint,pk" gqlgen:"id"`
	Name           string    `json:"name" gqlgen:"name"`
	Tag            string    `json:"tag" gqlgen:"tag"`
	Exists         *bool     `json:"exists" gqlgen:"exists" pg:",use_zero"`
	TotalMembers   int       `json:"totalMembers" gqlgen:"totalMembers" pg:",use_zero"`
	TotalVillages  int       `json:"totalVillages" gqlgen:"totalVillages" pg:",use_zero"`
	Points         int       `json:"points" gqlgen:"points" pg:",use_zero"`
	AllPoints      int       `json:"allPoints" gqlgen:"allPoints" pg:",use_zero"`
	Rank           int       `json:"rank" gqlgen:"rank" pg:",use_zero"`
	Dominance      float64   `json:"dominance" gqlgen:"dominance" pg:",use_zero"`
	BestRank       int       `json:"bestRank" pg:",use_zero" gqlgen:"bestRank"`
	BestRankAt     time.Time `json:"bestRankAt" pg:"default:now(),use_zero" gqlgen:"bestRankAt"`
	MostPoints     int       `json:"mostPoints" pg:",use_zero" gqlgen:"mostPoints"`
	MostPointsAt   time.Time `json:"mostPointsAt" pg:"default:now(),use_zero" gqlgen:"mostPointsAt"`
	MostVillages   int       `json:"mostVillages" pg:",use_zero" gqlgen:"mostVillages"`
	MostVillagesAt time.Time `json:"mostVillagesAt" pg:"default:now(),use_zero" gqlgen:"mostVillagesAt"`
	CreatedAt      time.Time `json:"createdAt" pg:"default:now(),use_zero" gqlgen:"createdAt"`
	DeletedAt      time.Time `json:"deletedAt" gqlgen:"deletedAt"`

	OpponentsDefeated
}

type TribeFilterOr struct {
	ID      []int    `json:"id" xml:"id" gqlgen:"id"`
	Tag     []string `json:"tag" xml:"tag" gqlgen:"tag"`
	TagIEQ  string   `json:"tagIEQ" xml:"tagIEQ" gqlgen:"tagIEQ"`
	Name    []string `json:"name" xml:"name" gqlgen:"name"`
	NameIEQ string   `json:"nameIEQ" xml:"nameIEQ" gqlgen:"nameIEQ"`
}

func (f *TribeFilterOr) WhereWithAlias(q *orm.Query, alias string) *orm.Query {
	q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
		if !isZero(f.ID) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("id", alias)), pg.Array(f.ID))
		}
		if !isZero(f.Tag) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("tag", alias)), pg.Array(f.Tag))
		}
		if !isZero(f.TagIEQ) {
			q = q.WhereOr(buildConditionIEQ(addAliasToColumnName("tag", alias)), f.TagIEQ)
		}
		if !isZero(f.Name) {
			q = q.WhereOr(buildConditionArray(addAliasToColumnName("name", alias)), pg.Array(f.Name))
		}
		if !isZero(f.NameIEQ) {
			q = q.WhereOr(buildConditionIEQ(addAliasToColumnName("name", alias)), f.NameIEQ)
		}
		return q, nil
	})
	return q
}

type TribeFilter struct {
	ID    []int `json:"id" gqlgen:"id" xml:"id"`
	IDNEQ []int `json:"idNEQ" gqlgen:"idNEQ" xml:"idneq"`

	Exists *bool `json:"exists" gqlgen:"exists" xml:"exists"`

	Tag      []string `json:"tag" gqlgen:"tag" xml:"tag"`
	TagNEQ   []string `json:"tagNEQ" gqlgen:"tagNEQ" xml:"tagNEQ"`
	TagMATCH string   `json:"tagMATCH" gqlgen:"tagMATCH" xml:"tagMATCH"`
	TagIEQ   string   `json:"tagIEQ" gqlgen:"tagIEQ" xml:"tagIEQ"`

	Name      []string `json:"name" gqlgen:"name" xml:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ" xml:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH" xml:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ" xml:"nameIEQ"`

	TotalMembers    int `json:"totalMembers" gqlgen:"totalMembers" xml:"totalMembers"`
	TotalMembersGT  int `json:"totalMembersGT" gqlgen:"totalMembersGT" xml:"totalMembersGT"`
	TotalMembersGTE int `json:"totalMembersGTE" gqlgen:"totalMembersGTE" xml:"totalMembersGTE"`
	TotalMembersLT  int `json:"totalMembersLT" gqlgen:"totalMembersLT" xml:"totalMembersLT"`
	TotalMembersLTE int `json:"totalMembersLTE" gqlgen:"totalMembersLTE" xml:"totalMembersLTE"`

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

	AllPoints    int `json:"allPoints" gqlgen:"allPoints" xml:"allPoints"`
	AllPointsGT  int `json:"allPointsGT" gqlgen:"allPointsGT" xml:"allPointsGT"`
	AllPointsGTE int `json:"allPointsGTE" gqlgen:"allPointsGTE" xml:"allPointsGTE"`
	AllPointsLT  int `json:"allPointsLT" gqlgen:"allPointsLT" xml:"allPointsLT"`
	AllPointsLTE int `json:"allPointsLTE" gqlgen:"allPointsLTE" xml:"allPointsLTE"`

	Rank    int `json:"rank" gqlgen:"rank" xml:"rank"`
	RankGT  int `json:"rankGT" gqlgen:"rankGT" xml:"rankGT"`
	RankGTE int `json:"rankGTE" gqlgen:"rankGTE" xml:"rankGTE"`
	RankLT  int `json:"rankLT" gqlgen:"rankLT" xml:"rankLT"`
	RankLTE int `json:"rankLTE" gqlgen:"rankLTE" xml:"rankLTE"`

	Dominance    int `json:"dominance" gqlgen:"dominance" xml:"dominance"`
	DominanceGT  int `json:"dominanceGT" gqlgen:"dominanceGT" xml:"dominanceGT"`
	DominanceGTE int `json:"dominanceGTE" gqlgen:"dominanceGTE" xml:"dominanceGTE"`
	DominanceLT  int `json:"dominanceLT" gqlgen:"dominanceLT" xml:"dominanceLT"`
	DominanceLTE int `json:"dominanceLTE" gqlgen:"dominanceLTE" xml:"dominanceLTE"`

	CreatedAt    time.Time `json:"createdAt" gqlgen:"createdAt" xml:"createdAt"`
	CreatedAtGT  time.Time `json:"createdAtGT" gqlgen:"createdAtGT" xml:"createdAtGT"`
	CreatedAtGTE time.Time `json:"createdAtGTE" gqlgen:"createdAtGTE" xml:"createdAtGTE"`
	CreatedAtLT  time.Time `json:"createdAtLT" gqlgen:"createdAtLT" xml:"createdAtLT"`
	CreatedAtLTE time.Time `json:"createdAtLTE" gqlgen:"createdAtLTE" xml:"createdAtLTE"`

	DeletedAt    time.Time `json:"deletedAt" gqlgen:"deletedAt" xml:"deletedAt"`
	DeletedAtGT  time.Time `json:"deletedAtGT" gqlgen:"deletedAtGT" xml:"deletedAtGT"`
	DeletedAtGTE time.Time `json:"deletedAtGTE" gqlgen:"deletedAtGTE" xml:"deletedAtGTE"`
	DeletedAtLT  time.Time `json:"deletedAtLT" gqlgen:"deletedAtLT" xml:"deletedAtLT"`
	DeletedAtLTE time.Time `json:"deletedAtLTE" gqlgen:"deletedAtLTE" xml:"deletedAtLTE"`

	OpponentsDefeatedFilter
	Or *TribeFilterOr `json:"or" xml:"or" gqlgen:"or"`
}

func (f *TribeFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if !isZero(f.ID) {
		q = q.Where(buildConditionArray(addAliasToColumnName("id", alias)), pg.Array(f.ID))
	}
	if !isZero(f.IDNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("id", alias)), pg.Array(f.IDNEQ))
	}

	if !isZero(f.Exists) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("exists", alias)), f.Exists)
	}

	if !isZero(f.Tag) {
		q = q.Where(buildConditionArray(addAliasToColumnName("tag", alias)), pg.Array(f.Tag))
	}
	if !isZero(f.TagNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("tag", alias)), pg.Array(f.TagNEQ))
	}
	if !isZero(f.TagMATCH) {
		q = q.Where(buildConditionMatch(addAliasToColumnName("tag", alias)), f.TagMATCH)
	}
	if !isZero(f.TagIEQ) {
		q = q.Where(buildConditionIEQ(addAliasToColumnName("tag", alias)), f.TagIEQ)
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

	if !isZero(f.TotalMembers) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("total_members", alias)), f.TotalMembers)
	}
	if !isZero(f.TotalMembersGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("total_members", alias)), f.TotalMembersGT)
	}
	if !isZero(f.TotalMembersGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("total_members", alias)), f.TotalMembersGTE)
	}
	if !isZero(f.TotalMembersLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("total_members", alias)), f.TotalMembersLT)
	}
	if !isZero(f.TotalMembersLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("total_members", alias)), f.TotalMembersLTE)
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

	if !isZero(f.AllPoints) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("all_points", alias)), f.AllPoints)
	}
	if !isZero(f.AllPointsGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("all_points", alias)), f.AllPointsGT)
	}
	if !isZero(f.AllPointsGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("all_points", alias)), f.AllPointsGTE)
	}
	if !isZero(f.AllPointsLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("all_points", alias)), f.AllPointsLT)
	}
	if !isZero(f.AllPointsLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("all_points", alias)), f.AllPointsLTE)
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

	if !isZero(f.Dominance) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("dominance", alias)), f.Dominance)
	}
	if !isZero(f.DominanceGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("dominance", alias)), f.DominanceGT)
	}
	if !isZero(f.DominanceGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("dominance", alias)), f.DominanceGTE)
	}
	if !isZero(f.DominanceLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("dominance", alias)), f.DominanceLT)
	}
	if !isZero(f.DominanceLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("dominance", alias)), f.DominanceLTE)
	}

	if !isZero(f.CreatedAt) {
		q = q.Where(buildConditionEquals(addAliasToColumnName("created_at", alias)), f.CreatedAt)
	}
	if !isZero(f.CreatedAtGT) {
		q = q.Where(buildConditionGT(addAliasToColumnName("created_at", alias)), f.CreatedAtGT)
	}
	if !isZero(f.CreatedAtGTE) {
		q = q.Where(buildConditionGTE(addAliasToColumnName("created_at", alias)), f.CreatedAtGTE)
	}
	if !isZero(f.CreatedAtLT) {
		q = q.Where(buildConditionLT(addAliasToColumnName("created_at", alias)), f.CreatedAtLT)
	}
	if !isZero(f.CreatedAtLTE) {
		q = q.Where(buildConditionLTE(addAliasToColumnName("created_at", alias)), f.CreatedAtLTE)
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

	if f.Or != nil {
		q = f.Or.WhereWithAlias(q, alias)
	}

	return f.OpponentsDefeatedFilter.WhereWithAlias(q, alias)
}

func (f *TribeFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "tribe")
}

type FoundTribe struct {
	Server       string `json:"server" xml:"server" gqlgen:"server"`
	ID           int    `json:"id" gqlgen:"id" xml:"id"`
	Tag          string `json:"tag" xml:"tag" gqlgen:"tag"`
	Name         string `json:"name" gqlgen:"name" xml:"name"`
	BestRank     int    `json:"bestRank" pg:",use_zero" gqlgen:"bestRank" xml:"bestRank"`
	MostPoints   int    `json:"mostPoints" pg:",use_zero" gqlgen:"mostPoints" xml:"mostPoints"`
	MostVillages int    `json:"mostVillages" pg:",use_zero" gqlgen:"mostVillages" xml:"mostVillages"`
}
