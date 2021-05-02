package twmodel

import (
	"fmt"
	"github.com/Kichiyaki/gopgutil/v10"
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
	if f == nil {
		return q, nil
	}

	if !isZero(f.Key) {
		q = q.Where(gopgutil.BuildConditionArray(gopgutil.AddAliasToColumnName("key", alias)), pg.Array(f.Key))
	}
	if !isZero(f.KeyNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray(gopgutil.AddAliasToColumnName("key", alias)), pg.Array(f.KeyNEQ))
	}
	if !isZero(f.KeyMATCH) {
		q = q.Where(gopgutil.BuildConditionMatch(gopgutil.AddAliasToColumnName("key", alias)), f.KeyMATCH)
	}
	if !isZero(f.KeyIEQ) {
		q = q.Where(gopgutil.BuildConditionIEQ(gopgutil.AddAliasToColumnName("key", alias)), f.KeyIEQ)
	}

	if !isZero(f.Status) {
		q = q.Where(gopgutil.BuildConditionArray(gopgutil.AddAliasToColumnName("status", alias)), pg.Array(f.Status))
	}
	if !isZero(f.StatusNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray(gopgutil.AddAliasToColumnName("status", alias)), pg.Array(f.StatusNEQ))
	}

	if !isZero(f.VersionCode) {
		q = q.Where(gopgutil.BuildConditionArray(gopgutil.AddAliasToColumnName("version_code", alias)), pg.Array(f.VersionCode))
	}
	if !isZero(f.VersionCodeNEQ) {
		q = q.Where(gopgutil.BuildConditionNotInArray(gopgutil.AddAliasToColumnName("version_code", alias)), pg.Array(f.VersionCodeNEQ))
	}

	return q, nil
}

func (f *ServerFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "server")
}
