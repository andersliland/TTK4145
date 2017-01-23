package main

import (
	"fmt"
	"net"
	"os"
)

// const constants
const (
	ConnectionHost = "localhost"
	ConnectionPort = "34933"
	ConnectionType = "tcp"
)

// CheckError prints error
func CheckError(msg string, err error) {
	if err != nil {
		fmt.Println("Error: ", msg, err)
		os.Exit(0)
	}
}

func main() {
	//Listen for incomming connections
	listener, err := net.Listen(ConnectionType, ConnectionHost+":"+ConnectionPort)
	CheckError("Listener", err)

	// Close the listener when the application closes
	defer listener.Close()

	fmt.Println("Listening on " + ConnectionHost + ":" + ConnectionPort)

	for {
		// Listen for an incomming connections
		conn, err := listener.Accept()
		CheckError("acception:", err)

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
	CheckError("Read", err)

	//Send a response back to preson contacting us
	conn.Write([]byte("Message recieved."))
	//Close the connection when you are done with it
	conn.Close()

}
