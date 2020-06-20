package models

type Player struct {
	tableName struct{} `pg:"?SERVER.players,alias:player"`

	ID            int    `json:"id" pg:",pk" gqlgen:"id"`
	Name          string `json:"name" gqlgen:"name"`
	TotalVillages int    `json:"totalVillages" pg:",use_zero" gqlgen:"totalVillages"`
	Points        int    `json:"points" pg:",use_zero" gqlgen:"points"`
	Rank          int    `json:"rank" pg:",use_zero" gqlgen:"rank"`
	Exist         *bool  `json:"exist" pg:",use_zero" gqlgen:"exist"`
	TribeID       int    `json:"-" pg:",use_zero" gqlgen:"tribeID"`
	Tribe         *Tribe `json:"tribe,omitempty" gqlgen:"-"`

	OpponentsDefeated
}

type PlayerFilter struct {
	tableName struct{} `urlstruct:"player"`

	ID    []int `json:"id" gqlgen:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Exist *bool `urlstruct:",nowhere" json:"exist" gqlgen:"exist"`

	Name      []string `json:"name" gqlgen:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ"`

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

	Rank    int `json:"rank" gqlgen:"rank"`
	RankGT  int `json:"rankGT" gqlgen:"rankGT"`
	RankGTE int `json:"rankGTE" gqlgen:"rankGTE"`
	RankLT  int `json:"rankLT" gqlgen:"rankLT"`
	RankLTE int `json:"rankLTE" gqlgen:"rankLTE"`

	TribeID []int `json:"tribeID" gqlgen:"tribeID"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`

	OpponentsDefeatedFilter
}
