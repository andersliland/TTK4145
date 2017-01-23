package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	host = "10.22.66.249"
	port = "20014"
)

// CheckError ...
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	socket, err := net.Listen("tcp", ":"+port)
	CheckError(err)
	connection, err := socket.Accept()
	CheckError(err)
	for {
		// Listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(connection).ReadString('\n')
		CheckError(err)
		// Output message received
		fmt.Print("Message received:", string(message))
		// Sample process for string received
		newmessage := strings.ToUpper(message)
		// Send new string back to client
		connection.Write([]byte(newmessage + "\n"))
	}
}
