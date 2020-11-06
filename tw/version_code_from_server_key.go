package tw

import (
	"github.com/tribalwarshelp/shared/models"
)

func VersionCodeFromServerKey(key string) models.VersionCode {
	if len(key) < 2 {
		return ""
	}
	return models.VersionCode(key[0:2])
}
