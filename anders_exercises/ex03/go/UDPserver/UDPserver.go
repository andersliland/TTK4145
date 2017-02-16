package main

import (
	"fmt"
	"net"
	"os"
)

const (
	remoteHost = "129.241.187.43"
	remotePort = "20014"
	localHost  = "129.241.187.142"
	localPort  = ""
)

//UDP SERVER 129.241.187.43

// CheckError prints error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+remotePort)
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buffer := make([]byte, 1024)

	fmt.Println("Listening for messages on port:", remotePort)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buffer)
		CheckError(err)
		fmt.Println("ReadFromUDP: ", string(buffer[0:n]), "from ", addr)

		//_, ReadFromAddr, err := ServerConn.ReadFrom(buffer)
		//	fmt.Println("ReadFrom: ", string(buffer[0:n]), "from ", ReadFromAddr)

		//fmt.Println("Message sendt")

	}

}
