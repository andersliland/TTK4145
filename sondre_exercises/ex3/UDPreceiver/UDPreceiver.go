package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host    = "10.22.66.249"
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
	buffer := make([]byte, 1024)
	address, _ := net.ResolveUDPAddr("udp", ":"+udpPort)
	socket, _ := net.ListenUDP("udp", address)
	for {
		_, _, err := socket.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(buffer[:]))
	}
}
