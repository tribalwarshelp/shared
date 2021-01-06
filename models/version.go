package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type VersionCode string

const (
	VersionCodePL  VersionCode = "pl"
	VersionCodeInt VersionCode = "en"
	VersionCodeDE  VersionCode = "de"
	VersionCodeUK  VersionCode = "uk"
	VersionCodeIT  VersionCode = "it"
	VersionCodeFR  VersionCode = "fr"
	VersionCodeUS  VersionCode = "us"
	VersionCodeNL  VersionCode = "nl"
	VersionCodeES  VersionCode = "es"
	VersionCodeRO  VersionCode = "ro"
	VersionCodeRU  VersionCode = "ru"
	VersionCodeGR  VersionCode = "gr"
	VersionCodeTR  VersionCode = "tr"
	VersionCodeCS  VersionCode = "cs"
	VersionCodeCH  VersionCode = "ch"
	VersionCodePT  VersionCode = "pt"
	VersionCodeBR  VersionCode = "br"
	VersionCodeHU  VersionCode = "hu"
	VersionCodeSK  VersionCode = "sk"
)

func (vc VersionCode) IsValid() bool {
	switch vc {
	case VersionCodePL,
		VersionCodeInt,
		VersionCodeDE,
		VersionCodeBR,
		VersionCodeCH,
		VersionCodeCS,
		VersionCodeES,
		VersionCodeFR,
		VersionCodeGR,
		VersionCodeHU,
		VersionCodeIT,
		VersionCodeNL,
		VersionCodePT,
		VersionCodeRO,
		VersionCodeRU,
		VersionCodeTR,
		VersionCodeUK,
		VersionCodeUS,
		VersionCodeSK:
		return true
	}
	return false
}

func (vc VersionCode) String() string {
	return string(vc)
}

func (vc *VersionCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*vc = VersionCode(strings.ToLower(str))
	if !vc.IsValid() {
		return fmt.Errorf("%s is not a valid VersionCode", str)
	}
	return nil
}

func (vc VersionCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(vc.String()))
}

type Version struct {
	tableName struct{} `pg:"alias:version"`

	Code           VersionCode    `pg:",pk" json:"code" gqlgen:"code"`
	Name           string         `json:"name" gqlgen:"name" pg:",unique"`
	Host           string         `json:"host" gqlgen:"host"`
	Timezone       string         `json:"timezone" gqlgen:"timezone"`
	SpecialServers SpecialServers `json:"specialServers" gqlgen:"specialServers" pg:"rel:has-many"`
}

type VersionFilter struct {
	Code    []VersionCode `json:"code" gqlgen:"code"`
	CodeNEQ []VersionCode `json:"codeNEQ" gqlgen:"codeNEQ"`

	Host      []string `json:"host" gqlgen:"host"`
	HostNEQ   []string `json:"hostNEQ" gqlgen:"hostNEQ"`
	HostMATCH string   `json:"hostMATCH" gqlgen:"hostMATCH"`
	HostIEQ   string   `json:"hostIEQ" gqlgen:"hostIEQ"`
}

func (f *VersionFilter) WhereWithAlias(q *orm.Query, alias string) (*orm.Query, error) {
	if !isZero(f.Code) {
		q = q.Where(buildConditionArray(addAliasToColumnName("code", alias)), pg.Array(f.Code))
	}
	if !isZero(f.CodeNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("code", alias)), pg.Array(f.CodeNEQ))
	}

	if !isZero(f.Host) {
		q = q.Where(buildConditionArray(addAliasToColumnName("host", alias)), pg.Array(f.Host))
	}
	if !isZero(f.HostNEQ) {
		q = q.Where(buildConditionNotInArray(addAliasToColumnName("host", alias)), pg.Array(f.HostNEQ))
	}
	if !isZero(f.HostMATCH) {
		q = q.Where(buildConditionMatch(addAliasToColumnName("host", alias)), f.HostMATCH)
	}
	if !isZero(f.HostIEQ) {
		q = q.Where(buildConditionIEQ(addAliasToColumnName("host", alias)), f.HostIEQ)
	}

	return q, nil
}

func (f *VersionFilter) Where(q *orm.Query) (*orm.Query, error) {
	return f.WhereWithAlias(q, "version")
}
