package models

import "time"

type Player struct {
	tableName struct{} `pg:"?SERVER.players,alias:player"`

	ID             int       `json:"id" pg:"type:bigint,pk" gqlgen:"id"`
	Name           string    `json:"name" gqlgen:"name"`
	Exists         *bool     `json:"exists" pg:",use_zero" gqlgen:"exists"`
	TotalVillages  int       `json:"totalVillages" pg:",use_zero" gqlgen:"totalVillages"`
	Points         int       `json:"points" pg:",use_zero" gqlgen:"points"`
	Rank           int       `json:"rank" pg:",use_zero" gqlgen:"rank"`
	TribeID        int       `json:"-" pg:",use_zero" gqlgen:"tribeID"`
	Tribe          *Tribe    `json:"tribe,omitempty" gqlgen:"-" pg:"rel:has-one"`
	DailyGrowth    int       `json:"dailyGrowth" pg:",use_zero" gqlgen:"dailyGrowth"`
	BestRank       int       `json:"bestRank" pg:",use_zero" gqlgen:"bestRank"`
	BestRankAt     time.Time `json:"bestRankAt" pg:"default:now(),use_zero" gqlgen:"bestRankAt"`
	MostPoints     int       `json:"mostPoints" pg:",use_zero" gqlgen:"mostPoints"`
	MostPointsAt   time.Time `json:"mostPointsAt" pg:"default:now(),use_zero" gqlgen:"mostPointsAt"`
	MostVillages   int       `json:"mostVillages" pg:",use_zero" gqlgen:"mostVillages"`
	MostVillagesAt time.Time `json:"mostVillagesAt" pg:"default:now(),use_zero" gqlgen:"mostVillagesAt"`
	JoinedAt       time.Time `json:"joinedAt" pg:"default:now(),use_zero" gqlgen:"joinedAt"`
	DeletedAt      time.Time `json:"deletedAt" pg:",use_zero" gqlgen:"deletedAt"`

	OpponentsDefeated
}

type PlayerFilter struct {
	tableName struct{} `urlstruct:"player"`

	ID    []int `json:"id" gqlgen:"id" xml:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ" xml:"idNEQ"`

	Exists *bool `urlstruct:",nowhere" json:"exists" gqlgen:"exists" xml:"exists"`

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

	DeletedAt    time.Time `json:"deletedAt" gqlgen:"deletedAt" xml:"deletedAt"`
	DeletedAtGT  time.Time `json:"deletedAtGT" gqlgen:"deletedAtGT" xml:"deletedAtGT"`
	DeletedAtGTE time.Time `json:"deletedAtGTE" gqlgen:"deletedAtGTE" xml:"deletedAtGTE"`
	DeletedAtLT  time.Time `json:"deletedAtLT" gqlgen:"deletedAtLT" xml:"deletedAtLT"`
	DeletedAtLTE time.Time `json:"deletedAtLTE" gqlgen:"deletedAtLTE" xml:"deletedAtLTE"`

	TribeID     []int        `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeFilter *TribeFilter `urlstruct:",nowhere" json:"tribeFilter" gqlgen:"tribeFilter" xml:"tribeFilter"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`

	OpponentsDefeatedFilter
}
