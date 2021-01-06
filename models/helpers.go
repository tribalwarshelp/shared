package models

import (
	"strings"
)

func addAliasToColumnName(column, prefix string) string {
	if prefix != "" && !strings.HasPrefix(column, prefix+".") {
		column = wrapStringInDoubleQuotes(prefix) + "." + wrapStringInDoubleQuotes(column)
	} else {
		column = wrapStringInDoubleQuotes(column)
	}
	return column
}

func wrapStringInDoubleQuotes(str string) string {
	return `"` + str + `"`
}

func buildConditionMatch(column string) string {
	return column + "LIKE ?"
}

func buildConditionIEQ(column string) string {
	return column + "ILIKE ?"
}

func buildConditionArray(column string) string {
	return column + " = ANY(?)"
}

func buildConditionNotInArray(column string) string {
	return "NOT (" + buildConditionArray(column) + ")"
}
