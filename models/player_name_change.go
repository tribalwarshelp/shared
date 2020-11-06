package models

import "time"

type PlayerNameChange struct {
	VersionCode VersionCode `pg:",unique:group_1" json:"versionCode" gqlgen:"versionCode" xml:"versionCode"`
	Version     *Version    `pg:"fk:version_code,rel:has-one" json:"version" gqlgen:"version" xml:"version"`
	PlayerID    int         `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
	OldName     string      `pg:",unique:group_1" json:"oldName" gqlgen:"oldName" xml:"oldName"`
	NewName     string      `pg:",unique:group_1" json:"newName" gqlgen:"newName" xml:"newName"`
	ChangeDate  time.Time   `pg:"default:CURRENT_DATE,type:DATE,use_zero,unique:group_1" json:"changeDate" gqlgen:"changeDate" xml:"changeDate"`
}
