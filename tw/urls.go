package tw

import "fmt"

const (
	EndpointTribeProfile   = "/game.php?screen=info_ally&id=%d"
	EndpointPlayerProfile  = "/game.php?screen=info_player&id=%d"
	EndpointVillageProfile = "/game.php?screen=info_village&id=%d"
)

func BuildVillageURL(server, host string, id int) string {
	return fmt.Sprintf("https://%s.%s"+EndpointVillageProfile, server, host, id)
}

func BuildTribeURL(server, host string, id int) string {
	return fmt.Sprintf("https://%s.%s"+EndpointTribeProfile, server, host, id)
}

func BuildPlayerURL(server, host string, id int) string {
	return fmt.Sprintf("https://%s.%s"+EndpointPlayerProfile, server, host, id)
}
