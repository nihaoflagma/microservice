package utils

import "log"

func LogAction(action string, id int) {
	log.Printf("ACTION=%s USER_ID=%d\n", action, id)
}
