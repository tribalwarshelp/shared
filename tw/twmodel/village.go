package twmodel

import (
	"fmt"
	"github.com/Kichiyaki/gopgutil/v10"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"strconv"
)

type Village struct {
	tableName struct{} `pg:"?SERVER.villages,alias:village"`

	ID     int    `json:"id" pg:"type:bigint,pk" gqlgen:"id"`
	Name   string `json:"name" gqlgen:"name"`
	Points int    `json:"points" pg:",use_zero" gqlgen:"points"`
	X      int    `json:"x" pg:",use_zero" gqlgen:"x"`
	Y      int    `json:"y" pg:",use_zero" gqlgen:"y"`
	Bonus  int    `json:"bonus" pg:",use_zero" gqlgen:"bonus"`

	PlayerID int     `json:"-" pg:",use_zero" gqlgen:"playerID"`
	Player   *Player `json:"player,omitempty" gqlgen:"-" pg:"rel:has-one"`
}

func (v *Village) Continent() string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("K%c%c", strconv.FormatInt(int64(v.Y), 10)[0], strconv.FormatInt(int64(v.X), 10)[0])
}

func (v *Village) FullName() string {
	return fmt.Sprintf("%s (%d|%d) %s",
		v.Name,
		v.X,
		v.Y,
		v.Continent())
}

type VillageFilter struct {
	ID    []int `json:"id" gqlgen:"id"`
	IDNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Name      []string `json:"name" gqlgen:"name"`
	NameNEQ   []string `json:"nameNEQ" gqlgen:"nameNEQ"`
	NameMATCH string   `json:"nameMATCH" gqlgen:"nameMATCH"`
	NameIEQ   string   `json:"nameIEQ" gqlgen:"nameIEQ"`

	XGT  int      `json:"xGT" gqlgen:"xGT"`
	XGTE int      `json:"xGTE" gqlgen:"xGTE"`
	XLT  int      `json:"xLT" gqlgen:"xLT"`
	XLTE int      `json:"xLTE" gqlgen:"xLTE"`
	YGT  int      `json:"yGT" gqlgen:"yGT"`
	YGTE int      `json:"yGTE" gqlgen:"yGTE"`
	YLT  int      `json:"yLT" gqlgen:"yLT"`
	YLTE int      `json:"yLTE" gqlgen:"yLTE"`
	XY   []string `json:"xy" gqlgen:"xy"`

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
	PlayerIDNEQ  []int         `json:"playerIDNEQ" gqlgen:"playerIDNEQ"`
	PlayerFilter *PlayerFilter `json:"playerFilter" gqlgen:"playerFilter"`
}

func (f *VillageFilter) WhereWithAlias(q *orm.Query, alias, playerRelationName, playerAlias, tribeRelationName, tribeAlias string) (*orm.Query, error) {
	if f == nil {
		return q, nil
	}

	if !isZero(f.ID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("id", alias), pg.Array(f.ID))
	}
	if !isZero(f.IDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("id", alias), pg.Array(f.IDNEQ))
	}

	if !isZero(f.Name) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("name", alias), pg.Array(f.Name))
	}
	if !isZero(f.NameNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("name", alias), pg.Array(f.NameNEQ))
	}
	if !isZero(f.NameMATCH) {
		q = q.Where(gopgutil.BuildConditionMatch("?"), gopgutil.AddAliasToColumnName("name", alias), f.NameMATCH)
	}
	if !isZero(f.NameIEQ) {
		q = q.Where(gopgutil.BuildConditionIEQ("?"), gopgutil.AddAliasToColumnName("name", alias), f.NameIEQ)
	}

	if !isZero(f.XGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("x", alias), f.XGT)
	}
	if !isZero(f.XGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("x", alias), f.XGTE)
	}
	if !isZero(f.XLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("x", alias), f.XLT)
	}
	if !isZero(f.XLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("x", alias), f.XLTE)
	}

	if !isZero(f.YGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("y", alias), f.YGT)
	}
	if !isZero(f.YGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("y", alias), f.YGTE)
	}
	if !isZero(f.YLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("y", alias), f.YLT)
	}
	if !isZero(f.YLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("y", alias), f.YLTE)
	}

	if !isZero(f.XY) {
		q = q.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			for _, xy := range f.XY {
				coords, err := ParseCoords(xy)
				if err != nil {
					continue
				}
				q = q.WhereOrGroup(func(q *orm.Query) (*orm.Query, error) {
					q = q.
						Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("x", alias), coords.X).
						Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("y", alias), coords.Y)
					return q, nil
				})
			}
			return q, nil
		})
	}

	if !isZero(f.Points) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("points", alias), f.Points)
	}
	if !isZero(f.PointsGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("points", alias), f.PointsGT)
	}
	if !isZero(f.PointsGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("points", alias), f.PointsGTE)
	}
	if !isZero(f.PointsLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("points", alias), f.PointsLT)
	}
	if !isZero(f.PointsLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("points", alias), f.PointsLTE)
	}

	if !isZero(f.Bonus) {
		q = q.Where(gopgutil.BuildConditionEquals("?"), gopgutil.AddAliasToColumnName("bonus", alias), f.Bonus)
	}
	if !isZero(f.BonusGT) {
		q = q.Where(gopgutil.BuildConditionGT("?"), gopgutil.AddAliasToColumnName("bonus", alias), f.BonusGT)
	}
	if !isZero(f.BonusGTE) {
		q = q.Where(gopgutil.BuildConditionGTE("?"), gopgutil.AddAliasToColumnName("bonus", alias), f.BonusGTE)
	}
	if !isZero(f.BonusLT) {
		q = q.Where(gopgutil.BuildConditionLT("?"), gopgutil.AddAliasToColumnName("bonus", alias), f.BonusLT)
	}
	if !isZero(f.BonusLTE) {
		q = q.Where(gopgutil.BuildConditionLTE("?"), gopgutil.AddAliasToColumnName("bonus", alias), f.BonusLTE)
	}

	if !isZero(f.PlayerID) {
		q = q.Where(gopgutil.BuildConditionArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerID))
	}
	if !isZero(f.PlayerIDNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray("?"), gopgutil.AddAliasToColumnName("player_id", alias), pg.Array(f.PlayerIDNEQ))
	}

	if f.PlayerFilter != nil && playerRelationName != "" && tribeRelationName != "" {
		return f.PlayerFilter.WhereWithAlias(q.Relation(playerRelationName), playerAlias, tribeRelationName, tribeAlias)
	}

	return q, nil
}

func (f *VillageFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(
		q,
		"village",
		"Player._",
		"player",
		"Player.Tribe._",
		"player__tribe",
	)
}
