package watchdog

import "time"

type LogInfo struct {
	IP         string
	Action     string
	ActionUser string
	CreateTime time.Timer
}
