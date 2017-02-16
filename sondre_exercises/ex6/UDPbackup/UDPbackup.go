package UDPbackup

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	host = "localhost"
	port = "20014"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func UDPlisten(listenChan chan byte) {
	buffer := make([]byte, 1)
	address, err := net.ResolveUDPAddr("udp", ":"+port)
	CheckError(err)
	socket, err := net.ListenUDP("udp", address)
	CheckError(err)
	for {
		_, _, err := socket.ReadFromUDP(buffer)
		CheckError(err)
		listenChan <- buffer[0]
	}
}

func UDPsend(counter byte) {
	ServerAddress, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, port))
	CheckError(err)
	connection, err := net.DialUDP("udp", nil, ServerAddress)
	CheckError(err)

	buffer := make([]byte, 1)
	for {
		time.Sleep(1000 * time.Millisecond)
		buffer[0] = counter
		connection.Write(buffer)
	}
}
