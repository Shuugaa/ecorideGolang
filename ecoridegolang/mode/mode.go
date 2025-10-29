package mode

import (
	"ecoride/database"
	"ecoride/userstructs"
)

func KnownUuid (uuid string) bool {
	result := false
	if uuid == "" {
		return result
	}
	_, result = database.CheckUuidExists(uuid)
	return result
}

func GetUsernameFromUuid (uuid string) string {
	var session userstructs.Session
	session, _ = database.CheckUuidExists(uuid)
	return session.Name
}