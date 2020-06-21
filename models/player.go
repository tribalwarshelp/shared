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

	ID    []int `json:"id" gqlgen:"id" xml:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ" xml:"idNEQ"`

	Exist *bool `urlstruct:",nowhere" json:"exist" gqlgen:"exist" xml:"exist"`

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

	TribeID     []int        `json:"tribeID" gqlgen:"tribeID" xml:"tribeID"`
	TribeFilter *TribeFilter `urlstruct:",nowhere" json:"tribeFilter" gqlgen:"tribeFilter" xml:"tribeFilter"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset" xml:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit" xml:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort" xml:"sort"`

	OpponentsDefeatedFilter
}
