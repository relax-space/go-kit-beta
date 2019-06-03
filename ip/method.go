package main

import (
	"net"
	"strings"
)

func getCurrentIp(ipParam string) (currentIp string, err error) {
	if len(ipParam) != 0 {
		currentIp = ipParam
		return
	}
	currentIp, err = getIp()
	return
}

func getIp() (currentIp string, err error) {
	ips, err := inIps()
	if err != nil {
		return
	}
	for _, ip := range ips {
		if strings.HasPrefix(ip, "10.202.101.") {
			currentIp = ip
			break
		}
	}
	return
}

func inIps() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	ips = make([]string, 0)
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return
}
