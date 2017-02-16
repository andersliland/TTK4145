package main

import (
	"fmt"
	"net"
	"os"
)

// const constants
const (
	remoteAddr = "localhost"
	remotePort = "34933"
	connType   = "tcp4"
	localAddr  = ""
	localPort  = ""
)

const messageSize = 1 * 1024

// CheckError prints error
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	remoteAddr, err := net.ResolveTCPAddr(connType, remoteAddr+":"+remotePort)
	CheckError(err)

	localAddr, err := net.ResolveTCPAddr(connType, localAddr+":"+localPort)
	CheckError(err)

	//Listen for incomming connections
	listener, err := net.Listen(ConnectionType, ConnectionHost+":"+ConnectionPort)
	CheckError(err)

	// Close the listener when the application closes
	defer listener.Close()

	fmt.Println("Listening on " + ConnectionHost + ":" + ConnectionPort)

	for {
		// Listen for an incomming connections
		conn, err := listener.Accept()
		CheckError(err)

		//Handle connection in a new goroutine
		go handleRequest(conn)
	}
}

// Handle inncomming requests
func handleRequest(conn net.Conn) {
	//Make buffer to hold incomming data
	buffer := make([]byte, 1024)
	//Read the incomming connection to buffer
	_, err := conn.Read(buffer)
	CheckError(err)

	//Send a response back to preson contacting us
	conn.Write([]byte("Message recieved."))
	//Close the connection when you are done with it
	conn.Close()

}
