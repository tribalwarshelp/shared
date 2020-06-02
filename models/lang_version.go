package models

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type LanguageTag string

const (
	LanguageTagPL  LanguageTag = "pl"
	LanguageTagInt LanguageTag = "en"
	LanguageTagDE  LanguageTag = "de"
)

func (lt LanguageTag) IsValid() bool {
	switch lt {
	case LanguageTagPL,
		LanguageTagInt,
		LanguageTagDE:
		return true
	}
	return false
}

func (lt LanguageTag) String() string {
	return string(lt)
}

func (lt *LanguageTag) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*lt = LanguageTag(strings.ToLower(str))
	if !lt.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageTag", str)
	}
	return nil
}

func (lt LanguageTag) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(lt.String()))
}

type LangVersion struct {
	Tag      LanguageTag `pg:",pk" json:"tag" gqlgen:"tag"`
	Name     string      `json:"name" gqlgen:"name" pg:",unique"`
	Host     string      `json:"host" gqlgen:"host"`
	Timezone string      `json:"timezone" gqlgen:"timezone"`
}
