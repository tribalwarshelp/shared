package models

type SpecialServer struct {
	tableName struct{} `pg:"special_servers,alias:special_server"`

	Tag         LanguageTag  `json:"tag" gqlgen:"tag" xml:"tag"`
	LangVersion *LangVersion `json:"omitempty,langVersion" gqlgen:"-" xml:"langVersion"`
	Key         string       `json:"key" gqlgen:"key" xml:"key"`
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
