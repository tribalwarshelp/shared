package models

type PlayerToServer struct {
	ServerKey string  `pg:",unique:group_1" json:"serverKey" gqlgen:"serverKey" xml:"serverKey"`
	Server    *Server `pg:"fk:server_key" json:"server" gqlgen:"server" xml:"server" pg:"rel:has-one"`
	PlayerID  int     `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
}
