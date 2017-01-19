package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host    = "127.0.0.255"
	udpPort = "20014"
)

// CheckError ...
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	buff := make([]byte, 1024)
	addr, _ := net.ResolveUDPAddr("udp", ":"+udpPort)
	sock, _ := net.ListenUDP("udp", addr)
	for {
		_, _, err := sock.ReadFromUDP(buff)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buff[:]))
	}
}
