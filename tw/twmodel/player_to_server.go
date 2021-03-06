package twmodel

type PlayerToServer struct {
	ID        int     `json:"id" gqlgen:"id" xml:"id"`
	ServerKey string  `pg:",unique:group_1" json:"serverKey" gqlgen:"serverKey" xml:"serverKey"`
	Server    *Server `pg:"fk:server_key,rel:has-one" json:"server" gqlgen:"server" xml:"server"`
	PlayerID  int     `pg:",unique:group_1" json:"playerID" gqlgen:"playerID" xml:"playerID"`
}
