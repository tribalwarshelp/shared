package models

import "time"

type Tribe struct {
	tableName struct{} `pg:"?SERVER.tribes,alias:tribe"`

	ID             int       `json:"id" gqlgen:"id"`
	Name           string    `json:"name" gqlgen:"name"`
	Tag            string    `json:"tag" gqlgen:"tag"`
	Exist          *bool     `json:"exist" gqlgen:"exist" pg:",use_zero"`
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
	DeletedAt      time.Time `json:"deletedAt" pg:",use_zero" gqlgen:"deletedAt"`

	OpponentsDefeated
}

type TribeFilter struct {
	tableName struct{} `urlstruct:"tribe"`

	ID    []int `json:"id" gqlgen:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Exist *bool `urlstruct:",nowhere" json:"exist" gqlgen:"exist"`

	Tag      []string `json:"tag" gqlgen:"tag"`
	TagNEQ   []string `json:"tagNEQ" gqlgen:"tagNEQ"`
	TagMATCH string   `json:"tagMATCH" gqlgen:"tagMATCH"`
	TagIEQ   string   `json:"tagIEQ" gqlgen:"tagIEQ"`

	Name      []string `json:"name" gqlgen:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ"`

	TotalMembers    int `json:"totalMembers" gqlgen:"totalMembers"`
	TotalMembersGT  int `json:"totalMembersGT" gqlgen:"totalMembersGT"`
	TotalMembersGTE int `json:"totalMembersGTE" gqlgen:"totalMembersGTE"`
	TotalMembersLT  int `json:"totalMembersLT" gqlgen:"totalMembersLT"`
	TotalMembersLTE int `json:"totalMembersLTE" gqlgen:"totalMembersLTE"`

	TotalVillages    int `json:"totalVillages" gqlgen:"totalVillages"`
	TotalVillagesGT  int `json:"totalVillagesGT" gqlgen:"totalVillagesGT"`
	TotalVillagesGTE int `json:"totalVillagesGTE" gqlgen:"totalVillagesGTE"`
	TotalVillagesLT  int `json:"totalVillagesLT" gqlgen:"totalVillagesLT"`
	TotalVillagesLTE int `json:"totalVillagesLTE" gqlgen:"totalVillagesLTE"`

	Points    int `json:"points" gqlgen:"points"`
	PointsGT  int `json:"pointsGT" gqlgen:"pointsGT"`
	PointsGTE int `json:"pointsGTE" gqlgen:"pointsGTE"`
	PointsLT  int `json:"pointsLT" gqlgen:"pointsLT"`
	PointsLTE int `json:"pointsLTE" gqlgen:"pointsLTE"`

	AllPoints    int `json:"allPoints" gqlgen:"allPoints"`
	AllPointsGT  int `json:"allPointsGT" gqlgen:"allPointsGT"`
	AllPointsGTE int `json:"allPointsGTE" gqlgen:"allPointsGTE"`
	AllPointsLT  int `json:"allPointsLT" gqlgen:"allPointsLT"`
	AllPointsLTE int `json:"allPointsLTE" gqlgen:"allPointsLTE"`

	Rank    int `json:"rank" gqlgen:"rank"`
	RankGT  int `json:"rankGT" gqlgen:"rankGT"`
	RankGTE int `json:"rankGTE" gqlgen:"rankGTE"`
	RankLT  int `json:"rankLT" gqlgen:"rankLT"`
	RankLTE int `json:"rankLTE" gqlgen:"rankLTE"`

	Dominance    int `json:"dominance" gqlgen:"dominance"`
	DominanceGT  int `json:"dominanceGT" gqlgen:"dominanceGT"`
	DominanceGTE int `json:"dominanceGTE" gqlgen:"dominanceGTE"`
	DominanceLT  int `json:"dominanceLT" gqlgen:"dominanceLT"`
	DominanceLTE int `json:"dominanceLTE" gqlgen:"dominanceLTE"`

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

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`

	OpponentsDefeatedFilter
}
