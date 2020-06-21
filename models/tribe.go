package models

type Tribe struct {
	tableName struct{} `pg:"?SERVER.tribes,alias:tribe"`

	ID            int     `json:"id" gqlgen:"id"`
	Name          string  `json:"name" gqlgen:"name"`
	Tag           string  `json:"tag" gqlgen:"tag"`
	TotalMembers  int     `json:"totalMembers" gqlgen:"totalMembers" pg:",use_zero"`
	TotalVillages int     `json:"totalVillages" gqlgen:"totalVillages" pg:",use_zero"`
	Points        int     `json:"points" gqlgen:"points" pg:",use_zero"`
	AllPoints     int     `json:"allPoints" gqlgen:"allPoints" pg:",use_zero"`
	Rank          int     `json:"rank" gqlgen:"rank" pg:",use_zero"`
	Dominance     float64 `json:"dominance" gqlgen:"dominance" pg:",use_zero"`
	Exist         *bool   `json:"exist" gqlgen:"exist" pg:",use_zero"`

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

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`

	OpponentsDefeatedFilter
}
