package tools

import (
	"log"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Print("获取ip 异常", err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return strings.Split(localAddr.String(), ":")[0]
}
