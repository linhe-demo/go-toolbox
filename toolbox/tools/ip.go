package tools

import (
	externalip "github.com/glendc/go-external-ip"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
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

func GetOutBoundIpNew(w http.ResponseWriter, r *http.Request) string {
	// 获取请求的IP地址
	ip := strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0]
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	// 将IP地址转换为net.IP类型
	netIP := net.ParseIP(ip)
	if netIP == nil {
		httpx.Error(w, status.Error(http.StatusBadRequest, "Invalid IP address"))
		return ""
	}
	return netIP.String()
}
