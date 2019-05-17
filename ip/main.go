package main

import (
	"fmt"
	"net"
)

func currentIp() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = localAddr.IP.String()
	return
}

func main() {
	fmt.Println(currentIp())
}
