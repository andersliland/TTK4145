package main

import (
	"fmt"
	"net"
	"os"
	"time"
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
	ServerAddress, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, udpPort))
	CheckError(err)

	connection, err := net.DialUDP("udp", nil, ServerAddress)
	CheckError(err)

	for {
		connection.Write([]byte("Hello, Mr. Liland"))
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("Message sent")
	}
}
