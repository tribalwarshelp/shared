package models

type PlayerToServer struct {
	ServerKey string  `json:"serverKey" gqlgen:"serverKey" xml:"serverKey"`
	Server    *Server `pg:"fk:server_key" json:"server" gqlgen:"server" xml:"server"`
	PlayerID  int     `json:"playerID" gqlgen:"playerID" xml:"playerID"`
}
