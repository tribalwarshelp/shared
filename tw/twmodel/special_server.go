package twmodel

type SpecialServer struct {
	tableName struct{} `pg:"special_servers,alias:special_server"`

	ID          int         `json:"id" gqlgen:"id" xml:"id"`
	VersionCode VersionCode `pg:",unique:group_1" json:"versionCode" gqlgen:"versionCode" xml:"versionCode"`
	Version     *Version    `pg:"fk:version_code,rel:has-one" json:"omitempty,version" gqlgen:"-" xml:"version"`
	Key         string      `pg:",unique:group_1" json:"key" gqlgen:"key" xml:"key"`
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
