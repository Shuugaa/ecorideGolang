package mode

import (
	"ecoride/database"
)

func KnownUuid (uuid string) bool {
	result := false
	if uuid == "" {
		return result
	}
	_, result = database.CheckUuidExists(uuid)
	return result
}