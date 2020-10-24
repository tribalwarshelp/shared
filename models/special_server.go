package models

type SpecialServer struct {
	tableName struct{} `pg:"special_servers,alias:special_server"`

	LangVersionTag LanguageTag  `pg:",unique:group_1" json:"langVersionTag" gqlgen:"langVersionTag" xml:"langVersionTag"`
	LangVersion    *LangVersion `pg:"fk:lang_version_tag" json:"omitempty,langVersion" gqlgen:"-" xml:"langVersion" pg:"rel:has-one"`
	Key            string       `pg:",unique:group_1" json:"key" gqlgen:"key" xml:"key"`
}

type SpecialServers []*SpecialServer

func (servers SpecialServers) Contains(key string) bool {
	for _, s := range servers {
		if s.Key == key {
			return true
		}
	}
	return false
}
