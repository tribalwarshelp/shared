package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
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

	LangVersionTag LanguageTag  `json:"langVersionTag" gqlgen:"langVersionTag" xml:"langVersionTag"`
	LangVersion    *LangVersion `json:"langVersion,omitempty" gqlgen:"-" xml:"langVersion"`

	DataUpdatedAt    time.Time `pg:"default:now(),use_zero" json:"dataUpdatedAt" gqlgen:"dataUpdatedAt" xml:"dataUpdatedAt"`
	HistoryUpdatedAt time.Time `pg:"default:now(),use_zero" json:"historyUpdatedAt" gqlgen:"historyUpdatedAt" xml:"historyUpdatedAt"`
	StatsUpdatedAt   time.Time `pg:"default:now(),use_zero" json:"statsUpdatedAt" gqlgen:"statsUpdatedAt" xml:"statsUpdatedAt"`
}

type ServerFilter struct {
	tableName struct{} `urlstruct:"server"`

	ID    []int `json:"id" gqlgen:"id"`
	IdNEQ []int `json:"idNEQ" gqlgen:"idNEQ"`

	Key      []string `json:"key" gqlgen:"key"`
	KeyNEQ   []string `json:"keyNEQ" gqlgen:"keyNEQ"`
	KeyMATCH string   `json:"keyMATCH" gqlgen:"keyMATCH"`
	KeyIEQ   string   `json:"keyIEQ" gqlgen:"keyIEQ"`

	Status    []ServerStatus `json:"status" gqlgen:"status"`
	StatusNEQ []ServerStatus `json:"statusNEQ" gqlgen:"statusNEQ"`

	LangVersionTag    []LanguageTag `json:"langVersionTag" gqlgen:"langVersionTag"`
	LangVersionTagNEQ []LanguageTag `json:"langVersionTagNEQ" gqlgen:"langVersionTagNEQ"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
