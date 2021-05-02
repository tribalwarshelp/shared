package urlbuilder

import "fmt"

const (
	EndpointTribeProfile   = "/game.php?screen=info_ally&id=%d"
	EndpointPlayerProfile  = "/game.php?screen=info_player&id=%d"
	EndpointVillageProfile = "/game.php?screen=info_village&id=%d"
)

func BuildServerURL(server, host string) string {
	return fmt.Sprintf("https://%s.%s", server, host)
}

func BuildPlayerURL(server, host string, id int) string {
	return BuildServerURL(server, host) + fmt.Sprintf(EndpointPlayerProfile, id)
}

func BuildVillageURL(server, host string, id int) string {
	return BuildServerURL(server, host) + fmt.Sprintf(EndpointVillageProfile, id)
}

func BuildTribeURL(server, host string, id int) string {
	return BuildServerURL(server, host) + fmt.Sprintf(EndpointTribeProfile, id)
}
