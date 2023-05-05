package tools

import (
	externalip "github.com/glendc/go-external-ip"
	"log"
)

func GetOutBoundIP() (ip string) {
	consensus := externalip.DefaultConsensus(nil, nil)
	tmpIp, err := consensus.ExternalIP()
	if err != nil {
		log.Print(err)
		return ""
	}
	return tmpIp.String()
}
