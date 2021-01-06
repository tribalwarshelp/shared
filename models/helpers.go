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

func buildConditionEquals(column string) string {
	return column + "= ?"
}

func buildConditionLT(column string) string {
	return column + "< ?"
}

func buildConditionLTE(column string) string {
	return column + "<= ?"
}

func buildConditionGT(column string) string {
	return column + "> ?"
}

func buildConditionGTE(column string) string {
	return column + ">= ?"
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
