package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"
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

	ID     int          `json:"id" gqlgen:"id"`
	Key    string       `json:"key" gqlgen:"key" pg:",unique"`
	Status ServerStatus `json:"status" gqlgen:"status"`

	LangVersionTag LanguageTag  `json:"langVersionTag" gqlgen:"langVersionTag"`
	LangVersion    *LangVersion `json:"langVersion,omitempty" gqlgen:"-"`
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

	LangVersionTag    []string `json:"langVersionTag" gqlgen:"langVersionTag"`
	LangVersionTagNEQ []string `json:"langVersionTagNEQ" gqlgen:"langVersionTagNEQ"`

	Offset int    `urlstruct:",nowhere" json:"offset" gqlgen:"offset"`
	Limit  int    `urlstruct:",nowhere" json:"limit" gqlgen:"limit"`
	Sort   string `urlstruct:",nowhere" json:"sort" gqlgen:"sort"`
}
