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
	tableName struct{} `urlstruct:"village" json:"tableName" gqlgen:"tableName"`

	ID    []int `json:"id" gqlgen:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Name      []string `json:"name" gqlgen:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ"`

	XCoordGT  int `urlstruct:"xGT" json:"xGT" gqlgen:"xGT"`
	XCoordGTE int `urlstruct:"xGTE" json:"xGTE" gqlgen:"xGTE"`
	XCoordLT  int `urlstruct:"xLT" json:"xLT" gqlgen:"xLT"`
	XCoordLTE int `urlstruct:"xLTE" json:"xLTE" gqlgen:"xLTE"`
	YCoordGT  int `urlstruct:"yGT" json:"yGT" gqlgen:"yGT"`
	YCoordGTE int `urlstruct:"yGTE" json:"yGTE" gqlgen:"yGTE"`
	YCoordLT  int `urlstruct:"yLT" json:"yLT" gqlgen:"yLT"`
	YCoordLTE int `urlstruct:"yLTE" json:"yLTE" gqlgen:"yLTE"`

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

	PlayerID []int `json:"playerID" gqlgen:"playerID"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
