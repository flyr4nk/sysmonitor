package confmgr

import (
	"fmt"
	"os"
	"net"
)

type HostInfo struct {
	HostName 	string
	Ipv4Addr	string
}

var hostInfo *HostInfo

func GetIpAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func GetHostInfo() *HostInfo{
	once.Do(func() {
		hostName, _ := os.Hostname()
		hostInfo = &HostInfo {
			HostName:hostName,
			Ipv4Addr:GetIpAddr(),
		}
	})
	return hostInfo
}