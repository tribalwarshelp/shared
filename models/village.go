package models

import (
	"fmt"
	"strconv"
)

type Village struct {
	tableName struct{} `pg:"?SERVER.villages,alias:village"`

	ID     int    `json:"id" pg:",pk" gqlgen:"id"`
	Name   string `json:"name" gqlgen:"name"`
	Points int    `json:"points" pg:",use_zero" gqlgen:"points"`
	X      int    `json:"x" pg:",use_zero" gqlgen:"x"`
	Y      int    `json:"y" pg:",use_zero" gqlgen:"y"`
	Bonus  int    `json:"bonus" pg:",use_zero" gqlgen:"bonus"`

	PlayerID int     `json:"-" pg:",use_zero" gqlgen:"playerID"`
	Player   *Player `json:"player,omitempty" gqlgen:"-"`
}

func (v *Village) Continent() string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("k%c%c", strconv.FormatInt(int64(v.Y), 10)[0], strconv.FormatInt(int64(v.X), 10)[0])
}

type VillageFilter struct {
	tableName struct{} `urlstruct:"village"`

	ID    []int `json:"id" gqlgen:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Name      []string `json:"name" gqlgen:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ"`

	XGT  int `urlstruct:",nowhere" json:"xGT" gqlgen:"xGT"`
	XGTE int `urlstruct:",nowhere" json:"xGTE" gqlgen:"xGTE"`
	XLT  int `urlstruct:",nowhere" json:"xLT" gqlgen:"xLT"`
	XLTE int `urlstruct:",nowhere" json:"xLTE" gqlgen:"xLTE"`
	YGT  int `urlstruct:",nowhere" json:"yGT" gqlgen:"yGT"`
	YGTE int `urlstruct:",nowhere" json:"yGTE" gqlgen:"yGTE"`
	YLT  int `urlstruct:",nowhere" json:"yLT" gqlgen:"yLT"`
	YLTE int `urlstruct:",nowhere" json:"yLTE" gqlgen:"yLTE"`

	Points    int `json:"points" gqlgen:"points"`
	PointsGT  int `json:"pointsGT" gqlgen:"pointsGT"`
	PointsGTE int `json:"pointsGTE" gqlgen:"pointsGTE"`
	PointsLT  int `json:"pointsLT" gqlgen:"pointsLT"`
	PointsLTE int `json:"pointsLTE" gqlgen:"pointsLTE"`

	Bonus    int `json:"bonus" gqlgen:"bonus"`
	BonusGT  int `json:"bonusGT" gqlgen:"bonusGT"`
	BonusGTE int `json:"bonusGTE" gqlgen:"bonusGTE"`
	BonusLT  int `json:"bonusLT" gqlgen:"bonusLT"`
	BonusLTE int `json:"bonusLTE" gqlgen:"bonusLTE"`

	PlayerID     []int         `json:"playerID" gqlgen:"playerID"`
	PlayerFilter *PlayerFilter `urlstruct:",nowhere" json:"playerFilter" gqlgen:"playerFilter"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
