package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type ServerStatus string

const (
	ServerStatusOpen   ServerStatus = "open"
	ServerStatusClosed ServerStatus = "closed"
)

func (ss ServerStatus) IsValid() bool {
	switch ss {
	case ServerStatusOpen,
		ServerStatusClosed:
		return true
	}
	return false
}

func (ss ServerStatus) String() string {
	return string(ss)
}

func (ss *ServerStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*ss = ServerStatus(strings.ToLower(str))
	if !ss.IsValid() {
		return fmt.Errorf("%s is not a valid ServerStatus", str)
	}
	return nil
}

func (ss ServerStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(ss.String()))
}

type Server struct {
	tableName struct{} `pg:"alias:server"`

	Key              string       `json:"key" gqlgen:"key" pg:",pk,unique" xml:"key"`
	Status           ServerStatus `json:"status" gqlgen:"status" xml:"status"`
	NumberOfPlayers  int          `pg:",use_zero" json:"numberOfPlayers" gqlgen:"numberOfPlayers" xml:"numberOfPlayers"`
	NumberOfTribes   int          `pg:",use_zero" json:"numberOfTribes" gqlgen:"numberOfTribes" xml:"numberOfTribes"`
	NumberOfVillages int          `pg:",use_zero" json:"numberOfVillages" gqlgen:"numberOfVillages" xml:"numberOfVillages"`

	Config         ServerConfig   `json:"config" gqlgen:"config" xml:"config"`
	BuildingConfig BuildingConfig `json:"buildingConfig" gqlgen:"buildingConfig" xml:"buildingConfig"`
	UnitConfig     UnitConfig     `json:"unitConfig" gqlgen:"unitConfig" xml:"unitConfig"`

	VersionCode VersionCode `json:"versionCode" gqlgen:"versionCode" xml:"versionCode"`
	Version     *Version    `json:"version,omitempty" gqlgen:"-" xml:"version" pg:"rel:has-one"`

	DataUpdatedAt    time.Time `pg:"default:now(),use_zero" json:"dataUpdatedAt" gqlgen:"dataUpdatedAt" xml:"dataUpdatedAt"`
	HistoryUpdatedAt time.Time `pg:"default:now(),use_zero" json:"historyUpdatedAt" gqlgen:"historyUpdatedAt" xml:"historyUpdatedAt"`
	StatsUpdatedAt   time.Time `pg:"default:now(),use_zero" json:"statsUpdatedAt" gqlgen:"statsUpdatedAt" xml:"statsUpdatedAt"`
}

type ServerFilter struct {
	Key      []string `json:"key" gqlgen:"key"`
	KeyNEQ   []string `json:"keyNEQ" gqlgen:"keyNEQ"`
	KeyMATCH string   `json:"keyMATCH" gqlgen:"keyMATCH"`
	KeyIEQ   string   `json:"keyIEQ" gqlgen:"keyIEQ"`

	Status    []ServerStatus `json:"status" gqlgen:"status"`
	StatusNEQ []ServerStatus `json:"statusNEQ" gqlgen:"statusNEQ"`

	VersionCode    []VersionCode `json:"versionCode" gqlgen:"versionCode"`
	VersionCodeNEQ []VersionCode `json:"versionCodeNEQ" gqlgen:"versionCodeNEQ"`
}

func (f *ServerFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if len(f.Key) > 0 {
		q = q.Where(buildConditionArray(addAliasToColumnName("key", alias)), pg.Array(f.Key))
	}
	if len(f.KeyNEQ) > 0 {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("key", alias)), pg.Array(f.KeyNEQ))
	}
	if f.KeyMATCH != "" {
		q = q.Where(buildConditionMatch(addAliasToColumnName("key", alias)), f.KeyMATCH)
	}
	if f.KeyIEQ != "" {
		q = q.Where(buildConditionIEQ(addAliasToColumnName("key", alias)), f.KeyIEQ)
	}

	if len(f.Status) > 0 {
		q = q.Where(buildConditionArray(addAliasToColumnName("status", alias)), pg.Array(f.Status))
	}
	if len(f.StatusNEQ) > 0 {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("status", alias)), pg.Array(f.StatusNEQ))
	}

	if len(f.VersionCode) > 0 {
		q = q.Where(buildConditionArray(addAliasToColumnName("version_code", alias)), pg.Array(f.VersionCode))
	}
	if len(f.VersionCodeNEQ) > 0 {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("version_code", alias)), pg.Array(f.VersionCodeNEQ))
	}

	return q, nil
}

func (f *ServerFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "server")
}
