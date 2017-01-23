package main

import (
	"fmt"
	"net"
	"os"
)

const (
	host = ""
	port = "30000"
)

// CheckError prints error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+port)
	CheckError(err)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buffer := make([]byte, 1024)

	fmt.Println("Listening for messages on port:", port)
	for {
		n, addr, err := ServerConn.ReadFromUDP(buffer)
		CheckError(err)
		fmt.Println("ReadFromUDP ", string(buffer[0:n]), "from ", addr)

		_, ReadFromAddr, err := ServerConn.ReadFrom(buffer)
		fmt.Println("ReadFrom ", string(buffer[0:n]), "from ", ReadFromAddr)

		//fmt.Println("Message sendt")

	}

}
