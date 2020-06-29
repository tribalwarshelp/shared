package models

import "time"

type PlayerNameChange struct {
	LangVersionTag LanguageTag  `pg:",unique:group_1" json:"langVersionTag" gqlgen:"langVersionTag" xml:"langVersionTag"`
	LangVersion    *LangVersion `pg:"fk:lang_version_tag" json:"langVersion" gqlgen:"langVersion" xml:"langVersion"`
	PlayerID       int          `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	OldName        string       `pg:",unique:group_1" json:"oldName" gqlgen:"oldName" xml:"oldName"`
	NewName        string       `pg:",unique:group_1" json:"newName" gqlgen:"newName" xml:"newName"`
	ChangeDate     time.Time    `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"changeDate" gqlgen:"changeDate" xml:"changeDate"`
}
