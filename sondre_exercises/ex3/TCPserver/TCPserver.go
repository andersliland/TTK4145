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
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(connection).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		connection.Write([]byte(newmessage + "\n"))
	}
}
